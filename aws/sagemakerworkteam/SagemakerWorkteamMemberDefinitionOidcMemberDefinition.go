// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerworkteam


type SagemakerWorkteamMemberDefinitionOidcMemberDefinition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/sagemaker_workteam#groups SagemakerWorkteam#groups}.
	Groups *[]*string `field:"required" json:"groups" yaml:"groups"`
}

