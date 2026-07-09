// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerhyperparametertuningjob


type SagemakerHyperParameterTuningJobTrainingJobDefinitionInputDataConfigDataSourceFileSystemDataSource struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.54.0/docs/resources/sagemaker_hyper_parameter_tuning_job#directory_path SagemakerHyperParameterTuningJob#directory_path}.
	DirectoryPath *string `field:"required" json:"directoryPath" yaml:"directoryPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.54.0/docs/resources/sagemaker_hyper_parameter_tuning_job#file_system_access_mode SagemakerHyperParameterTuningJob#file_system_access_mode}.
	FileSystemAccessMode *string `field:"required" json:"fileSystemAccessMode" yaml:"fileSystemAccessMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.54.0/docs/resources/sagemaker_hyper_parameter_tuning_job#file_system_id SagemakerHyperParameterTuningJob#file_system_id}.
	FileSystemId *string `field:"required" json:"fileSystemId" yaml:"fileSystemId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.54.0/docs/resources/sagemaker_hyper_parameter_tuning_job#file_system_type SagemakerHyperParameterTuningJob#file_system_type}.
	FileSystemType *string `field:"required" json:"fileSystemType" yaml:"fileSystemType"`
}

