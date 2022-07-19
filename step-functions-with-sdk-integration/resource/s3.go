package resource

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

func NewStepFunctionsOutputBucket(stack constructs.Construct, bucketName string) awss3.Bucket {
	return awss3.NewBucket(stack, jsii.String("StepFunctionsOutputBucket"), &awss3.BucketProps{
		BucketName:       jsii.String(bucketName),
		PublicReadAccess: jsii.Bool(false),
	})
}
