package resource

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctionstasks"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

var (
	stateMachineTimeout float64 = 300 * 1000 // 300 sec
)

// TODO: set by state machine input
const outputKey = "translate-result"

func NewSDKIntegrationExampleStateMachine(stack constructs.Construct, outBucket awss3.Bucket) {
	initSt := awsstepfunctions.NewPass(stack, jsii.String("init"), &awsstepfunctions.PassProps{
		Comment: jsii.String("init state"),
	})

	// Translate state
	translateResultSelector := map[string]interface{}{
		"translatedText.$": "$.TranslatedText",
	}
	translateState := awsstepfunctionstasks.NewCallAwsService(stack, jsii.String("TranslateSDKIntegration"), &awsstepfunctionstasks.CallAwsServiceProps{
		Service: jsii.String("Translate"),
		Action:  jsii.String("translateText"),
		IamResources: &[]*string{
			jsii.String("*"),
		},
		IamAction: jsii.String("translate:TranslateText"),
		Parameters: &map[string]interface{}{
			"SourceLanguageCode": awsstepfunctions.JsonPath_StringAt(jsii.String("$.sourceLang")),
			"TargetLanguageCode": awsstepfunctions.JsonPath_StringAt(jsii.String("$.targetLang")),
			"Text":               awsstepfunctions.JsonPath_StringAt(jsii.String("$.inputText")),
		},
		ResultSelector: &translateResultSelector,
	})

	// Output state
	outputState := awsstepfunctionstasks.NewCallAwsService(stack, jsii.String("S3SDKIntegration"), &awsstepfunctionstasks.CallAwsServiceProps{
		Service: jsii.String("S3"),
		Action:  jsii.String("putObject"),
		IamResources: &[]*string{
			jsii.String(fmt.Sprintf("%s/*", *outBucket.BucketArn())),
		},
		IamAction: jsii.String("s3:PutObject"),
		Parameters: &map[string]interface{}{
			"Bucket":      outBucket.BucketName(),
			"Key":         outputKey,
			"Body":        awsstepfunctions.JsonPath_StringAt(jsii.String("$.translatedText")),
			"ContentType": jsii.String("text/plain"),
		},
		ResultPath: awsstepfunctions.JsonPath_DISCARD(),
	})

	definition := initSt.Next(translateState).Next(outputState)

	logGroup := NewCloudWatchLogsLogGroup(stack, "step-functions-with-sdk-integration-example")

	awsstepfunctions.NewStateMachine(stack, jsii.String("SDKIntegrationExampleStateMachine"), &awsstepfunctions.StateMachineProps{
		StateMachineName: jsii.String("sdk-integration-example-state-machine"),
		StateMachineType: awsstepfunctions.StateMachineType_EXPRESS,
		Timeout:          awscdk.Duration_Millis(&stateMachineTimeout),
		Definition:       definition,
		Logs: &awsstepfunctions.LogOptions{
			Destination: logGroup,
			Level:       awsstepfunctions.LogLevel_ALL,
		},
	})
}
