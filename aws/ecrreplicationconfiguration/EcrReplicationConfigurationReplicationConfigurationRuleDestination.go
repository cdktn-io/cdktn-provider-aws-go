// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecrreplicationconfiguration


type EcrReplicationConfigurationReplicationConfigurationRuleDestination struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.31.0/docs/resources/ecr_replication_configuration#region EcrReplicationConfiguration#region}.
	Region *string `field:"required" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.31.0/docs/resources/ecr_replication_configuration#registry_id EcrReplicationConfiguration#registry_id}.
	RegistryId *string `field:"required" json:"registryId" yaml:"registryId"`
}

