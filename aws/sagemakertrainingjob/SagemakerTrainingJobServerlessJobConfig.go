// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakertrainingjob


type SagemakerTrainingJobServerlessJobConfig struct {
	// Base model ARN in SageMaker Public Hub. SageMaker always selects the latest version of the provided model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/sagemaker_training_job#base_model_arn SagemakerTrainingJob#base_model_arn}
	BaseModelArn *string `field:"required" json:"baseModelArn" yaml:"baseModelArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/sagemaker_training_job#job_type SagemakerTrainingJob#job_type}.
	JobType *string `field:"required" json:"jobType" yaml:"jobType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/sagemaker_training_job#accept_eula SagemakerTrainingJob#accept_eula}.
	AcceptEula interface{} `field:"optional" json:"acceptEula" yaml:"acceptEula"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/sagemaker_training_job#customization_technique SagemakerTrainingJob#customization_technique}.
	CustomizationTechnique *string `field:"optional" json:"customizationTechnique" yaml:"customizationTechnique"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/sagemaker_training_job#evaluation_type SagemakerTrainingJob#evaluation_type}.
	EvaluationType *string `field:"optional" json:"evaluationType" yaml:"evaluationType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/sagemaker_training_job#evaluator_arn SagemakerTrainingJob#evaluator_arn}.
	EvaluatorArn *string `field:"optional" json:"evaluatorArn" yaml:"evaluatorArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/sagemaker_training_job#peft SagemakerTrainingJob#peft}.
	Peft *string `field:"optional" json:"peft" yaml:"peft"`
}

