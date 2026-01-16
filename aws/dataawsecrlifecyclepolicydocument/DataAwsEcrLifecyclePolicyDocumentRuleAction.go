// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsecrlifecyclepolicydocument


type DataAwsEcrLifecyclePolicyDocumentRuleAction struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/data-sources/ecr_lifecycle_policy_document#type DataAwsEcrLifecyclePolicyDocument#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/data-sources/ecr_lifecycle_policy_document#target_storage_class DataAwsEcrLifecyclePolicyDocument#target_storage_class}.
	TargetStorageClass *string `field:"optional" json:"targetStorageClass" yaml:"targetStorageClass"`
}

