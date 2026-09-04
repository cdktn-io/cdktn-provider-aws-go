// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package arcregionswitchplan


type ArcregionswitchPlanWorkflowStepParallelConfigStepLambdaEventSourceMappingConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/arcregionswitch_plan#action ArcregionswitchPlan#action}.
	Action *string `field:"required" json:"action" yaml:"action"`
	// region_event_source_mapping block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/arcregionswitch_plan#region_event_source_mapping ArcregionswitchPlan#region_event_source_mapping}
	RegionEventSourceMapping interface{} `field:"optional" json:"regionEventSourceMapping" yaml:"regionEventSourceMapping"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/arcregionswitch_plan#timeout_minutes ArcregionswitchPlan#timeout_minutes}.
	TimeoutMinutes *float64 `field:"optional" json:"timeoutMinutes" yaml:"timeoutMinutes"`
	// ungraceful block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/arcregionswitch_plan#ungraceful ArcregionswitchPlan#ungraceful}
	Ungraceful interface{} `field:"optional" json:"ungraceful" yaml:"ungraceful"`
}

