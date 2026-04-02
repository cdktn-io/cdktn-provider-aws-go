// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lexv2modelsintent


type Lexv2ModelsIntentConfirmationSettingPromptSpecificationMessageGroup struct {
	// message block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.39.0/docs/resources/lexv2models_intent#message Lexv2ModelsIntent#message}
	Message interface{} `field:"optional" json:"message" yaml:"message"`
	// variation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.39.0/docs/resources/lexv2models_intent#variation Lexv2ModelsIntent#variation}
	Variation interface{} `field:"optional" json:"variation" yaml:"variation"`
}

