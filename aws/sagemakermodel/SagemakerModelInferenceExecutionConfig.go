// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakermodel


type SagemakerModelInferenceExecutionConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.35.1/docs/resources/sagemaker_model#mode SagemakerModel#mode}.
	Mode *string `field:"required" json:"mode" yaml:"mode"`
}

