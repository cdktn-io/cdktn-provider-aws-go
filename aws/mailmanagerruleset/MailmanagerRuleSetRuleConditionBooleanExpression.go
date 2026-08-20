// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mailmanagerruleset


type MailmanagerRuleSetRuleConditionBooleanExpression struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.61.0/docs/resources/mailmanager_rule_set#operator MailmanagerRuleSet#operator}.
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// evaluate block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.61.0/docs/resources/mailmanager_rule_set#evaluate MailmanagerRuleSet#evaluate}
	Evaluate interface{} `field:"optional" json:"evaluate" yaml:"evaluate"`
}

