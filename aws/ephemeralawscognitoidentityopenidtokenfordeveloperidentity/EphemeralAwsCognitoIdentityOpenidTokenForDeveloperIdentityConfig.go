// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ephemeralawscognitoidentityopenidtokenfordeveloperidentity

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type EphemeralAwsCognitoIdentityOpenidTokenForDeveloperIdentityConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/ephemeral-resources/cognito_identity_openid_token_for_developer_identity#identity_pool_id EphemeralAwsCognitoIdentityOpenidTokenForDeveloperIdentity#identity_pool_id}.
	IdentityPoolId *string `field:"required" json:"identityPoolId" yaml:"identityPoolId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/ephemeral-resources/cognito_identity_openid_token_for_developer_identity#logins EphemeralAwsCognitoIdentityOpenidTokenForDeveloperIdentity#logins}.
	Logins *map[string]*string `field:"required" json:"logins" yaml:"logins"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/ephemeral-resources/cognito_identity_openid_token_for_developer_identity#identity_id EphemeralAwsCognitoIdentityOpenidTokenForDeveloperIdentity#identity_id}.
	IdentityId *string `field:"optional" json:"identityId" yaml:"identityId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/ephemeral-resources/cognito_identity_openid_token_for_developer_identity#principal_tags EphemeralAwsCognitoIdentityOpenidTokenForDeveloperIdentity#principal_tags}.
	PrincipalTags *map[string]*string `field:"optional" json:"principalTags" yaml:"principalTags"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/ephemeral-resources/cognito_identity_openid_token_for_developer_identity#region EphemeralAwsCognitoIdentityOpenidTokenForDeveloperIdentity#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/ephemeral-resources/cognito_identity_openid_token_for_developer_identity#token_duration EphemeralAwsCognitoIdentityOpenidTokenForDeveloperIdentity#token_duration}.
	TokenDuration *float64 `field:"optional" json:"tokenDuration" yaml:"tokenDuration"`
}

