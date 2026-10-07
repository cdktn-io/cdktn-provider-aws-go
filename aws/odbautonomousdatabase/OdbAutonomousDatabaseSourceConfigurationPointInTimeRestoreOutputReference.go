// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package odbautonomousdatabase

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/jsii"

	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/odbautonomousdatabase/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type OdbAutonomousDatabaseSourceConfigurationPointInTimeRestoreOutputReference interface {
	cdktn.ComplexObject
	CloneTableSpaceList() *[]*float64
	SetCloneTableSpaceList(val *[]*float64)
	CloneTableSpaceListInput() *[]*float64
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
	Timestamp() *string
	SetTimestamp(val *string)
	TimestampInput() *string
	UseLatestAvailableBackupTimestamp() interface{}
	SetUseLatestAvailableBackupTimestamp(val interface{})
	UseLatestAvailableBackupTimestampInput() interface{}
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
	ResetCloneTableSpaceList()
	ResetTimestamp()
	ResetUseLatestAvailableBackupTimestamp()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OdbAutonomousDatabaseSourceConfigurationPointInTimeRestoreOutputReference
type jsiiProxy_OdbAutonomousDatabaseSourceConfigurationPointInTimeRestoreOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationPointInTimeRestoreOutputReference) CloneTableSpaceList() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"cloneTableSpaceList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationPointInTimeRestoreOutputReference) CloneTableSpaceListInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"cloneTableSpaceListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationPointInTimeRestoreOutputReference) CloneType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloneType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationPointInTimeRestoreOutputReference) CloneTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloneTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationPointInTimeRestoreOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationPointInTimeRestoreOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationPointInTimeRestoreOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationPointInTimeRestoreOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationPointInTimeRestoreOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationPointInTimeRestoreOutputReference) SourceAutonomousDatabaseId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceAutonomousDatabaseId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationPointInTimeRestoreOutputReference) SourceAutonomousDatabaseIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceAutonomousDatabaseIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationPointInTimeRestoreOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationPointInTimeRestoreOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationPointInTimeRestoreOutputReference) Timestamp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationPointInTimeRestoreOutputReference) TimestampInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationPointInTimeRestoreOutputReference) UseLatestAvailableBackupTimestamp() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useLatestAvailableBackupTimestamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationPointInTimeRestoreOutputReference) UseLatestAvailableBackupTimestampInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useLatestAvailableBackupTimestampInput",
		&returns,
	)
	return returns
}


func NewOdbAutonomousDatabaseSourceConfigurationPointInTimeRestoreOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) OdbAutonomousDatabaseSourceConfigurationPointInTimeRestoreOutputReference {
	_init_.Initialize()

	if err := validateNewOdbAutonomousDatabaseSourceConfigurationPointInTimeRestoreOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_OdbAutonomousDatabaseSourceConfigurationPointInTimeRestoreOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-aws.odbAutonomousDatabase.OdbAutonomousDatabaseSourceConfigurationPointInTimeRestoreOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewOdbAutonomousDatabaseSourceConfigurationPointInTimeRestoreOutputReference_Override(o OdbAutonomousDatabaseSourceConfigurationPointInTimeRestoreOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.odbAutonomousDatabase.OdbAutonomousDatabaseSourceConfigurationPointInTimeRestoreOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		o,
	)
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationPointInTimeRestoreOutputReference)SetCloneTableSpaceList(val *[]*float64) {
	if err := j.validateSetCloneTableSpaceListParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cloneTableSpaceList",
		val,
	)
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationPointInTimeRestoreOutputReference)SetCloneType(val *string) {
	if err := j.validateSetCloneTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cloneType",
		val,
	)
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationPointInTimeRestoreOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationPointInTimeRestoreOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationPointInTimeRestoreOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationPointInTimeRestoreOutputReference)SetSourceAutonomousDatabaseId(val *string) {
	if err := j.validateSetSourceAutonomousDatabaseIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceAutonomousDatabaseId",
		val,
	)
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationPointInTimeRestoreOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationPointInTimeRestoreOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationPointInTimeRestoreOutputReference)SetTimestamp(val *string) {
	if err := j.validateSetTimestampParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timestamp",
		val,
	)
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationPointInTimeRestoreOutputReference)SetUseLatestAvailableBackupTimestamp(val interface{}) {
	if err := j.validateSetUseLatestAvailableBackupTimestampParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useLatestAvailableBackupTimestamp",
		val,
	)
}

func (o *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationPointInTimeRestoreOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationPointInTimeRestoreOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationPointInTimeRestoreOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (o *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationPointInTimeRestoreOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationPointInTimeRestoreOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationPointInTimeRestoreOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationPointInTimeRestoreOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationPointInTimeRestoreOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationPointInTimeRestoreOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationPointInTimeRestoreOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationPointInTimeRestoreOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationPointInTimeRestoreOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (o *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationPointInTimeRestoreOutputReference) ResetCloneTableSpaceList() {
	_jsii_.InvokeVoid(
		o,
		"resetCloneTableSpaceList",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationPointInTimeRestoreOutputReference) ResetTimestamp() {
	_jsii_.InvokeVoid(
		o,
		"resetTimestamp",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationPointInTimeRestoreOutputReference) ResetUseLatestAvailableBackupTimestamp() {
	_jsii_.InvokeVoid(
		o,
		"resetUseLatestAvailableBackupTimestamp",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationPointInTimeRestoreOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (o *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationPointInTimeRestoreOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

