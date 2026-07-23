// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoregatewaytarget


type BedrockagentcoreGatewayTargetMetadataConfiguration struct {
	// A list of URL query parameters that are allowed to be propagated from incoming gateway URL to the target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.56.0/docs/resources/bedrockagentcore_gateway_target#allowed_query_parameters BedrockagentcoreGatewayTarget#allowed_query_parameters}
	AllowedQueryParameters *[]*string `field:"optional" json:"allowedQueryParameters" yaml:"allowedQueryParameters"`
	// A list of HTTP headers that are allowed to be propagated from incoming client requests to the target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.56.0/docs/resources/bedrockagentcore_gateway_target#allowed_request_headers BedrockagentcoreGatewayTarget#allowed_request_headers}
	AllowedRequestHeaders *[]*string `field:"optional" json:"allowedRequestHeaders" yaml:"allowedRequestHeaders"`
	// A list of HTTP headers that are allowed to be propagated from the target response back to the client.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.56.0/docs/resources/bedrockagentcore_gateway_target#allowed_response_headers BedrockagentcoreGatewayTarget#allowed_response_headers}
	AllowedResponseHeaders *[]*string `field:"optional" json:"allowedResponseHeaders" yaml:"allowedResponseHeaders"`
}

