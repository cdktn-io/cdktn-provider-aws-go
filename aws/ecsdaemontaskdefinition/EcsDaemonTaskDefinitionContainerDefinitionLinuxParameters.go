// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecsdaemontaskdefinition


type EcsDaemonTaskDefinitionContainerDefinitionLinuxParameters struct {
	// capabilities block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.56.0/docs/resources/ecs_daemon_task_definition#capabilities EcsDaemonTaskDefinition#capabilities}
	Capabilities interface{} `field:"optional" json:"capabilities" yaml:"capabilities"`
	// device block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.56.0/docs/resources/ecs_daemon_task_definition#device EcsDaemonTaskDefinition#device}
	Device interface{} `field:"optional" json:"device" yaml:"device"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.56.0/docs/resources/ecs_daemon_task_definition#init_process_enabled EcsDaemonTaskDefinition#init_process_enabled}.
	InitProcessEnabled interface{} `field:"optional" json:"initProcessEnabled" yaml:"initProcessEnabled"`
	// tmpfs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.56.0/docs/resources/ecs_daemon_task_definition#tmpfs EcsDaemonTaskDefinition#tmpfs}
	Tmpfs interface{} `field:"optional" json:"tmpfs" yaml:"tmpfs"`
}

