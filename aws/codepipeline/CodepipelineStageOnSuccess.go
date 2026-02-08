// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package codepipeline


type CodepipelineStageOnSuccess struct {
	// condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.31.0/docs/resources/codepipeline#condition Codepipeline#condition}
	Condition *CodepipelineStageOnSuccessCondition `field:"required" json:"condition" yaml:"condition"`
}

