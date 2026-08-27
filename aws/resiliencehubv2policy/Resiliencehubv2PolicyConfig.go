// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package resiliencehubv2policy

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Resiliencehubv2PolicyConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/resiliencehubv2_policy#name Resiliencehubv2Policy#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// availability_slo block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/resiliencehubv2_policy#availability_slo Resiliencehubv2Policy#availability_slo}
	AvailabilitySlo interface{} `field:"optional" json:"availabilitySlo" yaml:"availabilitySlo"`
	// data_recovery block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/resiliencehubv2_policy#data_recovery Resiliencehubv2Policy#data_recovery}
	DataRecovery interface{} `field:"optional" json:"dataRecovery" yaml:"dataRecovery"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/resiliencehubv2_policy#description Resiliencehubv2Policy#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/resiliencehubv2_policy#kms_key_id Resiliencehubv2Policy#kms_key_id}.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// multi_az block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/resiliencehubv2_policy#multi_az Resiliencehubv2Policy#multi_az}
	MultiAz interface{} `field:"optional" json:"multiAz" yaml:"multiAz"`
	// multi_region block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/resiliencehubv2_policy#multi_region Resiliencehubv2Policy#multi_region}
	MultiRegion interface{} `field:"optional" json:"multiRegion" yaml:"multiRegion"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/resiliencehubv2_policy#region Resiliencehubv2Policy#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/resiliencehubv2_policy#tags Resiliencehubv2Policy#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

