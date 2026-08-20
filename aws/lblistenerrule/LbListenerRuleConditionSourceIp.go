// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lblistenerrule


type LbListenerRuleConditionSourceIp struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.61.0/docs/resources/lb_listener_rule#ip_address_type LbListenerRule#ip_address_type}.
	IpAddressType *string `field:"optional" json:"ipAddressType" yaml:"ipAddressType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.61.0/docs/resources/lb_listener_rule#values LbListenerRule#values}.
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

