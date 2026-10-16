// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wafv2rulegroup


type Wafv2RuleGroupMonetizationConfigCryptoConfig struct {
	// payment_network block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.69.1/docs/resources/wafv2_rule_group#payment_network Wafv2RuleGroup#payment_network}
	PaymentNetwork interface{} `field:"required" json:"paymentNetwork" yaml:"paymentNetwork"`
}

