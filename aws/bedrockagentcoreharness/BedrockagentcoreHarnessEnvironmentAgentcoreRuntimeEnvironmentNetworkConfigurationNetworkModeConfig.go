// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreharness


type BedrockagentcoreHarnessEnvironmentAgentcoreRuntimeEnvironmentNetworkConfigurationNetworkModeConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/bedrockagentcore_harness#require_service_s3_endpoint BedrockagentcoreHarness#require_service_s3_endpoint}.
	RequireServiceS3Endpoint interface{} `field:"optional" json:"requireServiceS3Endpoint" yaml:"requireServiceS3Endpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/bedrockagentcore_harness#security_groups BedrockagentcoreHarness#security_groups}.
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/bedrockagentcore_harness#subnets BedrockagentcoreHarness#subnets}.
	Subnets *[]*string `field:"optional" json:"subnets" yaml:"subnets"`
}

