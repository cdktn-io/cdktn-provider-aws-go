// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package devopsagentasset


type DevopsagentAssetContentFileBody struct {
	// Base64-encoded binary file content.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.71.0/docs/resources/devopsagent_asset#bytes DevopsagentAsset#bytes}
	Bytes *string `field:"optional" json:"bytes" yaml:"bytes"`
	// Text file content.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.71.0/docs/resources/devopsagent_asset#text DevopsagentAsset#text}
	Text *string `field:"optional" json:"text" yaml:"text"`
}

