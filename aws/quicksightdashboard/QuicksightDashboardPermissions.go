// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksightdashboard


type QuicksightDashboardPermissions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/quicksight_dashboard#actions QuicksightDashboard#actions}.
	Actions *[]*string `field:"required" json:"actions" yaml:"actions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/quicksight_dashboard#principal QuicksightDashboard#principal}.
	Principal *string `field:"required" json:"principal" yaml:"principal"`
}

