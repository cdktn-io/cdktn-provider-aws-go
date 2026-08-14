// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package medialivechannel


type MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsFilterSettingsTemporalFilterSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/medialive_channel#post_filter_sharpening MedialiveChannel#post_filter_sharpening}.
	PostFilterSharpening *string `field:"optional" json:"postFilterSharpening" yaml:"postFilterSharpening"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/medialive_channel#strength MedialiveChannel#strength}.
	Strength *string `field:"optional" json:"strength" yaml:"strength"`
}

