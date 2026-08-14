// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package batchjobdefinition


type BatchJobDefinitionEksProperties struct {
	// pod_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/batch_job_definition#pod_properties BatchJobDefinition#pod_properties}
	PodProperties *BatchJobDefinitionEksPropertiesPodProperties `field:"required" json:"podProperties" yaml:"podProperties"`
}

