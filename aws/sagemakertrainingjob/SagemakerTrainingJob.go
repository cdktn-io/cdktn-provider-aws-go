// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakertrainingjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v23/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v23/sagemakertrainingjob/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.40.0/docs/resources/sagemaker_training_job aws_sagemaker_training_job}.
type SagemakerTrainingJob interface {
	cdktn.TerraformResource
	AlgorithmSpecification() SagemakerTrainingJobAlgorithmSpecificationList
	AlgorithmSpecificationInput() interface{}
	Arn() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	CheckpointConfig() SagemakerTrainingJobCheckpointConfigList
	CheckpointConfigInput() interface{}
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	DebugHookConfig() SagemakerTrainingJobDebugHookConfigList
	DebugHookConfigInput() interface{}
	DebugRuleConfigurations() SagemakerTrainingJobDebugRuleConfigurationsList
	DebugRuleConfigurationsInput() interface{}
	DeleteModelPackagesOnDestroy() interface{}
	SetDeleteModelPackagesOnDestroy(val interface{})
	DeleteModelPackagesOnDestroyInput() interface{}
	DeleteVpcEnisOnDestroy() interface{}
	SetDeleteVpcEnisOnDestroy(val interface{})
	DeleteVpcEnisOnDestroyInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EnableInterContainerTrafficEncryption() interface{}
	SetEnableInterContainerTrafficEncryption(val interface{})
	EnableInterContainerTrafficEncryptionInput() interface{}
	EnableManagedSpotTraining() interface{}
	SetEnableManagedSpotTraining(val interface{})
	EnableManagedSpotTrainingInput() interface{}
	EnableNetworkIsolation() interface{}
	SetEnableNetworkIsolation(val interface{})
	EnableNetworkIsolationInput() interface{}
	Environment() *map[string]*string
	SetEnvironment(val *map[string]*string)
	EnvironmentInput() *map[string]*string
	ExperimentConfig() SagemakerTrainingJobExperimentConfigList
	ExperimentConfigInput() interface{}
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HyperParameters() *map[string]*string
	SetHyperParameters(val *map[string]*string)
	HyperParametersInput() *map[string]*string
	InfraCheckConfig() SagemakerTrainingJobInfraCheckConfigList
	InfraCheckConfigInput() interface{}
	InputDataConfig() SagemakerTrainingJobInputDataConfigList
	InputDataConfigInput() interface{}
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	MlflowConfig() SagemakerTrainingJobMlflowConfigList
	MlflowConfigInput() interface{}
	ModelPackageConfig() SagemakerTrainingJobModelPackageConfigList
	ModelPackageConfigInput() interface{}
	// The tree node.
	Node() constructs.Node
	OutputDataConfig() SagemakerTrainingJobOutputDataConfigList
	OutputDataConfigInput() interface{}
	ProfilerConfig() SagemakerTrainingJobProfilerConfigList
	ProfilerConfigInput() interface{}
	ProfilerRuleConfigurations() SagemakerTrainingJobProfilerRuleConfigurationsList
	ProfilerRuleConfigurationsInput() interface{}
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	RemoteDebugConfig() SagemakerTrainingJobRemoteDebugConfigList
	RemoteDebugConfigInput() interface{}
	ResourceConfig() SagemakerTrainingJobResourceConfigList
	ResourceConfigInput() interface{}
	RetryStrategy() SagemakerTrainingJobRetryStrategyList
	RetryStrategyInput() interface{}
	RoleArn() *string
	SetRoleArn(val *string)
	RoleArnInput() *string
	ServerlessJobConfig() SagemakerTrainingJobServerlessJobConfigList
	ServerlessJobConfigInput() interface{}
	SessionChainingConfig() SagemakerTrainingJobSessionChainingConfigList
	SessionChainingConfigInput() interface{}
	StoppingCondition() SagemakerTrainingJobStoppingConditionList
	StoppingConditionInput() interface{}
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsAll() cdktn.StringMap
	TagsInput() *map[string]*string
	TensorBoardOutputConfig() SagemakerTrainingJobTensorBoardOutputConfigList
	TensorBoardOutputConfigInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() SagemakerTrainingJobTimeoutsOutputReference
	TimeoutsInput() interface{}
	TrainingJobName() *string
	SetTrainingJobName(val *string)
	TrainingJobNameInput() *string
	VpcConfig() SagemakerTrainingJobVpcConfigList
	VpcConfigInput() interface{}
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktn.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutAlgorithmSpecification(value interface{})
	PutCheckpointConfig(value interface{})
	PutDebugHookConfig(value interface{})
	PutDebugRuleConfigurations(value interface{})
	PutExperimentConfig(value interface{})
	PutInfraCheckConfig(value interface{})
	PutInputDataConfig(value interface{})
	PutMlflowConfig(value interface{})
	PutModelPackageConfig(value interface{})
	PutOutputDataConfig(value interface{})
	PutProfilerConfig(value interface{})
	PutProfilerRuleConfigurations(value interface{})
	PutRemoteDebugConfig(value interface{})
	PutResourceConfig(value interface{})
	PutRetryStrategy(value interface{})
	PutServerlessJobConfig(value interface{})
	PutSessionChainingConfig(value interface{})
	PutStoppingCondition(value interface{})
	PutTensorBoardOutputConfig(value interface{})
	PutTimeouts(value *SagemakerTrainingJobTimeouts)
	PutVpcConfig(value interface{})
	ResetAlgorithmSpecification()
	ResetCheckpointConfig()
	ResetDebugHookConfig()
	ResetDebugRuleConfigurations()
	ResetDeleteModelPackagesOnDestroy()
	ResetDeleteVpcEnisOnDestroy()
	ResetEnableInterContainerTrafficEncryption()
	ResetEnableManagedSpotTraining()
	ResetEnableNetworkIsolation()
	ResetEnvironment()
	ResetExperimentConfig()
	ResetHyperParameters()
	ResetInfraCheckConfig()
	ResetInputDataConfig()
	ResetMlflowConfig()
	ResetModelPackageConfig()
	ResetOutputDataConfig()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProfilerConfig()
	ResetProfilerRuleConfigurations()
	ResetRegion()
	ResetRemoteDebugConfig()
	ResetResourceConfig()
	ResetRetryStrategy()
	ResetServerlessJobConfig()
	ResetSessionChainingConfig()
	ResetStoppingCondition()
	ResetTags()
	ResetTensorBoardOutputConfig()
	ResetTimeouts()
	ResetVpcConfig()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
	// Applies one or more mixins to this construct.
	//
	// Mixins are applied in order. The list of constructs is captured at the
	// start of the call, so constructs added by a mixin will not be visited.
	// Use multiple `with()` calls if subsequent mixins should apply to added
	// constructs.
	//
	// Returns: This construct for chaining.
	With(mixins ...constructs.IMixin) constructs.IConstruct
}

// The jsii proxy struct for SagemakerTrainingJob
type jsiiProxy_SagemakerTrainingJob struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_SagemakerTrainingJob) AlgorithmSpecification() SagemakerTrainingJobAlgorithmSpecificationList {
	var returns SagemakerTrainingJobAlgorithmSpecificationList
	_jsii_.Get(
		j,
		"algorithmSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJob) AlgorithmSpecificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"algorithmSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJob) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJob) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJob) CheckpointConfig() SagemakerTrainingJobCheckpointConfigList {
	var returns SagemakerTrainingJobCheckpointConfigList
	_jsii_.Get(
		j,
		"checkpointConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJob) CheckpointConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"checkpointConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJob) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJob) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJob) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJob) DebugHookConfig() SagemakerTrainingJobDebugHookConfigList {
	var returns SagemakerTrainingJobDebugHookConfigList
	_jsii_.Get(
		j,
		"debugHookConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJob) DebugHookConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"debugHookConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJob) DebugRuleConfigurations() SagemakerTrainingJobDebugRuleConfigurationsList {
	var returns SagemakerTrainingJobDebugRuleConfigurationsList
	_jsii_.Get(
		j,
		"debugRuleConfigurations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJob) DebugRuleConfigurationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"debugRuleConfigurationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJob) DeleteModelPackagesOnDestroy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteModelPackagesOnDestroy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJob) DeleteModelPackagesOnDestroyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteModelPackagesOnDestroyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJob) DeleteVpcEnisOnDestroy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteVpcEnisOnDestroy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJob) DeleteVpcEnisOnDestroyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteVpcEnisOnDestroyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJob) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJob) EnableInterContainerTrafficEncryption() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableInterContainerTrafficEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJob) EnableInterContainerTrafficEncryptionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableInterContainerTrafficEncryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJob) EnableManagedSpotTraining() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableManagedSpotTraining",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJob) EnableManagedSpotTrainingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableManagedSpotTrainingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJob) EnableNetworkIsolation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableNetworkIsolation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJob) EnableNetworkIsolationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableNetworkIsolationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJob) Environment() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJob) EnvironmentInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"environmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJob) ExperimentConfig() SagemakerTrainingJobExperimentConfigList {
	var returns SagemakerTrainingJobExperimentConfigList
	_jsii_.Get(
		j,
		"experimentConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJob) ExperimentConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"experimentConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJob) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJob) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJob) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJob) HyperParameters() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"hyperParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJob) HyperParametersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"hyperParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJob) InfraCheckConfig() SagemakerTrainingJobInfraCheckConfigList {
	var returns SagemakerTrainingJobInfraCheckConfigList
	_jsii_.Get(
		j,
		"infraCheckConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJob) InfraCheckConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"infraCheckConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJob) InputDataConfig() SagemakerTrainingJobInputDataConfigList {
	var returns SagemakerTrainingJobInputDataConfigList
	_jsii_.Get(
		j,
		"inputDataConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJob) InputDataConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inputDataConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJob) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJob) MlflowConfig() SagemakerTrainingJobMlflowConfigList {
	var returns SagemakerTrainingJobMlflowConfigList
	_jsii_.Get(
		j,
		"mlflowConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJob) MlflowConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mlflowConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJob) ModelPackageConfig() SagemakerTrainingJobModelPackageConfigList {
	var returns SagemakerTrainingJobModelPackageConfigList
	_jsii_.Get(
		j,
		"modelPackageConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJob) ModelPackageConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"modelPackageConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJob) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJob) OutputDataConfig() SagemakerTrainingJobOutputDataConfigList {
	var returns SagemakerTrainingJobOutputDataConfigList
	_jsii_.Get(
		j,
		"outputDataConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJob) OutputDataConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"outputDataConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJob) ProfilerConfig() SagemakerTrainingJobProfilerConfigList {
	var returns SagemakerTrainingJobProfilerConfigList
	_jsii_.Get(
		j,
		"profilerConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJob) ProfilerConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"profilerConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJob) ProfilerRuleConfigurations() SagemakerTrainingJobProfilerRuleConfigurationsList {
	var returns SagemakerTrainingJobProfilerRuleConfigurationsList
	_jsii_.Get(
		j,
		"profilerRuleConfigurations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJob) ProfilerRuleConfigurationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"profilerRuleConfigurationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJob) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJob) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJob) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJob) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJob) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJob) RemoteDebugConfig() SagemakerTrainingJobRemoteDebugConfigList {
	var returns SagemakerTrainingJobRemoteDebugConfigList
	_jsii_.Get(
		j,
		"remoteDebugConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJob) RemoteDebugConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"remoteDebugConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJob) ResourceConfig() SagemakerTrainingJobResourceConfigList {
	var returns SagemakerTrainingJobResourceConfigList
	_jsii_.Get(
		j,
		"resourceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJob) ResourceConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJob) RetryStrategy() SagemakerTrainingJobRetryStrategyList {
	var returns SagemakerTrainingJobRetryStrategyList
	_jsii_.Get(
		j,
		"retryStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJob) RetryStrategyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"retryStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJob) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJob) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJob) ServerlessJobConfig() SagemakerTrainingJobServerlessJobConfigList {
	var returns SagemakerTrainingJobServerlessJobConfigList
	_jsii_.Get(
		j,
		"serverlessJobConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJob) ServerlessJobConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serverlessJobConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJob) SessionChainingConfig() SagemakerTrainingJobSessionChainingConfigList {
	var returns SagemakerTrainingJobSessionChainingConfigList
	_jsii_.Get(
		j,
		"sessionChainingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJob) SessionChainingConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sessionChainingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJob) StoppingCondition() SagemakerTrainingJobStoppingConditionList {
	var returns SagemakerTrainingJobStoppingConditionList
	_jsii_.Get(
		j,
		"stoppingCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJob) StoppingConditionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stoppingConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJob) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJob) TagsAll() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJob) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJob) TensorBoardOutputConfig() SagemakerTrainingJobTensorBoardOutputConfigList {
	var returns SagemakerTrainingJobTensorBoardOutputConfigList
	_jsii_.Get(
		j,
		"tensorBoardOutputConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJob) TensorBoardOutputConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tensorBoardOutputConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJob) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJob) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJob) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJob) Timeouts() SagemakerTrainingJobTimeoutsOutputReference {
	var returns SagemakerTrainingJobTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJob) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJob) TrainingJobName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trainingJobName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJob) TrainingJobNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trainingJobNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJob) VpcConfig() SagemakerTrainingJobVpcConfigList {
	var returns SagemakerTrainingJobVpcConfigList
	_jsii_.Get(
		j,
		"vpcConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJob) VpcConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vpcConfigInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.40.0/docs/resources/sagemaker_training_job aws_sagemaker_training_job} Resource.
func NewSagemakerTrainingJob(scope constructs.Construct, id *string, config *SagemakerTrainingJobConfig) SagemakerTrainingJob {
	_init_.Initialize()

	if err := validateNewSagemakerTrainingJobParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_SagemakerTrainingJob{}

	_jsii_.Create(
		"@cdktn/provider-aws.sagemakerTrainingJob.SagemakerTrainingJob",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.40.0/docs/resources/sagemaker_training_job aws_sagemaker_training_job} Resource.
func NewSagemakerTrainingJob_Override(s SagemakerTrainingJob, scope constructs.Construct, id *string, config *SagemakerTrainingJobConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.sagemakerTrainingJob.SagemakerTrainingJob",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SagemakerTrainingJob)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_SagemakerTrainingJob)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SagemakerTrainingJob)SetDeleteModelPackagesOnDestroy(val interface{}) {
	if err := j.validateSetDeleteModelPackagesOnDestroyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteModelPackagesOnDestroy",
		val,
	)
}

func (j *jsiiProxy_SagemakerTrainingJob)SetDeleteVpcEnisOnDestroy(val interface{}) {
	if err := j.validateSetDeleteVpcEnisOnDestroyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteVpcEnisOnDestroy",
		val,
	)
}

func (j *jsiiProxy_SagemakerTrainingJob)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SagemakerTrainingJob)SetEnableInterContainerTrafficEncryption(val interface{}) {
	if err := j.validateSetEnableInterContainerTrafficEncryptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableInterContainerTrafficEncryption",
		val,
	)
}

func (j *jsiiProxy_SagemakerTrainingJob)SetEnableManagedSpotTraining(val interface{}) {
	if err := j.validateSetEnableManagedSpotTrainingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableManagedSpotTraining",
		val,
	)
}

func (j *jsiiProxy_SagemakerTrainingJob)SetEnableNetworkIsolation(val interface{}) {
	if err := j.validateSetEnableNetworkIsolationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableNetworkIsolation",
		val,
	)
}

func (j *jsiiProxy_SagemakerTrainingJob)SetEnvironment(val *map[string]*string) {
	if err := j.validateSetEnvironmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"environment",
		val,
	)
}

func (j *jsiiProxy_SagemakerTrainingJob)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_SagemakerTrainingJob)SetHyperParameters(val *map[string]*string) {
	if err := j.validateSetHyperParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hyperParameters",
		val,
	)
}

func (j *jsiiProxy_SagemakerTrainingJob)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SagemakerTrainingJob)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SagemakerTrainingJob)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_SagemakerTrainingJob)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_SagemakerTrainingJob)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_SagemakerTrainingJob)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_SagemakerTrainingJob)SetTrainingJobName(val *string) {
	if err := j.validateSetTrainingJobNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trainingJobName",
		val,
	)
}

// Generates CDKTN code for importing a SagemakerTrainingJob resource upon running "cdktn plan <stack-name>".
func SagemakerTrainingJob_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateSagemakerTrainingJob_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-aws.sagemakerTrainingJob.SagemakerTrainingJob",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func SagemakerTrainingJob_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSagemakerTrainingJob_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-aws.sagemakerTrainingJob.SagemakerTrainingJob",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SagemakerTrainingJob_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSagemakerTrainingJob_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-aws.sagemakerTrainingJob.SagemakerTrainingJob",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SagemakerTrainingJob_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSagemakerTrainingJob_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-aws.sagemakerTrainingJob.SagemakerTrainingJob",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SagemakerTrainingJob_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-aws.sagemakerTrainingJob.SagemakerTrainingJob",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_SagemakerTrainingJob) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_SagemakerTrainingJob) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_SagemakerTrainingJob) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerTrainingJob) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerTrainingJob) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerTrainingJob) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerTrainingJob) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerTrainingJob) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerTrainingJob) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerTrainingJob) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerTrainingJob) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerTrainingJob) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerTrainingJob) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_SagemakerTrainingJob) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerTrainingJob) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SagemakerTrainingJob) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_SagemakerTrainingJob) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SagemakerTrainingJob) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SagemakerTrainingJob) PutAlgorithmSpecification(value interface{}) {
	if err := s.validatePutAlgorithmSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAlgorithmSpecification",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerTrainingJob) PutCheckpointConfig(value interface{}) {
	if err := s.validatePutCheckpointConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putCheckpointConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerTrainingJob) PutDebugHookConfig(value interface{}) {
	if err := s.validatePutDebugHookConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putDebugHookConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerTrainingJob) PutDebugRuleConfigurations(value interface{}) {
	if err := s.validatePutDebugRuleConfigurationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putDebugRuleConfigurations",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerTrainingJob) PutExperimentConfig(value interface{}) {
	if err := s.validatePutExperimentConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putExperimentConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerTrainingJob) PutInfraCheckConfig(value interface{}) {
	if err := s.validatePutInfraCheckConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putInfraCheckConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerTrainingJob) PutInputDataConfig(value interface{}) {
	if err := s.validatePutInputDataConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putInputDataConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerTrainingJob) PutMlflowConfig(value interface{}) {
	if err := s.validatePutMlflowConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putMlflowConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerTrainingJob) PutModelPackageConfig(value interface{}) {
	if err := s.validatePutModelPackageConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putModelPackageConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerTrainingJob) PutOutputDataConfig(value interface{}) {
	if err := s.validatePutOutputDataConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putOutputDataConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerTrainingJob) PutProfilerConfig(value interface{}) {
	if err := s.validatePutProfilerConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putProfilerConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerTrainingJob) PutProfilerRuleConfigurations(value interface{}) {
	if err := s.validatePutProfilerRuleConfigurationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putProfilerRuleConfigurations",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerTrainingJob) PutRemoteDebugConfig(value interface{}) {
	if err := s.validatePutRemoteDebugConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putRemoteDebugConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerTrainingJob) PutResourceConfig(value interface{}) {
	if err := s.validatePutResourceConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putResourceConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerTrainingJob) PutRetryStrategy(value interface{}) {
	if err := s.validatePutRetryStrategyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putRetryStrategy",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerTrainingJob) PutServerlessJobConfig(value interface{}) {
	if err := s.validatePutServerlessJobConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putServerlessJobConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerTrainingJob) PutSessionChainingConfig(value interface{}) {
	if err := s.validatePutSessionChainingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putSessionChainingConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerTrainingJob) PutStoppingCondition(value interface{}) {
	if err := s.validatePutStoppingConditionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putStoppingCondition",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerTrainingJob) PutTensorBoardOutputConfig(value interface{}) {
	if err := s.validatePutTensorBoardOutputConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTensorBoardOutputConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerTrainingJob) PutTimeouts(value *SagemakerTrainingJobTimeouts) {
	if err := s.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerTrainingJob) PutVpcConfig(value interface{}) {
	if err := s.validatePutVpcConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putVpcConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerTrainingJob) ResetAlgorithmSpecification() {
	_jsii_.InvokeVoid(
		s,
		"resetAlgorithmSpecification",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerTrainingJob) ResetCheckpointConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetCheckpointConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerTrainingJob) ResetDebugHookConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetDebugHookConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerTrainingJob) ResetDebugRuleConfigurations() {
	_jsii_.InvokeVoid(
		s,
		"resetDebugRuleConfigurations",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerTrainingJob) ResetDeleteModelPackagesOnDestroy() {
	_jsii_.InvokeVoid(
		s,
		"resetDeleteModelPackagesOnDestroy",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerTrainingJob) ResetDeleteVpcEnisOnDestroy() {
	_jsii_.InvokeVoid(
		s,
		"resetDeleteVpcEnisOnDestroy",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerTrainingJob) ResetEnableInterContainerTrafficEncryption() {
	_jsii_.InvokeVoid(
		s,
		"resetEnableInterContainerTrafficEncryption",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerTrainingJob) ResetEnableManagedSpotTraining() {
	_jsii_.InvokeVoid(
		s,
		"resetEnableManagedSpotTraining",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerTrainingJob) ResetEnableNetworkIsolation() {
	_jsii_.InvokeVoid(
		s,
		"resetEnableNetworkIsolation",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerTrainingJob) ResetEnvironment() {
	_jsii_.InvokeVoid(
		s,
		"resetEnvironment",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerTrainingJob) ResetExperimentConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetExperimentConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerTrainingJob) ResetHyperParameters() {
	_jsii_.InvokeVoid(
		s,
		"resetHyperParameters",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerTrainingJob) ResetInfraCheckConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetInfraCheckConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerTrainingJob) ResetInputDataConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetInputDataConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerTrainingJob) ResetMlflowConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetMlflowConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerTrainingJob) ResetModelPackageConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetModelPackageConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerTrainingJob) ResetOutputDataConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetOutputDataConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerTrainingJob) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerTrainingJob) ResetProfilerConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetProfilerConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerTrainingJob) ResetProfilerRuleConfigurations() {
	_jsii_.InvokeVoid(
		s,
		"resetProfilerRuleConfigurations",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerTrainingJob) ResetRegion() {
	_jsii_.InvokeVoid(
		s,
		"resetRegion",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerTrainingJob) ResetRemoteDebugConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetRemoteDebugConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerTrainingJob) ResetResourceConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetResourceConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerTrainingJob) ResetRetryStrategy() {
	_jsii_.InvokeVoid(
		s,
		"resetRetryStrategy",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerTrainingJob) ResetServerlessJobConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetServerlessJobConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerTrainingJob) ResetSessionChainingConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetSessionChainingConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerTrainingJob) ResetStoppingCondition() {
	_jsii_.InvokeVoid(
		s,
		"resetStoppingCondition",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerTrainingJob) ResetTags() {
	_jsii_.InvokeVoid(
		s,
		"resetTags",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerTrainingJob) ResetTensorBoardOutputConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetTensorBoardOutputConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerTrainingJob) ResetTimeouts() {
	_jsii_.InvokeVoid(
		s,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerTrainingJob) ResetVpcConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetVpcConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerTrainingJob) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerTrainingJob) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerTrainingJob) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerTrainingJob) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerTrainingJob) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerTrainingJob) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerTrainingJob) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		s,
		"with",
		args,
		&returns,
	)

	return returns
}

