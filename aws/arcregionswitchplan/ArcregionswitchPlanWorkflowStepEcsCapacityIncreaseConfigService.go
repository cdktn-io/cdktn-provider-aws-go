// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package arcregionswitchplan


type ArcregionswitchPlanWorkflowStepEcsCapacityIncreaseConfigService struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.56.0/docs/resources/arcregionswitch_plan#cluster_arn ArcregionswitchPlan#cluster_arn}.
	ClusterArn *string `field:"required" json:"clusterArn" yaml:"clusterArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.56.0/docs/resources/arcregionswitch_plan#service_arn ArcregionswitchPlan#service_arn}.
	ServiceArn *string `field:"required" json:"serviceArn" yaml:"serviceArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.56.0/docs/resources/arcregionswitch_plan#cross_account_role ArcregionswitchPlan#cross_account_role}.
	CrossAccountRole *string `field:"optional" json:"crossAccountRole" yaml:"crossAccountRole"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.56.0/docs/resources/arcregionswitch_plan#external_id ArcregionswitchPlan#external_id}.
	ExternalId *string `field:"optional" json:"externalId" yaml:"externalId"`
}

