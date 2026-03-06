// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appmeshvirtualnode


type AppmeshVirtualNodeSpecListenerTlsValidationSubjectAlternativeNames struct {
	// match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.35.1/docs/resources/appmesh_virtual_node#match AppmeshVirtualNode#match}
	Match *AppmeshVirtualNodeSpecListenerTlsValidationSubjectAlternativeNamesMatch `field:"required" json:"match" yaml:"match"`
}

