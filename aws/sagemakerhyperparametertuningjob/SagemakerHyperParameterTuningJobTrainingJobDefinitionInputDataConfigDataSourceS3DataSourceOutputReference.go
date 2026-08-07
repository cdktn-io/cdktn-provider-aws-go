// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerhyperparametertuningjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/jsii"

	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/sagemakerhyperparametertuningjob/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SagemakerHyperParameterTuningJobTrainingJobDefinitionInputDataConfigDataSourceS3DataSourceOutputReference interface {
	cdktn.ComplexObject
	AttributeNames() *[]*string
	SetAttributeNames(val *[]*string)
	AttributeNamesInput() *[]*string
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
	HubAccessConfig() SagemakerHyperParameterTuningJobTrainingJobDefinitionInputDataConfigDataSourceS3DataSourceHubAccessConfigList
	HubAccessConfigInput() interface{}
	InstanceGroupNames() *[]*string
	SetInstanceGroupNames(val *[]*string)
	InstanceGroupNamesInput() *[]*string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	ModelAccessConfig() SagemakerHyperParameterTuningJobTrainingJobDefinitionInputDataConfigDataSourceS3DataSourceModelAccessConfigList
	ModelAccessConfigInput() interface{}
	S3DataDistributionType() *string
	SetS3DataDistributionType(val *string)
	S3DataDistributionTypeInput() *string
	S3DataType() *string
	SetS3DataType(val *string)
	S3DataTypeInput() *string
	S3Uri() *string
	SetS3Uri(val *string)
	S3UriInput() *string
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
	PutHubAccessConfig(value interface{})
	PutModelAccessConfig(value interface{})
	ResetAttributeNames()
	ResetHubAccessConfig()
	ResetInstanceGroupNames()
	ResetModelAccessConfig()
	ResetS3DataDistributionType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SagemakerHyperParameterTuningJobTrainingJobDefinitionInputDataConfigDataSourceS3DataSourceOutputReference
type jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionInputDataConfigDataSourceS3DataSourceOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionInputDataConfigDataSourceS3DataSourceOutputReference) AttributeNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"attributeNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionInputDataConfigDataSourceS3DataSourceOutputReference) AttributeNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"attributeNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionInputDataConfigDataSourceS3DataSourceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionInputDataConfigDataSourceS3DataSourceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionInputDataConfigDataSourceS3DataSourceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionInputDataConfigDataSourceS3DataSourceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionInputDataConfigDataSourceS3DataSourceOutputReference) HubAccessConfig() SagemakerHyperParameterTuningJobTrainingJobDefinitionInputDataConfigDataSourceS3DataSourceHubAccessConfigList {
	var returns SagemakerHyperParameterTuningJobTrainingJobDefinitionInputDataConfigDataSourceS3DataSourceHubAccessConfigList
	_jsii_.Get(
		j,
		"hubAccessConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionInputDataConfigDataSourceS3DataSourceOutputReference) HubAccessConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hubAccessConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionInputDataConfigDataSourceS3DataSourceOutputReference) InstanceGroupNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"instanceGroupNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionInputDataConfigDataSourceS3DataSourceOutputReference) InstanceGroupNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"instanceGroupNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionInputDataConfigDataSourceS3DataSourceOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionInputDataConfigDataSourceS3DataSourceOutputReference) ModelAccessConfig() SagemakerHyperParameterTuningJobTrainingJobDefinitionInputDataConfigDataSourceS3DataSourceModelAccessConfigList {
	var returns SagemakerHyperParameterTuningJobTrainingJobDefinitionInputDataConfigDataSourceS3DataSourceModelAccessConfigList
	_jsii_.Get(
		j,
		"modelAccessConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionInputDataConfigDataSourceS3DataSourceOutputReference) ModelAccessConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"modelAccessConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionInputDataConfigDataSourceS3DataSourceOutputReference) S3DataDistributionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3DataDistributionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionInputDataConfigDataSourceS3DataSourceOutputReference) S3DataDistributionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3DataDistributionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionInputDataConfigDataSourceS3DataSourceOutputReference) S3DataType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3DataType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionInputDataConfigDataSourceS3DataSourceOutputReference) S3DataTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3DataTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionInputDataConfigDataSourceS3DataSourceOutputReference) S3Uri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3Uri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionInputDataConfigDataSourceS3DataSourceOutputReference) S3UriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3UriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionInputDataConfigDataSourceS3DataSourceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionInputDataConfigDataSourceS3DataSourceOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSagemakerHyperParameterTuningJobTrainingJobDefinitionInputDataConfigDataSourceS3DataSourceOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) SagemakerHyperParameterTuningJobTrainingJobDefinitionInputDataConfigDataSourceS3DataSourceOutputReference {
	_init_.Initialize()

	if err := validateNewSagemakerHyperParameterTuningJobTrainingJobDefinitionInputDataConfigDataSourceS3DataSourceOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionInputDataConfigDataSourceS3DataSourceOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-aws.sagemakerHyperParameterTuningJob.SagemakerHyperParameterTuningJobTrainingJobDefinitionInputDataConfigDataSourceS3DataSourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewSagemakerHyperParameterTuningJobTrainingJobDefinitionInputDataConfigDataSourceS3DataSourceOutputReference_Override(s SagemakerHyperParameterTuningJobTrainingJobDefinitionInputDataConfigDataSourceS3DataSourceOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.sagemakerHyperParameterTuningJob.SagemakerHyperParameterTuningJobTrainingJobDefinitionInputDataConfigDataSourceS3DataSourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionInputDataConfigDataSourceS3DataSourceOutputReference)SetAttributeNames(val *[]*string) {
	if err := j.validateSetAttributeNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"attributeNames",
		val,
	)
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionInputDataConfigDataSourceS3DataSourceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionInputDataConfigDataSourceS3DataSourceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionInputDataConfigDataSourceS3DataSourceOutputReference)SetInstanceGroupNames(val *[]*string) {
	if err := j.validateSetInstanceGroupNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceGroupNames",
		val,
	)
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionInputDataConfigDataSourceS3DataSourceOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionInputDataConfigDataSourceS3DataSourceOutputReference)SetS3DataDistributionType(val *string) {
	if err := j.validateSetS3DataDistributionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3DataDistributionType",
		val,
	)
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionInputDataConfigDataSourceS3DataSourceOutputReference)SetS3DataType(val *string) {
	if err := j.validateSetS3DataTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3DataType",
		val,
	)
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionInputDataConfigDataSourceS3DataSourceOutputReference)SetS3Uri(val *string) {
	if err := j.validateSetS3UriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3Uri",
		val,
	)
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionInputDataConfigDataSourceS3DataSourceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionInputDataConfigDataSourceS3DataSourceOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionInputDataConfigDataSourceS3DataSourceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionInputDataConfigDataSourceS3DataSourceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionInputDataConfigDataSourceS3DataSourceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionInputDataConfigDataSourceS3DataSourceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionInputDataConfigDataSourceS3DataSourceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionInputDataConfigDataSourceS3DataSourceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionInputDataConfigDataSourceS3DataSourceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionInputDataConfigDataSourceS3DataSourceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionInputDataConfigDataSourceS3DataSourceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionInputDataConfigDataSourceS3DataSourceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionInputDataConfigDataSourceS3DataSourceOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionInputDataConfigDataSourceS3DataSourceOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionInputDataConfigDataSourceS3DataSourceOutputReference) PutHubAccessConfig(value interface{}) {
	if err := s.validatePutHubAccessConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putHubAccessConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionInputDataConfigDataSourceS3DataSourceOutputReference) PutModelAccessConfig(value interface{}) {
	if err := s.validatePutModelAccessConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putModelAccessConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionInputDataConfigDataSourceS3DataSourceOutputReference) ResetAttributeNames() {
	_jsii_.InvokeVoid(
		s,
		"resetAttributeNames",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionInputDataConfigDataSourceS3DataSourceOutputReference) ResetHubAccessConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetHubAccessConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionInputDataConfigDataSourceS3DataSourceOutputReference) ResetInstanceGroupNames() {
	_jsii_.InvokeVoid(
		s,
		"resetInstanceGroupNames",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionInputDataConfigDataSourceS3DataSourceOutputReference) ResetModelAccessConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetModelAccessConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionInputDataConfigDataSourceS3DataSourceOutputReference) ResetS3DataDistributionType() {
	_jsii_.InvokeVoid(
		s,
		"resetS3DataDistributionType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionInputDataConfigDataSourceS3DataSourceOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (s *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionInputDataConfigDataSourceS3DataSourceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

