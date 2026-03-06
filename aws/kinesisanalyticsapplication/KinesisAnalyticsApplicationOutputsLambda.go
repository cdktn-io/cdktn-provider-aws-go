// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kinesisanalyticsapplication


type KinesisAnalyticsApplicationOutputsLambda struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.35.1/docs/resources/kinesis_analytics_application#resource_arn KinesisAnalyticsApplication#resource_arn}.
	ResourceArn *string `field:"required" json:"resourceArn" yaml:"resourceArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.35.1/docs/resources/kinesis_analytics_application#role_arn KinesisAnalyticsApplication#role_arn}.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
}

