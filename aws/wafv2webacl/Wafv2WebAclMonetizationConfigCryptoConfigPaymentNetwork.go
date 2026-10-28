// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wafv2webacl


type Wafv2WebAclMonetizationConfigCryptoConfigPaymentNetwork struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.70.0/docs/resources/wafv2_web_acl#chain Wafv2WebAcl#chain}.
	Chain *string `field:"required" json:"chain" yaml:"chain"`
	// prices block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.70.0/docs/resources/wafv2_web_acl#prices Wafv2WebAcl#prices}
	Prices *Wafv2WebAclMonetizationConfigCryptoConfigPaymentNetworkPrices `field:"required" json:"prices" yaml:"prices"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.70.0/docs/resources/wafv2_web_acl#wallet_address Wafv2WebAcl#wallet_address}.
	WalletAddress *string `field:"required" json:"walletAddress" yaml:"walletAddress"`
}

