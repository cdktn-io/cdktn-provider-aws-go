// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakeralgorithm


type SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/sagemaker_algorithm#training_input_mode SagemakerAlgorithm#training_input_mode}.
	TrainingInputMode *string `field:"required" json:"trainingInputMode" yaml:"trainingInputMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/sagemaker_algorithm#hyper_parameters SagemakerAlgorithm#hyper_parameters}.
	HyperParameters *map[string]*string `field:"optional" json:"hyperParameters" yaml:"hyperParameters"`
	// input_data_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/sagemaker_algorithm#input_data_config SagemakerAlgorithm#input_data_config}
	InputDataConfig interface{} `field:"optional" json:"inputDataConfig" yaml:"inputDataConfig"`
	// output_data_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/sagemaker_algorithm#output_data_config SagemakerAlgorithm#output_data_config}
	OutputDataConfig interface{} `field:"optional" json:"outputDataConfig" yaml:"outputDataConfig"`
	// resource_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/sagemaker_algorithm#resource_config SagemakerAlgorithm#resource_config}
	ResourceConfig interface{} `field:"optional" json:"resourceConfig" yaml:"resourceConfig"`
	// stopping_condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/sagemaker_algorithm#stopping_condition SagemakerAlgorithm#stopping_condition}
	StoppingCondition interface{} `field:"optional" json:"stoppingCondition" yaml:"stoppingCondition"`
}

