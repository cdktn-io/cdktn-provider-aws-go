// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerhyperparametertuningjob


type SagemakerHyperParameterTuningJobTrainingJobDefinitionInputDataConfigDataSource struct {
	// file_system_data_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.49.0/docs/resources/sagemaker_hyper_parameter_tuning_job#file_system_data_source SagemakerHyperParameterTuningJob#file_system_data_source}
	FileSystemDataSource interface{} `field:"optional" json:"fileSystemDataSource" yaml:"fileSystemDataSource"`
	// s3_data_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.49.0/docs/resources/sagemaker_hyper_parameter_tuning_job#s3_data_source SagemakerHyperParameterTuningJob#s3_data_source}
	S3DataSource interface{} `field:"optional" json:"s3DataSource" yaml:"s3DataSource"`
}

