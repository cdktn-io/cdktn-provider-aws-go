// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakeralgorithm


type SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigDataSourceFileSystemDataSource struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.58.0/docs/resources/sagemaker_algorithm#directory_path SagemakerAlgorithm#directory_path}.
	DirectoryPath *string `field:"required" json:"directoryPath" yaml:"directoryPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.58.0/docs/resources/sagemaker_algorithm#file_system_access_mode SagemakerAlgorithm#file_system_access_mode}.
	FileSystemAccessMode *string `field:"required" json:"fileSystemAccessMode" yaml:"fileSystemAccessMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.58.0/docs/resources/sagemaker_algorithm#file_system_id SagemakerAlgorithm#file_system_id}.
	FileSystemId *string `field:"required" json:"fileSystemId" yaml:"fileSystemId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.58.0/docs/resources/sagemaker_algorithm#file_system_type SagemakerAlgorithm#file_system_type}.
	FileSystemType *string `field:"required" json:"fileSystemType" yaml:"fileSystemType"`
}

