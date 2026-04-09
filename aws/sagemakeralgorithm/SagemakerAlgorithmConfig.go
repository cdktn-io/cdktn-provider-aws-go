// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakeralgorithm

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SagemakerAlgorithmConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.40.0/docs/resources/sagemaker_algorithm#algorithm_name SagemakerAlgorithm#algorithm_name}.
	AlgorithmName *string `field:"required" json:"algorithmName" yaml:"algorithmName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.40.0/docs/resources/sagemaker_algorithm#algorithm_description SagemakerAlgorithm#algorithm_description}.
	AlgorithmDescription *string `field:"optional" json:"algorithmDescription" yaml:"algorithmDescription"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.40.0/docs/resources/sagemaker_algorithm#certify_for_marketplace SagemakerAlgorithm#certify_for_marketplace}.
	CertifyForMarketplace interface{} `field:"optional" json:"certifyForMarketplace" yaml:"certifyForMarketplace"`
	// inference_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.40.0/docs/resources/sagemaker_algorithm#inference_specification SagemakerAlgorithm#inference_specification}
	InferenceSpecification interface{} `field:"optional" json:"inferenceSpecification" yaml:"inferenceSpecification"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.40.0/docs/resources/sagemaker_algorithm#region SagemakerAlgorithm#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.40.0/docs/resources/sagemaker_algorithm#tags SagemakerAlgorithm#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.40.0/docs/resources/sagemaker_algorithm#timeouts SagemakerAlgorithm#timeouts}
	Timeouts *SagemakerAlgorithmTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// training_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.40.0/docs/resources/sagemaker_algorithm#training_specification SagemakerAlgorithm#training_specification}
	TrainingSpecification interface{} `field:"optional" json:"trainingSpecification" yaml:"trainingSpecification"`
	// validation_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.40.0/docs/resources/sagemaker_algorithm#validation_specification SagemakerAlgorithm#validation_specification}
	ValidationSpecification interface{} `field:"optional" json:"validationSpecification" yaml:"validationSpecification"`
}

