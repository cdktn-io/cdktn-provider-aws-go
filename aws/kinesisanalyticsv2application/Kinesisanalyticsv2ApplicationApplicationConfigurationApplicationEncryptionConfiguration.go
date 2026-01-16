// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kinesisanalyticsv2application


type Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationEncryptionConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/kinesisanalyticsv2_application#key_type Kinesisanalyticsv2Application#key_type}.
	KeyType *string `field:"required" json:"keyType" yaml:"keyType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/kinesisanalyticsv2_application#key_id Kinesisanalyticsv2Application#key_id}.
	KeyId *string `field:"optional" json:"keyId" yaml:"keyId"`
}

