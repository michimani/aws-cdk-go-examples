package resource

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awseventstargets"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

func NewEventsRuleWithLambdaFunction(scope constructs.Construct, lfn awslambda.Function) {
	awsevents.NewRule(scope, jsii.String("LambdaCronRule"), &awsevents.RuleProps{
		// Daily, every 4 hours
		Schedule: awsevents.Schedule_Cron(&awsevents.CronOptions{
			// Day:     jsii.String("*"),
			Hour:   jsii.String("*/4"),
			Minute: jsii.String("0"),
			// Month:   jsii.String("*"),
			// WeekDay: jsii.String("?"),
			// Year:    jsii.String("*"),
		}),
		Targets: &[]awsevents.IRuleTarget{
			awseventstargets.NewLambdaFunction(lfn, &awseventstargets.LambdaFunctionProps{}),
		},
	})
}
