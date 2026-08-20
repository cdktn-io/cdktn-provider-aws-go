// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package directoryserviceradiussettings


type DirectoryServiceRadiusSettingsTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.61.0/docs/resources/directory_service_radius_settings#create DirectoryServiceRadiusSettings#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.61.0/docs/resources/directory_service_radius_settings#update DirectoryServiceRadiusSettings#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

