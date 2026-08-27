// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pinpointsmsvoicev2eventdestination


type Pinpointsmsvoicev2EventDestinationKinesisFirehoseDestination struct {
	// ARN of the Amazon Data Firehose delivery stream that receives the events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pinpointsmsvoicev2_event_destination#delivery_stream_arn Pinpointsmsvoicev2EventDestination#delivery_stream_arn}
	DeliveryStreamArn *string `field:"required" json:"deliveryStreamArn" yaml:"deliveryStreamArn"`
	// ARN of the IAM role that End User Messaging SMS assumes to write to the delivery stream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pinpointsmsvoicev2_event_destination#iam_role_arn Pinpointsmsvoicev2EventDestination#iam_role_arn}
	IamRoleArn *string `field:"required" json:"iamRoleArn" yaml:"iamRoleArn"`
}

