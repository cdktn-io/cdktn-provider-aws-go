// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoregatewaytarget


type BedrockagentcoreGatewayTargetCredentialProviderConfigurationGatewayIamRole struct {
	// AWS Region used for SigV4 signing of upstream requests.
	//
	// Defaults to the gateway's Region when omitted. Only meaningful when `service` is set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/bedrockagentcore_gateway_target#region BedrockagentcoreGatewayTarget#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// The target AWS service name used for SigV4 signing of upstream requests.
	//
	// Required when calling SigV4-protected endpoints such as another Bedrock AgentCore Runtime (use `bedrock-agentcore`). Omit for non-SigV4 IAM-role-based authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/bedrockagentcore_gateway_target#service BedrockagentcoreGatewayTarget#service}
	Service *string `field:"optional" json:"service" yaml:"service"`
}

