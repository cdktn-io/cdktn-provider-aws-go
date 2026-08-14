// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kinesisfirehosedeliverystream


type KinesisFirehoseDeliveryStreamServerSideEncryption struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/kinesis_firehose_delivery_stream#enabled KinesisFirehoseDeliveryStream#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/kinesis_firehose_delivery_stream#key_arn KinesisFirehoseDeliveryStream#key_arn}.
	KeyArn *string `field:"optional" json:"keyArn" yaml:"keyArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/kinesis_firehose_delivery_stream#key_type KinesisFirehoseDeliveryStream#key_type}.
	KeyType *string `field:"optional" json:"keyType" yaml:"keyType"`
}

