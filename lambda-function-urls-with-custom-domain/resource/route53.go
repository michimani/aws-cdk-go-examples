package resource

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53targets"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

const (
	route53IDPrefix string = "AWSCDKGoExampleFunctionURLFunctionRoute53"
)

// ChangeRoute53ResourceRecordSetForCloudFront create Route53:RecordSet that connect custom domain to CloudFront domain.
func ChangeRoute53ResourceRecordSetForCloudFront(scope constructs.Construct, dist awscloudfront.Distribution, hostZoneID string, domainName string, subDomain string) {
	hostZone := awsroute53.HostedZone_FromHostedZoneAttributes(scope, jsii.String(fmt.Sprintf("%sMyHostZone", route53IDPrefix)), &awsroute53.HostedZoneAttributes{
		HostedZoneId: &hostZoneID,
		ZoneName:     &domainName,
	})
	aliasTarget := awsroute53.RecordTarget_FromAlias(
		awsroute53targets.NewCloudFrontTarget(dist),
	)

	props := &awsroute53.RecordSetProps{
		RecordName: jsii.String(subDomain),
		RecordType: awsroute53.RecordType_A,
		Zone:       hostZone,
		Target:     aliasTarget,
	}

	awsroute53.NewRecordSet(scope, jsii.String(fmt.Sprintf("%sRecordSetForCloudFront", route53IDPrefix)), props)
}
