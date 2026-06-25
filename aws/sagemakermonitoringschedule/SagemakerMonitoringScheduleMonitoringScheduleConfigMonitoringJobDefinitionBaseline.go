// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakermonitoringschedule


type SagemakerMonitoringScheduleMonitoringScheduleConfigMonitoringJobDefinitionBaseline struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/sagemaker_monitoring_schedule#baselining_job_name SagemakerMonitoringSchedule#baselining_job_name}.
	BaseliningJobName *string `field:"optional" json:"baseliningJobName" yaml:"baseliningJobName"`
	// constraints_resource block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/sagemaker_monitoring_schedule#constraints_resource SagemakerMonitoringSchedule#constraints_resource}
	ConstraintsResource *SagemakerMonitoringScheduleMonitoringScheduleConfigMonitoringJobDefinitionBaselineConstraintsResource `field:"optional" json:"constraintsResource" yaml:"constraintsResource"`
	// statistics_resource block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/sagemaker_monitoring_schedule#statistics_resource SagemakerMonitoringSchedule#statistics_resource}
	StatisticsResource *SagemakerMonitoringScheduleMonitoringScheduleConfigMonitoringJobDefinitionBaselineStatisticsResource `field:"optional" json:"statisticsResource" yaml:"statisticsResource"`
}

