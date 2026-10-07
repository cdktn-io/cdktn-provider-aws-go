// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ephemeralawsbedrockruntimeapplyguardrail


type EphemeralAwsBedrockruntimeApplyGuardrailContentText struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.67.0/docs/ephemeral-resources/bedrockruntime_apply_guardrail#text EphemeralAwsBedrockruntimeApplyGuardrail#text}.
	Text *string `field:"required" json:"text" yaml:"text"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.67.0/docs/ephemeral-resources/bedrockruntime_apply_guardrail#qualifiers EphemeralAwsBedrockruntimeApplyGuardrail#qualifiers}.
	Qualifiers *[]*string `field:"optional" json:"qualifiers" yaml:"qualifiers"`
}

