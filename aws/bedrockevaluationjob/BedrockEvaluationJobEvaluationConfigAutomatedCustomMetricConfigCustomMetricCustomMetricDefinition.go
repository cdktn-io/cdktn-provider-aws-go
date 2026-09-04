// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockevaluationjob


type BedrockEvaluationJobEvaluationConfigAutomatedCustomMetricConfigCustomMetricCustomMetricDefinition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/bedrock_evaluation_job#instructions BedrockEvaluationJob#instructions}.
	Instructions *string `field:"required" json:"instructions" yaml:"instructions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/bedrock_evaluation_job#name BedrockEvaluationJob#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// rating_scale block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/bedrock_evaluation_job#rating_scale BedrockEvaluationJob#rating_scale}
	RatingScale interface{} `field:"optional" json:"ratingScale" yaml:"ratingScale"`
}

