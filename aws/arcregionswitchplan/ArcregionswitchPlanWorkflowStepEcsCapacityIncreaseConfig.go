// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package arcregionswitchplan


type ArcregionswitchPlanWorkflowStepEcsCapacityIncreaseConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.54.0/docs/resources/arcregionswitch_plan#capacity_monitoring_approach ArcregionswitchPlan#capacity_monitoring_approach}.
	CapacityMonitoringApproach *string `field:"required" json:"capacityMonitoringApproach" yaml:"capacityMonitoringApproach"`
	// service block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.54.0/docs/resources/arcregionswitch_plan#service ArcregionswitchPlan#service}
	Service interface{} `field:"optional" json:"service" yaml:"service"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.54.0/docs/resources/arcregionswitch_plan#target_percent ArcregionswitchPlan#target_percent}.
	TargetPercent *float64 `field:"optional" json:"targetPercent" yaml:"targetPercent"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.54.0/docs/resources/arcregionswitch_plan#timeout_minutes ArcregionswitchPlan#timeout_minutes}.
	TimeoutMinutes *float64 `field:"optional" json:"timeoutMinutes" yaml:"timeoutMinutes"`
	// ungraceful block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.54.0/docs/resources/arcregionswitch_plan#ungraceful ArcregionswitchPlan#ungraceful}
	Ungraceful interface{} `field:"optional" json:"ungraceful" yaml:"ungraceful"`
}

