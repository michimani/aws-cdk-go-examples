package main

import (
	"testing"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	assertions "github.com/aws/aws-cdk-go/awscdk/v2/assertions"
	"github.com/aws/jsii-runtime-go"
)

func TestScheduledLambdaFunctionStack_HelloFunction(t *testing.T) {
	app := awscdk.NewApp(nil)
	stack := NewScheduledLambdaFunctionStack(app, "TestStack", nil)
	template := assertions.Template_FromStack(stack)

	// Lambda Function
	template.HasResourceProperties(jsii.String("AWS::Lambda::Function"), map[string]interface{}{
		"FunctionName": "hello-aws-cdk-golang-function",
		"Description":  "Hello AWS CDK with Golang!",
		"Runtime":      "go1.x",
		"Handler":      "main",
		"MemorySize":   128,
		"Timeout":      30,
	})
}

func TestScheduledLambdaFunctionStack_EventsRule(t *testing.T) {
	app := awscdk.NewApp(nil)
	stack := NewScheduledLambdaFunctionStack(app, "TestStack", nil)
	template := assertions.Template_FromStack(stack)

	// EventBridge Rule
	template.HasResourceProperties(jsii.String("AWS::Events::Rule"), map[string]interface{}{
		"ScheduleExpression": "cron(0 */4 * * ? *)",
		"State":              "ENABLED",
	})
}
