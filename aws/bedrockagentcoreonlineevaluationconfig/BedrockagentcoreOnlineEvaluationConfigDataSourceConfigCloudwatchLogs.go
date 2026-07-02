// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreonlineevaluationconfig


type BedrockagentcoreOnlineEvaluationConfigDataSourceConfigCloudwatchLogs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.53.0/docs/resources/bedrockagentcore_online_evaluation_config#log_group_names BedrockagentcoreOnlineEvaluationConfig#log_group_names}.
	LogGroupNames *[]*string `field:"required" json:"logGroupNames" yaml:"logGroupNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.53.0/docs/resources/bedrockagentcore_online_evaluation_config#service_names BedrockagentcoreOnlineEvaluationConfig#service_names}.
	ServiceNames *[]*string `field:"required" json:"serviceNames" yaml:"serviceNames"`
}

