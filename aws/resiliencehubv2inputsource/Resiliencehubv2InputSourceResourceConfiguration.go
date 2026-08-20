// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package resiliencehubv2inputsource


type Resiliencehubv2InputSourceResourceConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.61.0/docs/resources/resiliencehubv2_input_source#cfn_stack_arn Resiliencehubv2InputSource#cfn_stack_arn}.
	CfnStackArn *string `field:"optional" json:"cfnStackArn" yaml:"cfnStackArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.61.0/docs/resources/resiliencehubv2_input_source#design_file_s3_url Resiliencehubv2InputSource#design_file_s3_url}.
	DesignFileS3Url *string `field:"optional" json:"designFileS3Url" yaml:"designFileS3Url"`
	// eks block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.61.0/docs/resources/resiliencehubv2_input_source#eks Resiliencehubv2InputSource#eks}
	Eks interface{} `field:"optional" json:"eks" yaml:"eks"`
	// resource_tag block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.61.0/docs/resources/resiliencehubv2_input_source#resource_tag Resiliencehubv2InputSource#resource_tag}
	ResourceTag interface{} `field:"optional" json:"resourceTag" yaml:"resourceTag"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.61.0/docs/resources/resiliencehubv2_input_source#tf_state_file_url Resiliencehubv2InputSource#tf_state_file_url}.
	TfStateFileUrl *string `field:"optional" json:"tfStateFileUrl" yaml:"tfStateFileUrl"`
}

