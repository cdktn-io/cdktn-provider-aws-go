// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerhyperparametertuningjob


type SagemakerHyperParameterTuningJobTrainingJobDefinitionsInputDataConfigDataSourceS3DataSource struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.47.0/docs/resources/sagemaker_hyper_parameter_tuning_job#s3_data_type SagemakerHyperParameterTuningJob#s3_data_type}.
	S3DataType *string `field:"required" json:"s3DataType" yaml:"s3DataType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.47.0/docs/resources/sagemaker_hyper_parameter_tuning_job#s3_uri SagemakerHyperParameterTuningJob#s3_uri}.
	S3Uri *string `field:"required" json:"s3Uri" yaml:"s3Uri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.47.0/docs/resources/sagemaker_hyper_parameter_tuning_job#attribute_names SagemakerHyperParameterTuningJob#attribute_names}.
	AttributeNames *[]*string `field:"optional" json:"attributeNames" yaml:"attributeNames"`
	// hub_access_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.47.0/docs/resources/sagemaker_hyper_parameter_tuning_job#hub_access_config SagemakerHyperParameterTuningJob#hub_access_config}
	HubAccessConfig interface{} `field:"optional" json:"hubAccessConfig" yaml:"hubAccessConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.47.0/docs/resources/sagemaker_hyper_parameter_tuning_job#instance_group_names SagemakerHyperParameterTuningJob#instance_group_names}.
	InstanceGroupNames *[]*string `field:"optional" json:"instanceGroupNames" yaml:"instanceGroupNames"`
	// model_access_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.47.0/docs/resources/sagemaker_hyper_parameter_tuning_job#model_access_config SagemakerHyperParameterTuningJob#model_access_config}
	ModelAccessConfig interface{} `field:"optional" json:"modelAccessConfig" yaml:"modelAccessConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.47.0/docs/resources/sagemaker_hyper_parameter_tuning_job#s3_data_distribution_type SagemakerHyperParameterTuningJob#s3_data_distribution_type}.
	S3DataDistributionType *string `field:"optional" json:"s3DataDistributionType" yaml:"s3DataDistributionType"`
}

