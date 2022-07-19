package main

import (
	"scheduled-lambda-function/resource"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
)

type ScheduledLambdaFunctionStackProps struct {
	awscdk.StackProps
}

func NewScheduledLambdaFunctionStack(scope constructs.Construct, id string, props *ScheduledLambdaFunctionStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	// Lambda Function
	lfn := resource.HelloAWSCDKGolangFunction(stack)

	// EventBridge Event Rule
	resource.NewEventsRuleWithLambdaFunction(stack, lfn)

	return stack
}

func main() {
	app := awscdk.NewApp(nil)

	NewScheduledLambdaFunctionStack(app, "ScheduledLambdaFunctionStack", &ScheduledLambdaFunctionStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

// env determines the AWS environment (account+region) in which our stack is to
// be deployed. For more information see: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func env() *awscdk.Environment {
	return nil
}
