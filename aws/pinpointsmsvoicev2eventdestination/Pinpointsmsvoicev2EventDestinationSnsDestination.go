// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pinpointsmsvoicev2eventdestination


type Pinpointsmsvoicev2EventDestinationSnsDestination struct {
	// ARN of the Amazon SNS topic that receives the events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/pinpointsmsvoicev2_event_destination#topic_arn Pinpointsmsvoicev2EventDestination#topic_arn}
	TopicArn *string `field:"required" json:"topicArn" yaml:"topicArn"`
}

