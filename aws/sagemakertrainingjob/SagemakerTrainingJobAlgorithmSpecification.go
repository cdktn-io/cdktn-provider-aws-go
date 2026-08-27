// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakertrainingjob


type SagemakerTrainingJobAlgorithmSpecification struct {
	// Name or ARN of a SageMaker algorithm resource. Exactly one of `algorithm_name` or `training_image` must be set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#algorithm_name SagemakerTrainingJob#algorithm_name}
	AlgorithmName *string `field:"optional" json:"algorithmName" yaml:"algorithmName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#container_arguments SagemakerTrainingJob#container_arguments}.
	ContainerArguments *[]*string `field:"optional" json:"containerArguments" yaml:"containerArguments"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#container_entrypoint SagemakerTrainingJob#container_entrypoint}.
	ContainerEntrypoint *[]*string `field:"optional" json:"containerEntrypoint" yaml:"containerEntrypoint"`
	// Whether SageMaker AI should publish time-series metrics.
	//
	// SageMaker enables this automatically for built-in algorithms, supported prebuilt images, and jobs with explicit `metric_definitions`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#enable_sagemaker_metrics_time_series SagemakerTrainingJob#enable_sagemaker_metrics_time_series}
	EnableSagemakerMetricsTimeSeries interface{} `field:"optional" json:"enableSagemakerMetricsTimeSeries" yaml:"enableSagemakerMetricsTimeSeries"`
	// metric_definitions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#metric_definitions SagemakerTrainingJob#metric_definitions}
	MetricDefinitions interface{} `field:"optional" json:"metricDefinitions" yaml:"metricDefinitions"`
	// Registry path of the training image.
	//
	// Exactly one of `algorithm_name` or `training_image` must be set. Use `metric_definitions` only when you need to extract custom metrics from your own training container logs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#training_image SagemakerTrainingJob#training_image}
	TrainingImage *string `field:"optional" json:"trainingImage" yaml:"trainingImage"`
	// training_image_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#training_image_config SagemakerTrainingJob#training_image_config}
	TrainingImageConfig interface{} `field:"optional" json:"trainingImageConfig" yaml:"trainingImageConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#training_input_mode SagemakerTrainingJob#training_input_mode}.
	TrainingInputMode *string `field:"optional" json:"trainingInputMode" yaml:"trainingInputMode"`
}

