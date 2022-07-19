scheduled-lambda-function
===

This is an example implementation of building a Lambda function that is executed periodically by an EventBridge Rule.

# Resources

## Lambda::function

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

## Events::Rule

- A schedule rule specified by the following CRON expression.

  ```bash
  0 */4 * * ? *
  ```

- The target of the rule is the above Lambda function.

# Usage

## Build the function

```bash
make build
```

## Test

```bash
go test .
```

## Deploy

```bash
cdk synth
```

```bash
cdk deploy
```

## Invoke manually

```bash
aws lambda invoke \
--function-name hello-aws-cdk-golang-function \
--invocation-type RequestResponse \
--region ap-northeast-1 \
out && cat out
```

# Author

[michimani210](https://twitter.com/michimani210)