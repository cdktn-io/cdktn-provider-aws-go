// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mskconnectconnector


type MskconnectConnectorPlugin struct {
	// custom_plugin block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.55.0/docs/resources/mskconnect_connector#custom_plugin MskconnectConnector#custom_plugin}
	CustomPlugin *MskconnectConnectorPluginCustomPlugin `field:"required" json:"customPlugin" yaml:"customPlugin"`
}

