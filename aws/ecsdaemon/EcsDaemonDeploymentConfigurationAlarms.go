// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecsdaemon


type EcsDaemonDeploymentConfigurationAlarms struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.55.0/docs/resources/ecs_daemon#alarm_names EcsDaemon#alarm_names}.
	AlarmNames *[]*string `field:"optional" json:"alarmNames" yaml:"alarmNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.55.0/docs/resources/ecs_daemon#enable EcsDaemon#enable}.
	Enable interface{} `field:"optional" json:"enable" yaml:"enable"`
}

