// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakeralgorithm


type SagemakerAlgorithmInferenceSpecificationContainers struct {
	// additional_s3_data_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.45.0/docs/resources/sagemaker_algorithm#additional_s3_data_source SagemakerAlgorithm#additional_s3_data_source}
	AdditionalS3DataSource interface{} `field:"optional" json:"additionalS3DataSource" yaml:"additionalS3DataSource"`
	// base_model block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.45.0/docs/resources/sagemaker_algorithm#base_model SagemakerAlgorithm#base_model}
	BaseModel interface{} `field:"optional" json:"baseModel" yaml:"baseModel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.45.0/docs/resources/sagemaker_algorithm#container_hostname SagemakerAlgorithm#container_hostname}.
	ContainerHostname *string `field:"optional" json:"containerHostname" yaml:"containerHostname"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.45.0/docs/resources/sagemaker_algorithm#environment SagemakerAlgorithm#environment}.
	Environment *map[string]*string `field:"optional" json:"environment" yaml:"environment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.45.0/docs/resources/sagemaker_algorithm#framework SagemakerAlgorithm#framework}.
	Framework *string `field:"optional" json:"framework" yaml:"framework"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.45.0/docs/resources/sagemaker_algorithm#framework_version SagemakerAlgorithm#framework_version}.
	FrameworkVersion *string `field:"optional" json:"frameworkVersion" yaml:"frameworkVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.45.0/docs/resources/sagemaker_algorithm#image SagemakerAlgorithm#image}.
	Image *string `field:"optional" json:"image" yaml:"image"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.45.0/docs/resources/sagemaker_algorithm#image_digest SagemakerAlgorithm#image_digest}.
	ImageDigest *string `field:"optional" json:"imageDigest" yaml:"imageDigest"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.45.0/docs/resources/sagemaker_algorithm#is_checkpoint SagemakerAlgorithm#is_checkpoint}.
	IsCheckpoint interface{} `field:"optional" json:"isCheckpoint" yaml:"isCheckpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.45.0/docs/resources/sagemaker_algorithm#model_data_etag SagemakerAlgorithm#model_data_etag}.
	ModelDataEtag *string `field:"optional" json:"modelDataEtag" yaml:"modelDataEtag"`
	// model_data_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.45.0/docs/resources/sagemaker_algorithm#model_data_source SagemakerAlgorithm#model_data_source}
	ModelDataSource interface{} `field:"optional" json:"modelDataSource" yaml:"modelDataSource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.45.0/docs/resources/sagemaker_algorithm#model_data_url SagemakerAlgorithm#model_data_url}.
	ModelDataUrl *string `field:"optional" json:"modelDataUrl" yaml:"modelDataUrl"`
	// model_input block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.45.0/docs/resources/sagemaker_algorithm#model_input SagemakerAlgorithm#model_input}
	ModelInput interface{} `field:"optional" json:"modelInput" yaml:"modelInput"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.45.0/docs/resources/sagemaker_algorithm#nearest_model_name SagemakerAlgorithm#nearest_model_name}.
	NearestModelName *string `field:"optional" json:"nearestModelName" yaml:"nearestModelName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.45.0/docs/resources/sagemaker_algorithm#product_id SagemakerAlgorithm#product_id}.
	ProductId *string `field:"optional" json:"productId" yaml:"productId"`
}

