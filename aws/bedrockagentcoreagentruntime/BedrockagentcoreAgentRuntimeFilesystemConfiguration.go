// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreagentruntime


type BedrockagentcoreAgentRuntimeFilesystemConfiguration struct {
	// efs_access_point block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/bedrockagentcore_agent_runtime#efs_access_point BedrockagentcoreAgentRuntime#efs_access_point}
	EfsAccessPoint interface{} `field:"optional" json:"efsAccessPoint" yaml:"efsAccessPoint"`
	// s3_files_access_point block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/bedrockagentcore_agent_runtime#s3_files_access_point BedrockagentcoreAgentRuntime#s3_files_access_point}
	S3FilesAccessPoint interface{} `field:"optional" json:"s3FilesAccessPoint" yaml:"s3FilesAccessPoint"`
	// session_storage block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/bedrockagentcore_agent_runtime#session_storage BedrockagentcoreAgentRuntime#session_storage}
	SessionStorage interface{} `field:"optional" json:"sessionStorage" yaml:"sessionStorage"`
}

