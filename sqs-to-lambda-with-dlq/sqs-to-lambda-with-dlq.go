package main

import (
	"sqs-to-lambda-with-dlq/resource"

	"github.com/aws/aws-cdk-go/awscdk/v2"

	// "github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type SqsToLambdaWithDlqStackProps struct {
	awscdk.StackProps
}

func NewSqsToLambdaWithDlqStack(scope constructs.Construct, id string, props *SqsToLambdaWithDlqStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	// dlq to Lambda
	dlq := resource.NewSQSQueue(stack, &resource.NewSQSQueueInput{
		QueueName: "dlq",
	})
	dlqFunction := resource.NewLambdaFunction(stack, &resource.NewLambdaFunctionInput{
		FunctionName: "dlq-function",
		CodePath:     "./src/lambda/dlq-function/bin",
	})
	resource.NewSQSToLambdaEventSourceMapping(stack, "DLQToLambda", dlqFunction, dlq)

	// main SQS to Lambda
	mainQueue := resource.NewSQSQueue(stack, &resource.NewSQSQueueInput{
		QueueName:  "main-queue",
		EnabledDLQ: true,
		DLQ:        dlq,
	})
	mainFunction := resource.NewLambdaFunction(stack, &resource.NewLambdaFunctionInput{
		FunctionName: "main-function",
		CodePath:     "./src/lambda/hello-function/bin",
	})
	resource.NewSQSToLambdaEventSourceMapping(stack, "MainToLambda", mainFunction, mainQueue)

	return stack
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	NewSqsToLambdaWithDlqStack(app, "SqsToLambdaWithDlqStack", &SqsToLambdaWithDlqStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

func env() *awscdk.Environment {
	return nil
}
