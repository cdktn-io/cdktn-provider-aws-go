// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package resiliencehubv2service


type Resiliencehubv2ServicePermissionModel struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.61.0/docs/resources/resiliencehubv2_service#invoker_role_name Resiliencehubv2Service#invoker_role_name}.
	InvokerRoleName *string `field:"required" json:"invokerRoleName" yaml:"invokerRoleName"`
	// cross_account_role block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.61.0/docs/resources/resiliencehubv2_service#cross_account_role Resiliencehubv2Service#cross_account_role}
	CrossAccountRole interface{} `field:"optional" json:"crossAccountRole" yaml:"crossAccountRole"`
}

