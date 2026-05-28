// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakeralgorithm


type SagemakerAlgorithmTrainingSpecificationSupportedHyperParametersRange struct {
	// categorical_parameter_range_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.47.0/docs/resources/sagemaker_algorithm#categorical_parameter_range_specification SagemakerAlgorithm#categorical_parameter_range_specification}
	CategoricalParameterRangeSpecification interface{} `field:"optional" json:"categoricalParameterRangeSpecification" yaml:"categoricalParameterRangeSpecification"`
	// continuous_parameter_range_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.47.0/docs/resources/sagemaker_algorithm#continuous_parameter_range_specification SagemakerAlgorithm#continuous_parameter_range_specification}
	ContinuousParameterRangeSpecification interface{} `field:"optional" json:"continuousParameterRangeSpecification" yaml:"continuousParameterRangeSpecification"`
	// integer_parameter_range_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.47.0/docs/resources/sagemaker_algorithm#integer_parameter_range_specification SagemakerAlgorithm#integer_parameter_range_specification}
	IntegerParameterRangeSpecification interface{} `field:"optional" json:"integerParameterRangeSpecification" yaml:"integerParameterRangeSpecification"`
}

