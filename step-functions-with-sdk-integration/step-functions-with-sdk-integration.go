package main

import (
	"os"
	"step-functions-with-sdk-integration/resource"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
)

type StepFunctionsWithSdkIntegrationStackProps struct {
	awscdk.StackProps
}

const bucketNameEnvKey = "OUTPUT_BUCKET_NAME"

func NewStepFunctionsWithSdkIntegrationStack(scope constructs.Construct, id string, props *StepFunctionsWithSdkIntegrationStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	bucketName := os.Getenv(bucketNameEnvKey)

	// S3 Bucket for output
	outputBucket := resource.NewStepFunctionsOutputBucket(stack, bucketName)

	// StepFunctions State Machine
	resource.NewSDKIntegrationExampleStateMachine(stack, outputBucket)

	return stack
}

func main() {
	app := awscdk.NewApp(nil)

	NewStepFunctionsWithSdkIntegrationStack(app, "StepFunctionsWithSdkIntegrationStack", &StepFunctionsWithSdkIntegrationStackProps{
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
