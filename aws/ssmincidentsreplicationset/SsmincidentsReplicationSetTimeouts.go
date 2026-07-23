// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ssmincidentsreplicationset


type SsmincidentsReplicationSetTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.56.0/docs/resources/ssmincidents_replication_set#create SsmincidentsReplicationSet#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.56.0/docs/resources/ssmincidents_replication_set#delete SsmincidentsReplicationSet#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.56.0/docs/resources/ssmincidents_replication_set#update SsmincidentsReplicationSet#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

