// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudwatchlogs3tableintegrationsource


type CloudwatchLogS3TableIntegrationSourceDataSource struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/cloudwatch_log_s3_table_integration_source#name CloudwatchLogS3TableIntegrationSource#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/cloudwatch_log_s3_table_integration_source#type CloudwatchLogS3TableIntegrationSource#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

