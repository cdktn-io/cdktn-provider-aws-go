// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakertrainingjob


type SagemakerTrainingJobExperimentConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.45.0/docs/resources/sagemaker_training_job#experiment_name SagemakerTrainingJob#experiment_name}.
	ExperimentName *string `field:"optional" json:"experimentName" yaml:"experimentName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.45.0/docs/resources/sagemaker_training_job#run_name SagemakerTrainingJob#run_name}.
	RunName *string `field:"optional" json:"runName" yaml:"runName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.45.0/docs/resources/sagemaker_training_job#trial_component_display_name SagemakerTrainingJob#trial_component_display_name}.
	TrialComponentDisplayName *string `field:"optional" json:"trialComponentDisplayName" yaml:"trialComponentDisplayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.45.0/docs/resources/sagemaker_training_job#trial_name SagemakerTrainingJob#trial_name}.
	TrialName *string `field:"optional" json:"trialName" yaml:"trialName"`
}

