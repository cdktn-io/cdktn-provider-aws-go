// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package networkfirewallcontainerassociation


type NetworkfirewallContainerAssociationContainerMonitoringConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.61.0/docs/resources/networkfirewall_container_association#cluster_arn NetworkfirewallContainerAssociation#cluster_arn}.
	ClusterArn *string `field:"required" json:"clusterArn" yaml:"clusterArn"`
	// attribute_filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.61.0/docs/resources/networkfirewall_container_association#attribute_filter NetworkfirewallContainerAssociation#attribute_filter}
	AttributeFilter interface{} `field:"optional" json:"attributeFilter" yaml:"attributeFilter"`
}

