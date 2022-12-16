package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdacontext"
)

type QueueMessage struct {
	ThrowError bool `json:"throwError"`
}

func (qm *QueueMessage) fromString(body string) error {
	return json.Unmarshal([]byte(body), qm)
}

func handleRequest(ctx context.Context, sqsEvent events.SQSEvent) error {
	lc, _ := lambdacontext.FromContext(ctx)
	reqID := lc.AwsRequestID

	for _, message := range sqsEvent.Records {
		fmt.Printf("The message %s for event source %s\n", message.MessageId, message.EventSource)

		qmsg := &QueueMessage{}
		if err := qmsg.fromString(message.Body); err != nil {
			fmt.Printf("Invalid message struct: %s err=%v\n", message.Body, err)
			continue
		}

		if qmsg.ThrowError {
			return errors.New("This function returns an error.")
		}
	}

	fmt.Printf("Current AWS Request ID is '%s'", reqID)

	return nil
}

func main() {
	lambda.Start(handleRequest)
}
