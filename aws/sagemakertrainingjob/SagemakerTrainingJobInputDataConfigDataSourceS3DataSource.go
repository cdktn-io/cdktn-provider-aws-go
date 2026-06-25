// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakertrainingjob


type SagemakerTrainingJobInputDataConfigDataSourceS3DataSource struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/sagemaker_training_job#s3_data_type SagemakerTrainingJob#s3_data_type}.
	S3DataType *string `field:"required" json:"s3DataType" yaml:"s3DataType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/sagemaker_training_job#s3_uri SagemakerTrainingJob#s3_uri}.
	S3Uri *string `field:"required" json:"s3Uri" yaml:"s3Uri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/sagemaker_training_job#attribute_names SagemakerTrainingJob#attribute_names}.
	AttributeNames *[]*string `field:"optional" json:"attributeNames" yaml:"attributeNames"`
	// hub_access_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/sagemaker_training_job#hub_access_config SagemakerTrainingJob#hub_access_config}
	HubAccessConfig interface{} `field:"optional" json:"hubAccessConfig" yaml:"hubAccessConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/sagemaker_training_job#instance_group_names SagemakerTrainingJob#instance_group_names}.
	InstanceGroupNames *[]*string `field:"optional" json:"instanceGroupNames" yaml:"instanceGroupNames"`
	// model_access_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/sagemaker_training_job#model_access_config SagemakerTrainingJob#model_access_config}
	ModelAccessConfig interface{} `field:"optional" json:"modelAccessConfig" yaml:"modelAccessConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/sagemaker_training_job#s3_data_distribution_type SagemakerTrainingJob#s3_data_distribution_type}.
	S3DataDistributionType *string `field:"optional" json:"s3DataDistributionType" yaml:"s3DataDistributionType"`
}

