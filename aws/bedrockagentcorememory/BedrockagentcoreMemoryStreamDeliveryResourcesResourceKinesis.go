// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcorememory


type BedrockagentcoreMemoryStreamDeliveryResourcesResourceKinesis struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/bedrockagentcore_memory#data_stream_arn BedrockagentcoreMemory#data_stream_arn}.
	DataStreamArn *string `field:"required" json:"dataStreamArn" yaml:"dataStreamArn"`
	// content_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/bedrockagentcore_memory#content_configuration BedrockagentcoreMemory#content_configuration}
	ContentConfiguration interface{} `field:"optional" json:"contentConfiguration" yaml:"contentConfiguration"`
}

