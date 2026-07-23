// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksightdashboard


type QuicksightDashboardSourceEntity struct {
	// source_template block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.56.0/docs/resources/quicksight_dashboard#source_template QuicksightDashboard#source_template}
	SourceTemplate *QuicksightDashboardSourceEntitySourceTemplate `field:"optional" json:"sourceTemplate" yaml:"sourceTemplate"`
}

