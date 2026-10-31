// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package devopsagentasset


type DevopsagentAssetContentZip struct {
	// Base64-encoded ZIP archive. Use filebase64() to load a local ZIP file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.71.0/docs/resources/devopsagent_asset#zip_file DevopsagentAsset#zip_file}
	ZipFile *string `field:"required" json:"zipFile" yaml:"zipFile"`
}

