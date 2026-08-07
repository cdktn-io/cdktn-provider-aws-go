// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerhyperparametertuningjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/jsii"

	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/sagemakerhyperparametertuningjob/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference interface {
	cdktn.ComplexObject
	AlgorithmSpecification() SagemakerHyperParameterTuningJobTrainingJobDefinitionAlgorithmSpecificationList
	AlgorithmSpecificationInput() interface{}
	CheckpointConfig() SagemakerHyperParameterTuningJobTrainingJobDefinitionCheckpointConfigList
	CheckpointConfigInput() interface{}
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DefinitionName() *string
	SetDefinitionName(val *string)
	DefinitionNameInput() *string
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
	// Experimental.
	Fqn() *string
	HyperParameterRanges() SagemakerHyperParameterTuningJobTrainingJobDefinitionHyperParameterRangesList
	HyperParameterRangesInput() interface{}
	HyperParameterTuningResourceConfig() SagemakerHyperParameterTuningJobTrainingJobDefinitionHyperParameterTuningResourceConfigList
	HyperParameterTuningResourceConfigInput() interface{}
	InputDataConfig() SagemakerHyperParameterTuningJobTrainingJobDefinitionInputDataConfigList
	InputDataConfigInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	OutputDataConfig() SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputDataConfigList
	OutputDataConfigInput() interface{}
	ResourceConfig() SagemakerHyperParameterTuningJobTrainingJobDefinitionResourceConfigList
	ResourceConfigInput() interface{}
	RetryStrategy() SagemakerHyperParameterTuningJobTrainingJobDefinitionRetryStrategyList
	RetryStrategyInput() interface{}
	RoleArn() *string
	SetRoleArn(val *string)
	RoleArnInput() *string
	StaticHyperParameters() *map[string]*string
	SetStaticHyperParameters(val *map[string]*string)
	StaticHyperParametersInput() *map[string]*string
	StoppingCondition() SagemakerHyperParameterTuningJobTrainingJobDefinitionStoppingConditionList
	StoppingConditionInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TuningObjective() SagemakerHyperParameterTuningJobTrainingJobDefinitionTuningObjectiveList
	TuningObjectiveInput() interface{}
	VpcConfig() SagemakerHyperParameterTuningJobTrainingJobDefinitionVpcConfigList
	VpcConfigInput() interface{}
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktn.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	PutAlgorithmSpecification(value interface{})
	PutCheckpointConfig(value interface{})
	PutHyperParameterRanges(value interface{})
	PutHyperParameterTuningResourceConfig(value interface{})
	PutInputDataConfig(value interface{})
	PutOutputDataConfig(value interface{})
	PutResourceConfig(value interface{})
	PutRetryStrategy(value interface{})
	PutStoppingCondition(value interface{})
	PutTuningObjective(value interface{})
	PutVpcConfig(value interface{})
	ResetAlgorithmSpecification()
	ResetCheckpointConfig()
	ResetDefinitionName()
	ResetEnableInterContainerTrafficEncryption()
	ResetEnableManagedSpotTraining()
	ResetEnableNetworkIsolation()
	ResetEnvironment()
	ResetHyperParameterRanges()
	ResetHyperParameterTuningResourceConfig()
	ResetInputDataConfig()
	ResetOutputDataConfig()
	ResetResourceConfig()
	ResetRetryStrategy()
	ResetStaticHyperParameters()
	ResetStoppingCondition()
	ResetTuningObjective()
	ResetVpcConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference
type jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference) AlgorithmSpecification() SagemakerHyperParameterTuningJobTrainingJobDefinitionAlgorithmSpecificationList {
	var returns SagemakerHyperParameterTuningJobTrainingJobDefinitionAlgorithmSpecificationList
	_jsii_.Get(
		j,
		"algorithmSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference) AlgorithmSpecificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"algorithmSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference) CheckpointConfig() SagemakerHyperParameterTuningJobTrainingJobDefinitionCheckpointConfigList {
	var returns SagemakerHyperParameterTuningJobTrainingJobDefinitionCheckpointConfigList
	_jsii_.Get(
		j,
		"checkpointConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference) CheckpointConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"checkpointConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference) DefinitionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"definitionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference) DefinitionNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"definitionNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference) EnableInterContainerTrafficEncryption() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableInterContainerTrafficEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference) EnableInterContainerTrafficEncryptionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableInterContainerTrafficEncryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference) EnableManagedSpotTraining() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableManagedSpotTraining",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference) EnableManagedSpotTrainingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableManagedSpotTrainingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference) EnableNetworkIsolation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableNetworkIsolation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference) EnableNetworkIsolationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableNetworkIsolationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference) Environment() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference) EnvironmentInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"environmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference) HyperParameterRanges() SagemakerHyperParameterTuningJobTrainingJobDefinitionHyperParameterRangesList {
	var returns SagemakerHyperParameterTuningJobTrainingJobDefinitionHyperParameterRangesList
	_jsii_.Get(
		j,
		"hyperParameterRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference) HyperParameterRangesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hyperParameterRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference) HyperParameterTuningResourceConfig() SagemakerHyperParameterTuningJobTrainingJobDefinitionHyperParameterTuningResourceConfigList {
	var returns SagemakerHyperParameterTuningJobTrainingJobDefinitionHyperParameterTuningResourceConfigList
	_jsii_.Get(
		j,
		"hyperParameterTuningResourceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference) HyperParameterTuningResourceConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hyperParameterTuningResourceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference) InputDataConfig() SagemakerHyperParameterTuningJobTrainingJobDefinitionInputDataConfigList {
	var returns SagemakerHyperParameterTuningJobTrainingJobDefinitionInputDataConfigList
	_jsii_.Get(
		j,
		"inputDataConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference) InputDataConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inputDataConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference) OutputDataConfig() SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputDataConfigList {
	var returns SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputDataConfigList
	_jsii_.Get(
		j,
		"outputDataConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference) OutputDataConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"outputDataConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference) ResourceConfig() SagemakerHyperParameterTuningJobTrainingJobDefinitionResourceConfigList {
	var returns SagemakerHyperParameterTuningJobTrainingJobDefinitionResourceConfigList
	_jsii_.Get(
		j,
		"resourceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference) ResourceConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference) RetryStrategy() SagemakerHyperParameterTuningJobTrainingJobDefinitionRetryStrategyList {
	var returns SagemakerHyperParameterTuningJobTrainingJobDefinitionRetryStrategyList
	_jsii_.Get(
		j,
		"retryStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference) RetryStrategyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"retryStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference) StaticHyperParameters() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"staticHyperParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference) StaticHyperParametersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"staticHyperParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference) StoppingCondition() SagemakerHyperParameterTuningJobTrainingJobDefinitionStoppingConditionList {
	var returns SagemakerHyperParameterTuningJobTrainingJobDefinitionStoppingConditionList
	_jsii_.Get(
		j,
		"stoppingCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference) StoppingConditionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stoppingConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference) TuningObjective() SagemakerHyperParameterTuningJobTrainingJobDefinitionTuningObjectiveList {
	var returns SagemakerHyperParameterTuningJobTrainingJobDefinitionTuningObjectiveList
	_jsii_.Get(
		j,
		"tuningObjective",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference) TuningObjectiveInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tuningObjectiveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference) VpcConfig() SagemakerHyperParameterTuningJobTrainingJobDefinitionVpcConfigList {
	var returns SagemakerHyperParameterTuningJobTrainingJobDefinitionVpcConfigList
	_jsii_.Get(
		j,
		"vpcConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference) VpcConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vpcConfigInput",
		&returns,
	)
	return returns
}


func NewSagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference {
	_init_.Initialize()

	if err := validateNewSagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-aws.sagemakerHyperParameterTuningJob.SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewSagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference_Override(s SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.sagemakerHyperParameterTuningJob.SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference)SetDefinitionName(val *string) {
	if err := j.validateSetDefinitionNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"definitionName",
		val,
	)
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference)SetEnableInterContainerTrafficEncryption(val interface{}) {
	if err := j.validateSetEnableInterContainerTrafficEncryptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableInterContainerTrafficEncryption",
		val,
	)
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference)SetEnableManagedSpotTraining(val interface{}) {
	if err := j.validateSetEnableManagedSpotTrainingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableManagedSpotTraining",
		val,
	)
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference)SetEnableNetworkIsolation(val interface{}) {
	if err := j.validateSetEnableNetworkIsolationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableNetworkIsolation",
		val,
	)
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference)SetEnvironment(val *map[string]*string) {
	if err := j.validateSetEnvironmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"environment",
		val,
	)
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference)SetStaticHyperParameters(val *map[string]*string) {
	if err := j.validateSetStaticHyperParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"staticHyperParameters",
		val,
	)
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference) PutAlgorithmSpecification(value interface{}) {
	if err := s.validatePutAlgorithmSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAlgorithmSpecification",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference) PutCheckpointConfig(value interface{}) {
	if err := s.validatePutCheckpointConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putCheckpointConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference) PutHyperParameterRanges(value interface{}) {
	if err := s.validatePutHyperParameterRangesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putHyperParameterRanges",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference) PutHyperParameterTuningResourceConfig(value interface{}) {
	if err := s.validatePutHyperParameterTuningResourceConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putHyperParameterTuningResourceConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference) PutInputDataConfig(value interface{}) {
	if err := s.validatePutInputDataConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putInputDataConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference) PutOutputDataConfig(value interface{}) {
	if err := s.validatePutOutputDataConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putOutputDataConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference) PutResourceConfig(value interface{}) {
	if err := s.validatePutResourceConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putResourceConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference) PutRetryStrategy(value interface{}) {
	if err := s.validatePutRetryStrategyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putRetryStrategy",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference) PutStoppingCondition(value interface{}) {
	if err := s.validatePutStoppingConditionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putStoppingCondition",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference) PutTuningObjective(value interface{}) {
	if err := s.validatePutTuningObjectiveParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTuningObjective",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference) PutVpcConfig(value interface{}) {
	if err := s.validatePutVpcConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putVpcConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference) ResetAlgorithmSpecification() {
	_jsii_.InvokeVoid(
		s,
		"resetAlgorithmSpecification",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference) ResetCheckpointConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetCheckpointConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference) ResetDefinitionName() {
	_jsii_.InvokeVoid(
		s,
		"resetDefinitionName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference) ResetEnableInterContainerTrafficEncryption() {
	_jsii_.InvokeVoid(
		s,
		"resetEnableInterContainerTrafficEncryption",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference) ResetEnableManagedSpotTraining() {
	_jsii_.InvokeVoid(
		s,
		"resetEnableManagedSpotTraining",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference) ResetEnableNetworkIsolation() {
	_jsii_.InvokeVoid(
		s,
		"resetEnableNetworkIsolation",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference) ResetEnvironment() {
	_jsii_.InvokeVoid(
		s,
		"resetEnvironment",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference) ResetHyperParameterRanges() {
	_jsii_.InvokeVoid(
		s,
		"resetHyperParameterRanges",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference) ResetHyperParameterTuningResourceConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetHyperParameterTuningResourceConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference) ResetInputDataConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetInputDataConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference) ResetOutputDataConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetOutputDataConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference) ResetResourceConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetResourceConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference) ResetRetryStrategy() {
	_jsii_.InvokeVoid(
		s,
		"resetRetryStrategy",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference) ResetStaticHyperParameters() {
	_jsii_.InvokeVoid(
		s,
		"resetStaticHyperParameters",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference) ResetStoppingCondition() {
	_jsii_.InvokeVoid(
		s,
		"resetStoppingCondition",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference) ResetTuningObjective() {
	_jsii_.InvokeVoid(
		s,
		"resetTuningObjective",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference) ResetVpcConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetVpcConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := s.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

