// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package imagebuilderimage


type ImagebuilderImageLoggingConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/imagebuilder_image#log_group_name ImagebuilderImage#log_group_name}.
	LogGroupName *string `field:"required" json:"logGroupName" yaml:"logGroupName"`
}

