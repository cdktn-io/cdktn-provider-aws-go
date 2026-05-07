// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package networkmanagerattachmentroutingpolicylabel

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type NetworkmanagerAttachmentRoutingPolicyLabelConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.44.0/docs/resources/networkmanager_attachment_routing_policy_label#attachment_id NetworkmanagerAttachmentRoutingPolicyLabel#attachment_id}.
	AttachmentId *string `field:"required" json:"attachmentId" yaml:"attachmentId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.44.0/docs/resources/networkmanager_attachment_routing_policy_label#core_network_id NetworkmanagerAttachmentRoutingPolicyLabel#core_network_id}.
	CoreNetworkId *string `field:"required" json:"coreNetworkId" yaml:"coreNetworkId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.44.0/docs/resources/networkmanager_attachment_routing_policy_label#routing_policy_label NetworkmanagerAttachmentRoutingPolicyLabel#routing_policy_label}.
	RoutingPolicyLabel *string `field:"required" json:"routingPolicyLabel" yaml:"routingPolicyLabel"`
}

