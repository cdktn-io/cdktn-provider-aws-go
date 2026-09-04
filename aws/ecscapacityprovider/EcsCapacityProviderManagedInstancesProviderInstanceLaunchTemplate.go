// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecscapacityprovider


type EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplate struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/ecs_capacity_provider#ec2_instance_profile_arn EcsCapacityProvider#ec2_instance_profile_arn}.
	Ec2InstanceProfileArn *string `field:"required" json:"ec2InstanceProfileArn" yaml:"ec2InstanceProfileArn"`
	// network_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/ecs_capacity_provider#network_configuration EcsCapacityProvider#network_configuration}
	NetworkConfiguration *EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateNetworkConfiguration `field:"required" json:"networkConfiguration" yaml:"networkConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/ecs_capacity_provider#capacity_option_type EcsCapacityProvider#capacity_option_type}.
	CapacityOptionType *string `field:"optional" json:"capacityOptionType" yaml:"capacityOptionType"`
	// capacity_reservations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/ecs_capacity_provider#capacity_reservations EcsCapacityProvider#capacity_reservations}
	CapacityReservations *EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateCapacityReservations `field:"optional" json:"capacityReservations" yaml:"capacityReservations"`
	// instance_requirements block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/ecs_capacity_provider#instance_requirements EcsCapacityProvider#instance_requirements}
	InstanceRequirements *EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateInstanceRequirements `field:"optional" json:"instanceRequirements" yaml:"instanceRequirements"`
	// local_storage_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/ecs_capacity_provider#local_storage_configuration EcsCapacityProvider#local_storage_configuration}
	LocalStorageConfiguration *EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateLocalStorageConfiguration `field:"optional" json:"localStorageConfiguration" yaml:"localStorageConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/ecs_capacity_provider#monitoring EcsCapacityProvider#monitoring}.
	Monitoring *string `field:"optional" json:"monitoring" yaml:"monitoring"`
	// storage_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/ecs_capacity_provider#storage_configuration EcsCapacityProvider#storage_configuration}
	StorageConfiguration *EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateStorageConfiguration `field:"optional" json:"storageConfiguration" yaml:"storageConfiguration"`
}

