// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wafv2webaclrulegroupassociation


type Wafv2WebAclRuleGroupAssociationManagedRuleGroupManagedRuleGroupConfigsAwsManagedRulesAntiDdosRuleSet struct {
	// client_side_action_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.54.0/docs/resources/wafv2_web_acl_rule_group_association#client_side_action_config Wafv2WebAclRuleGroupAssociation#client_side_action_config}
	ClientSideActionConfig interface{} `field:"optional" json:"clientSideActionConfig" yaml:"clientSideActionConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.54.0/docs/resources/wafv2_web_acl_rule_group_association#sensitivity_to_block Wafv2WebAclRuleGroupAssociation#sensitivity_to_block}.
	SensitivityToBlock *string `field:"optional" json:"sensitivityToBlock" yaml:"sensitivityToBlock"`
}

