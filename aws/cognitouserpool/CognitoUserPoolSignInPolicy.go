// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cognitouserpool


type CognitoUserPoolSignInPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.35.1/docs/resources/cognito_user_pool#allowed_first_auth_factors CognitoUserPool#allowed_first_auth_factors}.
	AllowedFirstAuthFactors *[]*string `field:"optional" json:"allowedFirstAuthFactors" yaml:"allowedFirstAuthFactors"`
}

