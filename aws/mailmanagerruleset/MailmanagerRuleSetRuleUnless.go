// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mailmanagerruleset


type MailmanagerRuleSetRuleUnless struct {
	// boolean_expression block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_rule_set#boolean_expression MailmanagerRuleSet#boolean_expression}
	BooleanExpression interface{} `field:"optional" json:"booleanExpression" yaml:"booleanExpression"`
	// dmarc_expression block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_rule_set#dmarc_expression MailmanagerRuleSet#dmarc_expression}
	DmarcExpression interface{} `field:"optional" json:"dmarcExpression" yaml:"dmarcExpression"`
	// ip_expression block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_rule_set#ip_expression MailmanagerRuleSet#ip_expression}
	IpExpression interface{} `field:"optional" json:"ipExpression" yaml:"ipExpression"`
	// number_expression block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_rule_set#number_expression MailmanagerRuleSet#number_expression}
	NumberExpression interface{} `field:"optional" json:"numberExpression" yaml:"numberExpression"`
	// string_expression block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_rule_set#string_expression MailmanagerRuleSet#string_expression}
	StringExpression interface{} `field:"optional" json:"stringExpression" yaml:"stringExpression"`
	// verdict_expression block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_rule_set#verdict_expression MailmanagerRuleSet#verdict_expression}
	VerdictExpression interface{} `field:"optional" json:"verdictExpression" yaml:"verdictExpression"`
}

