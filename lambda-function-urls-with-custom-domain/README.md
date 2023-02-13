lambda-function-urls-with-custom-domain
===

# Resources

## Lambda::Function

- Runtime: `go1.x`
- Timeout: 10 sec
- Memory: 128 MB
- Operation:
  - returns simple JSON response

## AWS::Lambda::Url

## CloudFront::Distribution

## CloudFront::OriginRequestPolicy

## AWS::CloudFront::CachePolicy

# Components

TBD

# Usage

## Build lambda function

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

# Author

[michimani210](https://twitter.com/michimani210)