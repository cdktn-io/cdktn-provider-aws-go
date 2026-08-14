// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakertrainingjob


type SagemakerTrainingJobCheckpointConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/sagemaker_training_job#s3_uri SagemakerTrainingJob#s3_uri}.
	S3Uri *string `field:"required" json:"s3Uri" yaml:"s3Uri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/sagemaker_training_job#local_path SagemakerTrainingJob#local_path}.
	LocalPath *string `field:"optional" json:"localPath" yaml:"localPath"`
}

