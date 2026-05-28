// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package networkmanagerdxgatewayattachment

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type NetworkmanagerDxGatewayAttachmentConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.47.0/docs/resources/networkmanager_dx_gateway_attachment#core_network_id NetworkmanagerDxGatewayAttachment#core_network_id}.
	CoreNetworkId *string `field:"required" json:"coreNetworkId" yaml:"coreNetworkId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.47.0/docs/resources/networkmanager_dx_gateway_attachment#direct_connect_gateway_arn NetworkmanagerDxGatewayAttachment#direct_connect_gateway_arn}.
	DirectConnectGatewayArn *string `field:"required" json:"directConnectGatewayArn" yaml:"directConnectGatewayArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.47.0/docs/resources/networkmanager_dx_gateway_attachment#edge_locations NetworkmanagerDxGatewayAttachment#edge_locations}.
	EdgeLocations *[]*string `field:"required" json:"edgeLocations" yaml:"edgeLocations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.47.0/docs/resources/networkmanager_dx_gateway_attachment#routing_policy_label NetworkmanagerDxGatewayAttachment#routing_policy_label}.
	RoutingPolicyLabel *string `field:"optional" json:"routingPolicyLabel" yaml:"routingPolicyLabel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.47.0/docs/resources/networkmanager_dx_gateway_attachment#tags NetworkmanagerDxGatewayAttachment#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.47.0/docs/resources/networkmanager_dx_gateway_attachment#timeouts NetworkmanagerDxGatewayAttachment#timeouts}
	Timeouts *NetworkmanagerDxGatewayAttachmentTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

