eventbridge-scheduler-cfn
===

This is an example implementation of building a Lambda Function that is executed periodically by EventBridge Scheduler. 
Scheduler resources are defined by `CfnResource`.

# Resources

## Lambda::Function

- Runtime: `go1.x`
- Timeout: 30 sec
- Memory: 128 MB
- Operation:
  - Just output `Hello AWS CDK with Golang.` with Request ID and current time to standard output.
  - If invoked as a "RequestResponse", it returns a response similar to the following:

    ```json
    {
      "message": "Hello AWS CDK with Golang.",
      "requestId": "71bd4fd1-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
      "timestamp": 1658236882
    }
    ```

## Scheduler::Schedule

- ScheduleExpression: `cron(* * */2 * ? *)`
- Target: Lambda::Function

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

## Check the execution log

```bash
aws logs get-log-events \
--log-group-name '/aws/lambda/hello-aws-cdk-golang-scheduler-function' \
--log-stream-name "$(
  aws logs describe-log-streams \
  --log-group-name '/aws/lambda/hello-aws-cdk-golang-scheduler-function' \
  --query 'max_by(logStreams[], &lastEventTimestamp).logStreamName' \
  --output text)" \
--limit 20 \
--query 'events[].[join(``, [ to_string(timestamp) ,`: `,message])]' \
--output text
```

# Author

[michimani210](https://twitter.com/michimani210)