// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakertrainingjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/jsii"

	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/sagemakertrainingjob/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SagemakerTrainingJobInputDataConfigOutputReference interface {
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
	DataSource() SagemakerTrainingJobInputDataConfigDataSourceList
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
	ShuffleConfig() SagemakerTrainingJobInputDataConfigShuffleConfigList
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

// The jsii proxy struct for SagemakerTrainingJobInputDataConfigOutputReference
type jsiiProxy_SagemakerTrainingJobInputDataConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_SagemakerTrainingJobInputDataConfigOutputReference) ChannelName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"channelName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJobInputDataConfigOutputReference) ChannelNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"channelNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJobInputDataConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJobInputDataConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJobInputDataConfigOutputReference) CompressionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compressionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJobInputDataConfigOutputReference) CompressionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compressionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJobInputDataConfigOutputReference) ContentType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJobInputDataConfigOutputReference) ContentTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJobInputDataConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJobInputDataConfigOutputReference) DataSource() SagemakerTrainingJobInputDataConfigDataSourceList {
	var returns SagemakerTrainingJobInputDataConfigDataSourceList
	_jsii_.Get(
		j,
		"dataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJobInputDataConfigOutputReference) DataSourceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJobInputDataConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJobInputDataConfigOutputReference) InputMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJobInputDataConfigOutputReference) InputModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJobInputDataConfigOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJobInputDataConfigOutputReference) RecordWrapperType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recordWrapperType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJobInputDataConfigOutputReference) RecordWrapperTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recordWrapperTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJobInputDataConfigOutputReference) ShuffleConfig() SagemakerTrainingJobInputDataConfigShuffleConfigList {
	var returns SagemakerTrainingJobInputDataConfigShuffleConfigList
	_jsii_.Get(
		j,
		"shuffleConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJobInputDataConfigOutputReference) ShuffleConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shuffleConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJobInputDataConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJobInputDataConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSagemakerTrainingJobInputDataConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) SagemakerTrainingJobInputDataConfigOutputReference {
	_init_.Initialize()

	if err := validateNewSagemakerTrainingJobInputDataConfigOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_SagemakerTrainingJobInputDataConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-aws.sagemakerTrainingJob.SagemakerTrainingJobInputDataConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewSagemakerTrainingJobInputDataConfigOutputReference_Override(s SagemakerTrainingJobInputDataConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.sagemakerTrainingJob.SagemakerTrainingJobInputDataConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_SagemakerTrainingJobInputDataConfigOutputReference)SetChannelName(val *string) {
	if err := j.validateSetChannelNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"channelName",
		val,
	)
}

func (j *jsiiProxy_SagemakerTrainingJobInputDataConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SagemakerTrainingJobInputDataConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SagemakerTrainingJobInputDataConfigOutputReference)SetCompressionType(val *string) {
	if err := j.validateSetCompressionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"compressionType",
		val,
	)
}

func (j *jsiiProxy_SagemakerTrainingJobInputDataConfigOutputReference)SetContentType(val *string) {
	if err := j.validateSetContentTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contentType",
		val,
	)
}

func (j *jsiiProxy_SagemakerTrainingJobInputDataConfigOutputReference)SetInputMode(val *string) {
	if err := j.validateSetInputModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputMode",
		val,
	)
}

func (j *jsiiProxy_SagemakerTrainingJobInputDataConfigOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SagemakerTrainingJobInputDataConfigOutputReference)SetRecordWrapperType(val *string) {
	if err := j.validateSetRecordWrapperTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recordWrapperType",
		val,
	)
}

func (j *jsiiProxy_SagemakerTrainingJobInputDataConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SagemakerTrainingJobInputDataConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SagemakerTrainingJobInputDataConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerTrainingJobInputDataConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SagemakerTrainingJobInputDataConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SagemakerTrainingJobInputDataConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SagemakerTrainingJobInputDataConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SagemakerTrainingJobInputDataConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SagemakerTrainingJobInputDataConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SagemakerTrainingJobInputDataConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SagemakerTrainingJobInputDataConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SagemakerTrainingJobInputDataConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SagemakerTrainingJobInputDataConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerTrainingJobInputDataConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SagemakerTrainingJobInputDataConfigOutputReference) PutDataSource(value interface{}) {
	if err := s.validatePutDataSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putDataSource",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerTrainingJobInputDataConfigOutputReference) PutShuffleConfig(value interface{}) {
	if err := s.validatePutShuffleConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putShuffleConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerTrainingJobInputDataConfigOutputReference) ResetCompressionType() {
	_jsii_.InvokeVoid(
		s,
		"resetCompressionType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerTrainingJobInputDataConfigOutputReference) ResetContentType() {
	_jsii_.InvokeVoid(
		s,
		"resetContentType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerTrainingJobInputDataConfigOutputReference) ResetDataSource() {
	_jsii_.InvokeVoid(
		s,
		"resetDataSource",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerTrainingJobInputDataConfigOutputReference) ResetInputMode() {
	_jsii_.InvokeVoid(
		s,
		"resetInputMode",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerTrainingJobInputDataConfigOutputReference) ResetRecordWrapperType() {
	_jsii_.InvokeVoid(
		s,
		"resetRecordWrapperType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerTrainingJobInputDataConfigOutputReference) ResetShuffleConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetShuffleConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerTrainingJobInputDataConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (s *jsiiProxy_SagemakerTrainingJobInputDataConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

