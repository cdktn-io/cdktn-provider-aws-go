// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package networkmanagercustomergatewayassociation


type NetworkmanagerCustomerGatewayAssociationTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/networkmanager_customer_gateway_association#create NetworkmanagerCustomerGatewayAssociation#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/networkmanager_customer_gateway_association#delete NetworkmanagerCustomerGatewayAssociation#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

