// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreharness


type BedrockagentcoreHarnessEnvironmentAgentcoreRuntimeEnvironmentFilesystemConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.47.0/docs/resources/bedrockagentcore_harness#efs_access_point BedrockagentcoreHarness#efs_access_point}.
	EfsAccessPoint interface{} `field:"optional" json:"efsAccessPoint" yaml:"efsAccessPoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.47.0/docs/resources/bedrockagentcore_harness#s3_files_access_point BedrockagentcoreHarness#s3_files_access_point}.
	S3FilesAccessPoint interface{} `field:"optional" json:"s3FilesAccessPoint" yaml:"s3FilesAccessPoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.47.0/docs/resources/bedrockagentcore_harness#session_storage BedrockagentcoreHarness#session_storage}.
	SessionStorage interface{} `field:"optional" json:"sessionStorage" yaml:"sessionStorage"`
}

