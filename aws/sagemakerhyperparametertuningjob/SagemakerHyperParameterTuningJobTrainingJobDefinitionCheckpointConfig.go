// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerhyperparametertuningjob


type SagemakerHyperParameterTuningJobTrainingJobDefinitionCheckpointConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.46.0/docs/resources/sagemaker_hyper_parameter_tuning_job#s3_uri SagemakerHyperParameterTuningJob#s3_uri}.
	S3Uri *string `field:"required" json:"s3Uri" yaml:"s3Uri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.46.0/docs/resources/sagemaker_hyper_parameter_tuning_job#local_path SagemakerHyperParameterTuningJob#local_path}.
	LocalPath *string `field:"optional" json:"localPath" yaml:"localPath"`
}

