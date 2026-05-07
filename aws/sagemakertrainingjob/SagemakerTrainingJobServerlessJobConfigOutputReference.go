// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakertrainingjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v24/jsii"

	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v24/sagemakertrainingjob/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SagemakerTrainingJobServerlessJobConfigOutputReference interface {
	cdktn.ComplexObject
	AcceptEula() interface{}
	SetAcceptEula(val interface{})
	AcceptEulaInput() interface{}
	BaseModelArn() *string
	SetBaseModelArn(val *string)
	BaseModelArnInput() *string
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
	CustomizationTechnique() *string
	SetCustomizationTechnique(val *string)
	CustomizationTechniqueInput() *string
	EvaluationType() *string
	SetEvaluationType(val *string)
	EvaluationTypeInput() *string
	EvaluatorArn() *string
	SetEvaluatorArn(val *string)
	EvaluatorArnInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	JobType() *string
	SetJobType(val *string)
	JobTypeInput() *string
	Peft() *string
	SetPeft(val *string)
	PeftInput() *string
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
	ResetAcceptEula()
	ResetCustomizationTechnique()
	ResetEvaluationType()
	ResetEvaluatorArn()
	ResetPeft()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SagemakerTrainingJobServerlessJobConfigOutputReference
type jsiiProxy_SagemakerTrainingJobServerlessJobConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_SagemakerTrainingJobServerlessJobConfigOutputReference) AcceptEula() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"acceptEula",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJobServerlessJobConfigOutputReference) AcceptEulaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"acceptEulaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJobServerlessJobConfigOutputReference) BaseModelArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseModelArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJobServerlessJobConfigOutputReference) BaseModelArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseModelArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJobServerlessJobConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJobServerlessJobConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJobServerlessJobConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJobServerlessJobConfigOutputReference) CustomizationTechnique() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customizationTechnique",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJobServerlessJobConfigOutputReference) CustomizationTechniqueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customizationTechniqueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJobServerlessJobConfigOutputReference) EvaluationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"evaluationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJobServerlessJobConfigOutputReference) EvaluationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"evaluationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJobServerlessJobConfigOutputReference) EvaluatorArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"evaluatorArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJobServerlessJobConfigOutputReference) EvaluatorArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"evaluatorArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJobServerlessJobConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJobServerlessJobConfigOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJobServerlessJobConfigOutputReference) JobType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJobServerlessJobConfigOutputReference) JobTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJobServerlessJobConfigOutputReference) Peft() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peft",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJobServerlessJobConfigOutputReference) PeftInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peftInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJobServerlessJobConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJobServerlessJobConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSagemakerTrainingJobServerlessJobConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) SagemakerTrainingJobServerlessJobConfigOutputReference {
	_init_.Initialize()

	if err := validateNewSagemakerTrainingJobServerlessJobConfigOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_SagemakerTrainingJobServerlessJobConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-aws.sagemakerTrainingJob.SagemakerTrainingJobServerlessJobConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewSagemakerTrainingJobServerlessJobConfigOutputReference_Override(s SagemakerTrainingJobServerlessJobConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.sagemakerTrainingJob.SagemakerTrainingJobServerlessJobConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_SagemakerTrainingJobServerlessJobConfigOutputReference)SetAcceptEula(val interface{}) {
	if err := j.validateSetAcceptEulaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"acceptEula",
		val,
	)
}

func (j *jsiiProxy_SagemakerTrainingJobServerlessJobConfigOutputReference)SetBaseModelArn(val *string) {
	if err := j.validateSetBaseModelArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"baseModelArn",
		val,
	)
}

func (j *jsiiProxy_SagemakerTrainingJobServerlessJobConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SagemakerTrainingJobServerlessJobConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SagemakerTrainingJobServerlessJobConfigOutputReference)SetCustomizationTechnique(val *string) {
	if err := j.validateSetCustomizationTechniqueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customizationTechnique",
		val,
	)
}

func (j *jsiiProxy_SagemakerTrainingJobServerlessJobConfigOutputReference)SetEvaluationType(val *string) {
	if err := j.validateSetEvaluationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"evaluationType",
		val,
	)
}

func (j *jsiiProxy_SagemakerTrainingJobServerlessJobConfigOutputReference)SetEvaluatorArn(val *string) {
	if err := j.validateSetEvaluatorArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"evaluatorArn",
		val,
	)
}

func (j *jsiiProxy_SagemakerTrainingJobServerlessJobConfigOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SagemakerTrainingJobServerlessJobConfigOutputReference)SetJobType(val *string) {
	if err := j.validateSetJobTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jobType",
		val,
	)
}

func (j *jsiiProxy_SagemakerTrainingJobServerlessJobConfigOutputReference)SetPeft(val *string) {
	if err := j.validateSetPeftParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"peft",
		val,
	)
}

func (j *jsiiProxy_SagemakerTrainingJobServerlessJobConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SagemakerTrainingJobServerlessJobConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SagemakerTrainingJobServerlessJobConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerTrainingJobServerlessJobConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SagemakerTrainingJobServerlessJobConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SagemakerTrainingJobServerlessJobConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SagemakerTrainingJobServerlessJobConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SagemakerTrainingJobServerlessJobConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SagemakerTrainingJobServerlessJobConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SagemakerTrainingJobServerlessJobConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SagemakerTrainingJobServerlessJobConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SagemakerTrainingJobServerlessJobConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SagemakerTrainingJobServerlessJobConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerTrainingJobServerlessJobConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SagemakerTrainingJobServerlessJobConfigOutputReference) ResetAcceptEula() {
	_jsii_.InvokeVoid(
		s,
		"resetAcceptEula",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerTrainingJobServerlessJobConfigOutputReference) ResetCustomizationTechnique() {
	_jsii_.InvokeVoid(
		s,
		"resetCustomizationTechnique",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerTrainingJobServerlessJobConfigOutputReference) ResetEvaluationType() {
	_jsii_.InvokeVoid(
		s,
		"resetEvaluationType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerTrainingJobServerlessJobConfigOutputReference) ResetEvaluatorArn() {
	_jsii_.InvokeVoid(
		s,
		"resetEvaluatorArn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerTrainingJobServerlessJobConfigOutputReference) ResetPeft() {
	_jsii_.InvokeVoid(
		s,
		"resetPeft",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerTrainingJobServerlessJobConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (s *jsiiProxy_SagemakerTrainingJobServerlessJobConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

