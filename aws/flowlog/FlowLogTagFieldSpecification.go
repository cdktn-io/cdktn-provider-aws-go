// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package flowlog


type FlowLogTagFieldSpecification struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.58.0/docs/resources/flow_log#resource_type FlowLog#resource_type}.
	ResourceType *string `field:"required" json:"resourceType" yaml:"resourceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.58.0/docs/resources/flow_log#tag_keys FlowLog#tag_keys}.
	TagKeys *[]*string `field:"required" json:"tagKeys" yaml:"tagKeys"`
}

