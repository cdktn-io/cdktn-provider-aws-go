// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreharness


type BedrockagentcoreHarnessEnvironmentArtifact struct {
	// container_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.53.0/docs/resources/bedrockagentcore_harness#container_configuration BedrockagentcoreHarness#container_configuration}
	ContainerConfiguration interface{} `field:"optional" json:"containerConfiguration" yaml:"containerConfiguration"`
}

