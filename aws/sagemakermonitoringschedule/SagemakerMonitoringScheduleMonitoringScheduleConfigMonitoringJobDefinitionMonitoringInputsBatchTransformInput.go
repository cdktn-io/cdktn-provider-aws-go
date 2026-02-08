// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakermonitoringschedule


type SagemakerMonitoringScheduleMonitoringScheduleConfigMonitoringJobDefinitionMonitoringInputsBatchTransformInput struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.31.0/docs/resources/sagemaker_monitoring_schedule#data_captured_destination_s3_uri SagemakerMonitoringSchedule#data_captured_destination_s3_uri}.
	DataCapturedDestinationS3Uri *string `field:"required" json:"dataCapturedDestinationS3Uri" yaml:"dataCapturedDestinationS3Uri"`
	// dataset_format block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.31.0/docs/resources/sagemaker_monitoring_schedule#dataset_format SagemakerMonitoringSchedule#dataset_format}
	DatasetFormat *SagemakerMonitoringScheduleMonitoringScheduleConfigMonitoringJobDefinitionMonitoringInputsBatchTransformInputDatasetFormat `field:"required" json:"datasetFormat" yaml:"datasetFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.31.0/docs/resources/sagemaker_monitoring_schedule#local_path SagemakerMonitoringSchedule#local_path}.
	LocalPath *string `field:"required" json:"localPath" yaml:"localPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.31.0/docs/resources/sagemaker_monitoring_schedule#end_time_offset SagemakerMonitoringSchedule#end_time_offset}.
	EndTimeOffset *string `field:"optional" json:"endTimeOffset" yaml:"endTimeOffset"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.31.0/docs/resources/sagemaker_monitoring_schedule#exclude_features_attribute SagemakerMonitoringSchedule#exclude_features_attribute}.
	ExcludeFeaturesAttribute *string `field:"optional" json:"excludeFeaturesAttribute" yaml:"excludeFeaturesAttribute"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.31.0/docs/resources/sagemaker_monitoring_schedule#features_attribute SagemakerMonitoringSchedule#features_attribute}.
	FeaturesAttribute *string `field:"optional" json:"featuresAttribute" yaml:"featuresAttribute"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.31.0/docs/resources/sagemaker_monitoring_schedule#inference_attribute SagemakerMonitoringSchedule#inference_attribute}.
	InferenceAttribute *string `field:"optional" json:"inferenceAttribute" yaml:"inferenceAttribute"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.31.0/docs/resources/sagemaker_monitoring_schedule#probability_attribute SagemakerMonitoringSchedule#probability_attribute}.
	ProbabilityAttribute *string `field:"optional" json:"probabilityAttribute" yaml:"probabilityAttribute"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.31.0/docs/resources/sagemaker_monitoring_schedule#probability_threshold_attribute SagemakerMonitoringSchedule#probability_threshold_attribute}.
	ProbabilityThresholdAttribute *float64 `field:"optional" json:"probabilityThresholdAttribute" yaml:"probabilityThresholdAttribute"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.31.0/docs/resources/sagemaker_monitoring_schedule#s3_data_distribution_type SagemakerMonitoringSchedule#s3_data_distribution_type}.
	S3DataDistributionType *string `field:"optional" json:"s3DataDistributionType" yaml:"s3DataDistributionType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.31.0/docs/resources/sagemaker_monitoring_schedule#s3_input_mode SagemakerMonitoringSchedule#s3_input_mode}.
	S3InputMode *string `field:"optional" json:"s3InputMode" yaml:"s3InputMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.31.0/docs/resources/sagemaker_monitoring_schedule#start_time_offset SagemakerMonitoringSchedule#start_time_offset}.
	StartTimeOffset *string `field:"optional" json:"startTimeOffset" yaml:"startTimeOffset"`
}

