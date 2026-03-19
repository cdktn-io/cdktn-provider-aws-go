// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package route53record


type Route53RecordLatencyRoutingPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.37.0/docs/resources/route53_record#region Route53Record#region}.
	Region *string `field:"required" json:"region" yaml:"region"`
}

