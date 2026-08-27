// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package rekognitionstreamprocessor


type RekognitionStreamProcessorOutputKinesisDataStream struct {
	// ARN of the output Amazon Kinesis Data Streams stream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rekognition_stream_processor#arn RekognitionStreamProcessor#arn}
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
}

