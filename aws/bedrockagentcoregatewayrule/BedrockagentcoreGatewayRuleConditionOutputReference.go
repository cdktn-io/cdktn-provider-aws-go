// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoregatewayrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/jsii"

	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/bedrockagentcoregatewayrule/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BedrockagentcoreGatewayRuleConditionOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MatchPaths() BedrockagentcoreGatewayRuleConditionMatchPathsList
	MatchPathsInput() interface{}
	MatchPrincipals() BedrockagentcoreGatewayRuleConditionMatchPrincipalsList
	MatchPrincipalsInput() interface{}
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
	PutMatchPaths(value interface{})
	PutMatchPrincipals(value interface{})
	ResetMatchPaths()
	ResetMatchPrincipals()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BedrockagentcoreGatewayRuleConditionOutputReference
type jsiiProxy_BedrockagentcoreGatewayRuleConditionOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_BedrockagentcoreGatewayRuleConditionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreGatewayRuleConditionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreGatewayRuleConditionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreGatewayRuleConditionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreGatewayRuleConditionOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreGatewayRuleConditionOutputReference) MatchPaths() BedrockagentcoreGatewayRuleConditionMatchPathsList {
	var returns BedrockagentcoreGatewayRuleConditionMatchPathsList
	_jsii_.Get(
		j,
		"matchPaths",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreGatewayRuleConditionOutputReference) MatchPathsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"matchPathsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreGatewayRuleConditionOutputReference) MatchPrincipals() BedrockagentcoreGatewayRuleConditionMatchPrincipalsList {
	var returns BedrockagentcoreGatewayRuleConditionMatchPrincipalsList
	_jsii_.Get(
		j,
		"matchPrincipals",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreGatewayRuleConditionOutputReference) MatchPrincipalsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"matchPrincipalsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreGatewayRuleConditionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreGatewayRuleConditionOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBedrockagentcoreGatewayRuleConditionOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) BedrockagentcoreGatewayRuleConditionOutputReference {
	_init_.Initialize()

	if err := validateNewBedrockagentcoreGatewayRuleConditionOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_BedrockagentcoreGatewayRuleConditionOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-aws.bedrockagentcoreGatewayRule.BedrockagentcoreGatewayRuleConditionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewBedrockagentcoreGatewayRuleConditionOutputReference_Override(b BedrockagentcoreGatewayRuleConditionOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.bedrockagentcoreGatewayRule.BedrockagentcoreGatewayRuleConditionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		b,
	)
}

func (j *jsiiProxy_BedrockagentcoreGatewayRuleConditionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreGatewayRuleConditionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreGatewayRuleConditionOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreGatewayRuleConditionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreGatewayRuleConditionOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BedrockagentcoreGatewayRuleConditionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreGatewayRuleConditionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (b *jsiiProxy_BedrockagentcoreGatewayRuleConditionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (b *jsiiProxy_BedrockagentcoreGatewayRuleConditionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (b *jsiiProxy_BedrockagentcoreGatewayRuleConditionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (b *jsiiProxy_BedrockagentcoreGatewayRuleConditionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (b *jsiiProxy_BedrockagentcoreGatewayRuleConditionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (b *jsiiProxy_BedrockagentcoreGatewayRuleConditionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (b *jsiiProxy_BedrockagentcoreGatewayRuleConditionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (b *jsiiProxy_BedrockagentcoreGatewayRuleConditionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (b *jsiiProxy_BedrockagentcoreGatewayRuleConditionOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreGatewayRuleConditionOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (b *jsiiProxy_BedrockagentcoreGatewayRuleConditionOutputReference) PutMatchPaths(value interface{}) {
	if err := b.validatePutMatchPathsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putMatchPaths",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentcoreGatewayRuleConditionOutputReference) PutMatchPrincipals(value interface{}) {
	if err := b.validatePutMatchPrincipalsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putMatchPrincipals",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentcoreGatewayRuleConditionOutputReference) ResetMatchPaths() {
	_jsii_.InvokeVoid(
		b,
		"resetMatchPaths",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreGatewayRuleConditionOutputReference) ResetMatchPrincipals() {
	_jsii_.InvokeVoid(
		b,
		"resetMatchPrincipals",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreGatewayRuleConditionOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (b *jsiiProxy_BedrockagentcoreGatewayRuleConditionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

