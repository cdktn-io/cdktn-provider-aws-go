// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerhyperparametertuningjob


type SagemakerHyperParameterTuningJobWarmStartConfig struct {
	// parent_hyper_parameter_tuning_jobs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.45.0/docs/resources/sagemaker_hyper_parameter_tuning_job#parent_hyper_parameter_tuning_jobs SagemakerHyperParameterTuningJob#parent_hyper_parameter_tuning_jobs}
	ParentHyperParameterTuningJobs interface{} `field:"optional" json:"parentHyperParameterTuningJobs" yaml:"parentHyperParameterTuningJobs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.45.0/docs/resources/sagemaker_hyper_parameter_tuning_job#warm_start_type SagemakerHyperParameterTuningJob#warm_start_type}.
	WarmStartType *string `field:"optional" json:"warmStartType" yaml:"warmStartType"`
}

