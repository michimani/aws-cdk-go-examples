package main

import (
	"lambda-function-urls-with-custom-domain/resource"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type LambdaFunctionUrlsWithCustomDomainStackProps struct {
	awscdk.StackProps
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
		CodePath:     "./src/lambda/simple-response/bin",
		Memory:       128,
		Timeout:      10,
	})

	// create function url
	furl := resource.NewFunctionURL(stack, lfn)

	// cloudfront distribution
	resource.CreateCloudFrontDistributionForFunctionURLs(stack, &resource.CreateCloudFrontDistributionInput{
		FunctionURL: furl,
	})

	return stack
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	NewLambdaFunctionUrlsWithCustomDomainStack(app, "LambdaFunctionUrlsWithCustomDomainStack", &LambdaFunctionUrlsWithCustomDomainStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

func env() *awscdk.Environment {
	return nil
}
