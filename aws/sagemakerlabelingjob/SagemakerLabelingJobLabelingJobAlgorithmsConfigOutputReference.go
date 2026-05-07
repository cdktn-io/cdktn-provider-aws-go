// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerlabelingjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v24/jsii"

	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v24/sagemakerlabelingjob/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SagemakerLabelingJobLabelingJobAlgorithmsConfigOutputReference interface {
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
	InitialActiveLearningModelArn() *string
	SetInitialActiveLearningModelArn(val *string)
	InitialActiveLearningModelArnInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LabelingJobAlgorithmSpecificationArn() *string
	SetLabelingJobAlgorithmSpecificationArn(val *string)
	LabelingJobAlgorithmSpecificationArnInput() *string
	LabelingJobResourceConfig() SagemakerLabelingJobLabelingJobAlgorithmsConfigLabelingJobResourceConfigList
	LabelingJobResourceConfigInput() interface{}
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
	PutLabelingJobResourceConfig(value interface{})
	ResetInitialActiveLearningModelArn()
	ResetLabelingJobResourceConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SagemakerLabelingJobLabelingJobAlgorithmsConfigOutputReference
type jsiiProxy_SagemakerLabelingJobLabelingJobAlgorithmsConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_SagemakerLabelingJobLabelingJobAlgorithmsConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerLabelingJobLabelingJobAlgorithmsConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerLabelingJobLabelingJobAlgorithmsConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerLabelingJobLabelingJobAlgorithmsConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerLabelingJobLabelingJobAlgorithmsConfigOutputReference) InitialActiveLearningModelArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"initialActiveLearningModelArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerLabelingJobLabelingJobAlgorithmsConfigOutputReference) InitialActiveLearningModelArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"initialActiveLearningModelArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerLabelingJobLabelingJobAlgorithmsConfigOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerLabelingJobLabelingJobAlgorithmsConfigOutputReference) LabelingJobAlgorithmSpecificationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"labelingJobAlgorithmSpecificationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerLabelingJobLabelingJobAlgorithmsConfigOutputReference) LabelingJobAlgorithmSpecificationArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"labelingJobAlgorithmSpecificationArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerLabelingJobLabelingJobAlgorithmsConfigOutputReference) LabelingJobResourceConfig() SagemakerLabelingJobLabelingJobAlgorithmsConfigLabelingJobResourceConfigList {
	var returns SagemakerLabelingJobLabelingJobAlgorithmsConfigLabelingJobResourceConfigList
	_jsii_.Get(
		j,
		"labelingJobResourceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerLabelingJobLabelingJobAlgorithmsConfigOutputReference) LabelingJobResourceConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"labelingJobResourceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerLabelingJobLabelingJobAlgorithmsConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerLabelingJobLabelingJobAlgorithmsConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSagemakerLabelingJobLabelingJobAlgorithmsConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) SagemakerLabelingJobLabelingJobAlgorithmsConfigOutputReference {
	_init_.Initialize()

	if err := validateNewSagemakerLabelingJobLabelingJobAlgorithmsConfigOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_SagemakerLabelingJobLabelingJobAlgorithmsConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-aws.sagemakerLabelingJob.SagemakerLabelingJobLabelingJobAlgorithmsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewSagemakerLabelingJobLabelingJobAlgorithmsConfigOutputReference_Override(s SagemakerLabelingJobLabelingJobAlgorithmsConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.sagemakerLabelingJob.SagemakerLabelingJobLabelingJobAlgorithmsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_SagemakerLabelingJobLabelingJobAlgorithmsConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SagemakerLabelingJobLabelingJobAlgorithmsConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SagemakerLabelingJobLabelingJobAlgorithmsConfigOutputReference)SetInitialActiveLearningModelArn(val *string) {
	if err := j.validateSetInitialActiveLearningModelArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"initialActiveLearningModelArn",
		val,
	)
}

func (j *jsiiProxy_SagemakerLabelingJobLabelingJobAlgorithmsConfigOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SagemakerLabelingJobLabelingJobAlgorithmsConfigOutputReference)SetLabelingJobAlgorithmSpecificationArn(val *string) {
	if err := j.validateSetLabelingJobAlgorithmSpecificationArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labelingJobAlgorithmSpecificationArn",
		val,
	)
}

func (j *jsiiProxy_SagemakerLabelingJobLabelingJobAlgorithmsConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SagemakerLabelingJobLabelingJobAlgorithmsConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SagemakerLabelingJobLabelingJobAlgorithmsConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerLabelingJobLabelingJobAlgorithmsConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SagemakerLabelingJobLabelingJobAlgorithmsConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SagemakerLabelingJobLabelingJobAlgorithmsConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SagemakerLabelingJobLabelingJobAlgorithmsConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SagemakerLabelingJobLabelingJobAlgorithmsConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SagemakerLabelingJobLabelingJobAlgorithmsConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SagemakerLabelingJobLabelingJobAlgorithmsConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SagemakerLabelingJobLabelingJobAlgorithmsConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SagemakerLabelingJobLabelingJobAlgorithmsConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SagemakerLabelingJobLabelingJobAlgorithmsConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerLabelingJobLabelingJobAlgorithmsConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SagemakerLabelingJobLabelingJobAlgorithmsConfigOutputReference) PutLabelingJobResourceConfig(value interface{}) {
	if err := s.validatePutLabelingJobResourceConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putLabelingJobResourceConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerLabelingJobLabelingJobAlgorithmsConfigOutputReference) ResetInitialActiveLearningModelArn() {
	_jsii_.InvokeVoid(
		s,
		"resetInitialActiveLearningModelArn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerLabelingJobLabelingJobAlgorithmsConfigOutputReference) ResetLabelingJobResourceConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetLabelingJobResourceConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerLabelingJobLabelingJobAlgorithmsConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (s *jsiiProxy_SagemakerLabelingJobLabelingJobAlgorithmsConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

