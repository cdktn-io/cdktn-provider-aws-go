// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wafv2webaclrulegroupassociation


type Wafv2WebAclRuleGroupAssociationManagedRuleGroupManagedRuleGroupConfigs struct {
	// aws_managed_rules_acfp_rule_set block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.40.0/docs/resources/wafv2_web_acl_rule_group_association#aws_managed_rules_acfp_rule_set Wafv2WebAclRuleGroupAssociation#aws_managed_rules_acfp_rule_set}
	AwsManagedRulesAcfpRuleSet interface{} `field:"optional" json:"awsManagedRulesAcfpRuleSet" yaml:"awsManagedRulesAcfpRuleSet"`
	// aws_managed_rules_anti_ddos_rule_set block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.40.0/docs/resources/wafv2_web_acl_rule_group_association#aws_managed_rules_anti_ddos_rule_set Wafv2WebAclRuleGroupAssociation#aws_managed_rules_anti_ddos_rule_set}
	AwsManagedRulesAntiDdosRuleSet interface{} `field:"optional" json:"awsManagedRulesAntiDdosRuleSet" yaml:"awsManagedRulesAntiDdosRuleSet"`
	// aws_managed_rules_atp_rule_set block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.40.0/docs/resources/wafv2_web_acl_rule_group_association#aws_managed_rules_atp_rule_set Wafv2WebAclRuleGroupAssociation#aws_managed_rules_atp_rule_set}
	AwsManagedRulesAtpRuleSet interface{} `field:"optional" json:"awsManagedRulesAtpRuleSet" yaml:"awsManagedRulesAtpRuleSet"`
	// aws_managed_rules_bot_control_rule_set block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.40.0/docs/resources/wafv2_web_acl_rule_group_association#aws_managed_rules_bot_control_rule_set Wafv2WebAclRuleGroupAssociation#aws_managed_rules_bot_control_rule_set}
	AwsManagedRulesBotControlRuleSet interface{} `field:"optional" json:"awsManagedRulesBotControlRuleSet" yaml:"awsManagedRulesBotControlRuleSet"`
}

