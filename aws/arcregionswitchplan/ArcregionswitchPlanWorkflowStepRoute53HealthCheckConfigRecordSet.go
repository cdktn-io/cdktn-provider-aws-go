// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package arcregionswitchplan


type ArcregionswitchPlanWorkflowStepRoute53HealthCheckConfigRecordSet struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.55.0/docs/resources/arcregionswitch_plan#record_set_identifier ArcregionswitchPlan#record_set_identifier}.
	RecordSetIdentifier *string `field:"required" json:"recordSetIdentifier" yaml:"recordSetIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.55.0/docs/resources/arcregionswitch_plan#region ArcregionswitchPlan#region}.
	Region *string `field:"required" json:"region" yaml:"region"`
}

