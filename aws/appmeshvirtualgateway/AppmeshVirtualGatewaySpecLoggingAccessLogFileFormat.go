// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appmeshvirtualgateway


type AppmeshVirtualGatewaySpecLoggingAccessLogFileFormat struct {
	// json block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.56.0/docs/resources/appmesh_virtual_gateway#json AppmeshVirtualGateway#json}
	Json interface{} `field:"optional" json:"json" yaml:"json"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.56.0/docs/resources/appmesh_virtual_gateway#text AppmeshVirtualGateway#text}.
	Text *string `field:"optional" json:"text" yaml:"text"`
}

