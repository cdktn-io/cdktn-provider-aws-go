// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package codepipeline


type CodepipelineStageOnFailure struct {
	// condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/codepipeline#condition Codepipeline#condition}
	Condition *CodepipelineStageOnFailureCondition `field:"optional" json:"condition" yaml:"condition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/codepipeline#result Codepipeline#result}.
	Result *string `field:"optional" json:"result" yaml:"result"`
	// retry_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/codepipeline#retry_configuration Codepipeline#retry_configuration}
	RetryConfiguration *CodepipelineStageOnFailureRetryConfiguration `field:"optional" json:"retryConfiguration" yaml:"retryConfiguration"`
}

