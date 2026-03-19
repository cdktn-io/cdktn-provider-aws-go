// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package arcregionswitchplan


type ArcregionswitchPlanWorkflow struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.37.0/docs/resources/arcregionswitch_plan#workflow_target_action ArcregionswitchPlan#workflow_target_action}.
	WorkflowTargetAction *string `field:"required" json:"workflowTargetAction" yaml:"workflowTargetAction"`
	// step block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.37.0/docs/resources/arcregionswitch_plan#step ArcregionswitchPlan#step}
	Step interface{} `field:"optional" json:"step" yaml:"step"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.37.0/docs/resources/arcregionswitch_plan#workflow_description ArcregionswitchPlan#workflow_description}.
	WorkflowDescription *string `field:"optional" json:"workflowDescription" yaml:"workflowDescription"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.37.0/docs/resources/arcregionswitch_plan#workflow_target_region ArcregionswitchPlan#workflow_target_region}.
	WorkflowTargetRegion *string `field:"optional" json:"workflowTargetRegion" yaml:"workflowTargetRegion"`
}

