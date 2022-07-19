scheduled-lambda-function
===

This is an example implementation of building a Lambda function that is executed periodically by an EventBridge Rule.

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