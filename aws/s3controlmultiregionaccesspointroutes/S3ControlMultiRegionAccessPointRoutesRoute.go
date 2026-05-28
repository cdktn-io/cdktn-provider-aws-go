// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3controlmultiregionaccesspointroutes


type S3ControlMultiRegionAccessPointRoutesRoute struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.47.0/docs/resources/s3control_multi_region_access_point_routes#bucket S3ControlMultiRegionAccessPointRoutes#bucket}.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.47.0/docs/resources/s3control_multi_region_access_point_routes#region S3ControlMultiRegionAccessPointRoutes#region}.
	Region *string `field:"required" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.47.0/docs/resources/s3control_multi_region_access_point_routes#traffic_dial_percentage S3ControlMultiRegionAccessPointRoutes#traffic_dial_percentage}.
	TrafficDialPercentage *float64 `field:"required" json:"trafficDialPercentage" yaml:"trafficDialPercentage"`
}

