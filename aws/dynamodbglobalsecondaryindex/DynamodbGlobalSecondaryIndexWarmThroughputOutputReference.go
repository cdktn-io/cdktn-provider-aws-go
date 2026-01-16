// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dynamodbglobalsecondaryindex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v21/jsii"

	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v21/dynamodbglobalsecondaryindex/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DynamodbGlobalSecondaryIndexWarmThroughputOutputReference interface {
	cdktf.ComplexObject
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
	ReadUnitsPerSecond() *float64
	SetReadUnitsPerSecond(val *float64)
	ReadUnitsPerSecondInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WriteUnitsPerSecond() *float64
	SetWriteUnitsPerSecond(val *float64)
	WriteUnitsPerSecondInput() *float64
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
	ResetReadUnitsPerSecond()
	ResetWriteUnitsPerSecond()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DynamodbGlobalSecondaryIndexWarmThroughputOutputReference
type jsiiProxy_DynamodbGlobalSecondaryIndexWarmThroughputOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DynamodbGlobalSecondaryIndexWarmThroughputOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalSecondaryIndexWarmThroughputOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalSecondaryIndexWarmThroughputOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalSecondaryIndexWarmThroughputOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalSecondaryIndexWarmThroughputOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalSecondaryIndexWarmThroughputOutputReference) ReadUnitsPerSecond() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"readUnitsPerSecond",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalSecondaryIndexWarmThroughputOutputReference) ReadUnitsPerSecondInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"readUnitsPerSecondInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalSecondaryIndexWarmThroughputOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalSecondaryIndexWarmThroughputOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalSecondaryIndexWarmThroughputOutputReference) WriteUnitsPerSecond() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"writeUnitsPerSecond",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalSecondaryIndexWarmThroughputOutputReference) WriteUnitsPerSecondInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"writeUnitsPerSecondInput",
		&returns,
	)
	return returns
}


func NewDynamodbGlobalSecondaryIndexWarmThroughputOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DynamodbGlobalSecondaryIndexWarmThroughputOutputReference {
	_init_.Initialize()

	if err := validateNewDynamodbGlobalSecondaryIndexWarmThroughputOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DynamodbGlobalSecondaryIndexWarmThroughputOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-aws.dynamodbGlobalSecondaryIndex.DynamodbGlobalSecondaryIndexWarmThroughputOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDynamodbGlobalSecondaryIndexWarmThroughputOutputReference_Override(d DynamodbGlobalSecondaryIndexWarmThroughputOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.dynamodbGlobalSecondaryIndex.DynamodbGlobalSecondaryIndexWarmThroughputOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DynamodbGlobalSecondaryIndexWarmThroughputOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DynamodbGlobalSecondaryIndexWarmThroughputOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DynamodbGlobalSecondaryIndexWarmThroughputOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DynamodbGlobalSecondaryIndexWarmThroughputOutputReference)SetReadUnitsPerSecond(val *float64) {
	if err := j.validateSetReadUnitsPerSecondParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"readUnitsPerSecond",
		val,
	)
}

func (j *jsiiProxy_DynamodbGlobalSecondaryIndexWarmThroughputOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DynamodbGlobalSecondaryIndexWarmThroughputOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DynamodbGlobalSecondaryIndexWarmThroughputOutputReference)SetWriteUnitsPerSecond(val *float64) {
	if err := j.validateSetWriteUnitsPerSecondParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"writeUnitsPerSecond",
		val,
	)
}

func (d *jsiiProxy_DynamodbGlobalSecondaryIndexWarmThroughputOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamodbGlobalSecondaryIndexWarmThroughputOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DynamodbGlobalSecondaryIndexWarmThroughputOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamodbGlobalSecondaryIndexWarmThroughputOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DynamodbGlobalSecondaryIndexWarmThroughputOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DynamodbGlobalSecondaryIndexWarmThroughputOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DynamodbGlobalSecondaryIndexWarmThroughputOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DynamodbGlobalSecondaryIndexWarmThroughputOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DynamodbGlobalSecondaryIndexWarmThroughputOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DynamodbGlobalSecondaryIndexWarmThroughputOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DynamodbGlobalSecondaryIndexWarmThroughputOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamodbGlobalSecondaryIndexWarmThroughputOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamodbGlobalSecondaryIndexWarmThroughputOutputReference) ResetReadUnitsPerSecond() {
	_jsii_.InvokeVoid(
		d,
		"resetReadUnitsPerSecond",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbGlobalSecondaryIndexWarmThroughputOutputReference) ResetWriteUnitsPerSecond() {
	_jsii_.InvokeVoid(
		d,
		"resetWriteUnitsPerSecond",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbGlobalSecondaryIndexWarmThroughputOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DynamodbGlobalSecondaryIndexWarmThroughputOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

