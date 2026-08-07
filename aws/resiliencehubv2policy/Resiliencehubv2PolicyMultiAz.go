// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package resiliencehubv2policy


type Resiliencehubv2PolicyMultiAz struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.58.0/docs/resources/resiliencehubv2_policy#disaster_recovery_approach Resiliencehubv2Policy#disaster_recovery_approach}.
	DisasterRecoveryApproach *string `field:"required" json:"disasterRecoveryApproach" yaml:"disasterRecoveryApproach"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.58.0/docs/resources/resiliencehubv2_policy#rpo_in_minutes Resiliencehubv2Policy#rpo_in_minutes}.
	RpoInMinutes *float64 `field:"optional" json:"rpoInMinutes" yaml:"rpoInMinutes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.58.0/docs/resources/resiliencehubv2_policy#rto_in_minutes Resiliencehubv2Policy#rto_in_minutes}.
	RtoInMinutes *float64 `field:"optional" json:"rtoInMinutes" yaml:"rtoInMinutes"`
}

