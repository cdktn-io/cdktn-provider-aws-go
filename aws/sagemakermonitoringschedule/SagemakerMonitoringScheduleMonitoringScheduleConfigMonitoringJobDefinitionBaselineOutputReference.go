// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakermonitoringschedule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v23/jsii"

	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v23/sagemakermonitoringschedule/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SagemakerMonitoringScheduleMonitoringScheduleConfigMonitoringJobDefinitionBaselineOutputReference interface {
	cdktn.ComplexObject
	BaseliningJobName() *string
	SetBaseliningJobName(val *string)
	BaseliningJobNameInput() *string
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
	ConstraintsResource() SagemakerMonitoringScheduleMonitoringScheduleConfigMonitoringJobDefinitionBaselineConstraintsResourceOutputReference
	ConstraintsResourceInput() *SagemakerMonitoringScheduleMonitoringScheduleConfigMonitoringJobDefinitionBaselineConstraintsResource
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *SagemakerMonitoringScheduleMonitoringScheduleConfigMonitoringJobDefinitionBaseline
	SetInternalValue(val *SagemakerMonitoringScheduleMonitoringScheduleConfigMonitoringJobDefinitionBaseline)
	StatisticsResource() SagemakerMonitoringScheduleMonitoringScheduleConfigMonitoringJobDefinitionBaselineStatisticsResourceOutputReference
	StatisticsResourceInput() *SagemakerMonitoringScheduleMonitoringScheduleConfigMonitoringJobDefinitionBaselineStatisticsResource
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
	PutConstraintsResource(value *SagemakerMonitoringScheduleMonitoringScheduleConfigMonitoringJobDefinitionBaselineConstraintsResource)
	PutStatisticsResource(value *SagemakerMonitoringScheduleMonitoringScheduleConfigMonitoringJobDefinitionBaselineStatisticsResource)
	ResetBaseliningJobName()
	ResetConstraintsResource()
	ResetStatisticsResource()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SagemakerMonitoringScheduleMonitoringScheduleConfigMonitoringJobDefinitionBaselineOutputReference
type jsiiProxy_SagemakerMonitoringScheduleMonitoringScheduleConfigMonitoringJobDefinitionBaselineOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_SagemakerMonitoringScheduleMonitoringScheduleConfigMonitoringJobDefinitionBaselineOutputReference) BaseliningJobName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseliningJobName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerMonitoringScheduleMonitoringScheduleConfigMonitoringJobDefinitionBaselineOutputReference) BaseliningJobNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseliningJobNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerMonitoringScheduleMonitoringScheduleConfigMonitoringJobDefinitionBaselineOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerMonitoringScheduleMonitoringScheduleConfigMonitoringJobDefinitionBaselineOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerMonitoringScheduleMonitoringScheduleConfigMonitoringJobDefinitionBaselineOutputReference) ConstraintsResource() SagemakerMonitoringScheduleMonitoringScheduleConfigMonitoringJobDefinitionBaselineConstraintsResourceOutputReference {
	var returns SagemakerMonitoringScheduleMonitoringScheduleConfigMonitoringJobDefinitionBaselineConstraintsResourceOutputReference
	_jsii_.Get(
		j,
		"constraintsResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerMonitoringScheduleMonitoringScheduleConfigMonitoringJobDefinitionBaselineOutputReference) ConstraintsResourceInput() *SagemakerMonitoringScheduleMonitoringScheduleConfigMonitoringJobDefinitionBaselineConstraintsResource {
	var returns *SagemakerMonitoringScheduleMonitoringScheduleConfigMonitoringJobDefinitionBaselineConstraintsResource
	_jsii_.Get(
		j,
		"constraintsResourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerMonitoringScheduleMonitoringScheduleConfigMonitoringJobDefinitionBaselineOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerMonitoringScheduleMonitoringScheduleConfigMonitoringJobDefinitionBaselineOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerMonitoringScheduleMonitoringScheduleConfigMonitoringJobDefinitionBaselineOutputReference) InternalValue() *SagemakerMonitoringScheduleMonitoringScheduleConfigMonitoringJobDefinitionBaseline {
	var returns *SagemakerMonitoringScheduleMonitoringScheduleConfigMonitoringJobDefinitionBaseline
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerMonitoringScheduleMonitoringScheduleConfigMonitoringJobDefinitionBaselineOutputReference) StatisticsResource() SagemakerMonitoringScheduleMonitoringScheduleConfigMonitoringJobDefinitionBaselineStatisticsResourceOutputReference {
	var returns SagemakerMonitoringScheduleMonitoringScheduleConfigMonitoringJobDefinitionBaselineStatisticsResourceOutputReference
	_jsii_.Get(
		j,
		"statisticsResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerMonitoringScheduleMonitoringScheduleConfigMonitoringJobDefinitionBaselineOutputReference) StatisticsResourceInput() *SagemakerMonitoringScheduleMonitoringScheduleConfigMonitoringJobDefinitionBaselineStatisticsResource {
	var returns *SagemakerMonitoringScheduleMonitoringScheduleConfigMonitoringJobDefinitionBaselineStatisticsResource
	_jsii_.Get(
		j,
		"statisticsResourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerMonitoringScheduleMonitoringScheduleConfigMonitoringJobDefinitionBaselineOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerMonitoringScheduleMonitoringScheduleConfigMonitoringJobDefinitionBaselineOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSagemakerMonitoringScheduleMonitoringScheduleConfigMonitoringJobDefinitionBaselineOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) SagemakerMonitoringScheduleMonitoringScheduleConfigMonitoringJobDefinitionBaselineOutputReference {
	_init_.Initialize()

	if err := validateNewSagemakerMonitoringScheduleMonitoringScheduleConfigMonitoringJobDefinitionBaselineOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SagemakerMonitoringScheduleMonitoringScheduleConfigMonitoringJobDefinitionBaselineOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-aws.sagemakerMonitoringSchedule.SagemakerMonitoringScheduleMonitoringScheduleConfigMonitoringJobDefinitionBaselineOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSagemakerMonitoringScheduleMonitoringScheduleConfigMonitoringJobDefinitionBaselineOutputReference_Override(s SagemakerMonitoringScheduleMonitoringScheduleConfigMonitoringJobDefinitionBaselineOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.sagemakerMonitoringSchedule.SagemakerMonitoringScheduleMonitoringScheduleConfigMonitoringJobDefinitionBaselineOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SagemakerMonitoringScheduleMonitoringScheduleConfigMonitoringJobDefinitionBaselineOutputReference)SetBaseliningJobName(val *string) {
	if err := j.validateSetBaseliningJobNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"baseliningJobName",
		val,
	)
}

func (j *jsiiProxy_SagemakerMonitoringScheduleMonitoringScheduleConfigMonitoringJobDefinitionBaselineOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SagemakerMonitoringScheduleMonitoringScheduleConfigMonitoringJobDefinitionBaselineOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SagemakerMonitoringScheduleMonitoringScheduleConfigMonitoringJobDefinitionBaselineOutputReference)SetInternalValue(val *SagemakerMonitoringScheduleMonitoringScheduleConfigMonitoringJobDefinitionBaseline) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SagemakerMonitoringScheduleMonitoringScheduleConfigMonitoringJobDefinitionBaselineOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SagemakerMonitoringScheduleMonitoringScheduleConfigMonitoringJobDefinitionBaselineOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SagemakerMonitoringScheduleMonitoringScheduleConfigMonitoringJobDefinitionBaselineOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerMonitoringScheduleMonitoringScheduleConfigMonitoringJobDefinitionBaselineOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SagemakerMonitoringScheduleMonitoringScheduleConfigMonitoringJobDefinitionBaselineOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SagemakerMonitoringScheduleMonitoringScheduleConfigMonitoringJobDefinitionBaselineOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SagemakerMonitoringScheduleMonitoringScheduleConfigMonitoringJobDefinitionBaselineOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SagemakerMonitoringScheduleMonitoringScheduleConfigMonitoringJobDefinitionBaselineOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SagemakerMonitoringScheduleMonitoringScheduleConfigMonitoringJobDefinitionBaselineOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SagemakerMonitoringScheduleMonitoringScheduleConfigMonitoringJobDefinitionBaselineOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SagemakerMonitoringScheduleMonitoringScheduleConfigMonitoringJobDefinitionBaselineOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SagemakerMonitoringScheduleMonitoringScheduleConfigMonitoringJobDefinitionBaselineOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SagemakerMonitoringScheduleMonitoringScheduleConfigMonitoringJobDefinitionBaselineOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerMonitoringScheduleMonitoringScheduleConfigMonitoringJobDefinitionBaselineOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SagemakerMonitoringScheduleMonitoringScheduleConfigMonitoringJobDefinitionBaselineOutputReference) PutConstraintsResource(value *SagemakerMonitoringScheduleMonitoringScheduleConfigMonitoringJobDefinitionBaselineConstraintsResource) {
	if err := s.validatePutConstraintsResourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putConstraintsResource",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerMonitoringScheduleMonitoringScheduleConfigMonitoringJobDefinitionBaselineOutputReference) PutStatisticsResource(value *SagemakerMonitoringScheduleMonitoringScheduleConfigMonitoringJobDefinitionBaselineStatisticsResource) {
	if err := s.validatePutStatisticsResourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putStatisticsResource",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerMonitoringScheduleMonitoringScheduleConfigMonitoringJobDefinitionBaselineOutputReference) ResetBaseliningJobName() {
	_jsii_.InvokeVoid(
		s,
		"resetBaseliningJobName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerMonitoringScheduleMonitoringScheduleConfigMonitoringJobDefinitionBaselineOutputReference) ResetConstraintsResource() {
	_jsii_.InvokeVoid(
		s,
		"resetConstraintsResource",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerMonitoringScheduleMonitoringScheduleConfigMonitoringJobDefinitionBaselineOutputReference) ResetStatisticsResource() {
	_jsii_.InvokeVoid(
		s,
		"resetStatisticsResource",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerMonitoringScheduleMonitoringScheduleConfigMonitoringJobDefinitionBaselineOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (s *jsiiProxy_SagemakerMonitoringScheduleMonitoringScheduleConfigMonitoringJobDefinitionBaselineOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

