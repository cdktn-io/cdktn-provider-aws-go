// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package launchtemplate


type LaunchTemplateTagSpecifications struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/launch_template#resource_type LaunchTemplate#resource_type}.
	ResourceType *string `field:"optional" json:"resourceType" yaml:"resourceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/launch_template#tags LaunchTemplate#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

