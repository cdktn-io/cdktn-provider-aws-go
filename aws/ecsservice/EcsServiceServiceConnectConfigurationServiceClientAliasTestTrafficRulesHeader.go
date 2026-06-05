// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecsservice


type EcsServiceServiceConnectConfigurationServiceClientAliasTestTrafficRulesHeader struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.49.0/docs/resources/ecs_service#name EcsService#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// value block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.49.0/docs/resources/ecs_service#value EcsService#value}
	Value *EcsServiceServiceConnectConfigurationServiceClientAliasTestTrafficRulesHeaderValue `field:"required" json:"value" yaml:"value"`
}

