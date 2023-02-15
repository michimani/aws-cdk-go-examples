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
	cfIDPrefix string = "AWSCDKGoExampleFunctionURLFunctionCF"
)

// FunctionURLPattern is struct for config of not default behavior.
type FunctionURLPattern struct {
	FunctionURL awslambda.FunctionUrl
	Pattern     string
}

// NewCloudFrontDistributionInput is struct of input for create CloudFront:Distribution.
type NewCloudFrontDistributionInput struct {
	Certificate            awscertificatemanager.ICertificate
	DomainName             string
	LogBucket              awss3.Bucket
	DefaultFunctionURL     awslambda.FunctionUrl
	AdditionalFunctionURLs []FunctionURLPattern
}

// separator for function url
var slash = "/"

func functionURLDomain(furl awslambda.FunctionUrl) *string {
	// function url format: https://hoge.lambda-url.ap-northeast-1.on.aws/
	// split: ["https:", "", "hoge.lambda-url.ap-northeast-1.on.aws", ""]
	splitURL := awscdk.Fn_Split(&slash, furl.Url(), jsii.Number(4))
	return (*splitURL)[2]
}

// NewCloudFrontDistributionForFunctionURLs creates CloudFront:Distribution.
func NewCloudFrontDistributionForFunctionURLs(scope constructs.Construct, in *NewCloudFrontDistributionInput) awscloudfront.Distribution {
	customHeaderForFunction := map[string]*string{
		"x-aws-cdk-go-example-from": jsii.String("aws-cdk-go-example-cf"),
	}

	defaultFnURLDomain := functionURLDomain(in.DefaultFunctionURL)

	cachePolicy := createCachePolicy(scope)
	originRequestPolicy := createOriginRequestPolicy(scope)

	props := &awscloudfront.DistributionProps{
		Enabled: jsii.Bool(true),
		DefaultBehavior: &awscloudfront.BehaviorOptions{
			CachePolicy:         cachePolicy,
			OriginRequestPolicy: originRequestPolicy,
			Origin: awscloudfrontorigins.NewHttpOrigin(defaultFnURLDomain, &awscloudfrontorigins.HttpOriginProps{
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

	// other behavior
	if len(in.AdditionalFunctionURLs) > 0 {
		additionalBehaviors := map[string]*awscloudfront.BehaviorOptions{}
		for _, fp := range in.AdditionalFunctionURLs {
			furlDomain := functionURLDomain(fp.FunctionURL)

			additionalBehaviors[fp.Pattern] = &awscloudfront.BehaviorOptions{
				CachePolicy:         cachePolicy,
				OriginRequestPolicy: originRequestPolicy,
				Origin: awscloudfrontorigins.NewHttpOrigin(furlDomain, &awscloudfrontorigins.HttpOriginProps{
					ConnectionAttempts: jsii.Number(1),
					ConnectionTimeout:  awscdk.Duration_Seconds(jsii.Number(5)),
					CustomHeaders:      &customHeaderForFunction,
					ProtocolPolicy:     awscloudfront.OriginProtocolPolicy_HTTPS_ONLY,
					OriginSslProtocols: &[]awscloudfront.OriginSslPolicy{
						awscloudfront.OriginSslPolicy("TLS_V1_2"),
					},
				}),
				ViewerProtocolPolicy: awscloudfront.ViewerProtocolPolicy_HTTPS_ONLY,
			}
		}
		props.AdditionalBehaviors = &additionalBehaviors
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

	return awscloudfront.NewDistribution(scope, jsii.String(fmt.Sprintf("%sDistribution", cfIDPrefix)), props)
}

func createOriginRequestPolicy(scope constructs.Construct) awscloudfront.OriginRequestPolicy {
	return awscloudfront.NewOriginRequestPolicy(scope, jsii.String(fmt.Sprintf("%sOriginRequestPolicy", cfIDPrefix)), &awscloudfront.OriginRequestPolicyProps{
		OriginRequestPolicyName: jsii.String("aws-cdk-go-example-furl-cf-orp"),
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
	return awscloudfront.NewCachePolicy(scope, jsii.String(fmt.Sprintf("%sCachePolicy", cfIDPrefix)), &awscloudfront.CachePolicyProps{
		CachePolicyName:            jsii.String("aws-cdk-go-example-furl-cf-cp"),
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
