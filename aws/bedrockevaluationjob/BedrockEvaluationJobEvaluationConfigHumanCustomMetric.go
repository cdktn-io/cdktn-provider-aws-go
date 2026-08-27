// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockevaluationjob


type BedrockEvaluationJobEvaluationConfigHumanCustomMetric struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_evaluation_job#name BedrockEvaluationJob#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_evaluation_job#rating_method BedrockEvaluationJob#rating_method}.
	RatingMethod *string `field:"required" json:"ratingMethod" yaml:"ratingMethod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_evaluation_job#description BedrockEvaluationJob#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

