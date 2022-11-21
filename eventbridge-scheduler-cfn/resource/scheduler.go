package resource

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

func EventBridgeSchedulerForLambdaFunction(scope constructs.Construct, lfn awslambda.Function) {
	role := awsiam.NewRole(scope, jsii.String("IAMRoleForExecutingLambdaFunction"), &awsiam.RoleProps{
		AssumedBy: awsiam.NewServicePrincipal(jsii.String("scheduler.amazonaws.com"), &awsiam.ServicePrincipalOpts{}),
		RoleName:  jsii.String("scheduler-role-for-executing-lambda-function"),
	})

	policy := awsiam.NewPolicyStatement(&awsiam.PolicyStatementProps{
		Actions: &[]*string{
			jsii.String("lambda:InvokeFunction"),
		},
		Resources: &[]*string{
			lfn.FunctionArn(),
		},
	})

	role.AddToPolicy(policy)

	awscdk.NewCfnResource(scope, jsii.String("SchedulerForLambdaFunction"), &awscdk.CfnResourceProps{
		Type: jsii.String("AWS::Scheduler::Schedule"),
		Properties: &map[string]any{
			"Description": jsii.String("EventBridge Scheduler for Lambda Function"),
			"FlexibleTimeWindow": &map[string]any{
				"Mode": jsii.String("OFF"), // "OFF"|"FLEXIBLE"
			},
			"ScheduleExpression": jsii.String("cron(* * */2 * ? *)"), // every 5 minute
			"Target": &map[string]any{
				"Arn":     lfn.FunctionArn(),
				"RoleArn": role.RoleArn(),
			},
		},
	})
}
