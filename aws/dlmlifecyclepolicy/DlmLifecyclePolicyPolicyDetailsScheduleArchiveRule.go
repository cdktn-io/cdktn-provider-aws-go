// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dlmlifecyclepolicy


type DlmLifecyclePolicyPolicyDetailsScheduleArchiveRule struct {
	// archive_retain_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.53.0/docs/resources/dlm_lifecycle_policy#archive_retain_rule DlmLifecyclePolicy#archive_retain_rule}
	ArchiveRetainRule *DlmLifecyclePolicyPolicyDetailsScheduleArchiveRuleArchiveRetainRule `field:"required" json:"archiveRetainRule" yaml:"archiveRetainRule"`
}

