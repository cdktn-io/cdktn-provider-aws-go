// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakeralgorithm


type SagemakerAlgorithmInferenceSpecificationContainersModelDataSourceS3DataSource struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.40.0/docs/resources/sagemaker_algorithm#compression_type SagemakerAlgorithm#compression_type}.
	CompressionType *string `field:"required" json:"compressionType" yaml:"compressionType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.40.0/docs/resources/sagemaker_algorithm#s3_data_type SagemakerAlgorithm#s3_data_type}.
	S3DataType *string `field:"required" json:"s3DataType" yaml:"s3DataType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.40.0/docs/resources/sagemaker_algorithm#s3_uri SagemakerAlgorithm#s3_uri}.
	S3Uri *string `field:"required" json:"s3Uri" yaml:"s3Uri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.40.0/docs/resources/sagemaker_algorithm#etag SagemakerAlgorithm#etag}.
	Etag *string `field:"optional" json:"etag" yaml:"etag"`
	// hub_access_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.40.0/docs/resources/sagemaker_algorithm#hub_access_config SagemakerAlgorithm#hub_access_config}
	HubAccessConfig interface{} `field:"optional" json:"hubAccessConfig" yaml:"hubAccessConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.40.0/docs/resources/sagemaker_algorithm#manifest_etag SagemakerAlgorithm#manifest_etag}.
	ManifestEtag *string `field:"optional" json:"manifestEtag" yaml:"manifestEtag"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.40.0/docs/resources/sagemaker_algorithm#manifest_s3_uri SagemakerAlgorithm#manifest_s3_uri}.
	ManifestS3Uri *string `field:"optional" json:"manifestS3Uri" yaml:"manifestS3Uri"`
	// model_access_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.40.0/docs/resources/sagemaker_algorithm#model_access_config SagemakerAlgorithm#model_access_config}
	ModelAccessConfig interface{} `field:"optional" json:"modelAccessConfig" yaml:"modelAccessConfig"`
}

