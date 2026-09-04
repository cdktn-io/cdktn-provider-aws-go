// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mailmanageringresspoint


type MailmanagerIngressPointNetworkConfigurationPublicNetworkConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/mailmanager_ingress_point#ip_type MailmanagerIngressPoint#ip_type}.
	IpType *string `field:"required" json:"ipType" yaml:"ipType"`
}

