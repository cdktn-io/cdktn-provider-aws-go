// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mailmanagertrafficpolicy


type MailmanagerTrafficPolicyPolicyStatementConditionBooleanExpressionEvaluate struct {
	// analysis block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/mailmanager_traffic_policy#analysis MailmanagerTrafficPolicy#analysis}
	Analysis interface{} `field:"optional" json:"analysis" yaml:"analysis"`
	// is_in_address_list block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/mailmanager_traffic_policy#is_in_address_list MailmanagerTrafficPolicy#is_in_address_list}
	IsInAddressList interface{} `field:"optional" json:"isInAddressList" yaml:"isInAddressList"`
}

