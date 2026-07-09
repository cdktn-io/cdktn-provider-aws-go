// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcorebrowser


type BedrockagentcoreBrowserEnterprisePolicy struct {
	// location block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.54.0/docs/resources/bedrockagentcore_browser#location BedrockagentcoreBrowser#location}
	Location interface{} `field:"optional" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.54.0/docs/resources/bedrockagentcore_browser#type BedrockagentcoreBrowser#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

