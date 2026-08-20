// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mailmanageringresspoint


type MailmanagerIngressPointNetworkConfiguration struct {
	// private_network_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.61.0/docs/resources/mailmanager_ingress_point#private_network_configuration MailmanagerIngressPoint#private_network_configuration}
	PrivateNetworkConfiguration interface{} `field:"optional" json:"privateNetworkConfiguration" yaml:"privateNetworkConfiguration"`
	// public_network_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.61.0/docs/resources/mailmanager_ingress_point#public_network_configuration MailmanagerIngressPoint#public_network_configuration}
	PublicNetworkConfiguration interface{} `field:"optional" json:"publicNetworkConfiguration" yaml:"publicNetworkConfiguration"`
}

