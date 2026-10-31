// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package devopsagentasset


type DevopsagentAssetContentSourceUrl struct {
	// Source URL from which to import asset content.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.71.0/docs/resources/devopsagent_asset#url DevopsagentAsset#url}
	Url *string `field:"required" json:"url" yaml:"url"`
}

