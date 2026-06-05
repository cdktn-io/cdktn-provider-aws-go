// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package arczonalshiftzonalautoshiftconfiguration


type ArczonalshiftZonalAutoshiftConfigurationBlockingAlarms struct {
	// ARN of the CloudWatch alarm.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.49.0/docs/resources/arczonalshift_zonal_autoshift_configuration#alarm_identifier ArczonalshiftZonalAutoshiftConfiguration#alarm_identifier}
	AlarmIdentifier *string `field:"required" json:"alarmIdentifier" yaml:"alarmIdentifier"`
	// Type of control condition. Valid value: `CLOUDWATCH`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.49.0/docs/resources/arczonalshift_zonal_autoshift_configuration#type ArczonalshiftZonalAutoshiftConfiguration#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

