// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreharness


type BedrockagentcoreHarnessSkill struct {
	// aws_skills block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.65.0/docs/resources/bedrockagentcore_harness#aws_skills BedrockagentcoreHarness#aws_skills}
	AwsSkills interface{} `field:"optional" json:"awsSkills" yaml:"awsSkills"`
	// git block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.65.0/docs/resources/bedrockagentcore_harness#git BedrockagentcoreHarness#git}
	Git interface{} `field:"optional" json:"git" yaml:"git"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.65.0/docs/resources/bedrockagentcore_harness#path BedrockagentcoreHarness#path}.
	Path *string `field:"optional" json:"path" yaml:"path"`
	// s3 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.65.0/docs/resources/bedrockagentcore_harness#s3 BedrockagentcoreHarness#s3}
	S3 interface{} `field:"optional" json:"s3" yaml:"s3"`
}

