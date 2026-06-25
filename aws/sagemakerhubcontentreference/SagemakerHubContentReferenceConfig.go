// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerhubcontentreference

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SagemakerHubContentReferenceConfig struct {
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
	// Name of the hub content reference.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/sagemaker_hub_content_reference#hub_content_name SagemakerHubContentReference#hub_content_name}
	HubContentName *string `field:"required" json:"hubContentName" yaml:"hubContentName"`
	// Name of the private SageMaker Hub to add the content reference to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/sagemaker_hub_content_reference#hub_name SagemakerHubContentReference#hub_name}
	HubName *string `field:"required" json:"hubName" yaml:"hubName"`
	// ARN of the public SageMaker JumpStart hub content to reference. The ARN must not include a version suffix.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/sagemaker_hub_content_reference#sagemaker_public_hub_content_arn SagemakerHubContentReference#sagemaker_public_hub_content_arn}
	SagemakerPublicHubContentArn *string `field:"required" json:"sagemakerPublicHubContentArn" yaml:"sagemakerPublicHubContentArn"`
	// Minimum version of the hub content to reference.
	//
	// Use "1.0.0" to support all versions. Changing this value to an empty string forces replacement of the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/sagemaker_hub_content_reference#min_version SagemakerHubContentReference#min_version}
	MinVersion *string `field:"optional" json:"minVersion" yaml:"minVersion"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/sagemaker_hub_content_reference#region SagemakerHubContentReference#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/sagemaker_hub_content_reference#tags SagemakerHubContentReference#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/sagemaker_hub_content_reference#timeouts SagemakerHubContentReference#timeouts}
	Timeouts *SagemakerHubContentReferenceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

