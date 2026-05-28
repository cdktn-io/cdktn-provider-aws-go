// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreharness


type BedrockagentcoreHarnessToolConfigAgentcoreGateway struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.47.0/docs/resources/bedrockagentcore_harness#gateway_arn BedrockagentcoreHarness#gateway_arn}.
	GatewayArn *string `field:"required" json:"gatewayArn" yaml:"gatewayArn"`
	// outbound_auth block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.47.0/docs/resources/bedrockagentcore_harness#outbound_auth BedrockagentcoreHarness#outbound_auth}
	OutboundAuth interface{} `field:"optional" json:"outboundAuth" yaml:"outboundAuth"`
}

