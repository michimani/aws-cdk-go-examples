package main

import (
	"testing"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/assertions"
	"github.com/aws/jsii-runtime-go"
)

func TestEventBridgeSchedulerCfnStack(t *testing.T) {
	// GIVEN
	app := awscdk.NewApp(nil)

	// WHEN
	stack := NewEventBridgeSchedulerCfnStack(app, "MyStack", nil)

	// THEN
	template := assertions.Template_FromStack(stack, &assertions.TemplateParsingOptions{})

	// Scheduler Schedule
	template.HasResourceProperties(jsii.String("AWS::Scheduler::Schedule"), map[string]any{
		"Description": jsii.String("EventBridge Scheduler for Lambda Function"),
		"FlexibleTimeWindow": &map[string]any{
			"Mode": jsii.String("OFF"), // "OFF"|"FLEXIBLE"
		},
		"ScheduleExpression": jsii.String("cron(* * */2 * ? *)"),
	})

	// Lambda Function
	template.HasResourceProperties(jsii.String("AWS::Lambda::Function"), map[string]any{
		"FunctionName": jsii.String("hello-aws-cdk-golang-scheduler-function"),
		"Description":  jsii.String("Hello AWS CDK with Golang!"),
		"Runtime":      jsii.String("go1.x"),
		"Handler":      jsii.String("main"),
		"MemorySize":   128,
		"Timeout":      30,
	})
}
