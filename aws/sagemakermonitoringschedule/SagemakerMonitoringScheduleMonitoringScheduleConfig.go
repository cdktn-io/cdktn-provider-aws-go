// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakermonitoringschedule


type SagemakerMonitoringScheduleMonitoringScheduleConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.35.1/docs/resources/sagemaker_monitoring_schedule#monitoring_type SagemakerMonitoringSchedule#monitoring_type}.
	MonitoringType *string `field:"required" json:"monitoringType" yaml:"monitoringType"`
	// monitoring_job_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.35.1/docs/resources/sagemaker_monitoring_schedule#monitoring_job_definition SagemakerMonitoringSchedule#monitoring_job_definition}
	MonitoringJobDefinition *SagemakerMonitoringScheduleMonitoringScheduleConfigMonitoringJobDefinition `field:"optional" json:"monitoringJobDefinition" yaml:"monitoringJobDefinition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.35.1/docs/resources/sagemaker_monitoring_schedule#monitoring_job_definition_name SagemakerMonitoringSchedule#monitoring_job_definition_name}.
	MonitoringJobDefinitionName *string `field:"optional" json:"monitoringJobDefinitionName" yaml:"monitoringJobDefinitionName"`
	// schedule_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.35.1/docs/resources/sagemaker_monitoring_schedule#schedule_config SagemakerMonitoringSchedule#schedule_config}
	ScheduleConfig *SagemakerMonitoringScheduleMonitoringScheduleConfigScheduleConfig `field:"optional" json:"scheduleConfig" yaml:"scheduleConfig"`
}

