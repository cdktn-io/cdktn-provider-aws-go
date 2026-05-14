// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerhyperparametertuningjob

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SagemakerHyperParameterTuningJobConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.45.0/docs/resources/sagemaker_hyper_parameter_tuning_job#name SagemakerHyperParameterTuningJob#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// autotune block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.45.0/docs/resources/sagemaker_hyper_parameter_tuning_job#autotune SagemakerHyperParameterTuningJob#autotune}
	Autotune interface{} `field:"optional" json:"autotune" yaml:"autotune"`
	// config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.45.0/docs/resources/sagemaker_hyper_parameter_tuning_job#config SagemakerHyperParameterTuningJob#config}
	Config interface{} `field:"optional" json:"config" yaml:"config"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.45.0/docs/resources/sagemaker_hyper_parameter_tuning_job#region SagemakerHyperParameterTuningJob#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.45.0/docs/resources/sagemaker_hyper_parameter_tuning_job#tags SagemakerHyperParameterTuningJob#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.45.0/docs/resources/sagemaker_hyper_parameter_tuning_job#timeouts SagemakerHyperParameterTuningJob#timeouts}
	Timeouts *SagemakerHyperParameterTuningJobTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// training_job_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.45.0/docs/resources/sagemaker_hyper_parameter_tuning_job#training_job_definition SagemakerHyperParameterTuningJob#training_job_definition}
	TrainingJobDefinition interface{} `field:"optional" json:"trainingJobDefinition" yaml:"trainingJobDefinition"`
	// training_job_definitions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.45.0/docs/resources/sagemaker_hyper_parameter_tuning_job#training_job_definitions SagemakerHyperParameterTuningJob#training_job_definitions}
	TrainingJobDefinitions interface{} `field:"optional" json:"trainingJobDefinitions" yaml:"trainingJobDefinitions"`
	// warm_start_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.45.0/docs/resources/sagemaker_hyper_parameter_tuning_job#warm_start_config SagemakerHyperParameterTuningJob#warm_start_config}
	WarmStartConfig interface{} `field:"optional" json:"warmStartConfig" yaml:"warmStartConfig"`
}

