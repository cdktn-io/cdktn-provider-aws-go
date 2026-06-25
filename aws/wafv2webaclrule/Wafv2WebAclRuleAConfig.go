// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wafv2webaclrule

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Wafv2WebAclRuleAConfig struct {
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
	// Rule name, unique within the Web ACL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/wafv2_web_acl_rule#name Wafv2WebAclRuleA#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Rule priority. Rules with lower priority are evaluated first.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/wafv2_web_acl_rule#priority Wafv2WebAclRuleA#priority}
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// ARN of the Web ACL to add the rule to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/wafv2_web_acl_rule#web_acl_arn Wafv2WebAclRuleA#web_acl_arn}
	WebAclArn *string `field:"required" json:"webAclArn" yaml:"webAclArn"`
	// action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/wafv2_web_acl_rule#action Wafv2WebAclRuleA#action}
	Action interface{} `field:"optional" json:"action" yaml:"action"`
	// captcha_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/wafv2_web_acl_rule#captcha_config Wafv2WebAclRuleA#captcha_config}
	CaptchaConfig interface{} `field:"optional" json:"captchaConfig" yaml:"captchaConfig"`
	// challenge_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/wafv2_web_acl_rule#challenge_config Wafv2WebAclRuleA#challenge_config}
	ChallengeConfig interface{} `field:"optional" json:"challengeConfig" yaml:"challengeConfig"`
	// override_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/wafv2_web_acl_rule#override_action Wafv2WebAclRuleA#override_action}
	OverrideAction interface{} `field:"optional" json:"overrideAction" yaml:"overrideAction"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/wafv2_web_acl_rule#region Wafv2WebAclRuleA#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// rule_label block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/wafv2_web_acl_rule#rule_label Wafv2WebAclRuleA#rule_label}
	RuleLabel interface{} `field:"optional" json:"ruleLabel" yaml:"ruleLabel"`
	// statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/wafv2_web_acl_rule#statement Wafv2WebAclRuleA#statement}
	Statement interface{} `field:"optional" json:"statement" yaml:"statement"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/wafv2_web_acl_rule#timeouts Wafv2WebAclRuleA#timeouts}
	Timeouts *Wafv2WebAclRuleTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// visibility_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/wafv2_web_acl_rule#visibility_config Wafv2WebAclRuleA#visibility_config}
	VisibilityConfig interface{} `field:"optional" json:"visibilityConfig" yaml:"visibilityConfig"`
}

