// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockmodelinvocationjob


type BedrockModelInvocationJobOutputDataConfig struct {
	// s3_output_data_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.65.0/docs/resources/bedrock_model_invocation_job#s3_output_data_config BedrockModelInvocationJob#s3_output_data_config}
	S3OutputDataConfig interface{} `field:"optional" json:"s3OutputDataConfig" yaml:"s3OutputDataConfig"`
}

