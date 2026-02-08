// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerlabelingjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v22/jsii"

	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v22/sagemakerlabelingjob/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SagemakerLabelingJobHumanTaskConfigOutputReference interface {
	cdktf.ComplexObject
	AnnotationConsolidationConfig() SagemakerLabelingJobHumanTaskConfigAnnotationConsolidationConfigList
	AnnotationConsolidationConfigInput() interface{}
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
	MaxConcurrentTaskCount() *float64
	SetMaxConcurrentTaskCount(val *float64)
	MaxConcurrentTaskCountInput() *float64
	NumberOfHumanWorkersPerDataObject() *float64
	SetNumberOfHumanWorkersPerDataObject(val *float64)
	NumberOfHumanWorkersPerDataObjectInput() *float64
	PreHumanTaskLambdaArn() *string
	SetPreHumanTaskLambdaArn(val *string)
	PreHumanTaskLambdaArnInput() *string
	PublicWorkforceTaskPrice() SagemakerLabelingJobHumanTaskConfigPublicWorkforceTaskPriceList
	PublicWorkforceTaskPriceInput() interface{}
	TaskAvailabilityLifetimeInSeconds() *float64
	SetTaskAvailabilityLifetimeInSeconds(val *float64)
	TaskAvailabilityLifetimeInSecondsInput() *float64
	TaskDescription() *string
	SetTaskDescription(val *string)
	TaskDescriptionInput() *string
	TaskKeywords() *[]*string
	SetTaskKeywords(val *[]*string)
	TaskKeywordsInput() *[]*string
	TaskTimeLimitInSeconds() *float64
	SetTaskTimeLimitInSeconds(val *float64)
	TaskTimeLimitInSecondsInput() *float64
	TaskTitle() *string
	SetTaskTitle(val *string)
	TaskTitleInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UiConfig() SagemakerLabelingJobHumanTaskConfigUiConfigList
	UiConfigInput() interface{}
	WorkteamArn() *string
	SetWorkteamArn(val *string)
	WorkteamArnInput() *string
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	PutAnnotationConsolidationConfig(value interface{})
	PutPublicWorkforceTaskPrice(value interface{})
	PutUiConfig(value interface{})
	ResetAnnotationConsolidationConfig()
	ResetMaxConcurrentTaskCount()
	ResetPreHumanTaskLambdaArn()
	ResetPublicWorkforceTaskPrice()
	ResetTaskAvailabilityLifetimeInSeconds()
	ResetTaskKeywords()
	ResetUiConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SagemakerLabelingJobHumanTaskConfigOutputReference
type jsiiProxy_SagemakerLabelingJobHumanTaskConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SagemakerLabelingJobHumanTaskConfigOutputReference) AnnotationConsolidationConfig() SagemakerLabelingJobHumanTaskConfigAnnotationConsolidationConfigList {
	var returns SagemakerLabelingJobHumanTaskConfigAnnotationConsolidationConfigList
	_jsii_.Get(
		j,
		"annotationConsolidationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerLabelingJobHumanTaskConfigOutputReference) AnnotationConsolidationConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"annotationConsolidationConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerLabelingJobHumanTaskConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerLabelingJobHumanTaskConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerLabelingJobHumanTaskConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerLabelingJobHumanTaskConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerLabelingJobHumanTaskConfigOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerLabelingJobHumanTaskConfigOutputReference) MaxConcurrentTaskCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrentTaskCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerLabelingJobHumanTaskConfigOutputReference) MaxConcurrentTaskCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrentTaskCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerLabelingJobHumanTaskConfigOutputReference) NumberOfHumanWorkersPerDataObject() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfHumanWorkersPerDataObject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerLabelingJobHumanTaskConfigOutputReference) NumberOfHumanWorkersPerDataObjectInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfHumanWorkersPerDataObjectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerLabelingJobHumanTaskConfigOutputReference) PreHumanTaskLambdaArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preHumanTaskLambdaArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerLabelingJobHumanTaskConfigOutputReference) PreHumanTaskLambdaArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preHumanTaskLambdaArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerLabelingJobHumanTaskConfigOutputReference) PublicWorkforceTaskPrice() SagemakerLabelingJobHumanTaskConfigPublicWorkforceTaskPriceList {
	var returns SagemakerLabelingJobHumanTaskConfigPublicWorkforceTaskPriceList
	_jsii_.Get(
		j,
		"publicWorkforceTaskPrice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerLabelingJobHumanTaskConfigOutputReference) PublicWorkforceTaskPriceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicWorkforceTaskPriceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerLabelingJobHumanTaskConfigOutputReference) TaskAvailabilityLifetimeInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"taskAvailabilityLifetimeInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerLabelingJobHumanTaskConfigOutputReference) TaskAvailabilityLifetimeInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"taskAvailabilityLifetimeInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerLabelingJobHumanTaskConfigOutputReference) TaskDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerLabelingJobHumanTaskConfigOutputReference) TaskDescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskDescriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerLabelingJobHumanTaskConfigOutputReference) TaskKeywords() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"taskKeywords",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerLabelingJobHumanTaskConfigOutputReference) TaskKeywordsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"taskKeywordsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerLabelingJobHumanTaskConfigOutputReference) TaskTimeLimitInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"taskTimeLimitInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerLabelingJobHumanTaskConfigOutputReference) TaskTimeLimitInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"taskTimeLimitInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerLabelingJobHumanTaskConfigOutputReference) TaskTitle() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskTitle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerLabelingJobHumanTaskConfigOutputReference) TaskTitleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskTitleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerLabelingJobHumanTaskConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerLabelingJobHumanTaskConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerLabelingJobHumanTaskConfigOutputReference) UiConfig() SagemakerLabelingJobHumanTaskConfigUiConfigList {
	var returns SagemakerLabelingJobHumanTaskConfigUiConfigList
	_jsii_.Get(
		j,
		"uiConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerLabelingJobHumanTaskConfigOutputReference) UiConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"uiConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerLabelingJobHumanTaskConfigOutputReference) WorkteamArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workteamArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerLabelingJobHumanTaskConfigOutputReference) WorkteamArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workteamArnInput",
		&returns,
	)
	return returns
}


func NewSagemakerLabelingJobHumanTaskConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) SagemakerLabelingJobHumanTaskConfigOutputReference {
	_init_.Initialize()

	if err := validateNewSagemakerLabelingJobHumanTaskConfigOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_SagemakerLabelingJobHumanTaskConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-aws.sagemakerLabelingJob.SagemakerLabelingJobHumanTaskConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewSagemakerLabelingJobHumanTaskConfigOutputReference_Override(s SagemakerLabelingJobHumanTaskConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.sagemakerLabelingJob.SagemakerLabelingJobHumanTaskConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_SagemakerLabelingJobHumanTaskConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SagemakerLabelingJobHumanTaskConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SagemakerLabelingJobHumanTaskConfigOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SagemakerLabelingJobHumanTaskConfigOutputReference)SetMaxConcurrentTaskCount(val *float64) {
	if err := j.validateSetMaxConcurrentTaskCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxConcurrentTaskCount",
		val,
	)
}

func (j *jsiiProxy_SagemakerLabelingJobHumanTaskConfigOutputReference)SetNumberOfHumanWorkersPerDataObject(val *float64) {
	if err := j.validateSetNumberOfHumanWorkersPerDataObjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numberOfHumanWorkersPerDataObject",
		val,
	)
}

func (j *jsiiProxy_SagemakerLabelingJobHumanTaskConfigOutputReference)SetPreHumanTaskLambdaArn(val *string) {
	if err := j.validateSetPreHumanTaskLambdaArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preHumanTaskLambdaArn",
		val,
	)
}

func (j *jsiiProxy_SagemakerLabelingJobHumanTaskConfigOutputReference)SetTaskAvailabilityLifetimeInSeconds(val *float64) {
	if err := j.validateSetTaskAvailabilityLifetimeInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"taskAvailabilityLifetimeInSeconds",
		val,
	)
}

func (j *jsiiProxy_SagemakerLabelingJobHumanTaskConfigOutputReference)SetTaskDescription(val *string) {
	if err := j.validateSetTaskDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"taskDescription",
		val,
	)
}

func (j *jsiiProxy_SagemakerLabelingJobHumanTaskConfigOutputReference)SetTaskKeywords(val *[]*string) {
	if err := j.validateSetTaskKeywordsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"taskKeywords",
		val,
	)
}

func (j *jsiiProxy_SagemakerLabelingJobHumanTaskConfigOutputReference)SetTaskTimeLimitInSeconds(val *float64) {
	if err := j.validateSetTaskTimeLimitInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"taskTimeLimitInSeconds",
		val,
	)
}

func (j *jsiiProxy_SagemakerLabelingJobHumanTaskConfigOutputReference)SetTaskTitle(val *string) {
	if err := j.validateSetTaskTitleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"taskTitle",
		val,
	)
}

func (j *jsiiProxy_SagemakerLabelingJobHumanTaskConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SagemakerLabelingJobHumanTaskConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_SagemakerLabelingJobHumanTaskConfigOutputReference)SetWorkteamArn(val *string) {
	if err := j.validateSetWorkteamArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workteamArn",
		val,
	)
}

func (s *jsiiProxy_SagemakerLabelingJobHumanTaskConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerLabelingJobHumanTaskConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SagemakerLabelingJobHumanTaskConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerLabelingJobHumanTaskConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SagemakerLabelingJobHumanTaskConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SagemakerLabelingJobHumanTaskConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SagemakerLabelingJobHumanTaskConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SagemakerLabelingJobHumanTaskConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SagemakerLabelingJobHumanTaskConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SagemakerLabelingJobHumanTaskConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SagemakerLabelingJobHumanTaskConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerLabelingJobHumanTaskConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerLabelingJobHumanTaskConfigOutputReference) PutAnnotationConsolidationConfig(value interface{}) {
	if err := s.validatePutAnnotationConsolidationConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAnnotationConsolidationConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerLabelingJobHumanTaskConfigOutputReference) PutPublicWorkforceTaskPrice(value interface{}) {
	if err := s.validatePutPublicWorkforceTaskPriceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putPublicWorkforceTaskPrice",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerLabelingJobHumanTaskConfigOutputReference) PutUiConfig(value interface{}) {
	if err := s.validatePutUiConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putUiConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerLabelingJobHumanTaskConfigOutputReference) ResetAnnotationConsolidationConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetAnnotationConsolidationConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerLabelingJobHumanTaskConfigOutputReference) ResetMaxConcurrentTaskCount() {
	_jsii_.InvokeVoid(
		s,
		"resetMaxConcurrentTaskCount",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerLabelingJobHumanTaskConfigOutputReference) ResetPreHumanTaskLambdaArn() {
	_jsii_.InvokeVoid(
		s,
		"resetPreHumanTaskLambdaArn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerLabelingJobHumanTaskConfigOutputReference) ResetPublicWorkforceTaskPrice() {
	_jsii_.InvokeVoid(
		s,
		"resetPublicWorkforceTaskPrice",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerLabelingJobHumanTaskConfigOutputReference) ResetTaskAvailabilityLifetimeInSeconds() {
	_jsii_.InvokeVoid(
		s,
		"resetTaskAvailabilityLifetimeInSeconds",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerLabelingJobHumanTaskConfigOutputReference) ResetTaskKeywords() {
	_jsii_.InvokeVoid(
		s,
		"resetTaskKeywords",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerLabelingJobHumanTaskConfigOutputReference) ResetUiConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetUiConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerLabelingJobHumanTaskConfigOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_SagemakerLabelingJobHumanTaskConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

