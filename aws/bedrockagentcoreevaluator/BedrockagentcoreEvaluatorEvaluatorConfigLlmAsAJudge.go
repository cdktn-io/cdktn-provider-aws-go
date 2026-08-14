// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreevaluator


type BedrockagentcoreEvaluatorEvaluatorConfigLlmAsAJudge struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/bedrockagentcore_evaluator#instructions BedrockagentcoreEvaluator#instructions}.
	Instructions *string `field:"required" json:"instructions" yaml:"instructions"`
	// model_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/bedrockagentcore_evaluator#model_config BedrockagentcoreEvaluator#model_config}
	ModelConfig interface{} `field:"optional" json:"modelConfig" yaml:"modelConfig"`
	// rating_scale block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/bedrockagentcore_evaluator#rating_scale BedrockagentcoreEvaluator#rating_scale}
	RatingScale interface{} `field:"optional" json:"ratingScale" yaml:"ratingScale"`
}

