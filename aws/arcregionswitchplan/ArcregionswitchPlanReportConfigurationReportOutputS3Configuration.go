// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package arcregionswitchplan


type ArcregionswitchPlanReportConfigurationReportOutputS3Configuration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/arcregionswitch_plan#bucket_owner ArcregionswitchPlan#bucket_owner}.
	BucketOwner *string `field:"required" json:"bucketOwner" yaml:"bucketOwner"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/arcregionswitch_plan#bucket_path ArcregionswitchPlan#bucket_path}.
	BucketPath *string `field:"required" json:"bucketPath" yaml:"bucketPath"`
}

