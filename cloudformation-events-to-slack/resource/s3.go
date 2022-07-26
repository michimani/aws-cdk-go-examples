package resource

import (
	"cloudformation-events-to-slack/util"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

func NewTemporaryBucket(stack constructs.Construct, bucketName string) awss3.Bucket {
	return awss3.NewBucket(stack, jsii.String(util.ToUpperCamelCase(bucketName)), &awss3.BucketProps{
		BucketName:       jsii.String(bucketName),
		PublicReadAccess: jsii.Bool(false),
		RemovalPolicy:    awscdk.RemovalPolicy_DESTROY,
	})
}
