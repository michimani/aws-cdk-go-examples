package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdacontext"
)

type response struct {
	StatusCode int               `json:"statusCode"`
	Headers    map[string]string `json:"headers"`
	Body       string            `json:"body"`
}

type errorBody struct {
	Error string `json:"error"`
}

type okBody struct {
	RequestID string `json:"requestId"`
	Message   string `json:"message"`
	Time      string `json:"time"`
}

func handleRequest(ctx context.Context, httpRequest events.APIGatewayProxyRequest) (response, error) {
	lctx, ok := lambdacontext.FromContext(ctx)
	if !ok {
		return jsonResponse(http.StatusInternalServerError, errorBody{Error: "failed to parse lambda context"}, nil)
	}

	if !isAvailableAccess(httpRequest) {
		return jsonResponse(http.StatusForbidden, errorBody{Error: "forbidden"}, nil)
	}

	body := okBody{
		RequestID: lctx.AwsRequestID,
		Message:   "Hello Lambda!",
		Time:      time.Now().Format(time.RFC3339Nano),
	}

	customHeader := map[string]string{
		"x-aws-cdk-example": "lambda-function-urls-with-custom-domain",
	}

	return jsonResponse(http.StatusOK, body, customHeader)
}

const (
	customHeaderKeyFromCloudFront   string = "x-aws-cdk-go-example-from"
	customHeaderValueFromCloudFront string = "aws-cdk-go-example-cf"
)

func isAvailableAccess(req events.APIGatewayProxyRequest) bool {
	// invoke from specified CloudFront
	if h, ok := req.Headers[customHeaderKeyFromCloudFront]; !ok {
		fmt.Printf("custom header %s does not exists. req:%v", customHeaderKeyFromCloudFront, req)
		return false
	} else if h != customHeaderValueFromCloudFront {
		fmt.Printf("custom header value is invalid. req:%v", req)
		return false
	}

	return true
}

func jsonResponse(statusCode int, body any, additionalHeaders map[string]string) (response, error) {
	h := map[string]string{
		"Content-Type": "application/json",
	}

	for k, v := range additionalHeaders {
		h[k] = v
	}

	b, err := json.Marshal(body)
	if err != nil {
		fmt.Println(err.Error())

		return response{
			StatusCode: http.StatusInternalServerError,
			Headers:    h,
			Body:       `{"error": "Internal Server Error"}`,
		}, err
	}

	return response{
		StatusCode: statusCode,
		Headers:    h,
		Body:       string(b),
	}, nil
}

func main() {
	lambda.Start(handleRequest)
}
