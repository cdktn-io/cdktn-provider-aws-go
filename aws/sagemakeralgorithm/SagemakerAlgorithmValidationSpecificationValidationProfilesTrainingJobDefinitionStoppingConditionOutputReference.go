// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakeralgorithm

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v24/jsii"

	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v24/sagemakeralgorithm/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionStoppingConditionOutputReference interface {
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
	MaxPendingTimeInSeconds() *float64
	SetMaxPendingTimeInSeconds(val *float64)
	MaxPendingTimeInSecondsInput() *float64
	MaxRuntimeInSeconds() *float64
	SetMaxRuntimeInSeconds(val *float64)
	MaxRuntimeInSecondsInput() *float64
	MaxWaitTimeInSeconds() *float64
	SetMaxWaitTimeInSeconds(val *float64)
	MaxWaitTimeInSecondsInput() *float64
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
	ResetMaxPendingTimeInSeconds()
	ResetMaxRuntimeInSeconds()
	ResetMaxWaitTimeInSeconds()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionStoppingConditionOutputReference
type jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionStoppingConditionOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionStoppingConditionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionStoppingConditionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionStoppingConditionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionStoppingConditionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionStoppingConditionOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionStoppingConditionOutputReference) MaxPendingTimeInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxPendingTimeInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionStoppingConditionOutputReference) MaxPendingTimeInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxPendingTimeInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionStoppingConditionOutputReference) MaxRuntimeInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRuntimeInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionStoppingConditionOutputReference) MaxRuntimeInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRuntimeInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionStoppingConditionOutputReference) MaxWaitTimeInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxWaitTimeInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionStoppingConditionOutputReference) MaxWaitTimeInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxWaitTimeInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionStoppingConditionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionStoppingConditionOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionStoppingConditionOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionStoppingConditionOutputReference {
	_init_.Initialize()

	if err := validateNewSagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionStoppingConditionOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionStoppingConditionOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-aws.sagemakerAlgorithm.SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionStoppingConditionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewSagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionStoppingConditionOutputReference_Override(s SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionStoppingConditionOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.sagemakerAlgorithm.SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionStoppingConditionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionStoppingConditionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionStoppingConditionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionStoppingConditionOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionStoppingConditionOutputReference)SetMaxPendingTimeInSeconds(val *float64) {
	if err := j.validateSetMaxPendingTimeInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxPendingTimeInSeconds",
		val,
	)
}

func (j *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionStoppingConditionOutputReference)SetMaxRuntimeInSeconds(val *float64) {
	if err := j.validateSetMaxRuntimeInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxRuntimeInSeconds",
		val,
	)
}

func (j *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionStoppingConditionOutputReference)SetMaxWaitTimeInSeconds(val *float64) {
	if err := j.validateSetMaxWaitTimeInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxWaitTimeInSeconds",
		val,
	)
}

func (j *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionStoppingConditionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionStoppingConditionOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionStoppingConditionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionStoppingConditionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionStoppingConditionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionStoppingConditionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionStoppingConditionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionStoppingConditionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionStoppingConditionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionStoppingConditionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionStoppingConditionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionStoppingConditionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionStoppingConditionOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionStoppingConditionOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionStoppingConditionOutputReference) ResetMaxPendingTimeInSeconds() {
	_jsii_.InvokeVoid(
		s,
		"resetMaxPendingTimeInSeconds",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionStoppingConditionOutputReference) ResetMaxRuntimeInSeconds() {
	_jsii_.InvokeVoid(
		s,
		"resetMaxRuntimeInSeconds",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionStoppingConditionOutputReference) ResetMaxWaitTimeInSeconds() {
	_jsii_.InvokeVoid(
		s,
		"resetMaxWaitTimeInSeconds",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionStoppingConditionOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (s *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionStoppingConditionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

