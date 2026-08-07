// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerhyperparametertuningjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/jsii"

	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/sagemakerhyperparametertuningjob/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SagemakerHyperParameterTuningJobConfigTuningJobCompletionCriteriaOutputReference interface {
	cdktn.ComplexObject
	BestObjectiveNotImproving() SagemakerHyperParameterTuningJobConfigTuningJobCompletionCriteriaBestObjectiveNotImprovingList
	BestObjectiveNotImprovingInput() interface{}
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
	ConvergenceDetected() SagemakerHyperParameterTuningJobConfigTuningJobCompletionCriteriaConvergenceDetectedList
	ConvergenceDetectedInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	TargetObjectiveMetricValue() *float64
	SetTargetObjectiveMetricValue(val *float64)
	TargetObjectiveMetricValueInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
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
	PutBestObjectiveNotImproving(value interface{})
	PutConvergenceDetected(value interface{})
	ResetBestObjectiveNotImproving()
	ResetConvergenceDetected()
	ResetTargetObjectiveMetricValue()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SagemakerHyperParameterTuningJobConfigTuningJobCompletionCriteriaOutputReference
type jsiiProxy_SagemakerHyperParameterTuningJobConfigTuningJobCompletionCriteriaOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobConfigTuningJobCompletionCriteriaOutputReference) BestObjectiveNotImproving() SagemakerHyperParameterTuningJobConfigTuningJobCompletionCriteriaBestObjectiveNotImprovingList {
	var returns SagemakerHyperParameterTuningJobConfigTuningJobCompletionCriteriaBestObjectiveNotImprovingList
	_jsii_.Get(
		j,
		"bestObjectiveNotImproving",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobConfigTuningJobCompletionCriteriaOutputReference) BestObjectiveNotImprovingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bestObjectiveNotImprovingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobConfigTuningJobCompletionCriteriaOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobConfigTuningJobCompletionCriteriaOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobConfigTuningJobCompletionCriteriaOutputReference) ConvergenceDetected() SagemakerHyperParameterTuningJobConfigTuningJobCompletionCriteriaConvergenceDetectedList {
	var returns SagemakerHyperParameterTuningJobConfigTuningJobCompletionCriteriaConvergenceDetectedList
	_jsii_.Get(
		j,
		"convergenceDetected",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobConfigTuningJobCompletionCriteriaOutputReference) ConvergenceDetectedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"convergenceDetectedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobConfigTuningJobCompletionCriteriaOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobConfigTuningJobCompletionCriteriaOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobConfigTuningJobCompletionCriteriaOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobConfigTuningJobCompletionCriteriaOutputReference) TargetObjectiveMetricValue() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetObjectiveMetricValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobConfigTuningJobCompletionCriteriaOutputReference) TargetObjectiveMetricValueInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetObjectiveMetricValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobConfigTuningJobCompletionCriteriaOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobConfigTuningJobCompletionCriteriaOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSagemakerHyperParameterTuningJobConfigTuningJobCompletionCriteriaOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) SagemakerHyperParameterTuningJobConfigTuningJobCompletionCriteriaOutputReference {
	_init_.Initialize()

	if err := validateNewSagemakerHyperParameterTuningJobConfigTuningJobCompletionCriteriaOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_SagemakerHyperParameterTuningJobConfigTuningJobCompletionCriteriaOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-aws.sagemakerHyperParameterTuningJob.SagemakerHyperParameterTuningJobConfigTuningJobCompletionCriteriaOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewSagemakerHyperParameterTuningJobConfigTuningJobCompletionCriteriaOutputReference_Override(s SagemakerHyperParameterTuningJobConfigTuningJobCompletionCriteriaOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.sagemakerHyperParameterTuningJob.SagemakerHyperParameterTuningJobConfigTuningJobCompletionCriteriaOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobConfigTuningJobCompletionCriteriaOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobConfigTuningJobCompletionCriteriaOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobConfigTuningJobCompletionCriteriaOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobConfigTuningJobCompletionCriteriaOutputReference)SetTargetObjectiveMetricValue(val *float64) {
	if err := j.validateSetTargetObjectiveMetricValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetObjectiveMetricValue",
		val,
	)
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobConfigTuningJobCompletionCriteriaOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobConfigTuningJobCompletionCriteriaOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SagemakerHyperParameterTuningJobConfigTuningJobCompletionCriteriaOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerHyperParameterTuningJobConfigTuningJobCompletionCriteriaOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SagemakerHyperParameterTuningJobConfigTuningJobCompletionCriteriaOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SagemakerHyperParameterTuningJobConfigTuningJobCompletionCriteriaOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SagemakerHyperParameterTuningJobConfigTuningJobCompletionCriteriaOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SagemakerHyperParameterTuningJobConfigTuningJobCompletionCriteriaOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SagemakerHyperParameterTuningJobConfigTuningJobCompletionCriteriaOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SagemakerHyperParameterTuningJobConfigTuningJobCompletionCriteriaOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SagemakerHyperParameterTuningJobConfigTuningJobCompletionCriteriaOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SagemakerHyperParameterTuningJobConfigTuningJobCompletionCriteriaOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SagemakerHyperParameterTuningJobConfigTuningJobCompletionCriteriaOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerHyperParameterTuningJobConfigTuningJobCompletionCriteriaOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SagemakerHyperParameterTuningJobConfigTuningJobCompletionCriteriaOutputReference) PutBestObjectiveNotImproving(value interface{}) {
	if err := s.validatePutBestObjectiveNotImprovingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putBestObjectiveNotImproving",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerHyperParameterTuningJobConfigTuningJobCompletionCriteriaOutputReference) PutConvergenceDetected(value interface{}) {
	if err := s.validatePutConvergenceDetectedParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putConvergenceDetected",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerHyperParameterTuningJobConfigTuningJobCompletionCriteriaOutputReference) ResetBestObjectiveNotImproving() {
	_jsii_.InvokeVoid(
		s,
		"resetBestObjectiveNotImproving",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerHyperParameterTuningJobConfigTuningJobCompletionCriteriaOutputReference) ResetConvergenceDetected() {
	_jsii_.InvokeVoid(
		s,
		"resetConvergenceDetected",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerHyperParameterTuningJobConfigTuningJobCompletionCriteriaOutputReference) ResetTargetObjectiveMetricValue() {
	_jsii_.InvokeVoid(
		s,
		"resetTargetObjectiveMetricValue",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerHyperParameterTuningJobConfigTuningJobCompletionCriteriaOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (s *jsiiProxy_SagemakerHyperParameterTuningJobConfigTuningJobCompletionCriteriaOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

