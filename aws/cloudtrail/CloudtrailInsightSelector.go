// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudtrail


type CloudtrailInsightSelector struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.35.0/docs/resources/cloudtrail#insight_type Cloudtrail#insight_type}.
	InsightType *string `field:"required" json:"insightType" yaml:"insightType"`
}

