// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakertrainingjob


type SagemakerTrainingJobInputDataConfigDataSource struct {
	// file_system_data_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/sagemaker_training_job#file_system_data_source SagemakerTrainingJob#file_system_data_source}
	FileSystemDataSource interface{} `field:"optional" json:"fileSystemDataSource" yaml:"fileSystemDataSource"`
	// s3_data_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/sagemaker_training_job#s3_data_source SagemakerTrainingJob#s3_data_source}
	S3DataSource interface{} `field:"optional" json:"s3DataSource" yaml:"s3DataSource"`
}

