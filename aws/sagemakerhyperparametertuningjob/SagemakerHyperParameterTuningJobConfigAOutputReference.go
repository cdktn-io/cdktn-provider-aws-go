// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerhyperparametertuningjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/jsii"

	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/sagemakerhyperparametertuningjob/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SagemakerHyperParameterTuningJobConfigAOutputReference interface {
	cdktn.ComplexObject
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
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Objective() SagemakerHyperParameterTuningJobConfigObjectiveList
	ObjectiveInput() interface{}
	ParameterRanges() SagemakerHyperParameterTuningJobConfigParameterRangesList
	ParameterRangesInput() interface{}
	RandomSeed() *float64
	SetRandomSeed(val *float64)
	RandomSeedInput() *float64
	ResourceLimits() SagemakerHyperParameterTuningJobConfigResourceLimitsList
	ResourceLimitsInput() interface{}
	Strategy() *string
	SetStrategy(val *string)
	StrategyConfig() SagemakerHyperParameterTuningJobConfigStrategyConfigList
	StrategyConfigInput() interface{}
	StrategyInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TrainingJobEarlyStoppingType() *string
	SetTrainingJobEarlyStoppingType(val *string)
	TrainingJobEarlyStoppingTypeInput() *string
	TuningJobCompletionCriteria() SagemakerHyperParameterTuningJobConfigTuningJobCompletionCriteriaList
	TuningJobCompletionCriteriaInput() interface{}
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
	PutObjective(value interface{})
	PutParameterRanges(value interface{})
	PutResourceLimits(value interface{})
	PutStrategyConfig(value interface{})
	PutTuningJobCompletionCriteria(value interface{})
	ResetObjective()
	ResetParameterRanges()
	ResetRandomSeed()
	ResetResourceLimits()
	ResetStrategyConfig()
	ResetTrainingJobEarlyStoppingType()
	ResetTuningJobCompletionCriteria()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SagemakerHyperParameterTuningJobConfigAOutputReference
type jsiiProxy_SagemakerHyperParameterTuningJobConfigAOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobConfigAOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobConfigAOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobConfigAOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobConfigAOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobConfigAOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobConfigAOutputReference) Objective() SagemakerHyperParameterTuningJobConfigObjectiveList {
	var returns SagemakerHyperParameterTuningJobConfigObjectiveList
	_jsii_.Get(
		j,
		"objective",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobConfigAOutputReference) ObjectiveInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"objectiveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobConfigAOutputReference) ParameterRanges() SagemakerHyperParameterTuningJobConfigParameterRangesList {
	var returns SagemakerHyperParameterTuningJobConfigParameterRangesList
	_jsii_.Get(
		j,
		"parameterRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobConfigAOutputReference) ParameterRangesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parameterRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobConfigAOutputReference) RandomSeed() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"randomSeed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobConfigAOutputReference) RandomSeedInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"randomSeedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobConfigAOutputReference) ResourceLimits() SagemakerHyperParameterTuningJobConfigResourceLimitsList {
	var returns SagemakerHyperParameterTuningJobConfigResourceLimitsList
	_jsii_.Get(
		j,
		"resourceLimits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobConfigAOutputReference) ResourceLimitsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceLimitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobConfigAOutputReference) Strategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobConfigAOutputReference) StrategyConfig() SagemakerHyperParameterTuningJobConfigStrategyConfigList {
	var returns SagemakerHyperParameterTuningJobConfigStrategyConfigList
	_jsii_.Get(
		j,
		"strategyConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobConfigAOutputReference) StrategyConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"strategyConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobConfigAOutputReference) StrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"strategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobConfigAOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobConfigAOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobConfigAOutputReference) TrainingJobEarlyStoppingType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trainingJobEarlyStoppingType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobConfigAOutputReference) TrainingJobEarlyStoppingTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trainingJobEarlyStoppingTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobConfigAOutputReference) TuningJobCompletionCriteria() SagemakerHyperParameterTuningJobConfigTuningJobCompletionCriteriaList {
	var returns SagemakerHyperParameterTuningJobConfigTuningJobCompletionCriteriaList
	_jsii_.Get(
		j,
		"tuningJobCompletionCriteria",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobConfigAOutputReference) TuningJobCompletionCriteriaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tuningJobCompletionCriteriaInput",
		&returns,
	)
	return returns
}


func NewSagemakerHyperParameterTuningJobConfigAOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) SagemakerHyperParameterTuningJobConfigAOutputReference {
	_init_.Initialize()

	if err := validateNewSagemakerHyperParameterTuningJobConfigAOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_SagemakerHyperParameterTuningJobConfigAOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-aws.sagemakerHyperParameterTuningJob.SagemakerHyperParameterTuningJobConfigAOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewSagemakerHyperParameterTuningJobConfigAOutputReference_Override(s SagemakerHyperParameterTuningJobConfigAOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.sagemakerHyperParameterTuningJob.SagemakerHyperParameterTuningJobConfigAOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobConfigAOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobConfigAOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobConfigAOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobConfigAOutputReference)SetRandomSeed(val *float64) {
	if err := j.validateSetRandomSeedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"randomSeed",
		val,
	)
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobConfigAOutputReference)SetStrategy(val *string) {
	if err := j.validateSetStrategyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"strategy",
		val,
	)
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobConfigAOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobConfigAOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobConfigAOutputReference)SetTrainingJobEarlyStoppingType(val *string) {
	if err := j.validateSetTrainingJobEarlyStoppingTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trainingJobEarlyStoppingType",
		val,
	)
}

func (s *jsiiProxy_SagemakerHyperParameterTuningJobConfigAOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerHyperParameterTuningJobConfigAOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SagemakerHyperParameterTuningJobConfigAOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SagemakerHyperParameterTuningJobConfigAOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SagemakerHyperParameterTuningJobConfigAOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SagemakerHyperParameterTuningJobConfigAOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SagemakerHyperParameterTuningJobConfigAOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SagemakerHyperParameterTuningJobConfigAOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SagemakerHyperParameterTuningJobConfigAOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SagemakerHyperParameterTuningJobConfigAOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SagemakerHyperParameterTuningJobConfigAOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerHyperParameterTuningJobConfigAOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SagemakerHyperParameterTuningJobConfigAOutputReference) PutObjective(value interface{}) {
	if err := s.validatePutObjectiveParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putObjective",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerHyperParameterTuningJobConfigAOutputReference) PutParameterRanges(value interface{}) {
	if err := s.validatePutParameterRangesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putParameterRanges",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerHyperParameterTuningJobConfigAOutputReference) PutResourceLimits(value interface{}) {
	if err := s.validatePutResourceLimitsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putResourceLimits",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerHyperParameterTuningJobConfigAOutputReference) PutStrategyConfig(value interface{}) {
	if err := s.validatePutStrategyConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putStrategyConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerHyperParameterTuningJobConfigAOutputReference) PutTuningJobCompletionCriteria(value interface{}) {
	if err := s.validatePutTuningJobCompletionCriteriaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTuningJobCompletionCriteria",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerHyperParameterTuningJobConfigAOutputReference) ResetObjective() {
	_jsii_.InvokeVoid(
		s,
		"resetObjective",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerHyperParameterTuningJobConfigAOutputReference) ResetParameterRanges() {
	_jsii_.InvokeVoid(
		s,
		"resetParameterRanges",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerHyperParameterTuningJobConfigAOutputReference) ResetRandomSeed() {
	_jsii_.InvokeVoid(
		s,
		"resetRandomSeed",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerHyperParameterTuningJobConfigAOutputReference) ResetResourceLimits() {
	_jsii_.InvokeVoid(
		s,
		"resetResourceLimits",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerHyperParameterTuningJobConfigAOutputReference) ResetStrategyConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetStrategyConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerHyperParameterTuningJobConfigAOutputReference) ResetTrainingJobEarlyStoppingType() {
	_jsii_.InvokeVoid(
		s,
		"resetTrainingJobEarlyStoppingType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerHyperParameterTuningJobConfigAOutputReference) ResetTuningJobCompletionCriteria() {
	_jsii_.InvokeVoid(
		s,
		"resetTuningJobCompletionCriteria",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerHyperParameterTuningJobConfigAOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (s *jsiiProxy_SagemakerHyperParameterTuningJobConfigAOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

