// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecsservice


type EcsServiceServiceConnectConfigurationAccessLogConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/ecs_service#format EcsService#format}.
	Format *string `field:"required" json:"format" yaml:"format"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/ecs_service#include_query_parameters EcsService#include_query_parameters}.
	IncludeQueryParameters *string `field:"optional" json:"includeQueryParameters" yaml:"includeQueryParameters"`
}

