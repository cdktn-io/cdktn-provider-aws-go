// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package notificationsmanagednotificationadditionalchannelassociation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NotificationsManagedNotificationAdditionalChannelAssociationConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.32.1/docs/resources/notifications_managed_notification_additional_channel_association#channel_arn NotificationsManagedNotificationAdditionalChannelAssociation#channel_arn}.
	ChannelArn *string `field:"required" json:"channelArn" yaml:"channelArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.32.1/docs/resources/notifications_managed_notification_additional_channel_association#managed_notification_arn NotificationsManagedNotificationAdditionalChannelAssociation#managed_notification_arn}.
	ManagedNotificationArn *string `field:"required" json:"managedNotificationArn" yaml:"managedNotificationArn"`
}

