// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreregistry


type BedrockagentcoreRegistryAuthorizerConfigurationCustomJwtAuthorizerAllowedWorkloadConfiguration struct {
	// hosting_environment block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_registry#hosting_environment BedrockagentcoreRegistry#hosting_environment}
	HostingEnvironment interface{} `field:"optional" json:"hostingEnvironment" yaml:"hostingEnvironment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_registry#workload_identities BedrockagentcoreRegistry#workload_identities}.
	WorkloadIdentities *[]*string `field:"optional" json:"workloadIdentities" yaml:"workloadIdentities"`
}

