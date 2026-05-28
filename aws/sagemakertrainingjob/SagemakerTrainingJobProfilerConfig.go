// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakertrainingjob


type SagemakerTrainingJobProfilerConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.47.0/docs/resources/sagemaker_training_job#disable_profiler SagemakerTrainingJob#disable_profiler}.
	DisableProfiler interface{} `field:"optional" json:"disableProfiler" yaml:"disableProfiler"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.47.0/docs/resources/sagemaker_training_job#profiling_interval_in_milliseconds SagemakerTrainingJob#profiling_interval_in_milliseconds}.
	ProfilingIntervalInMilliseconds *float64 `field:"optional" json:"profilingIntervalInMilliseconds" yaml:"profilingIntervalInMilliseconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.47.0/docs/resources/sagemaker_training_job#profiling_parameters SagemakerTrainingJob#profiling_parameters}.
	ProfilingParameters *map[string]*string `field:"optional" json:"profilingParameters" yaml:"profilingParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.47.0/docs/resources/sagemaker_training_job#s3_output_path SagemakerTrainingJob#s3_output_path}.
	S3OutputPath *string `field:"optional" json:"s3OutputPath" yaml:"s3OutputPath"`
}

