// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakertrainingjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v23/jsii"

	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v23/sagemakertrainingjob/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SagemakerTrainingJobAlgorithmSpecificationOutputReference interface {
	cdktn.ComplexObject
	AlgorithmName() *string
	SetAlgorithmName(val *string)
	AlgorithmNameInput() *string
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
	ContainerArguments() *[]*string
	SetContainerArguments(val *[]*string)
	ContainerArgumentsInput() *[]*string
	ContainerEntrypoint() *[]*string
	SetContainerEntrypoint(val *[]*string)
	ContainerEntrypointInput() *[]*string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	EnableSagemakerMetricsTimeSeries() interface{}
	SetEnableSagemakerMetricsTimeSeries(val interface{})
	EnableSagemakerMetricsTimeSeriesInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MetricDefinitions() SagemakerTrainingJobAlgorithmSpecificationMetricDefinitionsList
	MetricDefinitionsInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TrainingImage() *string
	SetTrainingImage(val *string)
	TrainingImageConfig() SagemakerTrainingJobAlgorithmSpecificationTrainingImageConfigList
	TrainingImageConfigInput() interface{}
	TrainingImageInput() *string
	TrainingInputMode() *string
	SetTrainingInputMode(val *string)
	TrainingInputModeInput() *string
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
	PutMetricDefinitions(value interface{})
	PutTrainingImageConfig(value interface{})
	ResetAlgorithmName()
	ResetContainerArguments()
	ResetContainerEntrypoint()
	ResetEnableSagemakerMetricsTimeSeries()
	ResetMetricDefinitions()
	ResetTrainingImage()
	ResetTrainingImageConfig()
	ResetTrainingInputMode()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SagemakerTrainingJobAlgorithmSpecificationOutputReference
type jsiiProxy_SagemakerTrainingJobAlgorithmSpecificationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_SagemakerTrainingJobAlgorithmSpecificationOutputReference) AlgorithmName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"algorithmName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJobAlgorithmSpecificationOutputReference) AlgorithmNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"algorithmNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJobAlgorithmSpecificationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJobAlgorithmSpecificationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJobAlgorithmSpecificationOutputReference) ContainerArguments() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"containerArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJobAlgorithmSpecificationOutputReference) ContainerArgumentsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"containerArgumentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJobAlgorithmSpecificationOutputReference) ContainerEntrypoint() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"containerEntrypoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJobAlgorithmSpecificationOutputReference) ContainerEntrypointInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"containerEntrypointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJobAlgorithmSpecificationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJobAlgorithmSpecificationOutputReference) EnableSagemakerMetricsTimeSeries() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableSagemakerMetricsTimeSeries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJobAlgorithmSpecificationOutputReference) EnableSagemakerMetricsTimeSeriesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableSagemakerMetricsTimeSeriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJobAlgorithmSpecificationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJobAlgorithmSpecificationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJobAlgorithmSpecificationOutputReference) MetricDefinitions() SagemakerTrainingJobAlgorithmSpecificationMetricDefinitionsList {
	var returns SagemakerTrainingJobAlgorithmSpecificationMetricDefinitionsList
	_jsii_.Get(
		j,
		"metricDefinitions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJobAlgorithmSpecificationOutputReference) MetricDefinitionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metricDefinitionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJobAlgorithmSpecificationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJobAlgorithmSpecificationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJobAlgorithmSpecificationOutputReference) TrainingImage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trainingImage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJobAlgorithmSpecificationOutputReference) TrainingImageConfig() SagemakerTrainingJobAlgorithmSpecificationTrainingImageConfigList {
	var returns SagemakerTrainingJobAlgorithmSpecificationTrainingImageConfigList
	_jsii_.Get(
		j,
		"trainingImageConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJobAlgorithmSpecificationOutputReference) TrainingImageConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"trainingImageConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJobAlgorithmSpecificationOutputReference) TrainingImageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trainingImageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJobAlgorithmSpecificationOutputReference) TrainingInputMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trainingInputMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerTrainingJobAlgorithmSpecificationOutputReference) TrainingInputModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trainingInputModeInput",
		&returns,
	)
	return returns
}


func NewSagemakerTrainingJobAlgorithmSpecificationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) SagemakerTrainingJobAlgorithmSpecificationOutputReference {
	_init_.Initialize()

	if err := validateNewSagemakerTrainingJobAlgorithmSpecificationOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_SagemakerTrainingJobAlgorithmSpecificationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-aws.sagemakerTrainingJob.SagemakerTrainingJobAlgorithmSpecificationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewSagemakerTrainingJobAlgorithmSpecificationOutputReference_Override(s SagemakerTrainingJobAlgorithmSpecificationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.sagemakerTrainingJob.SagemakerTrainingJobAlgorithmSpecificationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_SagemakerTrainingJobAlgorithmSpecificationOutputReference)SetAlgorithmName(val *string) {
	if err := j.validateSetAlgorithmNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"algorithmName",
		val,
	)
}

func (j *jsiiProxy_SagemakerTrainingJobAlgorithmSpecificationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SagemakerTrainingJobAlgorithmSpecificationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SagemakerTrainingJobAlgorithmSpecificationOutputReference)SetContainerArguments(val *[]*string) {
	if err := j.validateSetContainerArgumentsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerArguments",
		val,
	)
}

func (j *jsiiProxy_SagemakerTrainingJobAlgorithmSpecificationOutputReference)SetContainerEntrypoint(val *[]*string) {
	if err := j.validateSetContainerEntrypointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerEntrypoint",
		val,
	)
}

func (j *jsiiProxy_SagemakerTrainingJobAlgorithmSpecificationOutputReference)SetEnableSagemakerMetricsTimeSeries(val interface{}) {
	if err := j.validateSetEnableSagemakerMetricsTimeSeriesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableSagemakerMetricsTimeSeries",
		val,
	)
}

func (j *jsiiProxy_SagemakerTrainingJobAlgorithmSpecificationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SagemakerTrainingJobAlgorithmSpecificationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SagemakerTrainingJobAlgorithmSpecificationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_SagemakerTrainingJobAlgorithmSpecificationOutputReference)SetTrainingImage(val *string) {
	if err := j.validateSetTrainingImageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trainingImage",
		val,
	)
}

func (j *jsiiProxy_SagemakerTrainingJobAlgorithmSpecificationOutputReference)SetTrainingInputMode(val *string) {
	if err := j.validateSetTrainingInputModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trainingInputMode",
		val,
	)
}

func (s *jsiiProxy_SagemakerTrainingJobAlgorithmSpecificationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerTrainingJobAlgorithmSpecificationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SagemakerTrainingJobAlgorithmSpecificationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SagemakerTrainingJobAlgorithmSpecificationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SagemakerTrainingJobAlgorithmSpecificationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SagemakerTrainingJobAlgorithmSpecificationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SagemakerTrainingJobAlgorithmSpecificationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SagemakerTrainingJobAlgorithmSpecificationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SagemakerTrainingJobAlgorithmSpecificationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SagemakerTrainingJobAlgorithmSpecificationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SagemakerTrainingJobAlgorithmSpecificationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerTrainingJobAlgorithmSpecificationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SagemakerTrainingJobAlgorithmSpecificationOutputReference) PutMetricDefinitions(value interface{}) {
	if err := s.validatePutMetricDefinitionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putMetricDefinitions",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerTrainingJobAlgorithmSpecificationOutputReference) PutTrainingImageConfig(value interface{}) {
	if err := s.validatePutTrainingImageConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTrainingImageConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerTrainingJobAlgorithmSpecificationOutputReference) ResetAlgorithmName() {
	_jsii_.InvokeVoid(
		s,
		"resetAlgorithmName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerTrainingJobAlgorithmSpecificationOutputReference) ResetContainerArguments() {
	_jsii_.InvokeVoid(
		s,
		"resetContainerArguments",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerTrainingJobAlgorithmSpecificationOutputReference) ResetContainerEntrypoint() {
	_jsii_.InvokeVoid(
		s,
		"resetContainerEntrypoint",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerTrainingJobAlgorithmSpecificationOutputReference) ResetEnableSagemakerMetricsTimeSeries() {
	_jsii_.InvokeVoid(
		s,
		"resetEnableSagemakerMetricsTimeSeries",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerTrainingJobAlgorithmSpecificationOutputReference) ResetMetricDefinitions() {
	_jsii_.InvokeVoid(
		s,
		"resetMetricDefinitions",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerTrainingJobAlgorithmSpecificationOutputReference) ResetTrainingImage() {
	_jsii_.InvokeVoid(
		s,
		"resetTrainingImage",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerTrainingJobAlgorithmSpecificationOutputReference) ResetTrainingImageConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetTrainingImageConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerTrainingJobAlgorithmSpecificationOutputReference) ResetTrainingInputMode() {
	_jsii_.InvokeVoid(
		s,
		"resetTrainingInputMode",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerTrainingJobAlgorithmSpecificationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (s *jsiiProxy_SagemakerTrainingJobAlgorithmSpecificationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

