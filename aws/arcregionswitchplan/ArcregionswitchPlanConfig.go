// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package arcregionswitchplan

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ArcregionswitchPlanConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.49.0/docs/resources/arcregionswitch_plan#execution_role ArcregionswitchPlan#execution_role}.
	ExecutionRole *string `field:"required" json:"executionRole" yaml:"executionRole"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.49.0/docs/resources/arcregionswitch_plan#name ArcregionswitchPlan#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.49.0/docs/resources/arcregionswitch_plan#recovery_approach ArcregionswitchPlan#recovery_approach}.
	RecoveryApproach *string `field:"required" json:"recoveryApproach" yaml:"recoveryApproach"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.49.0/docs/resources/arcregionswitch_plan#regions ArcregionswitchPlan#regions}.
	Regions *[]*string `field:"required" json:"regions" yaml:"regions"`
	// associated_alarms block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.49.0/docs/resources/arcregionswitch_plan#associated_alarms ArcregionswitchPlan#associated_alarms}
	AssociatedAlarms interface{} `field:"optional" json:"associatedAlarms" yaml:"associatedAlarms"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.49.0/docs/resources/arcregionswitch_plan#description ArcregionswitchPlan#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.49.0/docs/resources/arcregionswitch_plan#primary_region ArcregionswitchPlan#primary_region}.
	PrimaryRegion *string `field:"optional" json:"primaryRegion" yaml:"primaryRegion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.49.0/docs/resources/arcregionswitch_plan#recovery_time_objective_minutes ArcregionswitchPlan#recovery_time_objective_minutes}.
	RecoveryTimeObjectiveMinutes *float64 `field:"optional" json:"recoveryTimeObjectiveMinutes" yaml:"recoveryTimeObjectiveMinutes"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.49.0/docs/resources/arcregionswitch_plan#region ArcregionswitchPlan#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// report_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.49.0/docs/resources/arcregionswitch_plan#report_configuration ArcregionswitchPlan#report_configuration}
	ReportConfiguration interface{} `field:"optional" json:"reportConfiguration" yaml:"reportConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.49.0/docs/resources/arcregionswitch_plan#tags ArcregionswitchPlan#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.49.0/docs/resources/arcregionswitch_plan#timeouts ArcregionswitchPlan#timeouts}
	Timeouts *ArcregionswitchPlanTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// triggers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.49.0/docs/resources/arcregionswitch_plan#triggers ArcregionswitchPlan#triggers}
	Triggers interface{} `field:"optional" json:"triggers" yaml:"triggers"`
	// workflow block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.49.0/docs/resources/arcregionswitch_plan#workflow ArcregionswitchPlan#workflow}
	Workflow interface{} `field:"optional" json:"workflow" yaml:"workflow"`
}

