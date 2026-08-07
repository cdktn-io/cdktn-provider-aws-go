// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakeralgorithm

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/jsii"

	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/sagemakeralgorithm/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SagemakerAlgorithmValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference interface {
	cdktn.ComplexObject
	BatchStrategy() *string
	SetBatchStrategy(val *string)
	BatchStrategyInput() *string
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
	Environment() *map[string]*string
	SetEnvironment(val *map[string]*string)
	EnvironmentInput() *map[string]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MaxConcurrentTransforms() *float64
	SetMaxConcurrentTransforms(val *float64)
	MaxConcurrentTransformsInput() *float64
	MaxPayloadInMb() *float64
	SetMaxPayloadInMb(val *float64)
	MaxPayloadInMbInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TransformInput() SagemakerAlgorithmValidationSpecificationValidationProfilesTransformJobDefinitionTransformInputList
	TransformInputInput() interface{}
	TransformOutput() SagemakerAlgorithmValidationSpecificationValidationProfilesTransformJobDefinitionTransformOutputList
	TransformOutputInput() interface{}
	TransformResources() SagemakerAlgorithmValidationSpecificationValidationProfilesTransformJobDefinitionTransformResourcesList
	TransformResourcesInput() interface{}
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
	PutTransformInput(value interface{})
	PutTransformOutput(value interface{})
	PutTransformResources(value interface{})
	ResetBatchStrategy()
	ResetEnvironment()
	ResetMaxConcurrentTransforms()
	ResetMaxPayloadInMb()
	ResetTransformInput()
	ResetTransformOutput()
	ResetTransformResources()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SagemakerAlgorithmValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference
type jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference) BatchStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"batchStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference) BatchStrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"batchStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference) Environment() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference) EnvironmentInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"environmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference) MaxConcurrentTransforms() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrentTransforms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference) MaxConcurrentTransformsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrentTransformsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference) MaxPayloadInMb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxPayloadInMb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference) MaxPayloadInMbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxPayloadInMbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference) TransformInput() SagemakerAlgorithmValidationSpecificationValidationProfilesTransformJobDefinitionTransformInputList {
	var returns SagemakerAlgorithmValidationSpecificationValidationProfilesTransformJobDefinitionTransformInputList
	_jsii_.Get(
		j,
		"transformInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference) TransformInputInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"transformInputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference) TransformOutput() SagemakerAlgorithmValidationSpecificationValidationProfilesTransformJobDefinitionTransformOutputList {
	var returns SagemakerAlgorithmValidationSpecificationValidationProfilesTransformJobDefinitionTransformOutputList
	_jsii_.Get(
		j,
		"transformOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference) TransformOutputInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"transformOutputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference) TransformResources() SagemakerAlgorithmValidationSpecificationValidationProfilesTransformJobDefinitionTransformResourcesList {
	var returns SagemakerAlgorithmValidationSpecificationValidationProfilesTransformJobDefinitionTransformResourcesList
	_jsii_.Get(
		j,
		"transformResources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference) TransformResourcesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"transformResourcesInput",
		&returns,
	)
	return returns
}


func NewSagemakerAlgorithmValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) SagemakerAlgorithmValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference {
	_init_.Initialize()

	if err := validateNewSagemakerAlgorithmValidationSpecificationValidationProfilesTransformJobDefinitionOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-aws.sagemakerAlgorithm.SagemakerAlgorithmValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewSagemakerAlgorithmValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference_Override(s SagemakerAlgorithmValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.sagemakerAlgorithm.SagemakerAlgorithmValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference)SetBatchStrategy(val *string) {
	if err := j.validateSetBatchStrategyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"batchStrategy",
		val,
	)
}

func (j *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference)SetEnvironment(val *map[string]*string) {
	if err := j.validateSetEnvironmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"environment",
		val,
	)
}

func (j *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference)SetMaxConcurrentTransforms(val *float64) {
	if err := j.validateSetMaxConcurrentTransformsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxConcurrentTransforms",
		val,
	)
}

func (j *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference)SetMaxPayloadInMb(val *float64) {
	if err := j.validateSetMaxPayloadInMbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxPayloadInMb",
		val,
	)
}

func (j *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference) PutTransformInput(value interface{}) {
	if err := s.validatePutTransformInputParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTransformInput",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference) PutTransformOutput(value interface{}) {
	if err := s.validatePutTransformOutputParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTransformOutput",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference) PutTransformResources(value interface{}) {
	if err := s.validatePutTransformResourcesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTransformResources",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference) ResetBatchStrategy() {
	_jsii_.InvokeVoid(
		s,
		"resetBatchStrategy",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference) ResetEnvironment() {
	_jsii_.InvokeVoid(
		s,
		"resetEnvironment",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference) ResetMaxConcurrentTransforms() {
	_jsii_.InvokeVoid(
		s,
		"resetMaxConcurrentTransforms",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference) ResetMaxPayloadInMb() {
	_jsii_.InvokeVoid(
		s,
		"resetMaxPayloadInMb",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference) ResetTransformInput() {
	_jsii_.InvokeVoid(
		s,
		"resetTransformInput",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference) ResetTransformOutput() {
	_jsii_.InvokeVoid(
		s,
		"resetTransformOutput",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference) ResetTransformResources() {
	_jsii_.InvokeVoid(
		s,
		"resetTransformResources",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (s *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

