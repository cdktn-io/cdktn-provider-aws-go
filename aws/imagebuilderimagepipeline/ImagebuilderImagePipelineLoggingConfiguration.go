// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package imagebuilderimagepipeline


type ImagebuilderImagePipelineLoggingConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/imagebuilder_image_pipeline#image_log_group_name ImagebuilderImagePipeline#image_log_group_name}.
	ImageLogGroupName *string `field:"optional" json:"imageLogGroupName" yaml:"imageLogGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/imagebuilder_image_pipeline#pipeline_log_group_name ImagebuilderImagePipeline#pipeline_log_group_name}.
	PipelineLogGroupName *string `field:"optional" json:"pipelineLogGroupName" yaml:"pipelineLogGroupName"`
}

