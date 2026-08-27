// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakermonitoringschedule


type SagemakerMonitoringScheduleMonitoringScheduleConfigMonitoringJobDefinition struct {
	// monitoring_app_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_monitoring_schedule#monitoring_app_specification SagemakerMonitoringSchedule#monitoring_app_specification}
	MonitoringAppSpecification *SagemakerMonitoringScheduleMonitoringScheduleConfigMonitoringJobDefinitionMonitoringAppSpecification `field:"required" json:"monitoringAppSpecification" yaml:"monitoringAppSpecification"`
	// monitoring_inputs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_monitoring_schedule#monitoring_inputs SagemakerMonitoringSchedule#monitoring_inputs}
	MonitoringInputs *SagemakerMonitoringScheduleMonitoringScheduleConfigMonitoringJobDefinitionMonitoringInputs `field:"required" json:"monitoringInputs" yaml:"monitoringInputs"`
	// monitoring_output_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_monitoring_schedule#monitoring_output_config SagemakerMonitoringSchedule#monitoring_output_config}
	MonitoringOutputConfig *SagemakerMonitoringScheduleMonitoringScheduleConfigMonitoringJobDefinitionMonitoringOutputConfig `field:"required" json:"monitoringOutputConfig" yaml:"monitoringOutputConfig"`
	// monitoring_resources block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_monitoring_schedule#monitoring_resources SagemakerMonitoringSchedule#monitoring_resources}
	MonitoringResources *SagemakerMonitoringScheduleMonitoringScheduleConfigMonitoringJobDefinitionMonitoringResources `field:"required" json:"monitoringResources" yaml:"monitoringResources"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_monitoring_schedule#role_arn SagemakerMonitoringSchedule#role_arn}.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// baseline block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_monitoring_schedule#baseline SagemakerMonitoringSchedule#baseline}
	Baseline *SagemakerMonitoringScheduleMonitoringScheduleConfigMonitoringJobDefinitionBaseline `field:"optional" json:"baseline" yaml:"baseline"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_monitoring_schedule#environment SagemakerMonitoringSchedule#environment}.
	Environment *map[string]*string `field:"optional" json:"environment" yaml:"environment"`
	// network_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_monitoring_schedule#network_config SagemakerMonitoringSchedule#network_config}
	NetworkConfig *SagemakerMonitoringScheduleMonitoringScheduleConfigMonitoringJobDefinitionNetworkConfig `field:"optional" json:"networkConfig" yaml:"networkConfig"`
	// stopping_condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_monitoring_schedule#stopping_condition SagemakerMonitoringSchedule#stopping_condition}
	StoppingCondition interface{} `field:"optional" json:"stoppingCondition" yaml:"stoppingCondition"`
}

