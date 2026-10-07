// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ephemeralawsbedrockruntimeapplyguardrail

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type EphemeralAwsBedrockruntimeApplyGuardrailConfig struct {
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformEphemeralResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.67.0/docs/ephemeral-resources/bedrockruntime_apply_guardrail#guardrail_identifier EphemeralAwsBedrockruntimeApplyGuardrail#guardrail_identifier}.
	GuardrailIdentifier *string `field:"required" json:"guardrailIdentifier" yaml:"guardrailIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.67.0/docs/ephemeral-resources/bedrockruntime_apply_guardrail#guardrail_version EphemeralAwsBedrockruntimeApplyGuardrail#guardrail_version}.
	GuardrailVersion *string `field:"required" json:"guardrailVersion" yaml:"guardrailVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.67.0/docs/ephemeral-resources/bedrockruntime_apply_guardrail#source EphemeralAwsBedrockruntimeApplyGuardrail#source}.
	Source *string `field:"required" json:"source" yaml:"source"`
	// content block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.67.0/docs/ephemeral-resources/bedrockruntime_apply_guardrail#content EphemeralAwsBedrockruntimeApplyGuardrail#content}
	Content interface{} `field:"optional" json:"content" yaml:"content"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.67.0/docs/ephemeral-resources/bedrockruntime_apply_guardrail#region EphemeralAwsBedrockruntimeApplyGuardrail#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
}

