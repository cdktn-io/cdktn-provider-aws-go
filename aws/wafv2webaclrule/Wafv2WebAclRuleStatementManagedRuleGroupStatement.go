// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wafv2webaclrule


type Wafv2WebAclRuleStatementManagedRuleGroupStatement struct {
	// Name of the managed rule group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/wafv2_web_acl_rule#name Wafv2WebAclRuleA#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Name of the managed rule group vendor.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/wafv2_web_acl_rule#vendor_name Wafv2WebAclRuleA#vendor_name}
	VendorName *string `field:"required" json:"vendorName" yaml:"vendorName"`
	// managed_rule_group_configs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/wafv2_web_acl_rule#managed_rule_group_configs Wafv2WebAclRuleA#managed_rule_group_configs}
	ManagedRuleGroupConfigs interface{} `field:"optional" json:"managedRuleGroupConfigs" yaml:"managedRuleGroupConfigs"`
	// rule_action_override block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/wafv2_web_acl_rule#rule_action_override Wafv2WebAclRuleA#rule_action_override}
	RuleActionOverride interface{} `field:"optional" json:"ruleActionOverride" yaml:"ruleActionOverride"`
	// scope_down_statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/wafv2_web_acl_rule#scope_down_statement Wafv2WebAclRuleA#scope_down_statement}
	ScopeDownStatement interface{} `field:"optional" json:"scopeDownStatement" yaml:"scopeDownStatement"`
	// Version of the managed rule group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/wafv2_web_acl_rule#version Wafv2WebAclRuleA#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}

