// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wafv2webaclrule


type Wafv2WebAclRuleStatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesBotControlRuleSet struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/wafv2_web_acl_rule#inspection_level Wafv2WebAclRuleA#inspection_level}.
	InspectionLevel *string `field:"required" json:"inspectionLevel" yaml:"inspectionLevel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/wafv2_web_acl_rule#enable_machine_learning Wafv2WebAclRuleA#enable_machine_learning}.
	EnableMachineLearning interface{} `field:"optional" json:"enableMachineLearning" yaml:"enableMachineLearning"`
}

