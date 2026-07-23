// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecsdaemontaskdefinition


type EcsDaemonTaskDefinitionContainerDefinitionLinuxParametersCapabilities struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.56.0/docs/resources/ecs_daemon_task_definition#add EcsDaemonTaskDefinition#add}.
	Add *[]*string `field:"optional" json:"add" yaml:"add"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.56.0/docs/resources/ecs_daemon_task_definition#drop EcsDaemonTaskDefinition#drop}.
	Drop *[]*string `field:"optional" json:"drop" yaml:"drop"`
}

