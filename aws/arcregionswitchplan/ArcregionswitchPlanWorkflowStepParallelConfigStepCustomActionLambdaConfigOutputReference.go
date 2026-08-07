// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package arcregionswitchplan

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/jsii"

	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/arcregionswitchplan/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ArcregionswitchPlanWorkflowStepParallelConfigStepCustomActionLambdaConfigOutputReference interface {
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
	Lambda() ArcregionswitchPlanWorkflowStepParallelConfigStepCustomActionLambdaConfigLambdaList
	LambdaInput() interface{}
	RegionToRun() *string
	SetRegionToRun(val *string)
	RegionToRunInput() *string
	RetryIntervalMinutes() *float64
	SetRetryIntervalMinutes(val *float64)
	RetryIntervalMinutesInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TimeoutMinutes() *float64
	SetTimeoutMinutes(val *float64)
	TimeoutMinutesInput() *float64
	Ungraceful() ArcregionswitchPlanWorkflowStepParallelConfigStepCustomActionLambdaConfigUngracefulList
	UngracefulInput() interface{}
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
	PutLambda(value interface{})
	PutUngraceful(value interface{})
	ResetLambda()
	ResetTimeoutMinutes()
	ResetUngraceful()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ArcregionswitchPlanWorkflowStepParallelConfigStepCustomActionLambdaConfigOutputReference
type jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepCustomActionLambdaConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepCustomActionLambdaConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepCustomActionLambdaConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepCustomActionLambdaConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepCustomActionLambdaConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepCustomActionLambdaConfigOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepCustomActionLambdaConfigOutputReference) Lambda() ArcregionswitchPlanWorkflowStepParallelConfigStepCustomActionLambdaConfigLambdaList {
	var returns ArcregionswitchPlanWorkflowStepParallelConfigStepCustomActionLambdaConfigLambdaList
	_jsii_.Get(
		j,
		"lambda",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepCustomActionLambdaConfigOutputReference) LambdaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lambdaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepCustomActionLambdaConfigOutputReference) RegionToRun() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionToRun",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepCustomActionLambdaConfigOutputReference) RegionToRunInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionToRunInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepCustomActionLambdaConfigOutputReference) RetryIntervalMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retryIntervalMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepCustomActionLambdaConfigOutputReference) RetryIntervalMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retryIntervalMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepCustomActionLambdaConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepCustomActionLambdaConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepCustomActionLambdaConfigOutputReference) TimeoutMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepCustomActionLambdaConfigOutputReference) TimeoutMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepCustomActionLambdaConfigOutputReference) Ungraceful() ArcregionswitchPlanWorkflowStepParallelConfigStepCustomActionLambdaConfigUngracefulList {
	var returns ArcregionswitchPlanWorkflowStepParallelConfigStepCustomActionLambdaConfigUngracefulList
	_jsii_.Get(
		j,
		"ungraceful",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepCustomActionLambdaConfigOutputReference) UngracefulInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ungracefulInput",
		&returns,
	)
	return returns
}


func NewArcregionswitchPlanWorkflowStepParallelConfigStepCustomActionLambdaConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ArcregionswitchPlanWorkflowStepParallelConfigStepCustomActionLambdaConfigOutputReference {
	_init_.Initialize()

	if err := validateNewArcregionswitchPlanWorkflowStepParallelConfigStepCustomActionLambdaConfigOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepCustomActionLambdaConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-aws.arcregionswitchPlan.ArcregionswitchPlanWorkflowStepParallelConfigStepCustomActionLambdaConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewArcregionswitchPlanWorkflowStepParallelConfigStepCustomActionLambdaConfigOutputReference_Override(a ArcregionswitchPlanWorkflowStepParallelConfigStepCustomActionLambdaConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.arcregionswitchPlan.ArcregionswitchPlanWorkflowStepParallelConfigStepCustomActionLambdaConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepCustomActionLambdaConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepCustomActionLambdaConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepCustomActionLambdaConfigOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepCustomActionLambdaConfigOutputReference)SetRegionToRun(val *string) {
	if err := j.validateSetRegionToRunParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"regionToRun",
		val,
	)
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepCustomActionLambdaConfigOutputReference)SetRetryIntervalMinutes(val *float64) {
	if err := j.validateSetRetryIntervalMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retryIntervalMinutes",
		val,
	)
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepCustomActionLambdaConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepCustomActionLambdaConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepCustomActionLambdaConfigOutputReference)SetTimeoutMinutes(val *float64) {
	if err := j.validateSetTimeoutMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeoutMinutes",
		val,
	)
}

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepCustomActionLambdaConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepCustomActionLambdaConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepCustomActionLambdaConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepCustomActionLambdaConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepCustomActionLambdaConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepCustomActionLambdaConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepCustomActionLambdaConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepCustomActionLambdaConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepCustomActionLambdaConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepCustomActionLambdaConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepCustomActionLambdaConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepCustomActionLambdaConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepCustomActionLambdaConfigOutputReference) PutLambda(value interface{}) {
	if err := a.validatePutLambdaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLambda",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepCustomActionLambdaConfigOutputReference) PutUngraceful(value interface{}) {
	if err := a.validatePutUngracefulParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putUngraceful",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepCustomActionLambdaConfigOutputReference) ResetLambda() {
	_jsii_.InvokeVoid(
		a,
		"resetLambda",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepCustomActionLambdaConfigOutputReference) ResetTimeoutMinutes() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeoutMinutes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepCustomActionLambdaConfigOutputReference) ResetUngraceful() {
	_jsii_.InvokeVoid(
		a,
		"resetUngraceful",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepCustomActionLambdaConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := a.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepCustomActionLambdaConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

