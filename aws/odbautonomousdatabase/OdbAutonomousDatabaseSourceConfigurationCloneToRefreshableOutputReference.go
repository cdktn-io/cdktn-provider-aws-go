// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package odbautonomousdatabase

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/jsii"

	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/odbautonomousdatabase/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type OdbAutonomousDatabaseSourceConfigurationCloneToRefreshableOutputReference interface {
	cdktn.ComplexObject
	AutoRefreshFrequencyInSeconds() *float64
	SetAutoRefreshFrequencyInSeconds(val *float64)
	AutoRefreshFrequencyInSecondsInput() *float64
	AutoRefreshPointLagInSeconds() *float64
	SetAutoRefreshPointLagInSeconds(val *float64)
	AutoRefreshPointLagInSecondsInput() *float64
	CloneType() *string
	SetCloneType(val *string)
	CloneTypeInput() *string
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
	OpenMode() *string
	SetOpenMode(val *string)
	OpenModeInput() *string
	RefreshableMode() *string
	SetRefreshableMode(val *string)
	RefreshableModeInput() *string
	SourceAutonomousDatabaseId() *string
	SetSourceAutonomousDatabaseId(val *string)
	SourceAutonomousDatabaseIdInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TimeOfAutoRefreshStart() *string
	SetTimeOfAutoRefreshStart(val *string)
	TimeOfAutoRefreshStartInput() *string
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
	ResetAutoRefreshFrequencyInSeconds()
	ResetAutoRefreshPointLagInSeconds()
	ResetCloneType()
	ResetOpenMode()
	ResetRefreshableMode()
	ResetTimeOfAutoRefreshStart()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OdbAutonomousDatabaseSourceConfigurationCloneToRefreshableOutputReference
type jsiiProxy_OdbAutonomousDatabaseSourceConfigurationCloneToRefreshableOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationCloneToRefreshableOutputReference) AutoRefreshFrequencyInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autoRefreshFrequencyInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationCloneToRefreshableOutputReference) AutoRefreshFrequencyInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autoRefreshFrequencyInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationCloneToRefreshableOutputReference) AutoRefreshPointLagInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autoRefreshPointLagInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationCloneToRefreshableOutputReference) AutoRefreshPointLagInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autoRefreshPointLagInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationCloneToRefreshableOutputReference) CloneType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloneType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationCloneToRefreshableOutputReference) CloneTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloneTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationCloneToRefreshableOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationCloneToRefreshableOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationCloneToRefreshableOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationCloneToRefreshableOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationCloneToRefreshableOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationCloneToRefreshableOutputReference) OpenMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"openMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationCloneToRefreshableOutputReference) OpenModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"openModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationCloneToRefreshableOutputReference) RefreshableMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"refreshableMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationCloneToRefreshableOutputReference) RefreshableModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"refreshableModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationCloneToRefreshableOutputReference) SourceAutonomousDatabaseId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceAutonomousDatabaseId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationCloneToRefreshableOutputReference) SourceAutonomousDatabaseIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceAutonomousDatabaseIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationCloneToRefreshableOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationCloneToRefreshableOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationCloneToRefreshableOutputReference) TimeOfAutoRefreshStart() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeOfAutoRefreshStart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationCloneToRefreshableOutputReference) TimeOfAutoRefreshStartInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeOfAutoRefreshStartInput",
		&returns,
	)
	return returns
}


func NewOdbAutonomousDatabaseSourceConfigurationCloneToRefreshableOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) OdbAutonomousDatabaseSourceConfigurationCloneToRefreshableOutputReference {
	_init_.Initialize()

	if err := validateNewOdbAutonomousDatabaseSourceConfigurationCloneToRefreshableOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_OdbAutonomousDatabaseSourceConfigurationCloneToRefreshableOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-aws.odbAutonomousDatabase.OdbAutonomousDatabaseSourceConfigurationCloneToRefreshableOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewOdbAutonomousDatabaseSourceConfigurationCloneToRefreshableOutputReference_Override(o OdbAutonomousDatabaseSourceConfigurationCloneToRefreshableOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.odbAutonomousDatabase.OdbAutonomousDatabaseSourceConfigurationCloneToRefreshableOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		o,
	)
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationCloneToRefreshableOutputReference)SetAutoRefreshFrequencyInSeconds(val *float64) {
	if err := j.validateSetAutoRefreshFrequencyInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoRefreshFrequencyInSeconds",
		val,
	)
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationCloneToRefreshableOutputReference)SetAutoRefreshPointLagInSeconds(val *float64) {
	if err := j.validateSetAutoRefreshPointLagInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoRefreshPointLagInSeconds",
		val,
	)
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationCloneToRefreshableOutputReference)SetCloneType(val *string) {
	if err := j.validateSetCloneTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cloneType",
		val,
	)
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationCloneToRefreshableOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationCloneToRefreshableOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationCloneToRefreshableOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationCloneToRefreshableOutputReference)SetOpenMode(val *string) {
	if err := j.validateSetOpenModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"openMode",
		val,
	)
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationCloneToRefreshableOutputReference)SetRefreshableMode(val *string) {
	if err := j.validateSetRefreshableModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"refreshableMode",
		val,
	)
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationCloneToRefreshableOutputReference)SetSourceAutonomousDatabaseId(val *string) {
	if err := j.validateSetSourceAutonomousDatabaseIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceAutonomousDatabaseId",
		val,
	)
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationCloneToRefreshableOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationCloneToRefreshableOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationCloneToRefreshableOutputReference)SetTimeOfAutoRefreshStart(val *string) {
	if err := j.validateSetTimeOfAutoRefreshStartParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeOfAutoRefreshStart",
		val,
	)
}

func (o *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationCloneToRefreshableOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationCloneToRefreshableOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := o.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationCloneToRefreshableOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := o.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationCloneToRefreshableOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := o.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationCloneToRefreshableOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := o.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationCloneToRefreshableOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := o.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationCloneToRefreshableOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := o.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationCloneToRefreshableOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := o.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationCloneToRefreshableOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := o.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationCloneToRefreshableOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := o.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationCloneToRefreshableOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationCloneToRefreshableOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := o.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationCloneToRefreshableOutputReference) ResetAutoRefreshFrequencyInSeconds() {
	_jsii_.InvokeVoid(
		o,
		"resetAutoRefreshFrequencyInSeconds",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationCloneToRefreshableOutputReference) ResetAutoRefreshPointLagInSeconds() {
	_jsii_.InvokeVoid(
		o,
		"resetAutoRefreshPointLagInSeconds",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationCloneToRefreshableOutputReference) ResetCloneType() {
	_jsii_.InvokeVoid(
		o,
		"resetCloneType",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationCloneToRefreshableOutputReference) ResetOpenMode() {
	_jsii_.InvokeVoid(
		o,
		"resetOpenMode",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationCloneToRefreshableOutputReference) ResetRefreshableMode() {
	_jsii_.InvokeVoid(
		o,
		"resetRefreshableMode",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationCloneToRefreshableOutputReference) ResetTimeOfAutoRefreshStart() {
	_jsii_.InvokeVoid(
		o,
		"resetTimeOfAutoRefreshStart",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationCloneToRefreshableOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := o.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		o,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationCloneToRefreshableOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

