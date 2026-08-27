// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerhyperparametertuningjob


type SagemakerHyperParameterTuningJobTrainingJobDefinitionsVpcConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_hyper_parameter_tuning_job#security_group_ids SagemakerHyperParameterTuningJob#security_group_ids}.
	SecurityGroupIds *[]*string `field:"required" json:"securityGroupIds" yaml:"securityGroupIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_hyper_parameter_tuning_job#subnets SagemakerHyperParameterTuningJob#subnets}.
	Subnets *[]*string `field:"required" json:"subnets" yaml:"subnets"`
}

