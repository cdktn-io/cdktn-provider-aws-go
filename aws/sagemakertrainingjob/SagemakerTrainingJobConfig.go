// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakertrainingjob

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SagemakerTrainingJobConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.44.0/docs/resources/sagemaker_training_job#role_arn SagemakerTrainingJob#role_arn}.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.44.0/docs/resources/sagemaker_training_job#training_job_name SagemakerTrainingJob#training_job_name}.
	TrainingJobName *string `field:"required" json:"trainingJobName" yaml:"trainingJobName"`
	// algorithm_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.44.0/docs/resources/sagemaker_training_job#algorithm_specification SagemakerTrainingJob#algorithm_specification}
	AlgorithmSpecification interface{} `field:"optional" json:"algorithmSpecification" yaml:"algorithmSpecification"`
	// checkpoint_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.44.0/docs/resources/sagemaker_training_job#checkpoint_config SagemakerTrainingJob#checkpoint_config}
	CheckpointConfig interface{} `field:"optional" json:"checkpointConfig" yaml:"checkpointConfig"`
	// debug_hook_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.44.0/docs/resources/sagemaker_training_job#debug_hook_config SagemakerTrainingJob#debug_hook_config}
	DebugHookConfig interface{} `field:"optional" json:"debugHookConfig" yaml:"debugHookConfig"`
	// debug_rule_configurations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.44.0/docs/resources/sagemaker_training_job#debug_rule_configurations SagemakerTrainingJob#debug_rule_configurations}
	DebugRuleConfigurations interface{} `field:"optional" json:"debugRuleConfigurations" yaml:"debugRuleConfigurations"`
	// Whether to delete model packages in the configured model package group when destroying the training job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.44.0/docs/resources/sagemaker_training_job#delete_model_packages_on_destroy SagemakerTrainingJob#delete_model_packages_on_destroy}
	DeleteModelPackagesOnDestroy interface{} `field:"optional" json:"deleteModelPackagesOnDestroy" yaml:"deleteModelPackagesOnDestroy"`
	// Whether to delete detached VPC ENIs that SageMaker may leave behind when destroying the training job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.44.0/docs/resources/sagemaker_training_job#delete_vpc_enis_on_destroy SagemakerTrainingJob#delete_vpc_enis_on_destroy}
	DeleteVpcEnisOnDestroy interface{} `field:"optional" json:"deleteVpcEnisOnDestroy" yaml:"deleteVpcEnisOnDestroy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.44.0/docs/resources/sagemaker_training_job#enable_inter_container_traffic_encryption SagemakerTrainingJob#enable_inter_container_traffic_encryption}.
	EnableInterContainerTrafficEncryption interface{} `field:"optional" json:"enableInterContainerTrafficEncryption" yaml:"enableInterContainerTrafficEncryption"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.44.0/docs/resources/sagemaker_training_job#enable_managed_spot_training SagemakerTrainingJob#enable_managed_spot_training}.
	EnableManagedSpotTraining interface{} `field:"optional" json:"enableManagedSpotTraining" yaml:"enableManagedSpotTraining"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.44.0/docs/resources/sagemaker_training_job#enable_network_isolation SagemakerTrainingJob#enable_network_isolation}.
	EnableNetworkIsolation interface{} `field:"optional" json:"enableNetworkIsolation" yaml:"enableNetworkIsolation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.44.0/docs/resources/sagemaker_training_job#environment SagemakerTrainingJob#environment}.
	Environment *map[string]*string `field:"optional" json:"environment" yaml:"environment"`
	// experiment_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.44.0/docs/resources/sagemaker_training_job#experiment_config SagemakerTrainingJob#experiment_config}
	ExperimentConfig interface{} `field:"optional" json:"experimentConfig" yaml:"experimentConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.44.0/docs/resources/sagemaker_training_job#hyper_parameters SagemakerTrainingJob#hyper_parameters}.
	HyperParameters *map[string]*string `field:"optional" json:"hyperParameters" yaml:"hyperParameters"`
	// infra_check_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.44.0/docs/resources/sagemaker_training_job#infra_check_config SagemakerTrainingJob#infra_check_config}
	InfraCheckConfig interface{} `field:"optional" json:"infraCheckConfig" yaml:"infraCheckConfig"`
	// input_data_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.44.0/docs/resources/sagemaker_training_job#input_data_config SagemakerTrainingJob#input_data_config}
	InputDataConfig interface{} `field:"optional" json:"inputDataConfig" yaml:"inputDataConfig"`
	// mlflow_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.44.0/docs/resources/sagemaker_training_job#mlflow_config SagemakerTrainingJob#mlflow_config}
	MlflowConfig interface{} `field:"optional" json:"mlflowConfig" yaml:"mlflowConfig"`
	// model_package_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.44.0/docs/resources/sagemaker_training_job#model_package_config SagemakerTrainingJob#model_package_config}
	ModelPackageConfig interface{} `field:"optional" json:"modelPackageConfig" yaml:"modelPackageConfig"`
	// output_data_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.44.0/docs/resources/sagemaker_training_job#output_data_config SagemakerTrainingJob#output_data_config}
	OutputDataConfig interface{} `field:"optional" json:"outputDataConfig" yaml:"outputDataConfig"`
	// profiler_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.44.0/docs/resources/sagemaker_training_job#profiler_config SagemakerTrainingJob#profiler_config}
	ProfilerConfig interface{} `field:"optional" json:"profilerConfig" yaml:"profilerConfig"`
	// profiler_rule_configurations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.44.0/docs/resources/sagemaker_training_job#profiler_rule_configurations SagemakerTrainingJob#profiler_rule_configurations}
	ProfilerRuleConfigurations interface{} `field:"optional" json:"profilerRuleConfigurations" yaml:"profilerRuleConfigurations"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.44.0/docs/resources/sagemaker_training_job#region SagemakerTrainingJob#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// remote_debug_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.44.0/docs/resources/sagemaker_training_job#remote_debug_config SagemakerTrainingJob#remote_debug_config}
	RemoteDebugConfig interface{} `field:"optional" json:"remoteDebugConfig" yaml:"remoteDebugConfig"`
	// resource_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.44.0/docs/resources/sagemaker_training_job#resource_config SagemakerTrainingJob#resource_config}
	ResourceConfig interface{} `field:"optional" json:"resourceConfig" yaml:"resourceConfig"`
	// retry_strategy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.44.0/docs/resources/sagemaker_training_job#retry_strategy SagemakerTrainingJob#retry_strategy}
	RetryStrategy interface{} `field:"optional" json:"retryStrategy" yaml:"retryStrategy"`
	// serverless_job_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.44.0/docs/resources/sagemaker_training_job#serverless_job_config SagemakerTrainingJob#serverless_job_config}
	ServerlessJobConfig interface{} `field:"optional" json:"serverlessJobConfig" yaml:"serverlessJobConfig"`
	// session_chaining_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.44.0/docs/resources/sagemaker_training_job#session_chaining_config SagemakerTrainingJob#session_chaining_config}
	SessionChainingConfig interface{} `field:"optional" json:"sessionChainingConfig" yaml:"sessionChainingConfig"`
	// stopping_condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.44.0/docs/resources/sagemaker_training_job#stopping_condition SagemakerTrainingJob#stopping_condition}
	StoppingCondition interface{} `field:"optional" json:"stoppingCondition" yaml:"stoppingCondition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.44.0/docs/resources/sagemaker_training_job#tags SagemakerTrainingJob#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// tensor_board_output_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.44.0/docs/resources/sagemaker_training_job#tensor_board_output_config SagemakerTrainingJob#tensor_board_output_config}
	TensorBoardOutputConfig interface{} `field:"optional" json:"tensorBoardOutputConfig" yaml:"tensorBoardOutputConfig"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.44.0/docs/resources/sagemaker_training_job#timeouts SagemakerTrainingJob#timeouts}
	Timeouts *SagemakerTrainingJobTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// vpc_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.44.0/docs/resources/sagemaker_training_job#vpc_config SagemakerTrainingJob#vpc_config}
	VpcConfig interface{} `field:"optional" json:"vpcConfig" yaml:"vpcConfig"`
}

