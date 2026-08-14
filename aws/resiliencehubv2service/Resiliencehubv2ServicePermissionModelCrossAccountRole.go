// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package resiliencehubv2service


type Resiliencehubv2ServicePermissionModelCrossAccountRole struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/resiliencehubv2_service#cross_account_role_arn Resiliencehubv2Service#cross_account_role_arn}.
	CrossAccountRoleArn *string `field:"required" json:"crossAccountRoleArn" yaml:"crossAccountRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/resiliencehubv2_service#external_id Resiliencehubv2Service#external_id}.
	ExternalId *string `field:"optional" json:"externalId" yaml:"externalId"`
}

