// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerhyperparametertuningjob


type SagemakerHyperParameterTuningJobConfigTuningJobCompletionCriteria struct {
	// best_objective_not_improving block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.51.0/docs/resources/sagemaker_hyper_parameter_tuning_job#best_objective_not_improving SagemakerHyperParameterTuningJob#best_objective_not_improving}
	BestObjectiveNotImproving interface{} `field:"optional" json:"bestObjectiveNotImproving" yaml:"bestObjectiveNotImproving"`
	// convergence_detected block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.51.0/docs/resources/sagemaker_hyper_parameter_tuning_job#convergence_detected SagemakerHyperParameterTuningJob#convergence_detected}
	ConvergenceDetected interface{} `field:"optional" json:"convergenceDetected" yaml:"convergenceDetected"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.51.0/docs/resources/sagemaker_hyper_parameter_tuning_job#target_objective_metric_value SagemakerHyperParameterTuningJob#target_objective_metric_value}.
	TargetObjectiveMetricValue *float64 `field:"optional" json:"targetObjectiveMetricValue" yaml:"targetObjectiveMetricValue"`
}

