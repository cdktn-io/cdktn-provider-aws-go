// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreharness


type BedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuth struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/bedrockagentcore_harness#aws_iam BedrockagentcoreHarness#aws_iam}.
	AwsIam interface{} `field:"optional" json:"awsIam" yaml:"awsIam"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/bedrockagentcore_harness#none BedrockagentcoreHarness#none}.
	None interface{} `field:"optional" json:"none" yaml:"none"`
	// oauth block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/bedrockagentcore_harness#oauth BedrockagentcoreHarness#oauth}
	Oauth interface{} `field:"optional" json:"oauth" yaml:"oauth"`
}

