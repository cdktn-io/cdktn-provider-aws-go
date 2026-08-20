// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package workspacespool


type WorkspacesPoolApplicationSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.61.0/docs/resources/workspaces_pool#settings_group WorkspacesPool#settings_group}.
	SettingsGroup *string `field:"optional" json:"settingsGroup" yaml:"settingsGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.61.0/docs/resources/workspaces_pool#status WorkspacesPool#status}.
	Status *string `field:"optional" json:"status" yaml:"status"`
}

