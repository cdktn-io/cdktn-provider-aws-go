// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wafv2webacl


type Wafv2WebAclMonetizationConfigCryptoConfigPaymentNetworkPrices struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.69.0/docs/resources/wafv2_web_acl#amount Wafv2WebAcl#amount}.
	Amount *string `field:"required" json:"amount" yaml:"amount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.69.0/docs/resources/wafv2_web_acl#currency Wafv2WebAcl#currency}.
	Currency *string `field:"required" json:"currency" yaml:"currency"`
}

