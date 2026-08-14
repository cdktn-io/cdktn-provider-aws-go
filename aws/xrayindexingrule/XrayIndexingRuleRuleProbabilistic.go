// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package xrayindexingrule


type XrayIndexingRuleRuleProbabilistic struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/xray_indexing_rule#desired_sampling_percentage XrayIndexingRule#desired_sampling_percentage}.
	DesiredSamplingPercentage *float64 `field:"required" json:"desiredSamplingPercentage" yaml:"desiredSamplingPercentage"`
}

