// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wafv2webaclrule


type Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUse struct {
	// allow block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.53.0/docs/resources/wafv2_web_acl_rule#allow Wafv2WebAclRuleA#allow}
	Allow interface{} `field:"optional" json:"allow" yaml:"allow"`
	// block block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.53.0/docs/resources/wafv2_web_acl_rule#block Wafv2WebAclRuleA#block}
	Block interface{} `field:"optional" json:"block" yaml:"block"`
	// captcha block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.53.0/docs/resources/wafv2_web_acl_rule#captcha Wafv2WebAclRuleA#captcha}
	Captcha interface{} `field:"optional" json:"captcha" yaml:"captcha"`
	// challenge block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.53.0/docs/resources/wafv2_web_acl_rule#challenge Wafv2WebAclRuleA#challenge}
	Challenge interface{} `field:"optional" json:"challenge" yaml:"challenge"`
	// count block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.53.0/docs/resources/wafv2_web_acl_rule#count Wafv2WebAclRuleA#count}
	Count interface{} `field:"optional" json:"count" yaml:"count"`
}

