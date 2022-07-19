package resource

import (
	"strings"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

func NewCloudWatchLogsLogGroup(stack constructs.Construct, groupName string) awslogs.LogGroup {
	return awslogs.NewLogGroup(stack, jsii.String(toUpperCamelCase(groupName)), &awslogs.LogGroupProps{
		LogGroupName: jsii.String(groupName),
	})
}

// toUpperCamelCase converts a string to an Upper Camel Case.
// e.g) "hello-aws-cdk-golang" -> "HelloAwsCdkGolang"
func toUpperCamelCase(base string) string {
	base = strings.ReplaceAll(base, "-", "_")

	sp := strings.Split(base, "_")
	res := ""
	for _, s := range sp {
		res += strings.ToUpper(s[0:1]) + s[1:]
	}

	return res
}
