// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package accountaccessentitlement


type AccountaccessEntitlementEntitlementPrincipalRole struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.65.0/docs/resources/accountaccess_entitlement#role_arn AccountaccessEntitlement#role_arn}.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// principal block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.65.0/docs/resources/accountaccess_entitlement#principal AccountaccessEntitlement#principal}
	Principal interface{} `field:"optional" json:"principal" yaml:"principal"`
}

