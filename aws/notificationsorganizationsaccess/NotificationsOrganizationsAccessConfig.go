// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package notificationsorganizationsaccess

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type NotificationsOrganizationsAccessConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/notifications_organizations_access#enabled NotificationsOrganizationsAccess#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/notifications_organizations_access#timeouts NotificationsOrganizationsAccess#timeouts}
	Timeouts *NotificationsOrganizationsAccessTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

