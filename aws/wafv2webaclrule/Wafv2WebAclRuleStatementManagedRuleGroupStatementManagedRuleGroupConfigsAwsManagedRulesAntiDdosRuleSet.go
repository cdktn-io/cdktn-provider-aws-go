// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wafv2webaclrule


type Wafv2WebAclRuleStatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAntiDdosRuleSet struct {
	// client_side_action_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/wafv2_web_acl_rule#client_side_action_config Wafv2WebAclRuleA#client_side_action_config}
	ClientSideActionConfig interface{} `field:"optional" json:"clientSideActionConfig" yaml:"clientSideActionConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/wafv2_web_acl_rule#sensitivity_to_block Wafv2WebAclRuleA#sensitivity_to_block}.
	SensitivityToBlock *string `field:"optional" json:"sensitivityToBlock" yaml:"sensitivityToBlock"`
}

