// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appmeshvirtualnode


type AppmeshVirtualNodeSpecListenerTls struct {
	// certificate block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.53.0/docs/resources/appmesh_virtual_node#certificate AppmeshVirtualNode#certificate}
	Certificate *AppmeshVirtualNodeSpecListenerTlsCertificate `field:"required" json:"certificate" yaml:"certificate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.53.0/docs/resources/appmesh_virtual_node#mode AppmeshVirtualNode#mode}.
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// validation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.53.0/docs/resources/appmesh_virtual_node#validation AppmeshVirtualNode#validation}
	Validation *AppmeshVirtualNodeSpecListenerTlsValidation `field:"optional" json:"validation" yaml:"validation"`
}

