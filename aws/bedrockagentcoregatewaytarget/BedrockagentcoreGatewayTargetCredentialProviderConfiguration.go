// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoregatewaytarget


type BedrockagentcoreGatewayTargetCredentialProviderConfiguration struct {
	// api_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.56.0/docs/resources/bedrockagentcore_gateway_target#api_key BedrockagentcoreGatewayTarget#api_key}
	ApiKey interface{} `field:"optional" json:"apiKey" yaml:"apiKey"`
	// caller_iam_credentials block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.56.0/docs/resources/bedrockagentcore_gateway_target#caller_iam_credentials BedrockagentcoreGatewayTarget#caller_iam_credentials}
	CallerIamCredentials interface{} `field:"optional" json:"callerIamCredentials" yaml:"callerIamCredentials"`
	// gateway_iam_role block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.56.0/docs/resources/bedrockagentcore_gateway_target#gateway_iam_role BedrockagentcoreGatewayTarget#gateway_iam_role}
	GatewayIamRole interface{} `field:"optional" json:"gatewayIamRole" yaml:"gatewayIamRole"`
	// jwt_passthrough block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.56.0/docs/resources/bedrockagentcore_gateway_target#jwt_passthrough BedrockagentcoreGatewayTarget#jwt_passthrough}
	JwtPassthrough interface{} `field:"optional" json:"jwtPassthrough" yaml:"jwtPassthrough"`
	// oauth block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.56.0/docs/resources/bedrockagentcore_gateway_target#oauth BedrockagentcoreGatewayTarget#oauth}
	Oauth interface{} `field:"optional" json:"oauth" yaml:"oauth"`
}

