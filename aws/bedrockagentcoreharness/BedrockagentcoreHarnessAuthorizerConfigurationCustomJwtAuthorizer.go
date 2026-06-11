// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreharness


type BedrockagentcoreHarnessAuthorizerConfigurationCustomJwtAuthorizer struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/bedrockagentcore_harness#discovery_url BedrockagentcoreHarness#discovery_url}.
	DiscoveryUrl *string `field:"required" json:"discoveryUrl" yaml:"discoveryUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/bedrockagentcore_harness#allowed_audience BedrockagentcoreHarness#allowed_audience}.
	AllowedAudience *[]*string `field:"optional" json:"allowedAudience" yaml:"allowedAudience"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/bedrockagentcore_harness#allowed_clients BedrockagentcoreHarness#allowed_clients}.
	AllowedClients *[]*string `field:"optional" json:"allowedClients" yaml:"allowedClients"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/bedrockagentcore_harness#allowed_scopes BedrockagentcoreHarness#allowed_scopes}.
	AllowedScopes *[]*string `field:"optional" json:"allowedScopes" yaml:"allowedScopes"`
	// custom_claim block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/bedrockagentcore_harness#custom_claim BedrockagentcoreHarness#custom_claim}
	CustomClaim interface{} `field:"optional" json:"customClaim" yaml:"customClaim"`
}

