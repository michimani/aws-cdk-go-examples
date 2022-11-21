package main

import (
	"eventbridge-scheduler-cfn/resource"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type EventBridgeSchedulerCfnStackProps struct {
	awscdk.StackProps
}

func NewEventBridgeSchedulerCfnStack(scope constructs.Construct, id string, props *EventBridgeSchedulerCfnStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	// Lambda Function
	lfn := resource.HelloAWSCDKGolangSchedulerFunction(stack)

	// EventBridge Scheduler
	resource.EventBridgeSchedulerForLambdaFunction(stack, lfn)

	return stack
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	NewEventBridgeSchedulerCfnStack(app, "EventBridgeSchedulerCfnStack", &EventBridgeSchedulerCfnStackProps{
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
