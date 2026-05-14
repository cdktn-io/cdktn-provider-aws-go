// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appmeshvirtualnode


type AppmeshVirtualNodeSpecLoggingAccessLogFileFormat struct {
	// json block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.45.0/docs/resources/appmesh_virtual_node#json AppmeshVirtualNode#json}
	Json interface{} `field:"optional" json:"json" yaml:"json"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.45.0/docs/resources/appmesh_virtual_node#text AppmeshVirtualNode#text}.
	Text *string `field:"optional" json:"text" yaml:"text"`
}

