// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mskreplicator


type MskReplicatorLogDeliveryReplicatorLogDelivery struct {
	// cloudwatch_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/msk_replicator#cloudwatch_logs MskReplicator#cloudwatch_logs}
	CloudwatchLogs *MskReplicatorLogDeliveryReplicatorLogDeliveryCloudwatchLogs `field:"optional" json:"cloudwatchLogs" yaml:"cloudwatchLogs"`
	// firehose block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/msk_replicator#firehose MskReplicator#firehose}
	Firehose *MskReplicatorLogDeliveryReplicatorLogDeliveryFirehose `field:"optional" json:"firehose" yaml:"firehose"`
	// s3 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/msk_replicator#s3 MskReplicator#s3}
	S3 *MskReplicatorLogDeliveryReplicatorLogDeliveryS3 `field:"optional" json:"s3" yaml:"s3"`
}

