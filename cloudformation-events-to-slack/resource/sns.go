package resource

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssnssubscriptions"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

func NewSNSTopicForSlackNotification(stack constructs.Construct, email string) awssns.Topic {
	topic := awssns.NewTopic(stack, jsii.String("SNSTopicForSlackNotification"), &awssns.TopicProps{
		TopicName: jsii.String("sns-topic-for-slack-notification"),
	})

	topic.AddSubscription(awssnssubscriptions.NewEmailSubscription(jsii.String(email), &awssnssubscriptions.EmailSubscriptionProps{}))

	return topic
}
