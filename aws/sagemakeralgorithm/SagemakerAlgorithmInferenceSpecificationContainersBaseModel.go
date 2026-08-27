// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakeralgorithm


type SagemakerAlgorithmInferenceSpecificationContainersBaseModel struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#hub_content_name SagemakerAlgorithm#hub_content_name}.
	HubContentName *string `field:"optional" json:"hubContentName" yaml:"hubContentName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#hub_content_version SagemakerAlgorithm#hub_content_version}.
	HubContentVersion *string `field:"optional" json:"hubContentVersion" yaml:"hubContentVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#recipe_name SagemakerAlgorithm#recipe_name}.
	RecipeName *string `field:"optional" json:"recipeName" yaml:"recipeName"`
}

