// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawsrdssnapshots

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/jsii"

	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/dataawsrdssnapshots/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAwsRdsSnapshotsSnapshotsOutputReference interface {
	cdktn.ComplexObject
	AllocatedStorage() *float64
	AvailabilityZone() *string
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
	DbInstanceIdentifier() *string
	DbSnapshotArn() *string
	DbSnapshotIdentifier() *string
	Encrypted() cdktn.IResolvable
	Engine() *string
	EngineVersion() *string
	// Experimental.
	Fqn() *string
	InternalValue() *DataAwsRdsSnapshotsSnapshots
	SetInternalValue(val *DataAwsRdsSnapshotsSnapshots)
	Iops() *float64
	KmsKeyId() *string
	LicenseModel() *string
	OptionGroupName() *string
	OriginalSnapshotCreateTime() *string
	Port() *float64
	SnapshotCreateTime() *string
	SnapshotType() *string
	SourceDbSnapshotIdentifier() *string
	SourceRegion() *string
	Status() *string
	StorageType() *string
	TagList() DataAwsRdsSnapshotsSnapshotsTagListStructList
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	VpcId() *string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataAwsRdsSnapshotsSnapshotsOutputReference
type jsiiProxy_DataAwsRdsSnapshotsSnapshotsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsRdsSnapshotsSnapshotsOutputReference) AllocatedStorage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"allocatedStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsSnapshotsSnapshotsOutputReference) AvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsSnapshotsSnapshotsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsSnapshotsSnapshotsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsSnapshotsSnapshotsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsSnapshotsSnapshotsOutputReference) DbInstanceIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbInstanceIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsSnapshotsSnapshotsOutputReference) DbSnapshotArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbSnapshotArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsSnapshotsSnapshotsOutputReference) DbSnapshotIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbSnapshotIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsSnapshotsSnapshotsOutputReference) Encrypted() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"encrypted",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsSnapshotsSnapshotsOutputReference) Engine() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsSnapshotsSnapshotsOutputReference) EngineVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsSnapshotsSnapshotsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsSnapshotsSnapshotsOutputReference) InternalValue() *DataAwsRdsSnapshotsSnapshots {
	var returns *DataAwsRdsSnapshotsSnapshots
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsSnapshotsSnapshotsOutputReference) Iops() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"iops",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsSnapshotsSnapshotsOutputReference) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsSnapshotsSnapshotsOutputReference) LicenseModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsSnapshotsSnapshotsOutputReference) OptionGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"optionGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsSnapshotsSnapshotsOutputReference) OriginalSnapshotCreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originalSnapshotCreateTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsSnapshotsSnapshotsOutputReference) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsSnapshotsSnapshotsOutputReference) SnapshotCreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotCreateTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsSnapshotsSnapshotsOutputReference) SnapshotType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsSnapshotsSnapshotsOutputReference) SourceDbSnapshotIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceDbSnapshotIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsSnapshotsSnapshotsOutputReference) SourceRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsSnapshotsSnapshotsOutputReference) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsSnapshotsSnapshotsOutputReference) StorageType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsSnapshotsSnapshotsOutputReference) TagList() DataAwsRdsSnapshotsSnapshotsTagListStructList {
	var returns DataAwsRdsSnapshotsSnapshotsTagListStructList
	_jsii_.Get(
		j,
		"tagList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsSnapshotsSnapshotsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsSnapshotsSnapshotsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsSnapshotsSnapshotsOutputReference) VpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcId",
		&returns,
	)
	return returns
}


func NewDataAwsRdsSnapshotsSnapshotsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataAwsRdsSnapshotsSnapshotsOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsRdsSnapshotsSnapshotsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsRdsSnapshotsSnapshotsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-aws.dataAwsRdsSnapshots.DataAwsRdsSnapshotsSnapshotsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataAwsRdsSnapshotsSnapshotsOutputReference_Override(d DataAwsRdsSnapshotsSnapshotsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.dataAwsRdsSnapshots.DataAwsRdsSnapshotsSnapshotsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsRdsSnapshotsSnapshotsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsRdsSnapshotsSnapshotsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsRdsSnapshotsSnapshotsOutputReference)SetInternalValue(val *DataAwsRdsSnapshotsSnapshots) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsRdsSnapshotsSnapshotsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsRdsSnapshotsSnapshotsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsRdsSnapshotsSnapshotsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsRdsSnapshotsSnapshotsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsRdsSnapshotsSnapshotsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsRdsSnapshotsSnapshotsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsRdsSnapshotsSnapshotsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsRdsSnapshotsSnapshotsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsRdsSnapshotsSnapshotsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsRdsSnapshotsSnapshotsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsRdsSnapshotsSnapshotsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsRdsSnapshotsSnapshotsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsRdsSnapshotsSnapshotsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsRdsSnapshotsSnapshotsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsRdsSnapshotsSnapshotsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsRdsSnapshotsSnapshotsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

