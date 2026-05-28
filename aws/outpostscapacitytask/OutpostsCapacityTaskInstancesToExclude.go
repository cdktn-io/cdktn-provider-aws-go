// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package outpostscapacitytask


type OutpostsCapacityTaskInstancesToExclude struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.47.0/docs/resources/outposts_capacity_task#instances OutpostsCapacityTask#instances}.
	Instances *[]*string `field:"required" json:"instances" yaml:"instances"`
}

