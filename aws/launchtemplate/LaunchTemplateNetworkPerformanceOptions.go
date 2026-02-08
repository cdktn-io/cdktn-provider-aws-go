// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package launchtemplate


type LaunchTemplateNetworkPerformanceOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.31.0/docs/resources/launch_template#bandwidth_weighting LaunchTemplate#bandwidth_weighting}.
	BandwidthWeighting *string `field:"optional" json:"bandwidthWeighting" yaml:"bandwidthWeighting"`
}

