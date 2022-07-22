package resource

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awseventstargets"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

func NewEventsRuleOfCloudFormationEvents(scope constructs.Construct, st awssns.Topic) {
	region := awscdk.Stack_Of(scope).Region()

	awsevents.NewRule(scope, jsii.String("CloudFormationEventsRule"), &awsevents.RuleProps{
		EventPattern: &awsevents.EventPattern{
			Source: &[]*string{
				jsii.String("aws.cloudformation"),
			},
			DetailType: &[]*string{
				jsii.String("CloudFormation Resource Status Change"),
				jsii.String("CloudFormation Stack Status Change"),
				jsii.String("CloudFormation Drift Detection Status Change"),
			},
			Region: &[]*string{region},
		},
		Targets: &[]awsevents.IRuleTarget{
			awseventstargets.NewSnsTopic(st, &awseventstargets.SnsTopicProps{}),
		},
	})
}
