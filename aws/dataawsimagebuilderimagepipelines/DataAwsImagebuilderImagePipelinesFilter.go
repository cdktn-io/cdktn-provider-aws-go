// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawsimagebuilderimagepipelines


type DataAwsImagebuilderImagePipelinesFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.35.0/docs/data-sources/imagebuilder_image_pipelines#name DataAwsImagebuilderImagePipelines#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.35.0/docs/data-sources/imagebuilder_image_pipelines#values DataAwsImagebuilderImagePipelines#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

