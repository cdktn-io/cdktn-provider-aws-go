// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ssoadmincustomermanagedpolicyattachmentsexclusive


type SsoadminCustomerManagedPolicyAttachmentsExclusiveCustomerManagedPolicyReference struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.54.0/docs/resources/ssoadmin_customer_managed_policy_attachments_exclusive#name SsoadminCustomerManagedPolicyAttachmentsExclusive#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.54.0/docs/resources/ssoadmin_customer_managed_policy_attachments_exclusive#path SsoadminCustomerManagedPolicyAttachmentsExclusive#path}.
	Path *string `field:"optional" json:"path" yaml:"path"`
}

