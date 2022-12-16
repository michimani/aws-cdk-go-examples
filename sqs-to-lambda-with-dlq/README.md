sqs-to-lambda-with-dlq
===

This is an example implementation of a Lambda function triggered by an SQS message; if the Lambda function fails, the message is sent to the Dead Letter Queue, and another Lambda function processes the message sent to the Dead Letter Queue.

# Resources

## Lambda::Function

### main

- Runtime: `go1.x`
- Timeout: 30 sec
- Memory: 128 MB
- Operation:
  - Invoked by sending a message to main SQS Queue
  - If `throwError` in the message is `true`, an error is returned. Otherwise, logs a simple message and exits.

### handling dlq message

- Runtime: `go1.x`
- Timeout: 30 sec
- Memory: 128 MB
- Operation:
  - Invoked by sending a message to SQS Queue acts as dead letter queue.
  - Logs a simple message and exits.

## SQS::Queue

### main

- Events the received message and invokes the main Lambda function.

### dlq

- Events the received message and invokes the Lambda function handling dead letter queue message.

# Usage

## Build the function

```bash
make build
```

## Deploy

```bash
cdk synth
```

```bash
cdk deploy
```

## Send message

- No error

  ```bash
  aws sqs send-message \
  --queue-url "$(aws sqs get-queue-url --queue-name 'aws-cdk-go-example-main-queue' --output text)" \
  --message-body "{\"throwError\":false}"
  ```

- Throw error

  ```bash
  aws sqs send-message \
  --queue-url "$(aws sqs get-queue-url --queue-name 'aws-cdk-go-example-main-queue' --output text)" \
  --message-body "{\"throwError\":true}"
  ```

## Check CloudWatch Logs

- Logs of main function

  ```bash
  aws logs get-log-events \
  --log-group-name /aws/lambda/aws-cdk-go-example-main-function \
  --log-stream-name "$(
    aws logs describe-log-streams \
    --log-group-name /aws/lambda/aws-cdk-go-example-main-function \
    --query 'max_by(logStreams[], &lastEventTimestamp).logStreamName' \
    --output text)" \
  --limit 10 \
  --query 'events[].join(``, [ to_string(timestamp) ,`: `,message])' \
  --output text
  ```

- Logs of dlq function

  ```bash
  aws logs get-log-events \
  --log-group-name /aws/lambda/aws-cdk-go-example-dlq-function \
  --log-stream-name "$(
    aws logs describe-log-streams \
    --log-group-name /aws/lambda/aws-cdk-go-example-dlq-function \
    --query 'max_by(logStreams[], &lastEventTimestamp).logStreamName' \
    --output text)" \
  --limit 10 \
  --query 'events[].join(``, [ to_string(timestamp) ,`: `,message])' \
  --output text
  ```

# Author

[michimani210](https://twitter.com/michimani210)