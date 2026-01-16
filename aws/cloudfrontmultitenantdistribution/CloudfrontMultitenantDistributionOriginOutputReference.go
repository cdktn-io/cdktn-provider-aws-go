// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudfrontmultitenantdistribution

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v21/jsii"

	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v21/cloudfrontmultitenantdistribution/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudfrontMultitenantDistributionOriginOutputReference interface {
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
	ConnectionAttempts() *float64
	SetConnectionAttempts(val *float64)
	ConnectionAttemptsInput() *float64
	ConnectionTimeout() *float64
	SetConnectionTimeout(val *float64)
	ConnectionTimeoutInput() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	CustomHeader() CloudfrontMultitenantDistributionOriginCustomHeaderList
	CustomHeaderInput() interface{}
	CustomOriginConfig() CloudfrontMultitenantDistributionOriginCustomOriginConfigList
	CustomOriginConfigInput() interface{}
	DomainName() *string
	SetDomainName(val *string)
	DomainNameInput() *string
	// Experimental.
	Fqn() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	OriginAccessControlId() *string
	SetOriginAccessControlId(val *string)
	OriginAccessControlIdInput() *string
	OriginPath() *string
	SetOriginPath(val *string)
	OriginPathInput() *string
	OriginShield() CloudfrontMultitenantDistributionOriginOriginShieldList
	OriginShieldInput() interface{}
	ResponseCompletionTimeout() *float64
	SetResponseCompletionTimeout(val *float64)
	ResponseCompletionTimeoutInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VpcOriginConfig() CloudfrontMultitenantDistributionOriginVpcOriginConfigList
	VpcOriginConfigInput() interface{}
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
	PutCustomHeader(value interface{})
	PutCustomOriginConfig(value interface{})
	PutOriginShield(value interface{})
	PutVpcOriginConfig(value interface{})
	ResetConnectionAttempts()
	ResetConnectionTimeout()
	ResetCustomHeader()
	ResetCustomOriginConfig()
	ResetOriginAccessControlId()
	ResetOriginPath()
	ResetOriginShield()
	ResetResponseCompletionTimeout()
	ResetVpcOriginConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CloudfrontMultitenantDistributionOriginOutputReference
type jsiiProxy_CloudfrontMultitenantDistributionOriginOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionOriginOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionOriginOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionOriginOutputReference) ConnectionAttempts() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"connectionAttempts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionOriginOutputReference) ConnectionAttemptsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"connectionAttemptsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionOriginOutputReference) ConnectionTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"connectionTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionOriginOutputReference) ConnectionTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"connectionTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionOriginOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionOriginOutputReference) CustomHeader() CloudfrontMultitenantDistributionOriginCustomHeaderList {
	var returns CloudfrontMultitenantDistributionOriginCustomHeaderList
	_jsii_.Get(
		j,
		"customHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionOriginOutputReference) CustomHeaderInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionOriginOutputReference) CustomOriginConfig() CloudfrontMultitenantDistributionOriginCustomOriginConfigList {
	var returns CloudfrontMultitenantDistributionOriginCustomOriginConfigList
	_jsii_.Get(
		j,
		"customOriginConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionOriginOutputReference) CustomOriginConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customOriginConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionOriginOutputReference) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionOriginOutputReference) DomainNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionOriginOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionOriginOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionOriginOutputReference) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionOriginOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionOriginOutputReference) OriginAccessControlId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originAccessControlId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionOriginOutputReference) OriginAccessControlIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originAccessControlIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionOriginOutputReference) OriginPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionOriginOutputReference) OriginPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionOriginOutputReference) OriginShield() CloudfrontMultitenantDistributionOriginOriginShieldList {
	var returns CloudfrontMultitenantDistributionOriginOriginShieldList
	_jsii_.Get(
		j,
		"originShield",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionOriginOutputReference) OriginShieldInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"originShieldInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionOriginOutputReference) ResponseCompletionTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"responseCompletionTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionOriginOutputReference) ResponseCompletionTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"responseCompletionTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionOriginOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionOriginOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionOriginOutputReference) VpcOriginConfig() CloudfrontMultitenantDistributionOriginVpcOriginConfigList {
	var returns CloudfrontMultitenantDistributionOriginVpcOriginConfigList
	_jsii_.Get(
		j,
		"vpcOriginConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionOriginOutputReference) VpcOriginConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vpcOriginConfigInput",
		&returns,
	)
	return returns
}


func NewCloudfrontMultitenantDistributionOriginOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) CloudfrontMultitenantDistributionOriginOutputReference {
	_init_.Initialize()

	if err := validateNewCloudfrontMultitenantDistributionOriginOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudfrontMultitenantDistributionOriginOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-aws.cloudfrontMultitenantDistribution.CloudfrontMultitenantDistributionOriginOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewCloudfrontMultitenantDistributionOriginOutputReference_Override(c CloudfrontMultitenantDistributionOriginOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.cloudfrontMultitenantDistribution.CloudfrontMultitenantDistributionOriginOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionOriginOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionOriginOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionOriginOutputReference)SetConnectionAttempts(val *float64) {
	if err := j.validateSetConnectionAttemptsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionAttempts",
		val,
	)
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionOriginOutputReference)SetConnectionTimeout(val *float64) {
	if err := j.validateSetConnectionTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionTimeout",
		val,
	)
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionOriginOutputReference)SetDomainName(val *string) {
	if err := j.validateSetDomainNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainName",
		val,
	)
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionOriginOutputReference)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionOriginOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionOriginOutputReference)SetOriginAccessControlId(val *string) {
	if err := j.validateSetOriginAccessControlIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"originAccessControlId",
		val,
	)
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionOriginOutputReference)SetOriginPath(val *string) {
	if err := j.validateSetOriginPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"originPath",
		val,
	)
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionOriginOutputReference)SetResponseCompletionTimeout(val *float64) {
	if err := j.validateSetResponseCompletionTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"responseCompletionTimeout",
		val,
	)
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionOriginOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionOriginOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CloudfrontMultitenantDistributionOriginOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontMultitenantDistributionOriginOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CloudfrontMultitenantDistributionOriginOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudfrontMultitenantDistributionOriginOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CloudfrontMultitenantDistributionOriginOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CloudfrontMultitenantDistributionOriginOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CloudfrontMultitenantDistributionOriginOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CloudfrontMultitenantDistributionOriginOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CloudfrontMultitenantDistributionOriginOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CloudfrontMultitenantDistributionOriginOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CloudfrontMultitenantDistributionOriginOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontMultitenantDistributionOriginOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudfrontMultitenantDistributionOriginOutputReference) PutCustomHeader(value interface{}) {
	if err := c.validatePutCustomHeaderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putCustomHeader",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontMultitenantDistributionOriginOutputReference) PutCustomOriginConfig(value interface{}) {
	if err := c.validatePutCustomOriginConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putCustomOriginConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontMultitenantDistributionOriginOutputReference) PutOriginShield(value interface{}) {
	if err := c.validatePutOriginShieldParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putOriginShield",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontMultitenantDistributionOriginOutputReference) PutVpcOriginConfig(value interface{}) {
	if err := c.validatePutVpcOriginConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putVpcOriginConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontMultitenantDistributionOriginOutputReference) ResetConnectionAttempts() {
	_jsii_.InvokeVoid(
		c,
		"resetConnectionAttempts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontMultitenantDistributionOriginOutputReference) ResetConnectionTimeout() {
	_jsii_.InvokeVoid(
		c,
		"resetConnectionTimeout",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontMultitenantDistributionOriginOutputReference) ResetCustomHeader() {
	_jsii_.InvokeVoid(
		c,
		"resetCustomHeader",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontMultitenantDistributionOriginOutputReference) ResetCustomOriginConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetCustomOriginConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontMultitenantDistributionOriginOutputReference) ResetOriginAccessControlId() {
	_jsii_.InvokeVoid(
		c,
		"resetOriginAccessControlId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontMultitenantDistributionOriginOutputReference) ResetOriginPath() {
	_jsii_.InvokeVoid(
		c,
		"resetOriginPath",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontMultitenantDistributionOriginOutputReference) ResetOriginShield() {
	_jsii_.InvokeVoid(
		c,
		"resetOriginShield",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontMultitenantDistributionOriginOutputReference) ResetResponseCompletionTimeout() {
	_jsii_.InvokeVoid(
		c,
		"resetResponseCompletionTimeout",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontMultitenantDistributionOriginOutputReference) ResetVpcOriginConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetVpcOriginConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontMultitenantDistributionOriginOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_CloudfrontMultitenantDistributionOriginOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

