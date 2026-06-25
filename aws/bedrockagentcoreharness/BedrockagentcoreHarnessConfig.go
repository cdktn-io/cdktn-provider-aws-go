// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreharness

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BedrockagentcoreHarnessConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/bedrockagentcore_harness#execution_role_arn BedrockagentcoreHarness#execution_role_arn}.
	ExecutionRoleArn *string `field:"required" json:"executionRoleArn" yaml:"executionRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/bedrockagentcore_harness#harness_name BedrockagentcoreHarness#harness_name}.
	HarnessName *string `field:"required" json:"harnessName" yaml:"harnessName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/bedrockagentcore_harness#allowed_tools BedrockagentcoreHarness#allowed_tools}.
	AllowedTools *[]*string `field:"optional" json:"allowedTools" yaml:"allowedTools"`
	// authorizer_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/bedrockagentcore_harness#authorizer_configuration BedrockagentcoreHarness#authorizer_configuration}
	AuthorizerConfiguration interface{} `field:"optional" json:"authorizerConfiguration" yaml:"authorizerConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/bedrockagentcore_harness#environment BedrockagentcoreHarness#environment}.
	Environment interface{} `field:"optional" json:"environment" yaml:"environment"`
	// environment_artifact block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/bedrockagentcore_harness#environment_artifact BedrockagentcoreHarness#environment_artifact}
	EnvironmentArtifact interface{} `field:"optional" json:"environmentArtifact" yaml:"environmentArtifact"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/bedrockagentcore_harness#environment_variables BedrockagentcoreHarness#environment_variables}.
	EnvironmentVariables *map[string]*string `field:"optional" json:"environmentVariables" yaml:"environmentVariables"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/bedrockagentcore_harness#max_iterations BedrockagentcoreHarness#max_iterations}.
	MaxIterations *float64 `field:"optional" json:"maxIterations" yaml:"maxIterations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/bedrockagentcore_harness#max_tokens BedrockagentcoreHarness#max_tokens}.
	MaxTokens *float64 `field:"optional" json:"maxTokens" yaml:"maxTokens"`
	// memory block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/bedrockagentcore_harness#memory BedrockagentcoreHarness#memory}
	Memory interface{} `field:"optional" json:"memory" yaml:"memory"`
	// model block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/bedrockagentcore_harness#model BedrockagentcoreHarness#model}
	Model interface{} `field:"optional" json:"model" yaml:"model"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/bedrockagentcore_harness#region BedrockagentcoreHarness#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// skill block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/bedrockagentcore_harness#skill BedrockagentcoreHarness#skill}
	Skill interface{} `field:"optional" json:"skill" yaml:"skill"`
	// system_prompt block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/bedrockagentcore_harness#system_prompt BedrockagentcoreHarness#system_prompt}
	SystemPrompt interface{} `field:"optional" json:"systemPrompt" yaml:"systemPrompt"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/bedrockagentcore_harness#tags BedrockagentcoreHarness#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/bedrockagentcore_harness#timeouts BedrockagentcoreHarness#timeouts}
	Timeouts *BedrockagentcoreHarnessTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/bedrockagentcore_harness#timeout_seconds BedrockagentcoreHarness#timeout_seconds}.
	TimeoutSeconds *float64 `field:"optional" json:"timeoutSeconds" yaml:"timeoutSeconds"`
	// tool block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/bedrockagentcore_harness#tool BedrockagentcoreHarness#tool}
	Tool interface{} `field:"optional" json:"tool" yaml:"tool"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/bedrockagentcore_harness#truncation BedrockagentcoreHarness#truncation}.
	Truncation interface{} `field:"optional" json:"truncation" yaml:"truncation"`
}

