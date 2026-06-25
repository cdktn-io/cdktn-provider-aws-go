// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksightdataset


type QuicksightDataSetRefreshProperties struct {
	// refresh_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/quicksight_data_set#refresh_configuration QuicksightDataSet#refresh_configuration}
	RefreshConfiguration *QuicksightDataSetRefreshPropertiesRefreshConfiguration `field:"required" json:"refreshConfiguration" yaml:"refreshConfiguration"`
}

