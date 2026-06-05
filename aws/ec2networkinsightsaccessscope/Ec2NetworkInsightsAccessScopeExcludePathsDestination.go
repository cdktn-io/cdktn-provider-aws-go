// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2networkinsightsaccessscope


type Ec2NetworkInsightsAccessScopeExcludePathsDestination struct {
	// packet_header_statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.49.0/docs/resources/ec2_network_insights_access_scope#packet_header_statement Ec2NetworkInsightsAccessScope#packet_header_statement}
	PacketHeaderStatement interface{} `field:"optional" json:"packetHeaderStatement" yaml:"packetHeaderStatement"`
	// resource_statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.49.0/docs/resources/ec2_network_insights_access_scope#resource_statement Ec2NetworkInsightsAccessScope#resource_statement}
	ResourceStatement interface{} `field:"optional" json:"resourceStatement" yaml:"resourceStatement"`
}

