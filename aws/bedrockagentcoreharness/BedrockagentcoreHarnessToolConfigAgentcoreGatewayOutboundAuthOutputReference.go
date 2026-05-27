// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreharness

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v24/jsii"

	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v24/bedrockagentcoreharness/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOutputReference interface {
	cdktn.ComplexObject
	AwsIam() interface{}
	SetAwsIam(val interface{})
	AwsIamInput() interface{}
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
	None() interface{}
	SetNone(val interface{})
	NoneInput() interface{}
	Oauth() BedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOauthList
	OauthInput() interface{}
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
	PutOauth(value interface{})
	ResetAwsIam()
	ResetNone()
	ResetOauth()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOutputReference
type jsiiProxy_BedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_BedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOutputReference) AwsIam() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"awsIam",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOutputReference) AwsIamInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"awsIamInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOutputReference) None() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"none",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOutputReference) NoneInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOutputReference) Oauth() BedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOauthList {
	var returns BedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOauthList
	_jsii_.Get(
		j,
		"oauth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOutputReference) OauthInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"oauthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) BedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOutputReference {
	_init_.Initialize()

	if err := validateNewBedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_BedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-aws.bedrockagentcoreHarness.BedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewBedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOutputReference_Override(b BedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.bedrockagentcoreHarness.BedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		b,
	)
}

func (j *jsiiProxy_BedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOutputReference)SetAwsIam(val interface{}) {
	if err := j.validateSetAwsIamParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"awsIam",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOutputReference)SetNone(val interface{}) {
	if err := j.validateSetNoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"none",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (b *jsiiProxy_BedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (b *jsiiProxy_BedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (b *jsiiProxy_BedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (b *jsiiProxy_BedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (b *jsiiProxy_BedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (b *jsiiProxy_BedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (b *jsiiProxy_BedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (b *jsiiProxy_BedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (b *jsiiProxy_BedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (b *jsiiProxy_BedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOutputReference) PutOauth(value interface{}) {
	if err := b.validatePutOauthParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putOauth",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOutputReference) ResetAwsIam() {
	_jsii_.InvokeVoid(
		b,
		"resetAwsIam",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOutputReference) ResetNone() {
	_jsii_.InvokeVoid(
		b,
		"resetNone",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOutputReference) ResetOauth() {
	_jsii_.InvokeVoid(
		b,
		"resetOauth",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (b *jsiiProxy_BedrockagentcoreHarnessToolConfigAgentcoreGatewayOutboundAuthOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

