// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package arcregionswitchplan


type ArcregionswitchPlanTriggers struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.47.0/docs/resources/arcregionswitch_plan#action ArcregionswitchPlan#action}.
	Action *string `field:"required" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.47.0/docs/resources/arcregionswitch_plan#min_delay_minutes_between_executions ArcregionswitchPlan#min_delay_minutes_between_executions}.
	MinDelayMinutesBetweenExecutions *float64 `field:"required" json:"minDelayMinutesBetweenExecutions" yaml:"minDelayMinutesBetweenExecutions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.47.0/docs/resources/arcregionswitch_plan#target_region ArcregionswitchPlan#target_region}.
	TargetRegion *string `field:"required" json:"targetRegion" yaml:"targetRegion"`
	// conditions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.47.0/docs/resources/arcregionswitch_plan#conditions ArcregionswitchPlan#conditions}
	Conditions interface{} `field:"optional" json:"conditions" yaml:"conditions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.47.0/docs/resources/arcregionswitch_plan#description ArcregionswitchPlan#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

