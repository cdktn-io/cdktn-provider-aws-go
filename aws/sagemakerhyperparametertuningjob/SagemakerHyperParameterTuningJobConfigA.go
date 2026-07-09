// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerhyperparametertuningjob


type SagemakerHyperParameterTuningJobConfigA struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.54.0/docs/resources/sagemaker_hyper_parameter_tuning_job#strategy SagemakerHyperParameterTuningJob#strategy}.
	Strategy *string `field:"required" json:"strategy" yaml:"strategy"`
	// objective block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.54.0/docs/resources/sagemaker_hyper_parameter_tuning_job#objective SagemakerHyperParameterTuningJob#objective}
	Objective interface{} `field:"optional" json:"objective" yaml:"objective"`
	// parameter_ranges block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.54.0/docs/resources/sagemaker_hyper_parameter_tuning_job#parameter_ranges SagemakerHyperParameterTuningJob#parameter_ranges}
	ParameterRanges interface{} `field:"optional" json:"parameterRanges" yaml:"parameterRanges"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.54.0/docs/resources/sagemaker_hyper_parameter_tuning_job#random_seed SagemakerHyperParameterTuningJob#random_seed}.
	RandomSeed *float64 `field:"optional" json:"randomSeed" yaml:"randomSeed"`
	// resource_limits block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.54.0/docs/resources/sagemaker_hyper_parameter_tuning_job#resource_limits SagemakerHyperParameterTuningJob#resource_limits}
	ResourceLimits interface{} `field:"optional" json:"resourceLimits" yaml:"resourceLimits"`
	// strategy_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.54.0/docs/resources/sagemaker_hyper_parameter_tuning_job#strategy_config SagemakerHyperParameterTuningJob#strategy_config}
	StrategyConfig interface{} `field:"optional" json:"strategyConfig" yaml:"strategyConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.54.0/docs/resources/sagemaker_hyper_parameter_tuning_job#training_job_early_stopping_type SagemakerHyperParameterTuningJob#training_job_early_stopping_type}.
	TrainingJobEarlyStoppingType *string `field:"optional" json:"trainingJobEarlyStoppingType" yaml:"trainingJobEarlyStoppingType"`
	// tuning_job_completion_criteria block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.54.0/docs/resources/sagemaker_hyper_parameter_tuning_job#tuning_job_completion_criteria SagemakerHyperParameterTuningJob#tuning_job_completion_criteria}
	TuningJobCompletionCriteria interface{} `field:"optional" json:"tuningJobCompletionCriteria" yaml:"tuningJobCompletionCriteria"`
}

