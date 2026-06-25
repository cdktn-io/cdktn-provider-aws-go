// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package shielddrtaccesslogbucketassociation

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ShieldDrtAccessLogBucketAssociationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/shield_drt_access_log_bucket_association#log_bucket ShieldDrtAccessLogBucketAssociation#log_bucket}.
	LogBucket *string `field:"required" json:"logBucket" yaml:"logBucket"`
	// Unused.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/shield_drt_access_log_bucket_association#role_arn_association_id ShieldDrtAccessLogBucketAssociation#role_arn_association_id}
	RoleArnAssociationId *string `field:"required" json:"roleArnAssociationId" yaml:"roleArnAssociationId"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/shield_drt_access_log_bucket_association#timeouts ShieldDrtAccessLogBucketAssociation#timeouts}
	Timeouts *ShieldDrtAccessLogBucketAssociationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

