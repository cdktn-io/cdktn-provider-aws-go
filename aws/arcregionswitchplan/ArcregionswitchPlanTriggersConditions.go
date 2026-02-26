// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package arcregionswitchplan


type ArcregionswitchPlanTriggersConditions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.34.0/docs/resources/arcregionswitch_plan#associated_alarm_name ArcregionswitchPlan#associated_alarm_name}.
	AssociatedAlarmName *string `field:"required" json:"associatedAlarmName" yaml:"associatedAlarmName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.34.0/docs/resources/arcregionswitch_plan#condition ArcregionswitchPlan#condition}.
	Condition *string `field:"required" json:"condition" yaml:"condition"`
}

