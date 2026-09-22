// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package fissafetyleverstate


type FisSafetyLeverStateState struct {
	// Reason for the current status of the safety lever.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.66.0/docs/resources/fis_safety_lever_state#reason FisSafetyLeverState#reason}
	Reason *string `field:"required" json:"reason" yaml:"reason"`
	// State of the safety lever. Valid values: engaged, disengaged.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.66.0/docs/resources/fis_safety_lever_state#status FisSafetyLeverState#status}
	Status *string `field:"required" json:"status" yaml:"status"`
}

