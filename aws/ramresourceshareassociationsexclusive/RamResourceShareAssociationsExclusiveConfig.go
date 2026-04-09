// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ramresourceshareassociationsexclusive

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type RamResourceShareAssociationsExclusiveConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.40.0/docs/resources/ram_resource_share_associations_exclusive#resource_share_arn RamResourceShareAssociationsExclusive#resource_share_arn}.
	ResourceShareArn *string `field:"required" json:"resourceShareArn" yaml:"resourceShareArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.40.0/docs/resources/ram_resource_share_associations_exclusive#principals RamResourceShareAssociationsExclusive#principals}.
	Principals *[]*string `field:"optional" json:"principals" yaml:"principals"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.40.0/docs/resources/ram_resource_share_associations_exclusive#region RamResourceShareAssociationsExclusive#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.40.0/docs/resources/ram_resource_share_associations_exclusive#resource_arns RamResourceShareAssociationsExclusive#resource_arns}.
	ResourceArns *[]*string `field:"optional" json:"resourceArns" yaml:"resourceArns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.40.0/docs/resources/ram_resource_share_associations_exclusive#sources RamResourceShareAssociationsExclusive#sources}.
	Sources *[]*string `field:"optional" json:"sources" yaml:"sources"`
}

