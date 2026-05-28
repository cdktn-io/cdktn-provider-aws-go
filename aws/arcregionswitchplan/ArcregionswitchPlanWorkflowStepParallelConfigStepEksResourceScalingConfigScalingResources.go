// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package arcregionswitchplan


type ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigScalingResources struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.47.0/docs/resources/arcregionswitch_plan#namespace ArcregionswitchPlan#namespace}.
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
	// resources block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.47.0/docs/resources/arcregionswitch_plan#resources ArcregionswitchPlan#resources}
	Resources interface{} `field:"optional" json:"resources" yaml:"resources"`
}

