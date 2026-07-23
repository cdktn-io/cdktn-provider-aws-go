// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecsdaemon


type EcsDaemonDeploymentConfiguration struct {
	// alarms block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.56.0/docs/resources/ecs_daemon#alarms EcsDaemon#alarms}
	Alarms interface{} `field:"optional" json:"alarms" yaml:"alarms"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.56.0/docs/resources/ecs_daemon#bake_time_in_minutes EcsDaemon#bake_time_in_minutes}.
	BakeTimeInMinutes *float64 `field:"optional" json:"bakeTimeInMinutes" yaml:"bakeTimeInMinutes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.56.0/docs/resources/ecs_daemon#drain_percent EcsDaemon#drain_percent}.
	DrainPercent *float64 `field:"optional" json:"drainPercent" yaml:"drainPercent"`
}

