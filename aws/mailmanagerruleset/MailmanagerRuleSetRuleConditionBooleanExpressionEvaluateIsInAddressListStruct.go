// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mailmanagerruleset


type MailmanagerRuleSetRuleConditionBooleanExpressionEvaluateIsInAddressListStruct struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/mailmanager_rule_set#address_lists MailmanagerRuleSet#address_lists}.
	AddressLists *[]*string `field:"required" json:"addressLists" yaml:"addressLists"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/mailmanager_rule_set#attribute MailmanagerRuleSet#attribute}.
	Attribute *string `field:"required" json:"attribute" yaml:"attribute"`
}

