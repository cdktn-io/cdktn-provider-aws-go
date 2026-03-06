// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securitylakecustomlogsource


type SecuritylakeCustomLogSourceConfigurationCrawlerConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.35.1/docs/resources/securitylake_custom_log_source#role_arn SecuritylakeCustomLogSource#role_arn}.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
}

