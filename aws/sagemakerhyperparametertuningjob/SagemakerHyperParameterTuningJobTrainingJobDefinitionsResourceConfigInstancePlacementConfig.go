// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerhyperparametertuningjob


type SagemakerHyperParameterTuningJobTrainingJobDefinitionsResourceConfigInstancePlacementConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/sagemaker_hyper_parameter_tuning_job#enable_multiple_jobs SagemakerHyperParameterTuningJob#enable_multiple_jobs}.
	EnableMultipleJobs interface{} `field:"optional" json:"enableMultipleJobs" yaml:"enableMultipleJobs"`
	// placement_specifications block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/sagemaker_hyper_parameter_tuning_job#placement_specifications SagemakerHyperParameterTuningJob#placement_specifications}
	PlacementSpecifications interface{} `field:"optional" json:"placementSpecifications" yaml:"placementSpecifications"`
}

