// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sesconfigurationset


type SesConfigurationSetDeliveryOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.46.0/docs/resources/ses_configuration_set#tls_policy SesConfigurationSet#tls_policy}.
	TlsPolicy *string `field:"optional" json:"tlsPolicy" yaml:"tlsPolicy"`
}

