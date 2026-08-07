// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package instance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/jsii"

	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/instance/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type InstanceSecondaryNetworkInterfaceOutputReference interface {
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
	DeleteOnTermination() interface{}
	SetDeleteOnTermination(val interface{})
	DeleteOnTerminationInput() interface{}
	DeviceIndex() *float64
	SetDeviceIndex(val *float64)
	DeviceIndexInput() *float64
	// Experimental.
	Fqn() *string
	InterfaceType() *string
	SetInterfaceType(val *string)
	InterfaceTypeInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MacAddress() *string
	NetworkCardIndex() *float64
	SetNetworkCardIndex(val *float64)
	NetworkCardIndexInput() *float64
	PrivateIpAddressCount() *float64
	SetPrivateIpAddressCount(val *float64)
	PrivateIpAddressCountInput() *float64
	PrivateIpAddresses() *[]*string
	SecondaryInterfaceId() *string
	SecondaryNetworkId() *string
	SecondarySubnetId() *string
	SetSecondarySubnetId(val *string)
	SecondarySubnetIdInput() *string
	SourceDestCheck() cdktn.IResolvable
	Status() *string
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
	ResetDeleteOnTermination()
	ResetDeviceIndex()
	ResetInterfaceType()
	ResetPrivateIpAddressCount()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for InstanceSecondaryNetworkInterfaceOutputReference
type jsiiProxy_InstanceSecondaryNetworkInterfaceOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_InstanceSecondaryNetworkInterfaceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstanceSecondaryNetworkInterfaceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstanceSecondaryNetworkInterfaceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstanceSecondaryNetworkInterfaceOutputReference) DeleteOnTermination() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteOnTermination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstanceSecondaryNetworkInterfaceOutputReference) DeleteOnTerminationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteOnTerminationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstanceSecondaryNetworkInterfaceOutputReference) DeviceIndex() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deviceIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstanceSecondaryNetworkInterfaceOutputReference) DeviceIndexInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deviceIndexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstanceSecondaryNetworkInterfaceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstanceSecondaryNetworkInterfaceOutputReference) InterfaceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interfaceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstanceSecondaryNetworkInterfaceOutputReference) InterfaceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interfaceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstanceSecondaryNetworkInterfaceOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstanceSecondaryNetworkInterfaceOutputReference) MacAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"macAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstanceSecondaryNetworkInterfaceOutputReference) NetworkCardIndex() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"networkCardIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstanceSecondaryNetworkInterfaceOutputReference) NetworkCardIndexInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"networkCardIndexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstanceSecondaryNetworkInterfaceOutputReference) PrivateIpAddressCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"privateIpAddressCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstanceSecondaryNetworkInterfaceOutputReference) PrivateIpAddressCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"privateIpAddressCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstanceSecondaryNetworkInterfaceOutputReference) PrivateIpAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"privateIpAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstanceSecondaryNetworkInterfaceOutputReference) SecondaryInterfaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryInterfaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstanceSecondaryNetworkInterfaceOutputReference) SecondaryNetworkId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryNetworkId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstanceSecondaryNetworkInterfaceOutputReference) SecondarySubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondarySubnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstanceSecondaryNetworkInterfaceOutputReference) SecondarySubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondarySubnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstanceSecondaryNetworkInterfaceOutputReference) SourceDestCheck() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"sourceDestCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstanceSecondaryNetworkInterfaceOutputReference) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstanceSecondaryNetworkInterfaceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstanceSecondaryNetworkInterfaceOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewInstanceSecondaryNetworkInterfaceOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) InstanceSecondaryNetworkInterfaceOutputReference {
	_init_.Initialize()

	if err := validateNewInstanceSecondaryNetworkInterfaceOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_InstanceSecondaryNetworkInterfaceOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-aws.instance.InstanceSecondaryNetworkInterfaceOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewInstanceSecondaryNetworkInterfaceOutputReference_Override(i InstanceSecondaryNetworkInterfaceOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.instance.InstanceSecondaryNetworkInterfaceOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		i,
	)
}

func (j *jsiiProxy_InstanceSecondaryNetworkInterfaceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_InstanceSecondaryNetworkInterfaceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_InstanceSecondaryNetworkInterfaceOutputReference)SetDeleteOnTermination(val interface{}) {
	if err := j.validateSetDeleteOnTerminationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteOnTermination",
		val,
	)
}

func (j *jsiiProxy_InstanceSecondaryNetworkInterfaceOutputReference)SetDeviceIndex(val *float64) {
	if err := j.validateSetDeviceIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deviceIndex",
		val,
	)
}

func (j *jsiiProxy_InstanceSecondaryNetworkInterfaceOutputReference)SetInterfaceType(val *string) {
	if err := j.validateSetInterfaceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"interfaceType",
		val,
	)
}

func (j *jsiiProxy_InstanceSecondaryNetworkInterfaceOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_InstanceSecondaryNetworkInterfaceOutputReference)SetNetworkCardIndex(val *float64) {
	if err := j.validateSetNetworkCardIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkCardIndex",
		val,
	)
}

func (j *jsiiProxy_InstanceSecondaryNetworkInterfaceOutputReference)SetPrivateIpAddressCount(val *float64) {
	if err := j.validateSetPrivateIpAddressCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateIpAddressCount",
		val,
	)
}

func (j *jsiiProxy_InstanceSecondaryNetworkInterfaceOutputReference)SetSecondarySubnetId(val *string) {
	if err := j.validateSetSecondarySubnetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secondarySubnetId",
		val,
	)
}

func (j *jsiiProxy_InstanceSecondaryNetworkInterfaceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_InstanceSecondaryNetworkInterfaceOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_InstanceSecondaryNetworkInterfaceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstanceSecondaryNetworkInterfaceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := i.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstanceSecondaryNetworkInterfaceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := i.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstanceSecondaryNetworkInterfaceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := i.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		i,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstanceSecondaryNetworkInterfaceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := i.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstanceSecondaryNetworkInterfaceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := i.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstanceSecondaryNetworkInterfaceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := i.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		i,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstanceSecondaryNetworkInterfaceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := i.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		i,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstanceSecondaryNetworkInterfaceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := i.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstanceSecondaryNetworkInterfaceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := i.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstanceSecondaryNetworkInterfaceOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstanceSecondaryNetworkInterfaceOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstanceSecondaryNetworkInterfaceOutputReference) ResetDeleteOnTermination() {
	_jsii_.InvokeVoid(
		i,
		"resetDeleteOnTermination",
		nil, // no parameters
	)
}

func (i *jsiiProxy_InstanceSecondaryNetworkInterfaceOutputReference) ResetDeviceIndex() {
	_jsii_.InvokeVoid(
		i,
		"resetDeviceIndex",
		nil, // no parameters
	)
}

func (i *jsiiProxy_InstanceSecondaryNetworkInterfaceOutputReference) ResetInterfaceType() {
	_jsii_.InvokeVoid(
		i,
		"resetInterfaceType",
		nil, // no parameters
	)
}

func (i *jsiiProxy_InstanceSecondaryNetworkInterfaceOutputReference) ResetPrivateIpAddressCount() {
	_jsii_.InvokeVoid(
		i,
		"resetPrivateIpAddressCount",
		nil, // no parameters
	)
}

func (i *jsiiProxy_InstanceSecondaryNetworkInterfaceOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := i.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		i,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstanceSecondaryNetworkInterfaceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

