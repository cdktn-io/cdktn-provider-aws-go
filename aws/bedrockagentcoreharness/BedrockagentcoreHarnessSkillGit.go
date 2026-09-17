// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreharness


type BedrockagentcoreHarnessSkillGit struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.65.0/docs/resources/bedrockagentcore_harness#url BedrockagentcoreHarness#url}.
	Url *string `field:"required" json:"url" yaml:"url"`
	// auth block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.65.0/docs/resources/bedrockagentcore_harness#auth BedrockagentcoreHarness#auth}
	Auth interface{} `field:"optional" json:"auth" yaml:"auth"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.65.0/docs/resources/bedrockagentcore_harness#path BedrockagentcoreHarness#path}.
	Path *string `field:"optional" json:"path" yaml:"path"`
}

