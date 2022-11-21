package main

import (
	"os"
	"testing"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/assertions"
	"github.com/aws/jsii-runtime-go"
)

func TestStepFunctionsWithSdkIntegrationStack_OutputBucket(t *testing.T) {
	app := awscdk.NewApp(nil)
	stack := NewStepFunctionsWithSdkIntegrationStack(app, "TestStack", nil)
	template := assertions.Template_FromStack(stack, &assertions.TemplateParsingOptions{})

	bucketName := os.Getenv(bucketNameEnvKey)

	// S3 Bucket
	template.HasResourceProperties(jsii.String("AWS::S3::Bucket"), map[string]interface{}{
		"BucketName": bucketName,
	})
}

func TestStepFunctionsWithSdkIntegrationStack_LogGroup(t *testing.T) {
	app := awscdk.NewApp(nil)
	stack := NewStepFunctionsWithSdkIntegrationStack(app, "TestStack", nil)
	template := assertions.Template_FromStack(stack, &assertions.TemplateParsingOptions{})

	// CloudWatch Logs LogGroup
	template.HasResourceProperties(jsii.String("AWS::Logs::LogGroup"), map[string]interface{}{
		"LogGroupName": "step-functions-with-sdk-integration-example",
	})
}

func TestStepFunctionsWithSdkIntegrationStack_StateMachine(t *testing.T) {
	app := awscdk.NewApp(nil)
	stack := NewStepFunctionsWithSdkIntegrationStack(app, "TestStack", nil)
	template := assertions.Template_FromStack(stack, &assertions.TemplateParsingOptions{})

	// StepFunctions StateMachine
	template.HasResourceProperties(jsii.String("AWS::StepFunctions::StateMachine"), map[string]interface{}{
		"StateMachineName": "sdk-integration-example-state-machine",
		"StateMachineType": "EXPRESS",
	})
}
