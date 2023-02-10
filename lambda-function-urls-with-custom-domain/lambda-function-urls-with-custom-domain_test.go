package main

import (
	"testing"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/assertions"
	"github.com/aws/jsii-runtime-go"
)

func TestLambdaFunction(t *testing.T) {
	// GIVEN
	app := awscdk.NewApp(nil)

	// WHEN
	stack := NewLambdaFunctionUrlsWithCustomDomainStack(app, "TestStack", nil)

	// THEN
	template := assertions.Template_FromStack(stack, nil)

	// Lambda Function
	type expect struct {
		functionName string
		runtime      string
		handler      string
		memory       float64
		timeout      float64
	}

	cases := []struct {
		name   string
		expect expect
	}{
		{
			name: "simple-response function",
			expect: expect{
				functionName: "aws-cdk-go-example-simple-response",
				runtime:      "go1.x",
				handler:      "main",
				memory:       128,
				timeout:      10,
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			template.HasResourceProperties(jsii.String("AWS::Lambda::Function"), map[string]any{
				"FunctionName": jsii.String(c.expect.functionName),
				"Runtime":      jsii.String(c.expect.runtime),
				"Handler":      jsii.String(c.expect.handler),
				"MemorySize":   c.expect.memory,
				"Timeout":      c.expect.timeout,
			})
		})
	}
}
