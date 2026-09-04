// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package arcregionswitchplan


type ArcregionswitchPlanWorkflowStepArcRoutingControlConfigRegionAndRoutingControls struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/arcregionswitch_plan#region ArcregionswitchPlan#region}.
	Region *string `field:"required" json:"region" yaml:"region"`
	// routing_control block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/arcregionswitch_plan#routing_control ArcregionswitchPlan#routing_control}
	RoutingControl interface{} `field:"optional" json:"routingControl" yaml:"routingControl"`
}

