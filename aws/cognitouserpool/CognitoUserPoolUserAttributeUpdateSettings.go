// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cognitouserpool


type CognitoUserPoolUserAttributeUpdateSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/cognito_user_pool#attributes_require_verification_before_update CognitoUserPool#attributes_require_verification_before_update}.
	AttributesRequireVerificationBeforeUpdate *[]*string `field:"required" json:"attributesRequireVerificationBeforeUpdate" yaml:"attributesRequireVerificationBeforeUpdate"`
}

