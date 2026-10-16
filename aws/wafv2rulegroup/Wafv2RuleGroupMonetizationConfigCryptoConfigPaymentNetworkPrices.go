// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wafv2rulegroup


type Wafv2RuleGroupMonetizationConfigCryptoConfigPaymentNetworkPrices struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.69.1/docs/resources/wafv2_rule_group#amount Wafv2RuleGroup#amount}.
	Amount *string `field:"required" json:"amount" yaml:"amount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.69.1/docs/resources/wafv2_rule_group#currency Wafv2RuleGroup#currency}.
	Currency *string `field:"required" json:"currency" yaml:"currency"`
}

