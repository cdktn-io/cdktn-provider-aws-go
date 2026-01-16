// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudfrontmultitenantdistribution

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v21/jsii"

	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v21/cloudfrontmultitenantdistribution/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudfrontMultitenantDistributionOriginCustomOriginConfigOutputReference interface {
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
	HttpPort() *float64
	SetHttpPort(val *float64)
	HttpPortInput() *float64
	HttpsPort() *float64
	SetHttpsPort(val *float64)
	HttpsPortInput() *float64
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IpAddressType() *string
	SetIpAddressType(val *string)
	IpAddressTypeInput() *string
	OriginKeepaliveTimeout() *float64
	SetOriginKeepaliveTimeout(val *float64)
	OriginKeepaliveTimeoutInput() *float64
	OriginProtocolPolicy() *string
	SetOriginProtocolPolicy(val *string)
	OriginProtocolPolicyInput() *string
	OriginReadTimeout() *float64
	SetOriginReadTimeout(val *float64)
	OriginReadTimeoutInput() *float64
	OriginSslProtocols() *[]*string
	SetOriginSslProtocols(val *[]*string)
	OriginSslProtocolsInput() *[]*string
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
	ResetIpAddressType()
	ResetOriginKeepaliveTimeout()
	ResetOriginReadTimeout()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CloudfrontMultitenantDistributionOriginCustomOriginConfigOutputReference
type jsiiProxy_CloudfrontMultitenantDistributionOriginCustomOriginConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionOriginCustomOriginConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionOriginCustomOriginConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionOriginCustomOriginConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionOriginCustomOriginConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionOriginCustomOriginConfigOutputReference) HttpPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionOriginCustomOriginConfigOutputReference) HttpPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionOriginCustomOriginConfigOutputReference) HttpsPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpsPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionOriginCustomOriginConfigOutputReference) HttpsPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpsPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionOriginCustomOriginConfigOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionOriginCustomOriginConfigOutputReference) IpAddressType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddressType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionOriginCustomOriginConfigOutputReference) IpAddressTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddressTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionOriginCustomOriginConfigOutputReference) OriginKeepaliveTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"originKeepaliveTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionOriginCustomOriginConfigOutputReference) OriginKeepaliveTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"originKeepaliveTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionOriginCustomOriginConfigOutputReference) OriginProtocolPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originProtocolPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionOriginCustomOriginConfigOutputReference) OriginProtocolPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originProtocolPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionOriginCustomOriginConfigOutputReference) OriginReadTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"originReadTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionOriginCustomOriginConfigOutputReference) OriginReadTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"originReadTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionOriginCustomOriginConfigOutputReference) OriginSslProtocols() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"originSslProtocols",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionOriginCustomOriginConfigOutputReference) OriginSslProtocolsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"originSslProtocolsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionOriginCustomOriginConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionOriginCustomOriginConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCloudfrontMultitenantDistributionOriginCustomOriginConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) CloudfrontMultitenantDistributionOriginCustomOriginConfigOutputReference {
	_init_.Initialize()

	if err := validateNewCloudfrontMultitenantDistributionOriginCustomOriginConfigOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudfrontMultitenantDistributionOriginCustomOriginConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-aws.cloudfrontMultitenantDistribution.CloudfrontMultitenantDistributionOriginCustomOriginConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewCloudfrontMultitenantDistributionOriginCustomOriginConfigOutputReference_Override(c CloudfrontMultitenantDistributionOriginCustomOriginConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.cloudfrontMultitenantDistribution.CloudfrontMultitenantDistributionOriginCustomOriginConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionOriginCustomOriginConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionOriginCustomOriginConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionOriginCustomOriginConfigOutputReference)SetHttpPort(val *float64) {
	if err := j.validateSetHttpPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpPort",
		val,
	)
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionOriginCustomOriginConfigOutputReference)SetHttpsPort(val *float64) {
	if err := j.validateSetHttpsPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpsPort",
		val,
	)
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionOriginCustomOriginConfigOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionOriginCustomOriginConfigOutputReference)SetIpAddressType(val *string) {
	if err := j.validateSetIpAddressTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipAddressType",
		val,
	)
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionOriginCustomOriginConfigOutputReference)SetOriginKeepaliveTimeout(val *float64) {
	if err := j.validateSetOriginKeepaliveTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"originKeepaliveTimeout",
		val,
	)
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionOriginCustomOriginConfigOutputReference)SetOriginProtocolPolicy(val *string) {
	if err := j.validateSetOriginProtocolPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"originProtocolPolicy",
		val,
	)
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionOriginCustomOriginConfigOutputReference)SetOriginReadTimeout(val *float64) {
	if err := j.validateSetOriginReadTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"originReadTimeout",
		val,
	)
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionOriginCustomOriginConfigOutputReference)SetOriginSslProtocols(val *[]*string) {
	if err := j.validateSetOriginSslProtocolsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"originSslProtocols",
		val,
	)
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionOriginCustomOriginConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionOriginCustomOriginConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CloudfrontMultitenantDistributionOriginCustomOriginConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontMultitenantDistributionOriginCustomOriginConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontMultitenantDistributionOriginCustomOriginConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontMultitenantDistributionOriginCustomOriginConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontMultitenantDistributionOriginCustomOriginConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontMultitenantDistributionOriginCustomOriginConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontMultitenantDistributionOriginCustomOriginConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontMultitenantDistributionOriginCustomOriginConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontMultitenantDistributionOriginCustomOriginConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontMultitenantDistributionOriginCustomOriginConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontMultitenantDistributionOriginCustomOriginConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontMultitenantDistributionOriginCustomOriginConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontMultitenantDistributionOriginCustomOriginConfigOutputReference) ResetIpAddressType() {
	_jsii_.InvokeVoid(
		c,
		"resetIpAddressType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontMultitenantDistributionOriginCustomOriginConfigOutputReference) ResetOriginKeepaliveTimeout() {
	_jsii_.InvokeVoid(
		c,
		"resetOriginKeepaliveTimeout",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontMultitenantDistributionOriginCustomOriginConfigOutputReference) ResetOriginReadTimeout() {
	_jsii_.InvokeVoid(
		c,
		"resetOriginReadTimeout",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontMultitenantDistributionOriginCustomOriginConfigOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := c.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontMultitenantDistributionOriginCustomOriginConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

