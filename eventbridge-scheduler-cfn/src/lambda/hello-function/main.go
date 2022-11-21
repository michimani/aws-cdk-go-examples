package main

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdacontext"
)

type HelloResponse struct {
	Message   string `json:"message"`
	RequestID string `json:"requestId"`
	Timestamp int64  `json:"timestamp"`
}

const message = "Hello AWS CDK with Golang."

func handleRequest(ctx context.Context) (HelloResponse, error) {
	now := time.Now()

	lc, _ := lambdacontext.FromContext(ctx)
	reqID := lc.AwsRequestID

	fmt.Printf("[%s][%s] %s", now.Format(time.RFC3339), reqID, message)

	return HelloResponse{
		Message:   message,
		RequestID: reqID,
		Timestamp: now.Unix(),
	}, nil
}

func main() {
	lambda.Start(handleRequest)
}
