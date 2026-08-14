// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcorecodeinterpreter


type BedrockagentcoreCodeInterpreterNetworkConfigurationVpcConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/bedrockagentcore_code_interpreter#security_groups BedrockagentcoreCodeInterpreter#security_groups}.
	SecurityGroups *[]*string `field:"required" json:"securityGroups" yaml:"securityGroups"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/bedrockagentcore_code_interpreter#subnets BedrockagentcoreCodeInterpreter#subnets}.
	Subnets *[]*string `field:"required" json:"subnets" yaml:"subnets"`
}

