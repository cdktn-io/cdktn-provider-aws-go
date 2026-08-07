// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakeralgorithm


type SagemakerAlgorithmValidationSpecificationValidationProfilesTransformJobDefinition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.58.0/docs/resources/sagemaker_algorithm#batch_strategy SagemakerAlgorithm#batch_strategy}.
	BatchStrategy *string `field:"optional" json:"batchStrategy" yaml:"batchStrategy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.58.0/docs/resources/sagemaker_algorithm#environment SagemakerAlgorithm#environment}.
	Environment *map[string]*string `field:"optional" json:"environment" yaml:"environment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.58.0/docs/resources/sagemaker_algorithm#max_concurrent_transforms SagemakerAlgorithm#max_concurrent_transforms}.
	MaxConcurrentTransforms *float64 `field:"optional" json:"maxConcurrentTransforms" yaml:"maxConcurrentTransforms"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.58.0/docs/resources/sagemaker_algorithm#max_payload_in_mb SagemakerAlgorithm#max_payload_in_mb}.
	MaxPayloadInMb *float64 `field:"optional" json:"maxPayloadInMb" yaml:"maxPayloadInMb"`
	// transform_input block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.58.0/docs/resources/sagemaker_algorithm#transform_input SagemakerAlgorithm#transform_input}
	TransformInput interface{} `field:"optional" json:"transformInput" yaml:"transformInput"`
	// transform_output block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.58.0/docs/resources/sagemaker_algorithm#transform_output SagemakerAlgorithm#transform_output}
	TransformOutput interface{} `field:"optional" json:"transformOutput" yaml:"transformOutput"`
	// transform_resources block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.58.0/docs/resources/sagemaker_algorithm#transform_resources SagemakerAlgorithm#transform_resources}
	TransformResources interface{} `field:"optional" json:"transformResources" yaml:"transformResources"`
}

