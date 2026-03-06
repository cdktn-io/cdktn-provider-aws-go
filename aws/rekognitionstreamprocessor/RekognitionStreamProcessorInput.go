// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package rekognitionstreamprocessor


type RekognitionStreamProcessorInput struct {
	// kinesis_video_stream block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.35.1/docs/resources/rekognition_stream_processor#kinesis_video_stream RekognitionStreamProcessor#kinesis_video_stream}
	KinesisVideoStream interface{} `field:"optional" json:"kinesisVideoStream" yaml:"kinesisVideoStream"`
}

