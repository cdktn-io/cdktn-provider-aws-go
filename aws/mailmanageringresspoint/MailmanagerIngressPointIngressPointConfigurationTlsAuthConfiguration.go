// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mailmanageringresspoint


type MailmanagerIngressPointIngressPointConfigurationTlsAuthConfiguration struct {
	// trust_store block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.61.0/docs/resources/mailmanager_ingress_point#trust_store MailmanagerIngressPoint#trust_store}
	TrustStore interface{} `field:"optional" json:"trustStore" yaml:"trustStore"`
}

