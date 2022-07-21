package main

import (
	"cloudformation-events-to-slack/resource"
	"os"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
)

type CloudformationEventsToSlackStackProps struct {
	awscdk.StackProps
}

type NotificationTestStack struct {
	awscdk.StackProps
}

const emailEnvKey = "EMAIL_FOR_SUBSCRIBE"
const temporaryBucketNameEnvKey = "TMP_BUCKET_NAME"

func NewCloudformationEventsToSlackStack(scope constructs.Construct, id string, props *CloudformationEventsToSlackStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	email := os.Getenv(emailEnvKey)
	topic := resource.NewSNSTopicForSlackNotification(stack, email)

	resource.NewEventsRuleOfCloudFormationEvents(stack, topic)

	return stack
}

// Create a stack for notification test.
func NewNotificationTestStack(scope constructs.Construct, id string, props *NotificationTestStack) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	bucketName := os.Getenv(temporaryBucketNameEnvKey)
	resource.NewTemporaryBucket(stack, bucketName)

	return stack
}

func main() {
	app := awscdk.NewApp(nil)

	NewCloudformationEventsToSlackStack(app, "CloudformationEventsToSlackStack", &CloudformationEventsToSlackStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	NewNotificationTestStack(app, "NotificationTestStack", &NotificationTestStack{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

// env determines the AWS environment (account+region) in which our stack is to
// be deployed. For more information see: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func env() *awscdk.Environment {
	return nil
}
