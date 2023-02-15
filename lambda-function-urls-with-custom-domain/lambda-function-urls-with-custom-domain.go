package main

import (
	"lambda-function-urls-with-custom-domain/resource"
	"os"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscertificatemanager"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type LambdaFunctionUrlsWithCustomDomainStackProps struct {
	StackProps       awscdk.StackProps
	CertificateARN   string
	HostZoneID       string
	DomainName       string
	CustomDomainName string
}

func NewLambdaFunctionUrlsWithCustomDomainStack(scope constructs.Construct, id string, props *LambdaFunctionUrlsWithCustomDomainStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	// create lambda function
	lfn := resource.NewLambdaFunction(stack, &resource.NewLambdaFunctionInput{
		FunctionName: "simple-response",
		CodePath:     "./src/lambda/simple-response/bin/default",
		Memory:       128,
		Timeout:      10,
	})

	// create function url for default behavior
	defaultFnURL := resource.NewFunctionURL(stack, lfn)

	// cloudfront distribution
	ceritificate := awscertificatemanager.Certificate_FromCertificateArn(stack, jsii.String("AWSCDKGoExampleFunctionURLFunctionACMCertificate"), &props.CertificateARN)
	dist := resource.NewCloudFrontDistributionForFunctionURLs(stack, &resource.NewCloudFrontDistributionInput{
		Certificate:        ceritificate,
		DomainName:         props.CustomDomainName,
		DefaultFunctionURL: defaultFnURL,
	})

	// route 53 record set
	// resource.ChangeRoute53ResourceRecordSetForCloudFront(stack, dist, props.HostZoneID, props.CustomDomainName)
	resource.ChangeRoute53ResourceRecordSetForCloudFront(stack, dist, props.HostZoneID, props.DomainName, props.CustomDomainName)

	return stack
}

const (
	accountIDEnvKey        string = "AWS_ACCOUNT_ID"
	regionEnvKey           string = "AWS_REGION"
	certificateARNEnvKey   string = "CERTIFICATE_ARN"
	customDomainNameEnvKey string = "CUSTOM_DOMAIN_NAME"
	domainNameEnvKey       string = "DOMAIN_NAME"
	hostZoneIDEnvKey       string = "HOST_ZONE_ID"
)

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	certificateARN := os.Getenv(certificateARNEnvKey)
	customDomainName := os.Getenv(customDomainNameEnvKey)
	domainName := os.Getenv(domainNameEnvKey)
	hostZoneID := os.Getenv(hostZoneIDEnvKey)

	NewLambdaFunctionUrlsWithCustomDomainStack(app, "LambdaFunctionUrlsWithCustomDomainStack", &LambdaFunctionUrlsWithCustomDomainStackProps{
		StackProps: awscdk.StackProps{
			Env: env(),
		},
		CertificateARN:   certificateARN,
		CustomDomainName: customDomainName,
		DomainName:       domainName,
		HostZoneID:       hostZoneID,
	})

	app.Synth(nil)
}

func env() *awscdk.Environment {
	// accountID := os.Getenv(accountIDEnvKey)
	// region := os.Getenv(regionEnvKey)
	// return &awscdk.Environment{
	// 	Account: &accountID,
	// 	Region:  &region,
	// }
	return nil
}
