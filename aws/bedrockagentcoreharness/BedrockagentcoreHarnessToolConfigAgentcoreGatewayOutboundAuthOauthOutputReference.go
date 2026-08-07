// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreharness

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/jsii"

	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/bedrockagentcoreharness/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOauthOutputReference interface {
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
	CustomParameters() *map[string]*string
	SetCustomParameters(val *map[string]*string)
	CustomParametersInput() *map[string]*string
	DefaultReturnUrl() *string
	SetDefaultReturnUrl(val *string)
	DefaultReturnUrlInput() *string
	// Experimental.
	Fqn() *string
	GrantType() *string
	SetGrantType(val *string)
	GrantTypeInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	ProviderArn() *string
	SetProviderArn(val *string)
	ProviderArnInput() *string
	Scopes() *[]*string
	SetScopes(val *[]*string)
	ScopesInput() *[]*string
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
	ResetCustomParameters()
	ResetDefaultReturnUrl()
	ResetGrantType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOauthOutputReference
type jsiiProxy_BedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOauthOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_BedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOauthOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOauthOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOauthOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOauthOutputReference) CustomParameters() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"customParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOauthOutputReference) CustomParametersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"customParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOauthOutputReference) DefaultReturnUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultReturnUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOauthOutputReference) DefaultReturnUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultReturnUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOauthOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOauthOutputReference) GrantType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"grantType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOauthOutputReference) GrantTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"grantTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOauthOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOauthOutputReference) ProviderArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOauthOutputReference) ProviderArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOauthOutputReference) Scopes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"scopes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOauthOutputReference) ScopesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"scopesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOauthOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOauthOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOauthOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) BedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOauthOutputReference {
	_init_.Initialize()

	if err := validateNewBedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOauthOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_BedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOauthOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-aws.bedrockagentcoreHarness.BedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOauthOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewBedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOauthOutputReference_Override(b BedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOauthOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.bedrockagentcoreHarness.BedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOauthOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		b,
	)
}

func (j *jsiiProxy_BedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOauthOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOauthOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOauthOutputReference)SetCustomParameters(val *map[string]*string) {
	if err := j.validateSetCustomParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customParameters",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOauthOutputReference)SetDefaultReturnUrl(val *string) {
	if err := j.validateSetDefaultReturnUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultReturnUrl",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOauthOutputReference)SetGrantType(val *string) {
	if err := j.validateSetGrantTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"grantType",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOauthOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOauthOutputReference)SetProviderArn(val *string) {
	if err := j.validateSetProviderArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"providerArn",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOauthOutputReference)SetScopes(val *[]*string) {
	if err := j.validateSetScopesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scopes",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOauthOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOauthOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOauthOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOauthOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := b.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOauthOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := b.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOauthOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := b.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOauthOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := b.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOauthOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := b.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOauthOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := b.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOauthOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := b.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOauthOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := b.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOauthOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := b.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOauthOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOauthOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOauthOutputReference) ResetCustomParameters() {
	_jsii_.InvokeVoid(
		b,
		"resetCustomParameters",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOauthOutputReference) ResetDefaultReturnUrl() {
	_jsii_.InvokeVoid(
		b,
		"resetDefaultReturnUrl",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOauthOutputReference) ResetGrantType() {
	_jsii_.InvokeVoid(
		b,
		"resetGrantType",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOauthOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := b.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOauthOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

