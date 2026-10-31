// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package devopsagentasset


type DevopsagentAssetContentFile struct {
	// Path of the file within the asset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.71.0/docs/resources/devopsagent_asset#path DevopsagentAsset#path}
	Path *string `field:"required" json:"path" yaml:"path"`
	// body block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.71.0/docs/resources/devopsagent_asset#body DevopsagentAsset#body}
	Body interface{} `field:"optional" json:"body" yaml:"body"`
	// JSON metadata for the file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.71.0/docs/resources/devopsagent_asset#metadata DevopsagentAsset#metadata}
	Metadata *string `field:"optional" json:"metadata" yaml:"metadata"`
}

