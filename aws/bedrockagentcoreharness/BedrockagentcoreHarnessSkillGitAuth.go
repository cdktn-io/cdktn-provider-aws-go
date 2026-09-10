// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreharness


type BedrockagentcoreHarnessSkillGitAuth struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.64.0/docs/resources/bedrockagentcore_harness#credential_arn BedrockagentcoreHarness#credential_arn}.
	CredentialArn *string `field:"required" json:"credentialArn" yaml:"credentialArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.64.0/docs/resources/bedrockagentcore_harness#username BedrockagentcoreHarness#username}.
	Username *string `field:"optional" json:"username" yaml:"username"`
}

