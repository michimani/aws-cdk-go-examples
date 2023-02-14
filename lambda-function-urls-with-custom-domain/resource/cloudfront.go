package resource

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscertificatemanager"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfrontorigins"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

const (
	cfnIDPrefix string = "AWSCDKGoExampleFunctionURLFunctionCFn"
)

type NewCloudFrontDistributionInput struct {
	Certificate awscertificatemanager.Certificate
	DomainName  string
	LogBucket   awss3.Bucket
	FunctionURL awslambda.FunctionUrl
}

// separator for function url
var slash = "/"

func NewCloudFrontDistributionForFunctionURLs(scope constructs.Construct, in *NewCloudFrontDistributionInput) awscloudfront.Distribution {
	customHeaderForFunction := map[string]*string{
		"x-aws-cdk-go-example-from": jsii.String("aws-cdk-go-example-cfn"),
	}

	// function url format: https://hoge.lambda-url.ap-northeast-1.on.aws/
	// split: ["https:", "", "hoge.lambda-url.ap-northeast-1.on.aws", ""]
	splitURL := awscdk.Fn_Split(&slash, in.FunctionURL.Url(), jsii.Number(4))
	functionURLDomain := (*splitURL)[2]

	props := &awscloudfront.DistributionProps{
		Enabled: jsii.Bool(true),
		DefaultBehavior: &awscloudfront.BehaviorOptions{
			CachePolicy:         createCachePolicy(scope),
			OriginRequestPolicy: createOriginRequestPolicy(scope),
			Origin: awscloudfrontorigins.NewHttpOrigin(functionURLDomain, &awscloudfrontorigins.HttpOriginProps{
				ConnectionAttempts: jsii.Number(1),
				ConnectionTimeout:  awscdk.Duration_Seconds(jsii.Number(5)),
				CustomHeaders:      &customHeaderForFunction,
				ProtocolPolicy:     awscloudfront.OriginProtocolPolicy_HTTPS_ONLY,
				OriginSslProtocols: &[]awscloudfront.OriginSslPolicy{
					awscloudfront.OriginSslPolicy("TLS_V1_2"),
				},
			}),
			ViewerProtocolPolicy: awscloudfront.ViewerProtocolPolicy_HTTPS_ONLY,
		},
		HttpVersion:   awscloudfront.HttpVersion_HTTP2,
		PriceClass:    awscloudfront.PriceClass_PRICE_CLASS_200,
		EnableLogging: jsii.Bool(false),
	}

	if len(in.DomainName) > 0 {
		props.Certificate = in.Certificate
		props.DomainNames = &[]*string{
			jsii.String(in.DomainName),
		}
	}

	if in.LogBucket != nil {
		props.LogBucket = in.LogBucket
		props.EnableLogging = jsii.Bool(true)
	}

	return awscloudfront.NewDistribution(scope, jsii.String(fmt.Sprintf("%sDistribution", cfnIDPrefix)), props)
}

func createOriginRequestPolicy(scope constructs.Construct) awscloudfront.OriginRequestPolicy {
	return awscloudfront.NewOriginRequestPolicy(scope, jsii.String(fmt.Sprintf("%sOriginRequestPolicy", cfnIDPrefix)), &awscloudfront.OriginRequestPolicyProps{
		OriginRequestPolicyName: jsii.String("aws-cdk-go-example-furl-cfn-orp"),
		HeaderBehavior:          awscloudfront.OriginRequestHeaderBehavior_None(),
	})
}

// cache TTL
const (
	cacheDefault float64 = 120
	cacheMax     float64 = 300
	cacheMin     float64 = 1
)

func createCachePolicy(scope constructs.Construct) awscloudfront.CachePolicy {
	return awscloudfront.NewCachePolicy(scope, jsii.String(fmt.Sprintf("%sCachePolicy", cfnIDPrefix)), &awscloudfront.CachePolicyProps{
		CachePolicyName:            jsii.String("aws-cdk-go-example-furl-cfn-cp"),
		DefaultTtl:                 awscdk.Duration_Seconds(jsii.Number(cacheDefault)),
		MaxTtl:                     awscdk.Duration_Seconds(jsii.Number(cacheMax)),
		MinTtl:                     awscdk.Duration_Seconds(jsii.Number(cacheMin)),
		EnableAcceptEncodingGzip:   jsii.Bool(true),
		EnableAcceptEncodingBrotli: jsii.Bool(true),
		CookieBehavior:             awscloudfront.CacheCookieBehavior_None(),
		HeaderBehavior:             awscloudfront.CacheHeaderBehavior_None(),
		QueryStringBehavior:        awscloudfront.CacheQueryStringBehavior_None(),
	})
}
