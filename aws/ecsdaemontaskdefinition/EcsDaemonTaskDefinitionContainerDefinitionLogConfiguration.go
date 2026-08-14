// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecsdaemontaskdefinition


type EcsDaemonTaskDefinitionContainerDefinitionLogConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/ecs_daemon_task_definition#log_driver EcsDaemonTaskDefinition#log_driver}.
	LogDriver *string `field:"required" json:"logDriver" yaml:"logDriver"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/ecs_daemon_task_definition#options EcsDaemonTaskDefinition#options}.
	Options *map[string]*string `field:"optional" json:"options" yaml:"options"`
	// secret_option block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/ecs_daemon_task_definition#secret_option EcsDaemonTaskDefinition#secret_option}
	SecretOption interface{} `field:"optional" json:"secretOption" yaml:"secretOption"`
}

