// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakeralgorithm


type SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/sagemaker_algorithm#channel_name SagemakerAlgorithm#channel_name}.
	ChannelName *string `field:"required" json:"channelName" yaml:"channelName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/sagemaker_algorithm#compression_type SagemakerAlgorithm#compression_type}.
	CompressionType *string `field:"optional" json:"compressionType" yaml:"compressionType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/sagemaker_algorithm#content_type SagemakerAlgorithm#content_type}.
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// data_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/sagemaker_algorithm#data_source SagemakerAlgorithm#data_source}
	DataSource interface{} `field:"optional" json:"dataSource" yaml:"dataSource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/sagemaker_algorithm#input_mode SagemakerAlgorithm#input_mode}.
	InputMode *string `field:"optional" json:"inputMode" yaml:"inputMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/sagemaker_algorithm#record_wrapper_type SagemakerAlgorithm#record_wrapper_type}.
	RecordWrapperType *string `field:"optional" json:"recordWrapperType" yaml:"recordWrapperType"`
	// shuffle_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/sagemaker_algorithm#shuffle_config SagemakerAlgorithm#shuffle_config}
	ShuffleConfig interface{} `field:"optional" json:"shuffleConfig" yaml:"shuffleConfig"`
}

