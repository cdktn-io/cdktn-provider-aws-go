// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakeralgorithm

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/jsii"

	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/sagemakeralgorithm/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigOutputReference interface {
	cdktn.ComplexObject
	ChannelName() *string
	SetChannelName(val *string)
	ChannelNameInput() *string
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
	CompressionType() *string
	SetCompressionType(val *string)
	CompressionTypeInput() *string
	ContentType() *string
	SetContentType(val *string)
	ContentTypeInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DataSource() SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigDataSourceList
	DataSourceInput() interface{}
	// Experimental.
	Fqn() *string
	InputMode() *string
	SetInputMode(val *string)
	InputModeInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	RecordWrapperType() *string
	SetRecordWrapperType(val *string)
	RecordWrapperTypeInput() *string
	ShuffleConfig() SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigShuffleConfigList
	ShuffleConfigInput() interface{}
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
	PutDataSource(value interface{})
	PutShuffleConfig(value interface{})
	ResetCompressionType()
	ResetContentType()
	ResetDataSource()
	ResetInputMode()
	ResetRecordWrapperType()
	ResetShuffleConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigOutputReference
type jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigOutputReference) ChannelName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"channelName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigOutputReference) ChannelNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"channelNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigOutputReference) CompressionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compressionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigOutputReference) CompressionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compressionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigOutputReference) ContentType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigOutputReference) ContentTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigOutputReference) DataSource() SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigDataSourceList {
	var returns SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigDataSourceList
	_jsii_.Get(
		j,
		"dataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigOutputReference) DataSourceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigOutputReference) InputMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigOutputReference) InputModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigOutputReference) RecordWrapperType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recordWrapperType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigOutputReference) RecordWrapperTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recordWrapperTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigOutputReference) ShuffleConfig() SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigShuffleConfigList {
	var returns SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigShuffleConfigList
	_jsii_.Get(
		j,
		"shuffleConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigOutputReference) ShuffleConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shuffleConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigOutputReference {
	_init_.Initialize()

	if err := validateNewSagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-aws.sagemakerAlgorithm.SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewSagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigOutputReference_Override(s SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.sagemakerAlgorithm.SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigOutputReference)SetChannelName(val *string) {
	if err := j.validateSetChannelNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"channelName",
		val,
	)
}

func (j *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigOutputReference)SetCompressionType(val *string) {
	if err := j.validateSetCompressionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"compressionType",
		val,
	)
}

func (j *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigOutputReference)SetContentType(val *string) {
	if err := j.validateSetContentTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contentType",
		val,
	)
}

func (j *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigOutputReference)SetInputMode(val *string) {
	if err := j.validateSetInputModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputMode",
		val,
	)
}

func (j *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigOutputReference)SetRecordWrapperType(val *string) {
	if err := j.validateSetRecordWrapperTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recordWrapperType",
		val,
	)
}

func (j *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigOutputReference) PutDataSource(value interface{}) {
	if err := s.validatePutDataSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putDataSource",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigOutputReference) PutShuffleConfig(value interface{}) {
	if err := s.validatePutShuffleConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putShuffleConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigOutputReference) ResetCompressionType() {
	_jsii_.InvokeVoid(
		s,
		"resetCompressionType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigOutputReference) ResetContentType() {
	_jsii_.InvokeVoid(
		s,
		"resetContentType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigOutputReference) ResetDataSource() {
	_jsii_.InvokeVoid(
		s,
		"resetDataSource",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigOutputReference) ResetInputMode() {
	_jsii_.InvokeVoid(
		s,
		"resetInputMode",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigOutputReference) ResetRecordWrapperType() {
	_jsii_.InvokeVoid(
		s,
		"resetRecordWrapperType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigOutputReference) ResetShuffleConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetShuffleConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (s *jsiiProxy_SagemakerAlgorithmValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

