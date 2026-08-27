// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mailmanageringresspoint


type MailmanagerIngressPointNetworkConfigurationPrivateNetworkConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_ingress_point#vpc_endpoint_id MailmanagerIngressPoint#vpc_endpoint_id}.
	VpcEndpointId *string `field:"required" json:"vpcEndpointId" yaml:"vpcEndpointId"`
}

