// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecsdaemontaskdefinition


type EcsDaemonTaskDefinitionContainerDefinitionEnvironmentFile struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.56.0/docs/resources/ecs_daemon_task_definition#type EcsDaemonTaskDefinition#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.56.0/docs/resources/ecs_daemon_task_definition#value EcsDaemonTaskDefinition#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

