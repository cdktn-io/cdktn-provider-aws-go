// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockmodelinvocationjob


type BedrockModelInvocationJobInputDataConfig struct {
	// s3_input_data_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.64.0/docs/resources/bedrock_model_invocation_job#s3_input_data_config BedrockModelInvocationJob#s3_input_data_config}
	S3InputDataConfig interface{} `field:"optional" json:"s3InputDataConfig" yaml:"s3InputDataConfig"`
}

