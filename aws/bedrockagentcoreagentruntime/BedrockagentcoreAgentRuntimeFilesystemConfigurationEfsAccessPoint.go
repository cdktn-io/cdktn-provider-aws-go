// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreagentruntime


type BedrockagentcoreAgentRuntimeFilesystemConfigurationEfsAccessPoint struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.53.0/docs/resources/bedrockagentcore_agent_runtime#access_point_arn BedrockagentcoreAgentRuntime#access_point_arn}.
	AccessPointArn *string `field:"required" json:"accessPointArn" yaml:"accessPointArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.53.0/docs/resources/bedrockagentcore_agent_runtime#mount_path BedrockagentcoreAgentRuntime#mount_path}.
	MountPath *string `field:"required" json:"mountPath" yaml:"mountPath"`
}

