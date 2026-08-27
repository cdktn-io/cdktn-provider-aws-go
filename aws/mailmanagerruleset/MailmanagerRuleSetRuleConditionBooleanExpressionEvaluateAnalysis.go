// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mailmanagerruleset


type MailmanagerRuleSetRuleConditionBooleanExpressionEvaluateAnalysis struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_rule_set#analyzer MailmanagerRuleSet#analyzer}.
	Analyzer *string `field:"required" json:"analyzer" yaml:"analyzer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_rule_set#result_field MailmanagerRuleSet#result_field}.
	ResultField *string `field:"required" json:"resultField" yaml:"resultField"`
}

