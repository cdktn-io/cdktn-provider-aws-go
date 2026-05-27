// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appmeshvirtualservice


type AppmeshVirtualServiceSpec struct {
	// provider block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.46.0/docs/resources/appmesh_virtual_service#provider AppmeshVirtualService#provider}
	Provider *AppmeshVirtualServiceSpecProvider `field:"optional" json:"provider" yaml:"provider"`
}

