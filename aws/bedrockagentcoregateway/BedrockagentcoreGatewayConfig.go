// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoregateway

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BedrockagentcoreGatewayConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/bedrockagentcore_gateway#authorizer_type BedrockagentcoreGateway#authorizer_type}.
	AuthorizerType *string `field:"required" json:"authorizerType" yaml:"authorizerType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/bedrockagentcore_gateway#name BedrockagentcoreGateway#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/bedrockagentcore_gateway#role_arn BedrockagentcoreGateway#role_arn}.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// authorizer_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/bedrockagentcore_gateway#authorizer_configuration BedrockagentcoreGateway#authorizer_configuration}
	AuthorizerConfiguration interface{} `field:"optional" json:"authorizerConfiguration" yaml:"authorizerConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/bedrockagentcore_gateway#description BedrockagentcoreGateway#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/bedrockagentcore_gateway#exception_level BedrockagentcoreGateway#exception_level}.
	ExceptionLevel *string `field:"optional" json:"exceptionLevel" yaml:"exceptionLevel"`
	// interceptor_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/bedrockagentcore_gateway#interceptor_configuration BedrockagentcoreGateway#interceptor_configuration}
	InterceptorConfiguration interface{} `field:"optional" json:"interceptorConfiguration" yaml:"interceptorConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/bedrockagentcore_gateway#kms_key_arn BedrockagentcoreGateway#kms_key_arn}.
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// policy_engine_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/bedrockagentcore_gateway#policy_engine_configuration BedrockagentcoreGateway#policy_engine_configuration}
	PolicyEngineConfiguration interface{} `field:"optional" json:"policyEngineConfiguration" yaml:"policyEngineConfiguration"`
	// protocol_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/bedrockagentcore_gateway#protocol_configuration BedrockagentcoreGateway#protocol_configuration}
	ProtocolConfiguration interface{} `field:"optional" json:"protocolConfiguration" yaml:"protocolConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/bedrockagentcore_gateway#protocol_type BedrockagentcoreGateway#protocol_type}.
	ProtocolType *string `field:"optional" json:"protocolType" yaml:"protocolType"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/bedrockagentcore_gateway#region BedrockagentcoreGateway#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/bedrockagentcore_gateway#tags BedrockagentcoreGateway#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/bedrockagentcore_gateway#timeouts BedrockagentcoreGateway#timeouts}
	Timeouts *BedrockagentcoreGatewayTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

