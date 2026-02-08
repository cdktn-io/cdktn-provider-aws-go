// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package evidentlylaunch


type EvidentlyLaunchScheduledSplitsConfig struct {
	// steps block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.31.0/docs/resources/evidently_launch#steps EvidentlyLaunch#steps}
	Steps interface{} `field:"required" json:"steps" yaml:"steps"`
}

