// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreharness


type BedrockagentcoreHarnessAuthorizerConfigurationCustomJwtAuthorizerAllowedWorkloadConfiguration struct {
	// hosting_environment block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/bedrockagentcore_harness#hosting_environment BedrockagentcoreHarness#hosting_environment}
	HostingEnvironment interface{} `field:"optional" json:"hostingEnvironment" yaml:"hostingEnvironment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/bedrockagentcore_harness#workload_identities BedrockagentcoreHarness#workload_identities}.
	WorkloadIdentities *[]*string `field:"optional" json:"workloadIdentities" yaml:"workloadIdentities"`
}

