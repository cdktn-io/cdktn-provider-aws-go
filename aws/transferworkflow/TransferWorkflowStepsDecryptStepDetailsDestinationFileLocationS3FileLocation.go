// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package transferworkflow


type TransferWorkflowStepsDecryptStepDetailsDestinationFileLocationS3FileLocation struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.38.0/docs/resources/transfer_workflow#bucket TransferWorkflow#bucket}.
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.38.0/docs/resources/transfer_workflow#key TransferWorkflow#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
}

