package resource

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssnssubscriptions"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

func NewSNSTopicForSlackNotification(stack constructs.Construct, email string) awssns.Topic {
	topic := awssns.NewTopic(stack, jsii.String("SNSTopicForSlackNotification"), &awssns.TopicProps{
		TopicName: jsii.String("sns-topic-for-slack-notification"),
	})

	// SNS:Topic
	topic.AddSubscription(awssnssubscriptions.NewEmailSubscription(jsii.String(email), &awssnssubscriptions.EmailSubscriptionProps{}))

	// SNS:TopicPolicy
	awssns.NewTopicPolicy(stack, jsii.String("SNSTopicPolicyForSlackNotification"), &awssns.TopicPolicyProps{
		PolicyDocument: awsiam.NewPolicyDocument(&awsiam.PolicyDocumentProps{
			Statements: &[]awsiam.PolicyStatement{
				awsiam.NewPolicyStatement(&awsiam.PolicyStatementProps{
					Actions: &[]*string{
						jsii.String("sns:Publish"),
					},
					Effect: awsiam.Effect_ALLOW,
					Principals: &[]awsiam.IPrincipal{
						awsiam.NewServicePrincipal(
							jsii.String("events.amazonaws.com"),
							&awsiam.ServicePrincipalOpts{},
						),
					},
					Resources: &[]*string{
						topic.TopicArn(),
					},
				}),
			},
		}),
		Topics: &[]awssns.ITopic{
			topic,
		},
	})

	return topic
}
