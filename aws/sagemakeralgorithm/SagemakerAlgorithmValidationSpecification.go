// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakeralgorithm


type SagemakerAlgorithmValidationSpecification struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.58.0/docs/resources/sagemaker_algorithm#validation_role SagemakerAlgorithm#validation_role}.
	ValidationRole *string `field:"required" json:"validationRole" yaml:"validationRole"`
	// validation_profiles block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.58.0/docs/resources/sagemaker_algorithm#validation_profiles SagemakerAlgorithm#validation_profiles}
	ValidationProfiles interface{} `field:"optional" json:"validationProfiles" yaml:"validationProfiles"`
}

