// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreonlineevaluationconfig


type BedrockagentcoreOnlineEvaluationConfigRule struct {
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_online_evaluation_config#filter BedrockagentcoreOnlineEvaluationConfig#filter}
	Filter interface{} `field:"optional" json:"filter" yaml:"filter"`
	// sampling_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_online_evaluation_config#sampling_config BedrockagentcoreOnlineEvaluationConfig#sampling_config}
	SamplingConfig interface{} `field:"optional" json:"samplingConfig" yaml:"samplingConfig"`
	// session_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_online_evaluation_config#session_config BedrockagentcoreOnlineEvaluationConfig#session_config}
	SessionConfig interface{} `field:"optional" json:"sessionConfig" yaml:"sessionConfig"`
}

