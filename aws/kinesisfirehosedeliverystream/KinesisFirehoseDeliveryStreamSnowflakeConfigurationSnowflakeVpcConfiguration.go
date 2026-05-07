// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kinesisfirehosedeliverystream


type KinesisFirehoseDeliveryStreamSnowflakeConfigurationSnowflakeVpcConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.44.0/docs/resources/kinesis_firehose_delivery_stream#private_link_vpce_id KinesisFirehoseDeliveryStream#private_link_vpce_id}.
	PrivateLinkVpceId *string `field:"required" json:"privateLinkVpceId" yaml:"privateLinkVpceId"`
}

