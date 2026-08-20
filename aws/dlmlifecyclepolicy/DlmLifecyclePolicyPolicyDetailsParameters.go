// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dlmlifecyclepolicy


type DlmLifecyclePolicyPolicyDetailsParameters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.61.0/docs/resources/dlm_lifecycle_policy#exclude_boot_volume DlmLifecyclePolicy#exclude_boot_volume}.
	ExcludeBootVolume interface{} `field:"optional" json:"excludeBootVolume" yaml:"excludeBootVolume"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.61.0/docs/resources/dlm_lifecycle_policy#exclude_data_volume_tags DlmLifecyclePolicy#exclude_data_volume_tags}.
	ExcludeDataVolumeTags *map[string]*string `field:"optional" json:"excludeDataVolumeTags" yaml:"excludeDataVolumeTags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.61.0/docs/resources/dlm_lifecycle_policy#no_reboot DlmLifecyclePolicy#no_reboot}.
	NoReboot interface{} `field:"optional" json:"noReboot" yaml:"noReboot"`
}

