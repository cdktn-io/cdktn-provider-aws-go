// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package arcregionswitchplan


type ArcregionswitchPlanWorkflowStepAuroraServerlessScalingConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#global_cluster_identifier ArcregionswitchPlan#global_cluster_identifier}.
	GlobalClusterIdentifier *string `field:"required" json:"globalClusterIdentifier" yaml:"globalClusterIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#region_database_cluster_arns ArcregionswitchPlan#region_database_cluster_arns}.
	RegionDatabaseClusterArns *map[string]*string `field:"required" json:"regionDatabaseClusterArns" yaml:"regionDatabaseClusterArns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#cross_account_role ArcregionswitchPlan#cross_account_role}.
	CrossAccountRole *string `field:"optional" json:"crossAccountRole" yaml:"crossAccountRole"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#external_id ArcregionswitchPlan#external_id}.
	ExternalId *string `field:"optional" json:"externalId" yaml:"externalId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#target_percent ArcregionswitchPlan#target_percent}.
	TargetPercent *float64 `field:"optional" json:"targetPercent" yaml:"targetPercent"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#timeout_minutes ArcregionswitchPlan#timeout_minutes}.
	TimeoutMinutes *float64 `field:"optional" json:"timeoutMinutes" yaml:"timeoutMinutes"`
}

