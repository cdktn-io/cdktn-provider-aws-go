// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wafv2rulegroup


type Wafv2RuleGroupMonetizationConfigCryptoConfigPaymentNetwork struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.70.0/docs/resources/wafv2_rule_group#chain Wafv2RuleGroup#chain}.
	Chain *string `field:"required" json:"chain" yaml:"chain"`
	// prices block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.70.0/docs/resources/wafv2_rule_group#prices Wafv2RuleGroup#prices}
	Prices *Wafv2RuleGroupMonetizationConfigCryptoConfigPaymentNetworkPrices `field:"required" json:"prices" yaml:"prices"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.70.0/docs/resources/wafv2_rule_group#wallet_address Wafv2RuleGroup#wallet_address}.
	WalletAddress *string `field:"required" json:"walletAddress" yaml:"walletAddress"`
}

