// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package odbautonomousdatabase

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/jsii"

	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/odbautonomousdatabase/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type OdbAutonomousDatabaseSourceConfigurationOutputReference interface {
	cdktn.ComplexObject
	CloneToRefreshable() OdbAutonomousDatabaseSourceConfigurationCloneToRefreshableList
	CloneToRefreshableInput() interface{}
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
	CrossRegionDataGuard() OdbAutonomousDatabaseSourceConfigurationCrossRegionDataGuardList
	CrossRegionDataGuardInput() interface{}
	CrossRegionDisasterRecovery() OdbAutonomousDatabaseSourceConfigurationCrossRegionDisasterRecoveryList
	CrossRegionDisasterRecoveryInput() interface{}
	DatabaseClone() OdbAutonomousDatabaseSourceConfigurationDatabaseCloneList
	DatabaseCloneInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	PointInTimeRestore() OdbAutonomousDatabaseSourceConfigurationPointInTimeRestoreList
	PointInTimeRestoreInput() interface{}
	RestoreFromBackup() OdbAutonomousDatabaseSourceConfigurationRestoreFromBackupList
	RestoreFromBackupInput() interface{}
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
	PutCloneToRefreshable(value interface{})
	PutCrossRegionDataGuard(value interface{})
	PutCrossRegionDisasterRecovery(value interface{})
	PutDatabaseClone(value interface{})
	PutPointInTimeRestore(value interface{})
	PutRestoreFromBackup(value interface{})
	ResetCloneToRefreshable()
	ResetCrossRegionDataGuard()
	ResetCrossRegionDisasterRecovery()
	ResetDatabaseClone()
	ResetPointInTimeRestore()
	ResetRestoreFromBackup()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OdbAutonomousDatabaseSourceConfigurationOutputReference
type jsiiProxy_OdbAutonomousDatabaseSourceConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationOutputReference) CloneToRefreshable() OdbAutonomousDatabaseSourceConfigurationCloneToRefreshableList {
	var returns OdbAutonomousDatabaseSourceConfigurationCloneToRefreshableList
	_jsii_.Get(
		j,
		"cloneToRefreshable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationOutputReference) CloneToRefreshableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cloneToRefreshableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationOutputReference) CrossRegionDataGuard() OdbAutonomousDatabaseSourceConfigurationCrossRegionDataGuardList {
	var returns OdbAutonomousDatabaseSourceConfigurationCrossRegionDataGuardList
	_jsii_.Get(
		j,
		"crossRegionDataGuard",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationOutputReference) CrossRegionDataGuardInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"crossRegionDataGuardInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationOutputReference) CrossRegionDisasterRecovery() OdbAutonomousDatabaseSourceConfigurationCrossRegionDisasterRecoveryList {
	var returns OdbAutonomousDatabaseSourceConfigurationCrossRegionDisasterRecoveryList
	_jsii_.Get(
		j,
		"crossRegionDisasterRecovery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationOutputReference) CrossRegionDisasterRecoveryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"crossRegionDisasterRecoveryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationOutputReference) DatabaseClone() OdbAutonomousDatabaseSourceConfigurationDatabaseCloneList {
	var returns OdbAutonomousDatabaseSourceConfigurationDatabaseCloneList
	_jsii_.Get(
		j,
		"databaseClone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationOutputReference) DatabaseCloneInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"databaseCloneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationOutputReference) PointInTimeRestore() OdbAutonomousDatabaseSourceConfigurationPointInTimeRestoreList {
	var returns OdbAutonomousDatabaseSourceConfigurationPointInTimeRestoreList
	_jsii_.Get(
		j,
		"pointInTimeRestore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationOutputReference) PointInTimeRestoreInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pointInTimeRestoreInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationOutputReference) RestoreFromBackup() OdbAutonomousDatabaseSourceConfigurationRestoreFromBackupList {
	var returns OdbAutonomousDatabaseSourceConfigurationRestoreFromBackupList
	_jsii_.Get(
		j,
		"restoreFromBackup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationOutputReference) RestoreFromBackupInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"restoreFromBackupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewOdbAutonomousDatabaseSourceConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) OdbAutonomousDatabaseSourceConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewOdbAutonomousDatabaseSourceConfigurationOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_OdbAutonomousDatabaseSourceConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-aws.odbAutonomousDatabase.OdbAutonomousDatabaseSourceConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewOdbAutonomousDatabaseSourceConfigurationOutputReference_Override(o OdbAutonomousDatabaseSourceConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.odbAutonomousDatabase.OdbAutonomousDatabaseSourceConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		o,
	)
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (o *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (o *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationOutputReference) PutCloneToRefreshable(value interface{}) {
	if err := o.validatePutCloneToRefreshableParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putCloneToRefreshable",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationOutputReference) PutCrossRegionDataGuard(value interface{}) {
	if err := o.validatePutCrossRegionDataGuardParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putCrossRegionDataGuard",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationOutputReference) PutCrossRegionDisasterRecovery(value interface{}) {
	if err := o.validatePutCrossRegionDisasterRecoveryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putCrossRegionDisasterRecovery",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationOutputReference) PutDatabaseClone(value interface{}) {
	if err := o.validatePutDatabaseCloneParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putDatabaseClone",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationOutputReference) PutPointInTimeRestore(value interface{}) {
	if err := o.validatePutPointInTimeRestoreParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putPointInTimeRestore",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationOutputReference) PutRestoreFromBackup(value interface{}) {
	if err := o.validatePutRestoreFromBackupParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putRestoreFromBackup",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationOutputReference) ResetCloneToRefreshable() {
	_jsii_.InvokeVoid(
		o,
		"resetCloneToRefreshable",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationOutputReference) ResetCrossRegionDataGuard() {
	_jsii_.InvokeVoid(
		o,
		"resetCrossRegionDataGuard",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationOutputReference) ResetCrossRegionDisasterRecovery() {
	_jsii_.InvokeVoid(
		o,
		"resetCrossRegionDisasterRecovery",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationOutputReference) ResetDatabaseClone() {
	_jsii_.InvokeVoid(
		o,
		"resetDatabaseClone",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationOutputReference) ResetPointInTimeRestore() {
	_jsii_.InvokeVoid(
		o,
		"resetPointInTimeRestore",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationOutputReference) ResetRestoreFromBackup() {
	_jsii_.InvokeVoid(
		o,
		"resetRestoreFromBackup",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (o *jsiiProxy_OdbAutonomousDatabaseSourceConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

