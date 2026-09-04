// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockevaluationjob


type BedrockEvaluationJobEvaluationConfigHumanHumanWorkflowConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/bedrock_evaluation_job#flow_definition_arn BedrockEvaluationJob#flow_definition_arn}.
	FlowDefinitionArn *string `field:"required" json:"flowDefinitionArn" yaml:"flowDefinitionArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/bedrock_evaluation_job#instructions BedrockEvaluationJob#instructions}.
	Instructions *string `field:"optional" json:"instructions" yaml:"instructions"`
}

