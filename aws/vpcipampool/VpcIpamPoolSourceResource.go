// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vpcipampool


type VpcIpamPoolSourceResource struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.51.0/docs/resources/vpc_ipam_pool#resource_id VpcIpamPool#resource_id}.
	ResourceId *string `field:"required" json:"resourceId" yaml:"resourceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.51.0/docs/resources/vpc_ipam_pool#resource_owner VpcIpamPool#resource_owner}.
	ResourceOwner *string `field:"required" json:"resourceOwner" yaml:"resourceOwner"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.51.0/docs/resources/vpc_ipam_pool#resource_region VpcIpamPool#resource_region}.
	ResourceRegion *string `field:"required" json:"resourceRegion" yaml:"resourceRegion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.51.0/docs/resources/vpc_ipam_pool#resource_type VpcIpamPool#resource_type}.
	ResourceType *string `field:"required" json:"resourceType" yaml:"resourceType"`
}

