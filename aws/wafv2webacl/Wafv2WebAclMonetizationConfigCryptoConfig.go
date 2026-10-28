// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wafv2webacl


type Wafv2WebAclMonetizationConfigCryptoConfig struct {
	// payment_network block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.70.0/docs/resources/wafv2_web_acl#payment_network Wafv2WebAcl#payment_network}
	PaymentNetwork interface{} `field:"required" json:"paymentNetwork" yaml:"paymentNetwork"`
}

