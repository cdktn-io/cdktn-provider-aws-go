// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mailmanagertrafficpolicy


type MailmanagerTrafficPolicyPolicyStatementConditionBooleanExpressionEvaluateIsInAddressListStruct struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_traffic_policy#address_lists MailmanagerTrafficPolicy#address_lists}.
	AddressLists *[]*string `field:"required" json:"addressLists" yaml:"addressLists"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_traffic_policy#attribute MailmanagerTrafficPolicy#attribute}.
	Attribute *string `field:"required" json:"attribute" yaml:"attribute"`
}

