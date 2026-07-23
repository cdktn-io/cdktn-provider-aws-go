// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreharness


type BedrockagentcoreHarnessAuthorizerConfigurationCustomJwtAuthorizerPrivateEndpoint struct {
	// managed_vpc_resource block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.56.0/docs/resources/bedrockagentcore_harness#managed_vpc_resource BedrockagentcoreHarness#managed_vpc_resource}
	ManagedVpcResource interface{} `field:"optional" json:"managedVpcResource" yaml:"managedVpcResource"`
	// self_managed_lattice_resource block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.56.0/docs/resources/bedrockagentcore_harness#self_managed_lattice_resource BedrockagentcoreHarness#self_managed_lattice_resource}
	SelfManagedLatticeResource interface{} `field:"optional" json:"selfManagedLatticeResource" yaml:"selfManagedLatticeResource"`
}

