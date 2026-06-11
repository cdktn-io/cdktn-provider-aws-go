// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ssoadmincustomermanagedpolicyattachmentsexclusive

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SsoadminCustomerManagedPolicyAttachmentsExclusiveConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/ssoadmin_customer_managed_policy_attachments_exclusive#instance_arn SsoadminCustomerManagedPolicyAttachmentsExclusive#instance_arn}.
	InstanceArn *string `field:"required" json:"instanceArn" yaml:"instanceArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/ssoadmin_customer_managed_policy_attachments_exclusive#permission_set_arn SsoadminCustomerManagedPolicyAttachmentsExclusive#permission_set_arn}.
	PermissionSetArn *string `field:"required" json:"permissionSetArn" yaml:"permissionSetArn"`
	// customer_managed_policy_reference block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/ssoadmin_customer_managed_policy_attachments_exclusive#customer_managed_policy_reference SsoadminCustomerManagedPolicyAttachmentsExclusive#customer_managed_policy_reference}
	CustomerManagedPolicyReference interface{} `field:"optional" json:"customerManagedPolicyReference" yaml:"customerManagedPolicyReference"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/ssoadmin_customer_managed_policy_attachments_exclusive#region SsoadminCustomerManagedPolicyAttachmentsExclusive#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/ssoadmin_customer_managed_policy_attachments_exclusive#timeouts SsoadminCustomerManagedPolicyAttachmentsExclusive#timeouts}
	Timeouts *SsoadminCustomerManagedPolicyAttachmentsExclusiveTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

