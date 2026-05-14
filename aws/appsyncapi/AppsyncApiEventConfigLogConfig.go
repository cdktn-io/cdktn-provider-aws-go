// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appsyncapi


type AppsyncApiEventConfigLogConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.45.0/docs/resources/appsync_api#cloudwatch_logs_role_arn AppsyncApi#cloudwatch_logs_role_arn}.
	CloudwatchLogsRoleArn *string `field:"required" json:"cloudwatchLogsRoleArn" yaml:"cloudwatchLogsRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.45.0/docs/resources/appsync_api#log_level AppsyncApi#log_level}.
	LogLevel *string `field:"required" json:"logLevel" yaml:"logLevel"`
}

