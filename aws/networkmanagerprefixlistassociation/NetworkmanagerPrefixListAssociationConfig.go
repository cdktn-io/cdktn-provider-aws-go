// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package networkmanagerprefixlistassociation

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type NetworkmanagerPrefixListAssociationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkmanager_prefix_list_association#core_network_id NetworkmanagerPrefixListAssociation#core_network_id}.
	CoreNetworkId *string `field:"required" json:"coreNetworkId" yaml:"coreNetworkId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkmanager_prefix_list_association#prefix_list_alias NetworkmanagerPrefixListAssociation#prefix_list_alias}.
	PrefixListAlias *string `field:"required" json:"prefixListAlias" yaml:"prefixListAlias"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkmanager_prefix_list_association#prefix_list_arn NetworkmanagerPrefixListAssociation#prefix_list_arn}.
	PrefixListArn *string `field:"required" json:"prefixListArn" yaml:"prefixListArn"`
}

