package resource

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

var (
	memory                float64 = 128
	lambdaFunctionTimeout float64 = 5 * 1000 // 30 sec
)

type NewLambdaFunctionInput struct {
	FunctionName string
	CodePath     string
}

func NewLambdaFunction(stack constructs.Construct, in *NewLambdaFunctionInput) awslambda.Function {
	props := awslambda.FunctionProps{
		FunctionName: jsii.String(fmt.Sprintf("aws-cdk-go-example-%s", in.FunctionName)),
		Runtime:      awslambda.Runtime_GO_1_X(),
		Handler:      jsii.String("main"),
		Code:         awslambda.AssetCode_FromAsset(jsii.String(in.CodePath), nil),
		MemorySize:   &memory,
		Timeout:      awscdk.Duration_Millis(&lambdaFunctionTimeout),
	}

	return awslambda.NewFunction(stack, jsii.String(in.FunctionName), &props)
}

func NewSQSToLambdaEventSourceMapping(stack constructs.Construct, id string, lf awslambda.Function, sq awssqs.Queue) {
	lf.AddToRolePolicy(awsiam.NewPolicyStatement(&awsiam.PolicyStatementProps{
		Actions: &[]*string{
			jsii.String("sqs:DeleteMessage"),
			jsii.String("sqs:ReceiveMessage"),
			jsii.String("sqs:GetQueueAttributes"),
		},
		Resources: &[]*string{
			sq.QueueArn(),
		},
	}))

	awslambda.NewEventSourceMapping(stack, jsii.String(id), &awslambda.EventSourceMappingProps{
		BatchSize:      jsii.Number(1),
		Enabled:        jsii.Bool(true),
		EventSourceArn: sq.QueueArn(),
		Target:         lf,
	})
}
