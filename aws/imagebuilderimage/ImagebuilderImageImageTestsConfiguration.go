// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package imagebuilderimage


type ImagebuilderImageImageTestsConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.31.0/docs/resources/imagebuilder_image#image_tests_enabled ImagebuilderImage#image_tests_enabled}.
	ImageTestsEnabled interface{} `field:"optional" json:"imageTestsEnabled" yaml:"imageTestsEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.31.0/docs/resources/imagebuilder_image#timeout_minutes ImagebuilderImage#timeout_minutes}.
	TimeoutMinutes *float64 `field:"optional" json:"timeoutMinutes" yaml:"timeoutMinutes"`
}

