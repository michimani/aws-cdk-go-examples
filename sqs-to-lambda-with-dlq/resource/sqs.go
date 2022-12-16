package resource

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type NewSQSQueueInput struct {
	QueueName  string
	EnabledDLQ bool
	DLQ        awssqs.Queue
}

func NewSQSQueue(stack constructs.Construct, in *NewSQSQueueInput) awssqs.Queue {
	props := awssqs.QueueProps{
		QueueName:         jsii.String(fmt.Sprintf("aws-cdk-go-example-%s", in.QueueName)),
		VisibilityTimeout: awscdk.Duration_Seconds(jsii.Number(31)),
	}

	if in.EnabledDLQ {
		props.DeadLetterQueue = &awssqs.DeadLetterQueue{
			MaxReceiveCount: jsii.Number(3),
			Queue:           in.DLQ,
		}
	}

	return awssqs.NewQueue(stack, jsii.String(in.QueueName), &props)
}
