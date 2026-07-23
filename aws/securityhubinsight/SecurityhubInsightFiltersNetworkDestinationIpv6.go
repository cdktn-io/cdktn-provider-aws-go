// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securityhubinsight


type SecurityhubInsightFiltersNetworkDestinationIpv6 struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.56.0/docs/resources/securityhub_insight#cidr SecurityhubInsight#cidr}.
	Cidr *string `field:"required" json:"cidr" yaml:"cidr"`
}

