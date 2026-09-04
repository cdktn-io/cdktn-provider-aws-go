// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package workspacespool


type WorkspacesPoolCapacity struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/workspaces_pool#desired_user_sessions WorkspacesPool#desired_user_sessions}.
	DesiredUserSessions *float64 `field:"required" json:"desiredUserSessions" yaml:"desiredUserSessions"`
}

