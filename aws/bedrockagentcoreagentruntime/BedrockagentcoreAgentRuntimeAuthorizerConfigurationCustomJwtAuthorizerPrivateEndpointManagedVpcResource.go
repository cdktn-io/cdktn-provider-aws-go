// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreagentruntime


type BedrockagentcoreAgentRuntimeAuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointManagedVpcResource struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_agent_runtime#endpoint_ip_address_type BedrockagentcoreAgentRuntime#endpoint_ip_address_type}.
	EndpointIpAddressType *string `field:"required" json:"endpointIpAddressType" yaml:"endpointIpAddressType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_agent_runtime#subnet_ids BedrockagentcoreAgentRuntime#subnet_ids}.
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_agent_runtime#vpc_identifier BedrockagentcoreAgentRuntime#vpc_identifier}.
	VpcIdentifier *string `field:"required" json:"vpcIdentifier" yaml:"vpcIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_agent_runtime#routing_domain BedrockagentcoreAgentRuntime#routing_domain}.
	RoutingDomain *string `field:"optional" json:"routingDomain" yaml:"routingDomain"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_agent_runtime#security_group_ids BedrockagentcoreAgentRuntime#security_group_ids}.
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_agent_runtime#tags BedrockagentcoreAgentRuntime#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

