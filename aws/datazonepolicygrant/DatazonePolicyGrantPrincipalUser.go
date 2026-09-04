// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datazonepolicygrant


type DatazonePolicyGrantPrincipalUser struct {
	// all_users_grant_filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/datazone_policy_grant#all_users_grant_filter DatazonePolicyGrant#all_users_grant_filter}
	AllUsersGrantFilter interface{} `field:"optional" json:"allUsersGrantFilter" yaml:"allUsersGrantFilter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/datazone_policy_grant#user_identifier DatazonePolicyGrant#user_identifier}.
	UserIdentifier *string `field:"optional" json:"userIdentifier" yaml:"userIdentifier"`
}

