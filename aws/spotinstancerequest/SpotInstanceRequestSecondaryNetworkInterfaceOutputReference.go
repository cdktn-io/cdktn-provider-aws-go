// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package spotinstancerequest

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v22/jsii"

	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v22/spotinstancerequest/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SpotInstanceRequestSecondaryNetworkInterfaceOutputReference interface {
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
	SourceDestCheck() cdktf.IResolvable
	Status() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	ResetDeleteOnTermination()
	ResetDeviceIndex()
	ResetInterfaceType()
	ResetPrivateIpAddressCount()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SpotInstanceRequestSecondaryNetworkInterfaceOutputReference
type jsiiProxy_SpotInstanceRequestSecondaryNetworkInterfaceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SpotInstanceRequestSecondaryNetworkInterfaceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequestSecondaryNetworkInterfaceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequestSecondaryNetworkInterfaceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequestSecondaryNetworkInterfaceOutputReference) DeleteOnTermination() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteOnTermination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequestSecondaryNetworkInterfaceOutputReference) DeleteOnTerminationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteOnTerminationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequestSecondaryNetworkInterfaceOutputReference) DeviceIndex() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deviceIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequestSecondaryNetworkInterfaceOutputReference) DeviceIndexInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deviceIndexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequestSecondaryNetworkInterfaceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequestSecondaryNetworkInterfaceOutputReference) InterfaceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interfaceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequestSecondaryNetworkInterfaceOutputReference) InterfaceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interfaceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequestSecondaryNetworkInterfaceOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequestSecondaryNetworkInterfaceOutputReference) MacAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"macAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequestSecondaryNetworkInterfaceOutputReference) NetworkCardIndex() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"networkCardIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequestSecondaryNetworkInterfaceOutputReference) NetworkCardIndexInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"networkCardIndexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequestSecondaryNetworkInterfaceOutputReference) PrivateIpAddressCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"privateIpAddressCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequestSecondaryNetworkInterfaceOutputReference) PrivateIpAddressCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"privateIpAddressCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequestSecondaryNetworkInterfaceOutputReference) PrivateIpAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"privateIpAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequestSecondaryNetworkInterfaceOutputReference) SecondaryInterfaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryInterfaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequestSecondaryNetworkInterfaceOutputReference) SecondaryNetworkId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryNetworkId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequestSecondaryNetworkInterfaceOutputReference) SecondarySubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondarySubnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequestSecondaryNetworkInterfaceOutputReference) SecondarySubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondarySubnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequestSecondaryNetworkInterfaceOutputReference) SourceDestCheck() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"sourceDestCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequestSecondaryNetworkInterfaceOutputReference) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequestSecondaryNetworkInterfaceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpotInstanceRequestSecondaryNetworkInterfaceOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSpotInstanceRequestSecondaryNetworkInterfaceOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) SpotInstanceRequestSecondaryNetworkInterfaceOutputReference {
	_init_.Initialize()

	if err := validateNewSpotInstanceRequestSecondaryNetworkInterfaceOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_SpotInstanceRequestSecondaryNetworkInterfaceOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-aws.spotInstanceRequest.SpotInstanceRequestSecondaryNetworkInterfaceOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewSpotInstanceRequestSecondaryNetworkInterfaceOutputReference_Override(s SpotInstanceRequestSecondaryNetworkInterfaceOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.spotInstanceRequest.SpotInstanceRequestSecondaryNetworkInterfaceOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_SpotInstanceRequestSecondaryNetworkInterfaceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SpotInstanceRequestSecondaryNetworkInterfaceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SpotInstanceRequestSecondaryNetworkInterfaceOutputReference)SetDeleteOnTermination(val interface{}) {
	if err := j.validateSetDeleteOnTerminationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteOnTermination",
		val,
	)
}

func (j *jsiiProxy_SpotInstanceRequestSecondaryNetworkInterfaceOutputReference)SetDeviceIndex(val *float64) {
	if err := j.validateSetDeviceIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deviceIndex",
		val,
	)
}

func (j *jsiiProxy_SpotInstanceRequestSecondaryNetworkInterfaceOutputReference)SetInterfaceType(val *string) {
	if err := j.validateSetInterfaceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"interfaceType",
		val,
	)
}

func (j *jsiiProxy_SpotInstanceRequestSecondaryNetworkInterfaceOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SpotInstanceRequestSecondaryNetworkInterfaceOutputReference)SetNetworkCardIndex(val *float64) {
	if err := j.validateSetNetworkCardIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkCardIndex",
		val,
	)
}

func (j *jsiiProxy_SpotInstanceRequestSecondaryNetworkInterfaceOutputReference)SetPrivateIpAddressCount(val *float64) {
	if err := j.validateSetPrivateIpAddressCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateIpAddressCount",
		val,
	)
}

func (j *jsiiProxy_SpotInstanceRequestSecondaryNetworkInterfaceOutputReference)SetSecondarySubnetId(val *string) {
	if err := j.validateSetSecondarySubnetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secondarySubnetId",
		val,
	)
}

func (j *jsiiProxy_SpotInstanceRequestSecondaryNetworkInterfaceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SpotInstanceRequestSecondaryNetworkInterfaceOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SpotInstanceRequestSecondaryNetworkInterfaceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpotInstanceRequestSecondaryNetworkInterfaceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SpotInstanceRequestSecondaryNetworkInterfaceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SpotInstanceRequestSecondaryNetworkInterfaceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SpotInstanceRequestSecondaryNetworkInterfaceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SpotInstanceRequestSecondaryNetworkInterfaceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SpotInstanceRequestSecondaryNetworkInterfaceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SpotInstanceRequestSecondaryNetworkInterfaceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SpotInstanceRequestSecondaryNetworkInterfaceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SpotInstanceRequestSecondaryNetworkInterfaceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SpotInstanceRequestSecondaryNetworkInterfaceOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpotInstanceRequestSecondaryNetworkInterfaceOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SpotInstanceRequestSecondaryNetworkInterfaceOutputReference) ResetDeleteOnTermination() {
	_jsii_.InvokeVoid(
		s,
		"resetDeleteOnTermination",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotInstanceRequestSecondaryNetworkInterfaceOutputReference) ResetDeviceIndex() {
	_jsii_.InvokeVoid(
		s,
		"resetDeviceIndex",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotInstanceRequestSecondaryNetworkInterfaceOutputReference) ResetInterfaceType() {
	_jsii_.InvokeVoid(
		s,
		"resetInterfaceType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotInstanceRequestSecondaryNetworkInterfaceOutputReference) ResetPrivateIpAddressCount() {
	_jsii_.InvokeVoid(
		s,
		"resetPrivateIpAddressCount",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpotInstanceRequestSecondaryNetworkInterfaceOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_SpotInstanceRequestSecondaryNetworkInterfaceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

