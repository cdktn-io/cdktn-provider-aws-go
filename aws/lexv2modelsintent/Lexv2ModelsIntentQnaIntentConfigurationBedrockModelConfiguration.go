// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lexv2modelsintent


type Lexv2ModelsIntentQnaIntentConfigurationBedrockModelConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.55.0/docs/resources/lexv2models_intent#model_arn Lexv2ModelsIntent#model_arn}.
	ModelArn *string `field:"required" json:"modelArn" yaml:"modelArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.55.0/docs/resources/lexv2models_intent#custom_prompt Lexv2ModelsIntent#custom_prompt}.
	CustomPrompt *string `field:"optional" json:"customPrompt" yaml:"customPrompt"`
	// guardrail block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.55.0/docs/resources/lexv2models_intent#guardrail Lexv2ModelsIntent#guardrail}
	Guardrail interface{} `field:"optional" json:"guardrail" yaml:"guardrail"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.55.0/docs/resources/lexv2models_intent#trace_status Lexv2ModelsIntent#trace_status}.
	TraceStatus *string `field:"optional" json:"traceStatus" yaml:"traceStatus"`
}

