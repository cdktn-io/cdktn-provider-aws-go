// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package codepipeline


type CodepipelineStageBeforeEntry struct {
	// condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.35.1/docs/resources/codepipeline#condition Codepipeline#condition}
	Condition *CodepipelineStageBeforeEntryCondition `field:"required" json:"condition" yaml:"condition"`
}

