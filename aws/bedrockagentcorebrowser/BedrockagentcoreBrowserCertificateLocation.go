// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcorebrowser


type BedrockagentcoreBrowserCertificateLocation struct {
	// secrets_manager block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.58.0/docs/resources/bedrockagentcore_browser#secrets_manager BedrockagentcoreBrowser#secrets_manager}
	SecretsManager interface{} `field:"optional" json:"secretsManager" yaml:"secretsManager"`
}

