// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package notificationsorganizationalunitassociation

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type NotificationsOrganizationalUnitAssociationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.35.1/docs/resources/notifications_organizational_unit_association#notification_configuration_arn NotificationsOrganizationalUnitAssociation#notification_configuration_arn}.
	NotificationConfigurationArn *string `field:"required" json:"notificationConfigurationArn" yaml:"notificationConfigurationArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.35.1/docs/resources/notifications_organizational_unit_association#organizational_unit_id NotificationsOrganizationalUnitAssociation#organizational_unit_id}.
	OrganizationalUnitId *string `field:"required" json:"organizationalUnitId" yaml:"organizationalUnitId"`
}

