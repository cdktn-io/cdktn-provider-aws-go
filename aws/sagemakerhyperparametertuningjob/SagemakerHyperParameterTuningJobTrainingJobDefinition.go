// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerhyperparametertuningjob


type SagemakerHyperParameterTuningJobTrainingJobDefinition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/sagemaker_hyper_parameter_tuning_job#role_arn SagemakerHyperParameterTuningJob#role_arn}.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// algorithm_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/sagemaker_hyper_parameter_tuning_job#algorithm_specification SagemakerHyperParameterTuningJob#algorithm_specification}
	AlgorithmSpecification interface{} `field:"optional" json:"algorithmSpecification" yaml:"algorithmSpecification"`
	// checkpoint_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/sagemaker_hyper_parameter_tuning_job#checkpoint_config SagemakerHyperParameterTuningJob#checkpoint_config}
	CheckpointConfig interface{} `field:"optional" json:"checkpointConfig" yaml:"checkpointConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/sagemaker_hyper_parameter_tuning_job#definition_name SagemakerHyperParameterTuningJob#definition_name}.
	DefinitionName *string `field:"optional" json:"definitionName" yaml:"definitionName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/sagemaker_hyper_parameter_tuning_job#enable_inter_container_traffic_encryption SagemakerHyperParameterTuningJob#enable_inter_container_traffic_encryption}.
	EnableInterContainerTrafficEncryption interface{} `field:"optional" json:"enableInterContainerTrafficEncryption" yaml:"enableInterContainerTrafficEncryption"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/sagemaker_hyper_parameter_tuning_job#enable_managed_spot_training SagemakerHyperParameterTuningJob#enable_managed_spot_training}.
	EnableManagedSpotTraining interface{} `field:"optional" json:"enableManagedSpotTraining" yaml:"enableManagedSpotTraining"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/sagemaker_hyper_parameter_tuning_job#enable_network_isolation SagemakerHyperParameterTuningJob#enable_network_isolation}.
	EnableNetworkIsolation interface{} `field:"optional" json:"enableNetworkIsolation" yaml:"enableNetworkIsolation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/sagemaker_hyper_parameter_tuning_job#environment SagemakerHyperParameterTuningJob#environment}.
	Environment *map[string]*string `field:"optional" json:"environment" yaml:"environment"`
	// hyper_parameter_ranges block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/sagemaker_hyper_parameter_tuning_job#hyper_parameter_ranges SagemakerHyperParameterTuningJob#hyper_parameter_ranges}
	HyperParameterRanges interface{} `field:"optional" json:"hyperParameterRanges" yaml:"hyperParameterRanges"`
	// hyper_parameter_tuning_resource_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/sagemaker_hyper_parameter_tuning_job#hyper_parameter_tuning_resource_config SagemakerHyperParameterTuningJob#hyper_parameter_tuning_resource_config}
	HyperParameterTuningResourceConfig interface{} `field:"optional" json:"hyperParameterTuningResourceConfig" yaml:"hyperParameterTuningResourceConfig"`
	// input_data_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/sagemaker_hyper_parameter_tuning_job#input_data_config SagemakerHyperParameterTuningJob#input_data_config}
	InputDataConfig interface{} `field:"optional" json:"inputDataConfig" yaml:"inputDataConfig"`
	// output_data_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/sagemaker_hyper_parameter_tuning_job#output_data_config SagemakerHyperParameterTuningJob#output_data_config}
	OutputDataConfig interface{} `field:"optional" json:"outputDataConfig" yaml:"outputDataConfig"`
	// resource_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/sagemaker_hyper_parameter_tuning_job#resource_config SagemakerHyperParameterTuningJob#resource_config}
	ResourceConfig interface{} `field:"optional" json:"resourceConfig" yaml:"resourceConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/sagemaker_hyper_parameter_tuning_job#retry_strategy SagemakerHyperParameterTuningJob#retry_strategy}.
	RetryStrategy interface{} `field:"optional" json:"retryStrategy" yaml:"retryStrategy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/sagemaker_hyper_parameter_tuning_job#static_hyper_parameters SagemakerHyperParameterTuningJob#static_hyper_parameters}.
	StaticHyperParameters *map[string]*string `field:"optional" json:"staticHyperParameters" yaml:"staticHyperParameters"`
	// stopping_condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/sagemaker_hyper_parameter_tuning_job#stopping_condition SagemakerHyperParameterTuningJob#stopping_condition}
	StoppingCondition interface{} `field:"optional" json:"stoppingCondition" yaml:"stoppingCondition"`
	// tuning_objective block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/sagemaker_hyper_parameter_tuning_job#tuning_objective SagemakerHyperParameterTuningJob#tuning_objective}
	TuningObjective interface{} `field:"optional" json:"tuningObjective" yaml:"tuningObjective"`
	// vpc_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/sagemaker_hyper_parameter_tuning_job#vpc_config SagemakerHyperParameterTuningJob#vpc_config}
	VpcConfig interface{} `field:"optional" json:"vpcConfig" yaml:"vpcConfig"`
}

