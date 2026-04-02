// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package arcregionswitchplan


type ArcregionswitchPlanWorkflowStepCustomActionLambdaConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.39.0/docs/resources/arcregionswitch_plan#region_to_run ArcregionswitchPlan#region_to_run}.
	RegionToRun *string `field:"required" json:"regionToRun" yaml:"regionToRun"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.39.0/docs/resources/arcregionswitch_plan#retry_interval_minutes ArcregionswitchPlan#retry_interval_minutes}.
	RetryIntervalMinutes *float64 `field:"required" json:"retryIntervalMinutes" yaml:"retryIntervalMinutes"`
	// lambda block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.39.0/docs/resources/arcregionswitch_plan#lambda ArcregionswitchPlan#lambda}
	Lambda interface{} `field:"optional" json:"lambda" yaml:"lambda"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.39.0/docs/resources/arcregionswitch_plan#timeout_minutes ArcregionswitchPlan#timeout_minutes}.
	TimeoutMinutes *float64 `field:"optional" json:"timeoutMinutes" yaml:"timeoutMinutes"`
	// ungraceful block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.39.0/docs/resources/arcregionswitch_plan#ungraceful ArcregionswitchPlan#ungraceful}
	Ungraceful interface{} `field:"optional" json:"ungraceful" yaml:"ungraceful"`
}

