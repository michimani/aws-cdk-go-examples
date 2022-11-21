package resource

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

var (
	memory                float64 = 128
	lambdaFunctionTimeout float64 = 30 * 1000 // 30 sec
)

func HelloAWSCDKGolangSchedulerFunction(stack constructs.Construct) awslambda.Function {
	return awslambda.NewFunction(stack, jsii.String("HelloAWSCDKGolangSchedulerFunction"), &awslambda.FunctionProps{
		FunctionName: jsii.String("hello-aws-cdk-golang-scheduler-function"),
		Description:  jsii.String("Hello AWS CDK with Golang!"),
		Runtime:      awslambda.Runtime_GO_1_X(),
		Handler:      jsii.String("main"),
		Code:         awslambda.AssetCode_FromAsset(jsii.String("./src/lambda/hello-function/bin"), nil),
		MemorySize:   &memory,
		Timeout:      awscdk.Duration_Millis(&lambdaFunctionTimeout),
	})
}
