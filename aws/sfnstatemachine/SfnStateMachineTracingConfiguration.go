// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sfnstatemachine


type SfnStateMachineTracingConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/sfn_state_machine#enabled SfnStateMachine#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

