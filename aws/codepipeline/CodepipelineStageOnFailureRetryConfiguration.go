// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package codepipeline


type CodepipelineStageOnFailureRetryConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.35.1/docs/resources/codepipeline#retry_mode Codepipeline#retry_mode}.
	RetryMode *string `field:"optional" json:"retryMode" yaml:"retryMode"`
}

