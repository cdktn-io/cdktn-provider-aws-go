// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package route53cidrlocation

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Route53CidrLocationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.46.0/docs/resources/route53_cidr_location#cidr_blocks Route53CidrLocation#cidr_blocks}.
	CidrBlocks *[]*string `field:"required" json:"cidrBlocks" yaml:"cidrBlocks"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.46.0/docs/resources/route53_cidr_location#cidr_collection_id Route53CidrLocation#cidr_collection_id}.
	CidrCollectionId *string `field:"required" json:"cidrCollectionId" yaml:"cidrCollectionId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.46.0/docs/resources/route53_cidr_location#name Route53CidrLocation#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
}

