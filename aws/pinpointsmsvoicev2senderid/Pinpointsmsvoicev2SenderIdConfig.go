// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pinpointsmsvoicev2senderid

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Pinpointsmsvoicev2SenderIdConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/pinpointsmsvoicev2_sender_id#iso_country_code Pinpointsmsvoicev2SenderId#iso_country_code}.
	IsoCountryCode *string `field:"required" json:"isoCountryCode" yaml:"isoCountryCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/pinpointsmsvoicev2_sender_id#sender_id Pinpointsmsvoicev2SenderId#sender_id}.
	SenderId *string `field:"required" json:"senderId" yaml:"senderId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/pinpointsmsvoicev2_sender_id#deletion_protection_enabled Pinpointsmsvoicev2SenderId#deletion_protection_enabled}.
	DeletionProtectionEnabled interface{} `field:"optional" json:"deletionProtectionEnabled" yaml:"deletionProtectionEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/pinpointsmsvoicev2_sender_id#message_types Pinpointsmsvoicev2SenderId#message_types}.
	MessageTypes *[]*string `field:"optional" json:"messageTypes" yaml:"messageTypes"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/pinpointsmsvoicev2_sender_id#region Pinpointsmsvoicev2SenderId#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/pinpointsmsvoicev2_sender_id#tags Pinpointsmsvoicev2SenderId#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/pinpointsmsvoicev2_sender_id#timeouts Pinpointsmsvoicev2SenderId#timeouts}
	Timeouts *Pinpointsmsvoicev2SenderIdTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

