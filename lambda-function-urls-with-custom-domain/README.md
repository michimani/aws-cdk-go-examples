lambda-function-urls-with-custom-domain
===

# Resources

## Lambda::Function

- Runtime: `go1.x`
- Timeout: 10 sec
- Memory: 128 MB
- Operation:
  - returns simple JSON response

### AWS::Lambda::Url

## CloudFront::Distribution

### CloudFront::OriginRequestPolicy

- HeaderBehavior: none

### AWS::CloudFront::CachePolicy

- DefaultTTL: 120
- MaxTTL: 300
- MinTTL: 1

## Route53:RecordSet

- Name: api.example.com (e.g)
- Type: `A`
- AliasTarget: xxxxxxxxxxx.cloudfront.net

# Components

```mermaid
C4Component


Container_Boundary(r53, "Route 53") {
  Component(r, "RecordSet", "custom domain record","e.g) api.example.com -> xxxxxx.cloudfront.net")
}

Container_Boundary(acm, "Certificate Manager") {
  Component(c, "Certificate", "SSL Certificate", "import by ARN")
}

Container_Boundary(cf,"CloudFront") {
  Component(oap, "Origin Access Policy", "", "")
  Component(dist, "Distribution", "", "")
  Component(cp, "Cache Policy", "", "")
}

Container_Boundary(lambda, "Lambda") {
  Component(url, "Function URL", "", "")
  Component(fn, "Function", "", "")
}

Rel_D(c, dist, "", "")
Rel_U(dist, r, "", "")
Rel_Up(cp, dist,"", "")
Rel_Up(oap, dist,"", "")
Rel(dist, url, "","")
Rel(url, fn, "","")

UpdateElementStyle(c, $bgColor="grey")
```

# Usage

## Create .env file and load

```bash
cp .env.sample .env
```

Fix values to your own.

```bash
source .env
```

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

## Call API

```bash
curl https://<your-own-custom-domain>
```

# Author

[michimani210](https://twitter.com/michimani210)
