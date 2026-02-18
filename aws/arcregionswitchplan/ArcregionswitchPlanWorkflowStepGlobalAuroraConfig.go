// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package arcregionswitchplan


type ArcregionswitchPlanWorkflowStepGlobalAuroraConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.32.1/docs/resources/arcregionswitch_plan#behavior ArcregionswitchPlan#behavior}.
	Behavior *string `field:"required" json:"behavior" yaml:"behavior"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.32.1/docs/resources/arcregionswitch_plan#database_cluster_arns ArcregionswitchPlan#database_cluster_arns}.
	DatabaseClusterArns *[]*string `field:"required" json:"databaseClusterArns" yaml:"databaseClusterArns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.32.1/docs/resources/arcregionswitch_plan#global_cluster_identifier ArcregionswitchPlan#global_cluster_identifier}.
	GlobalClusterIdentifier *string `field:"required" json:"globalClusterIdentifier" yaml:"globalClusterIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.32.1/docs/resources/arcregionswitch_plan#cross_account_role ArcregionswitchPlan#cross_account_role}.
	CrossAccountRole *string `field:"optional" json:"crossAccountRole" yaml:"crossAccountRole"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.32.1/docs/resources/arcregionswitch_plan#external_id ArcregionswitchPlan#external_id}.
	ExternalId *string `field:"optional" json:"externalId" yaml:"externalId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.32.1/docs/resources/arcregionswitch_plan#timeout_minutes ArcregionswitchPlan#timeout_minutes}.
	TimeoutMinutes *float64 `field:"optional" json:"timeoutMinutes" yaml:"timeoutMinutes"`
	// ungraceful block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.32.1/docs/resources/arcregionswitch_plan#ungraceful ArcregionswitchPlan#ungraceful}
	Ungraceful interface{} `field:"optional" json:"ungraceful" yaml:"ungraceful"`
}

