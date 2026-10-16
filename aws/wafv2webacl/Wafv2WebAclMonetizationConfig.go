// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wafv2webacl


type Wafv2WebAclMonetizationConfig struct {
	// crypto_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.69.1/docs/resources/wafv2_web_acl#crypto_config Wafv2WebAcl#crypto_config}
	CryptoConfig *Wafv2WebAclMonetizationConfigCryptoConfig `field:"optional" json:"cryptoConfig" yaml:"cryptoConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.69.1/docs/resources/wafv2_web_acl#currency_mode Wafv2WebAcl#currency_mode}.
	CurrencyMode *string `field:"optional" json:"currencyMode" yaml:"currencyMode"`
}

