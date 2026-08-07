// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package opensearchapplication


type OpensearchApplicationIamIdentityCenterOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.58.0/docs/resources/opensearch_application#enabled OpensearchApplication#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.58.0/docs/resources/opensearch_application#iam_identity_center_instance_arn OpensearchApplication#iam_identity_center_instance_arn}.
	IamIdentityCenterInstanceArn *string `field:"optional" json:"iamIdentityCenterInstanceArn" yaml:"iamIdentityCenterInstanceArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.58.0/docs/resources/opensearch_application#iam_role_for_identity_center_application_arn OpensearchApplication#iam_role_for_identity_center_application_arn}.
	IamRoleForIdentityCenterApplicationArn *string `field:"optional" json:"iamRoleForIdentityCenterApplicationArn" yaml:"iamRoleForIdentityCenterApplicationArn"`
}

