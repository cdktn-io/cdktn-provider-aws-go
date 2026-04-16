// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wafv2webaclrule


type Wafv2WebAclRuleStatementManagedRuleGroupStatementManagedRuleGroupConfigs struct {
	// aws_managed_rules_acfp_rule_set block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.41.0/docs/resources/wafv2_web_acl_rule#aws_managed_rules_acfp_rule_set Wafv2WebAclRuleA#aws_managed_rules_acfp_rule_set}
	AwsManagedRulesAcfpRuleSet interface{} `field:"optional" json:"awsManagedRulesAcfpRuleSet" yaml:"awsManagedRulesAcfpRuleSet"`
	// aws_managed_rules_anti_ddos_rule_set block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.41.0/docs/resources/wafv2_web_acl_rule#aws_managed_rules_anti_ddos_rule_set Wafv2WebAclRuleA#aws_managed_rules_anti_ddos_rule_set}
	AwsManagedRulesAntiDdosRuleSet interface{} `field:"optional" json:"awsManagedRulesAntiDdosRuleSet" yaml:"awsManagedRulesAntiDdosRuleSet"`
	// aws_managed_rules_atp_rule_set block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.41.0/docs/resources/wafv2_web_acl_rule#aws_managed_rules_atp_rule_set Wafv2WebAclRuleA#aws_managed_rules_atp_rule_set}
	AwsManagedRulesAtpRuleSet interface{} `field:"optional" json:"awsManagedRulesAtpRuleSet" yaml:"awsManagedRulesAtpRuleSet"`
	// aws_managed_rules_bot_control_rule_set block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.41.0/docs/resources/wafv2_web_acl_rule#aws_managed_rules_bot_control_rule_set Wafv2WebAclRuleA#aws_managed_rules_bot_control_rule_set}
	AwsManagedRulesBotControlRuleSet interface{} `field:"optional" json:"awsManagedRulesBotControlRuleSet" yaml:"awsManagedRulesBotControlRuleSet"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.41.0/docs/resources/wafv2_web_acl_rule#login_path Wafv2WebAclRuleA#login_path}.
	LoginPath *string `field:"optional" json:"loginPath" yaml:"loginPath"`
	// password_field block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.41.0/docs/resources/wafv2_web_acl_rule#password_field Wafv2WebAclRuleA#password_field}
	PasswordField interface{} `field:"optional" json:"passwordField" yaml:"passwordField"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.41.0/docs/resources/wafv2_web_acl_rule#payload_type Wafv2WebAclRuleA#payload_type}.
	PayloadType *string `field:"optional" json:"payloadType" yaml:"payloadType"`
	// username_field block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.41.0/docs/resources/wafv2_web_acl_rule#username_field Wafv2WebAclRuleA#username_field}
	UsernameField interface{} `field:"optional" json:"usernameField" yaml:"usernameField"`
}

