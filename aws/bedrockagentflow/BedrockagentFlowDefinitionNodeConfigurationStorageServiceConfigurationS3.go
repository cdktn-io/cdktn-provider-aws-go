// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentflow


type BedrockagentFlowDefinitionNodeConfigurationStorageServiceConfigurationS3 struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.35.0/docs/resources/bedrockagent_flow#bucket_name BedrockagentFlow#bucket_name}.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
}

