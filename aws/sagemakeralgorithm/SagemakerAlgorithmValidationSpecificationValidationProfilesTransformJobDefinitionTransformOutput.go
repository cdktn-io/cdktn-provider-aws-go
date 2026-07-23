// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakeralgorithm


type SagemakerAlgorithmValidationSpecificationValidationProfilesTransformJobDefinitionTransformOutput struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.56.0/docs/resources/sagemaker_algorithm#s3_output_path SagemakerAlgorithm#s3_output_path}.
	S3OutputPath *string `field:"required" json:"s3OutputPath" yaml:"s3OutputPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.56.0/docs/resources/sagemaker_algorithm#accept SagemakerAlgorithm#accept}.
	Accept *string `field:"optional" json:"accept" yaml:"accept"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.56.0/docs/resources/sagemaker_algorithm#assemble_with SagemakerAlgorithm#assemble_with}.
	AssembleWith *string `field:"optional" json:"assembleWith" yaml:"assembleWith"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.56.0/docs/resources/sagemaker_algorithm#kms_key_id SagemakerAlgorithm#kms_key_id}.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
}

