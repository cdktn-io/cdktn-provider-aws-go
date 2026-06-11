// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package outpostscapacitytask


type OutpostsCapacityTaskInstancePool struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/outposts_capacity_task#count OutpostsCapacityTask#count}.
	Count *float64 `field:"required" json:"count" yaml:"count"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/outposts_capacity_task#instance_type OutpostsCapacityTask#instance_type}.
	InstanceType *string `field:"required" json:"instanceType" yaml:"instanceType"`
}

