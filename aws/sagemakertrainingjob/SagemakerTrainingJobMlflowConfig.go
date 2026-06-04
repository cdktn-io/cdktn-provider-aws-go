// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakertrainingjob


type SagemakerTrainingJobMlflowConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.48.0/docs/resources/sagemaker_training_job#mlflow_resource_arn SagemakerTrainingJob#mlflow_resource_arn}.
	MlflowResourceArn *string `field:"required" json:"mlflowResourceArn" yaml:"mlflowResourceArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.48.0/docs/resources/sagemaker_training_job#mlflow_experiment_name SagemakerTrainingJob#mlflow_experiment_name}.
	MlflowExperimentName *string `field:"optional" json:"mlflowExperimentName" yaml:"mlflowExperimentName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.48.0/docs/resources/sagemaker_training_job#mlflow_run_name SagemakerTrainingJob#mlflow_run_name}.
	MlflowRunName *string `field:"optional" json:"mlflowRunName" yaml:"mlflowRunName"`
}

