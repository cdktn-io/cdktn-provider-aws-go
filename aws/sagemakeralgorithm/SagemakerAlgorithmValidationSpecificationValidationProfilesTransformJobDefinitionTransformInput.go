// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakeralgorithm


type SagemakerAlgorithmValidationSpecificationValidationProfilesTransformJobDefinitionTransformInput struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.46.0/docs/resources/sagemaker_algorithm#compression_type SagemakerAlgorithm#compression_type}.
	CompressionType *string `field:"optional" json:"compressionType" yaml:"compressionType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.46.0/docs/resources/sagemaker_algorithm#content_type SagemakerAlgorithm#content_type}.
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// data_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.46.0/docs/resources/sagemaker_algorithm#data_source SagemakerAlgorithm#data_source}
	DataSource interface{} `field:"optional" json:"dataSource" yaml:"dataSource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.46.0/docs/resources/sagemaker_algorithm#split_type SagemakerAlgorithm#split_type}.
	SplitType *string `field:"optional" json:"splitType" yaml:"splitType"`
}

