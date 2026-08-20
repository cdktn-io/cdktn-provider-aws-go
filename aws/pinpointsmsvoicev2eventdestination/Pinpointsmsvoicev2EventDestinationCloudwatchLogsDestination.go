// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pinpointsmsvoicev2eventdestination


type Pinpointsmsvoicev2EventDestinationCloudwatchLogsDestination struct {
	// ARN of the IAM role that End User Messaging SMS assumes to write to the log group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.61.0/docs/resources/pinpointsmsvoicev2_event_destination#iam_role_arn Pinpointsmsvoicev2EventDestination#iam_role_arn}
	IamRoleArn *string `field:"required" json:"iamRoleArn" yaml:"iamRoleArn"`
	// ARN of the Amazon CloudWatch log group that receives the events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.61.0/docs/resources/pinpointsmsvoicev2_event_destination#log_group_arn Pinpointsmsvoicev2EventDestination#log_group_arn}
	LogGroupArn *string `field:"required" json:"logGroupArn" yaml:"logGroupArn"`
}

