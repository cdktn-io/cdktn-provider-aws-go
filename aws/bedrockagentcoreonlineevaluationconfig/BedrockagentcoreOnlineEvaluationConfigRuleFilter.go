// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreonlineevaluationconfig


type BedrockagentcoreOnlineEvaluationConfigRuleFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/bedrockagentcore_online_evaluation_config#key BedrockagentcoreOnlineEvaluationConfig#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/bedrockagentcore_online_evaluation_config#operator BedrockagentcoreOnlineEvaluationConfig#operator}.
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// value block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/bedrockagentcore_online_evaluation_config#value BedrockagentcoreOnlineEvaluationConfig#value}
	Value interface{} `field:"optional" json:"value" yaml:"value"`
}

