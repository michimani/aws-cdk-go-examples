package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handleRequest(ctx context.Context, sqsEvent events.SQSEvent) error {
	for _, message := range sqsEvent.Records {
		fmt.Printf("The message %s for event source %s\n", message.MessageId, message.EventSource)
		fmt.Printf("Received message: %#+v", message)

		fmt.Printf("Message body: %s", message.Body)
	}

	return nil
}

func main() {
	lambda.Start(handleRequest)
}
