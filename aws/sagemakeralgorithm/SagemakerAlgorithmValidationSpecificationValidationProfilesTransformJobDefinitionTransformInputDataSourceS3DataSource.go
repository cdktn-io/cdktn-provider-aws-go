// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakeralgorithm


type SagemakerAlgorithmValidationSpecificationValidationProfilesTransformJobDefinitionTransformInputDataSourceS3DataSource struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.46.0/docs/resources/sagemaker_algorithm#s3_data_type SagemakerAlgorithm#s3_data_type}.
	S3DataType *string `field:"required" json:"s3DataType" yaml:"s3DataType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.46.0/docs/resources/sagemaker_algorithm#s3_uri SagemakerAlgorithm#s3_uri}.
	S3Uri *string `field:"required" json:"s3Uri" yaml:"s3Uri"`
}

