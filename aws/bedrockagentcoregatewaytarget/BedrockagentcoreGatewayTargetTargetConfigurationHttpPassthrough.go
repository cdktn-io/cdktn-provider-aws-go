// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoregatewaytarget


type BedrockagentcoreGatewayTargetTargetConfigurationHttpPassthrough struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/bedrockagentcore_gateway_target#endpoint BedrockagentcoreGatewayTarget#endpoint}.
	Endpoint *string `field:"required" json:"endpoint" yaml:"endpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/bedrockagentcore_gateway_target#protocol_type BedrockagentcoreGatewayTarget#protocol_type}.
	ProtocolType *string `field:"required" json:"protocolType" yaml:"protocolType"`
	// schema block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/bedrockagentcore_gateway_target#schema BedrockagentcoreGatewayTarget#schema}
	Schema interface{} `field:"optional" json:"schema" yaml:"schema"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/bedrockagentcore_gateway_target#static_query_parameter_conflict_resolution BedrockagentcoreGatewayTarget#static_query_parameter_conflict_resolution}.
	StaticQueryParameterConflictResolution *string `field:"optional" json:"staticQueryParameterConflictResolution" yaml:"staticQueryParameterConflictResolution"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/bedrockagentcore_gateway_target#static_query_parameters BedrockagentcoreGatewayTarget#static_query_parameters}.
	StaticQueryParameters *map[string]*string `field:"optional" json:"staticQueryParameters" yaml:"staticQueryParameters"`
	// stickiness_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/bedrockagentcore_gateway_target#stickiness_configuration BedrockagentcoreGatewayTarget#stickiness_configuration}
	StickinessConfiguration interface{} `field:"optional" json:"stickinessConfiguration" yaml:"stickinessConfiguration"`
}

