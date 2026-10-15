// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wafv2rulegroup


type Wafv2RuleGroupMonetizationConfig struct {
	// crypto_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.69.0/docs/resources/wafv2_rule_group#crypto_config Wafv2RuleGroup#crypto_config}
	CryptoConfig *Wafv2RuleGroupMonetizationConfigCryptoConfig `field:"optional" json:"cryptoConfig" yaml:"cryptoConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.69.0/docs/resources/wafv2_rule_group#currency_mode Wafv2RuleGroup#currency_mode}.
	CurrencyMode *string `field:"optional" json:"currencyMode" yaml:"currencyMode"`
}

