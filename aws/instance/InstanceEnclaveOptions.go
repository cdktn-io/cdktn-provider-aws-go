// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package instance


type InstanceEnclaveOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.46.0/docs/resources/instance#enabled Instance#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

