// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecscluster


type EcsClusterServiceConnectDefaults struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.37.0/docs/resources/ecs_cluster#namespace EcsCluster#namespace}.
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
}

