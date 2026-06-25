// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoregateway


type BedrockagentcoreGatewayAuthorizerConfigurationCustomJwtAuthorizerCustomClaim struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/bedrockagentcore_gateway#inbound_token_claim_name BedrockagentcoreGateway#inbound_token_claim_name}.
	InboundTokenClaimName *string `field:"required" json:"inboundTokenClaimName" yaml:"inboundTokenClaimName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/bedrockagentcore_gateway#inbound_token_claim_value_type BedrockagentcoreGateway#inbound_token_claim_value_type}.
	InboundTokenClaimValueType *string `field:"required" json:"inboundTokenClaimValueType" yaml:"inboundTokenClaimValueType"`
	// authorizing_claim_match_value block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/bedrockagentcore_gateway#authorizing_claim_match_value BedrockagentcoreGateway#authorizing_claim_match_value}
	AuthorizingClaimMatchValue interface{} `field:"optional" json:"authorizingClaimMatchValue" yaml:"authorizingClaimMatchValue"`
}

