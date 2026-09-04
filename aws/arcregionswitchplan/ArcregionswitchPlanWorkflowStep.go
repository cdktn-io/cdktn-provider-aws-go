// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package arcregionswitchplan


type ArcregionswitchPlanWorkflowStep struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/arcregionswitch_plan#execution_block_type ArcregionswitchPlan#execution_block_type}.
	ExecutionBlockType *string `field:"required" json:"executionBlockType" yaml:"executionBlockType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/arcregionswitch_plan#name ArcregionswitchPlan#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// arc_routing_control_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/arcregionswitch_plan#arc_routing_control_config ArcregionswitchPlan#arc_routing_control_config}
	ArcRoutingControlConfig interface{} `field:"optional" json:"arcRoutingControlConfig" yaml:"arcRoutingControlConfig"`
	// aurora_provisioned_scaling_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/arcregionswitch_plan#aurora_provisioned_scaling_config ArcregionswitchPlan#aurora_provisioned_scaling_config}
	AuroraProvisionedScalingConfig interface{} `field:"optional" json:"auroraProvisionedScalingConfig" yaml:"auroraProvisionedScalingConfig"`
	// aurora_serverless_scaling_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/arcregionswitch_plan#aurora_serverless_scaling_config ArcregionswitchPlan#aurora_serverless_scaling_config}
	AuroraServerlessScalingConfig interface{} `field:"optional" json:"auroraServerlessScalingConfig" yaml:"auroraServerlessScalingConfig"`
	// custom_action_lambda_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/arcregionswitch_plan#custom_action_lambda_config ArcregionswitchPlan#custom_action_lambda_config}
	CustomActionLambdaConfig interface{} `field:"optional" json:"customActionLambdaConfig" yaml:"customActionLambdaConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/arcregionswitch_plan#description ArcregionswitchPlan#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// document_db_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/arcregionswitch_plan#document_db_config ArcregionswitchPlan#document_db_config}
	DocumentDbConfig interface{} `field:"optional" json:"documentDbConfig" yaml:"documentDbConfig"`
	// ec2_asg_capacity_increase_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/arcregionswitch_plan#ec2_asg_capacity_increase_config ArcregionswitchPlan#ec2_asg_capacity_increase_config}
	Ec2AsgCapacityIncreaseConfig interface{} `field:"optional" json:"ec2AsgCapacityIncreaseConfig" yaml:"ec2AsgCapacityIncreaseConfig"`
	// ecs_capacity_increase_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/arcregionswitch_plan#ecs_capacity_increase_config ArcregionswitchPlan#ecs_capacity_increase_config}
	EcsCapacityIncreaseConfig interface{} `field:"optional" json:"ecsCapacityIncreaseConfig" yaml:"ecsCapacityIncreaseConfig"`
	// eks_resource_scaling_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/arcregionswitch_plan#eks_resource_scaling_config ArcregionswitchPlan#eks_resource_scaling_config}
	EksResourceScalingConfig interface{} `field:"optional" json:"eksResourceScalingConfig" yaml:"eksResourceScalingConfig"`
	// execution_approval_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/arcregionswitch_plan#execution_approval_config ArcregionswitchPlan#execution_approval_config}
	ExecutionApprovalConfig interface{} `field:"optional" json:"executionApprovalConfig" yaml:"executionApprovalConfig"`
	// global_aurora_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/arcregionswitch_plan#global_aurora_config ArcregionswitchPlan#global_aurora_config}
	GlobalAuroraConfig interface{} `field:"optional" json:"globalAuroraConfig" yaml:"globalAuroraConfig"`
	// lambda_event_source_mapping_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/arcregionswitch_plan#lambda_event_source_mapping_config ArcregionswitchPlan#lambda_event_source_mapping_config}
	LambdaEventSourceMappingConfig interface{} `field:"optional" json:"lambdaEventSourceMappingConfig" yaml:"lambdaEventSourceMappingConfig"`
	// neptune_global_database_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/arcregionswitch_plan#neptune_global_database_config ArcregionswitchPlan#neptune_global_database_config}
	NeptuneGlobalDatabaseConfig interface{} `field:"optional" json:"neptuneGlobalDatabaseConfig" yaml:"neptuneGlobalDatabaseConfig"`
	// parallel_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/arcregionswitch_plan#parallel_config ArcregionswitchPlan#parallel_config}
	ParallelConfig interface{} `field:"optional" json:"parallelConfig" yaml:"parallelConfig"`
	// rds_create_cross_region_read_replica_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/arcregionswitch_plan#rds_create_cross_region_read_replica_config ArcregionswitchPlan#rds_create_cross_region_read_replica_config}
	RdsCreateCrossRegionReadReplicaConfig interface{} `field:"optional" json:"rdsCreateCrossRegionReadReplicaConfig" yaml:"rdsCreateCrossRegionReadReplicaConfig"`
	// rds_promote_read_replica_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/arcregionswitch_plan#rds_promote_read_replica_config ArcregionswitchPlan#rds_promote_read_replica_config}
	RdsPromoteReadReplicaConfig interface{} `field:"optional" json:"rdsPromoteReadReplicaConfig" yaml:"rdsPromoteReadReplicaConfig"`
	// region_switch_plan_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/arcregionswitch_plan#region_switch_plan_config ArcregionswitchPlan#region_switch_plan_config}
	RegionSwitchPlanConfig interface{} `field:"optional" json:"regionSwitchPlanConfig" yaml:"regionSwitchPlanConfig"`
	// route53_health_check_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/arcregionswitch_plan#route53_health_check_config ArcregionswitchPlan#route53_health_check_config}
	Route53HealthCheckConfig interface{} `field:"optional" json:"route53HealthCheckConfig" yaml:"route53HealthCheckConfig"`
}

