package resource

import (
	"fmt"
	"lambda-function-urls-with-custom-domain/libs"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type NewLambdaFunctionInput struct {
	FunctionName string
	CodePath     string
	Memory       float64
	Timeout      float64
}

// NewLambdaFunction creates Lambda:Function.
func NewLambdaFunction(scope constructs.Construct, in *NewLambdaFunctionInput) awslambda.Function {
	props := &awslambda.FunctionProps{
		FunctionName: jsii.String(fmt.Sprintf("aws-cdk-go-example-%s", in.FunctionName)),
		Runtime:      awslambda.Runtime_GO_1_X(),
		Handler:      jsii.String("main"),
		Code:         awslambda.AssetCode_FromAsset(jsii.String(in.CodePath), nil),
		MemorySize:   jsii.Number(in.Memory),
		Timeout:      awscdk.Duration_Seconds(jsii.Number(in.Timeout)),
	}

	id := jsii.String(libs.ToUpperCamelCase(*props.FunctionName))

	return awslambda.NewFunction(scope, id, props)
}

// NewFunctionURL creates Lambda:URL of lfn.
func NewFunctionURL(scope constructs.Construct, lfn awslambda.Function) awslambda.FunctionUrl {
	props := &awslambda.FunctionUrlProps{
		Function: lfn,
		AuthType: awslambda.FunctionUrlAuthType_NONE,
	}

	id := jsii.String(fmt.Sprintf("%sUrl", libs.ToUpperCamelCase(*lfn.Node().Id())))

	return awslambda.NewFunctionUrl(scope, id, props)
}
