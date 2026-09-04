// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package workspacesdirectory


type WorkspacesDirectoryWorkspaceAccessPropertiesAccessEndpointConfig struct {
	// access_endpoints block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/workspaces_directory#access_endpoints WorkspacesDirectory#access_endpoints}
	AccessEndpoints interface{} `field:"required" json:"accessEndpoints" yaml:"accessEndpoints"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/workspaces_directory#internet_fallback_protocols WorkspacesDirectory#internet_fallback_protocols}.
	InternetFallbackProtocols *[]*string `field:"optional" json:"internetFallbackProtocols" yaml:"internetFallbackProtocols"`
}

