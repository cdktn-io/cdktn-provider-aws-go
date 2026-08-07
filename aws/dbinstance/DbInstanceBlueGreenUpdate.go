// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dbinstance


type DbInstanceBlueGreenUpdate struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.58.0/docs/resources/db_instance#enabled DbInstance#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

