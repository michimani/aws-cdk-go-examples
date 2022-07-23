package resource

import (
	"cloudformation-events-to-slack/util"
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awseventstargets"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

// Notify all Stack events.
func NewEventsRuleOfAllCloudFormationEvents(scope constructs.Construct, st awssns.Topic) {
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

// Notify only events for specific Stacks matching the resource prefix.
func NewEventsRuleOfSpecifiedCloudFormationEvents(scope constructs.Construct, st awssns.Topic) {
	region := awscdk.Stack_Of(scope).Region()
	accountID := awscdk.Stack_Of(scope).Account()

	awsevents.NewCfnRule(scope, jsii.String("CloudFormationEventsRule"), &awsevents.CfnRuleProps{
		Name:         jsii.String(util.ToKebabCase("events-rule-of-cloud-formation-events")),
		EventBusName: jsii.String("default"),
		State:        jsii.String("ENABLED"),
		EventPattern: &map[string]interface{}{
			"detail-type": &[]*string{
				jsii.String("CloudFormation Resource Status Change"),
				jsii.String("CloudFormation Stack Status Change"),
				jsii.String("CloudFormation Drift Detection Status Change"),
			},
			"region": &[]*string{
				region,
			},
			"resources": &[]interface{}{
				&map[string]*string{
					"prefix": jsii.String(fmt.Sprintf("arn:aws:cloudformation:%s:%s:stack/NotificationTest", *region, *accountID)),
				},
			},
		},
		Targets: &[]interface{}{
			&map[string]*string{
				"arn": st.TopicArn(),
				"id":  jsii.String("Target0"),
			},
		},
	})
}
