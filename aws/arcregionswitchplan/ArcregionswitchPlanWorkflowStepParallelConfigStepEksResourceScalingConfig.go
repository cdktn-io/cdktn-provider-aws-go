// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package arcregionswitchplan


type ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.33.0/docs/resources/arcregionswitch_plan#capacity_monitoring_approach ArcregionswitchPlan#capacity_monitoring_approach}.
	CapacityMonitoringApproach *string `field:"required" json:"capacityMonitoringApproach" yaml:"capacityMonitoringApproach"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.33.0/docs/resources/arcregionswitch_plan#target_percent ArcregionswitchPlan#target_percent}.
	TargetPercent *float64 `field:"required" json:"targetPercent" yaml:"targetPercent"`
	// eks_clusters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.33.0/docs/resources/arcregionswitch_plan#eks_clusters ArcregionswitchPlan#eks_clusters}
	EksClusters interface{} `field:"optional" json:"eksClusters" yaml:"eksClusters"`
	// kubernetes_resource_type block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.33.0/docs/resources/arcregionswitch_plan#kubernetes_resource_type ArcregionswitchPlan#kubernetes_resource_type}
	KubernetesResourceType interface{} `field:"optional" json:"kubernetesResourceType" yaml:"kubernetesResourceType"`
	// scaling_resources block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.33.0/docs/resources/arcregionswitch_plan#scaling_resources ArcregionswitchPlan#scaling_resources}
	ScalingResources interface{} `field:"optional" json:"scalingResources" yaml:"scalingResources"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.33.0/docs/resources/arcregionswitch_plan#timeout_minutes ArcregionswitchPlan#timeout_minutes}.
	TimeoutMinutes *float64 `field:"optional" json:"timeoutMinutes" yaml:"timeoutMinutes"`
	// ungraceful block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.33.0/docs/resources/arcregionswitch_plan#ungraceful ArcregionswitchPlan#ungraceful}
	Ungraceful interface{} `field:"optional" json:"ungraceful" yaml:"ungraceful"`
}

