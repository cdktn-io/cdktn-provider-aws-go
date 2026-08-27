// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerhyperparametertuningjob


type SagemakerHyperParameterTuningJobConfigStrategyConfigHyperbandStrategyConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_hyper_parameter_tuning_job#max_resource SagemakerHyperParameterTuningJob#max_resource}.
	MaxResource *float64 `field:"optional" json:"maxResource" yaml:"maxResource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_hyper_parameter_tuning_job#min_resource SagemakerHyperParameterTuningJob#min_resource}.
	MinResource *float64 `field:"optional" json:"minResource" yaml:"minResource"`
}

