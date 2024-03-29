aws-cdk-go-examples
===

This is a repository of example implementations of using AWS CDK with the Go language.

# Examples

## cloudformation-events-to-slack

This is an example implementation of a configuration in which specified CloudFormation stack events are notified via email by Amazon SNS.

Please read [cloudformation-events-to-slack/README.md](https://github.com/michimani/aws-cdk-go-examples/blob/main/cloudformation-events-to-slack).

## eventbridge-scheduler-cfn

This is an example implementation of building a Lambda Function that is executed periodically by EventBridge Scheduler. 
Scheduler resources are defined by `CfnResource`.

Please read [eventbridge-scheduler-cfn/README.md](https://github.com/michimani/aws-cdk-go-examples/blob/main/eventbridge-scheduler-cfn).

## lambda-function-urls-with-custom-domain

This is an example implementation of calling a Lambda Function URL with a custom domain.

Please read [lambda-function-urls-with-custom-domain/README.md](https://github.com/michimani/aws-cdk-go-examples/blob/main/lambda-function-urls-with-custom-domain).

## scheduled-lambda-function

This is an example implementation of a Lambda function that is executed periodically according to a schedule specified by a CRON expression.

Please read [scheduled-lambda-function/README.md](https://github.com/michimani/aws-cdk-go-examples/blob/main/scheduled-lambda-function).

## sqs-to-lambda-with-dlq

This is an example implementation of a Lambda function triggered by an SQS message; if the Lambda function fails, the message is sent to the Dead Letter Queue, and another Lambda function processes the message sent to the Dead Letter Queue.

Please read [sqs-to-lambda-with-dlq/README.md](https://github.com/michimani/aws-cdk-go-examples/blob/main/sqs-to-lambda-with-dlq).

## step-functions-with-sdk-integration

This is an example implementation of building a StepFunctions state machine that translates text using some AWS SDK integration.

Please read [step-functions-with-sdk-integration/README.md](https://github.com/michimani/aws-cdk-go-examples/blob/main/step-functions-with-sdk-integration).

# Memo

[AWS CDK を Go で書いたときのメモ - zenn](https://zenn.dev/michimani/scraps/3fb7f8675ef22e)

# License

[MIT](https://github.com/michimani/aws-cdk-go-examples/blob/main/LICENSE)

# Author

[michimani210](https://twitter.com/michimani210)