// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecsservice


type EcsServiceServiceConnectConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/ecs_service#enabled EcsService#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// access_log_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/ecs_service#access_log_configuration EcsService#access_log_configuration}
	AccessLogConfiguration *EcsServiceServiceConnectConfigurationAccessLogConfiguration `field:"optional" json:"accessLogConfiguration" yaml:"accessLogConfiguration"`
	// log_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/ecs_service#log_configuration EcsService#log_configuration}
	LogConfiguration *EcsServiceServiceConnectConfigurationLogConfiguration `field:"optional" json:"logConfiguration" yaml:"logConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/ecs_service#namespace EcsService#namespace}.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// service block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/ecs_service#service EcsService#service}
	Service interface{} `field:"optional" json:"service" yaml:"service"`
}

