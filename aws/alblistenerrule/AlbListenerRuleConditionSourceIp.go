// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package alblistenerrule


type AlbListenerRuleConditionSourceIp struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/alb_listener_rule#ip_address_type AlbListenerRule#ip_address_type}.
	IpAddressType *string `field:"optional" json:"ipAddressType" yaml:"ipAddressType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/alb_listener_rule#values AlbListenerRule#values}.
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

