// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcorememorystrategy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/jsii"

	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/bedrockagentcorememorystrategy/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BedrockagentcoreMemoryStrategyConfigurationSelfManagedConfigurationTriggerConditionsOutputReference interface {
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
	MessageBasedTrigger() BedrockagentcoreMemoryStrategyConfigurationSelfManagedConfigurationTriggerConditionsMessageBasedTriggerList
	MessageBasedTriggerInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TimeBasedTrigger() BedrockagentcoreMemoryStrategyConfigurationSelfManagedConfigurationTriggerConditionsTimeBasedTriggerList
	TimeBasedTriggerInput() interface{}
	TokenBasedTrigger() BedrockagentcoreMemoryStrategyConfigurationSelfManagedConfigurationTriggerConditionsTokenBasedTriggerList
	TokenBasedTriggerInput() interface{}
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
	PutMessageBasedTrigger(value interface{})
	PutTimeBasedTrigger(value interface{})
	PutTokenBasedTrigger(value interface{})
	ResetMessageBasedTrigger()
	ResetTimeBasedTrigger()
	ResetTokenBasedTrigger()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BedrockagentcoreMemoryStrategyConfigurationSelfManagedConfigurationTriggerConditionsOutputReference
type jsiiProxy_BedrockagentcoreMemoryStrategyConfigurationSelfManagedConfigurationTriggerConditionsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_BedrockagentcoreMemoryStrategyConfigurationSelfManagedConfigurationTriggerConditionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreMemoryStrategyConfigurationSelfManagedConfigurationTriggerConditionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreMemoryStrategyConfigurationSelfManagedConfigurationTriggerConditionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreMemoryStrategyConfigurationSelfManagedConfigurationTriggerConditionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreMemoryStrategyConfigurationSelfManagedConfigurationTriggerConditionsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreMemoryStrategyConfigurationSelfManagedConfigurationTriggerConditionsOutputReference) MessageBasedTrigger() BedrockagentcoreMemoryStrategyConfigurationSelfManagedConfigurationTriggerConditionsMessageBasedTriggerList {
	var returns BedrockagentcoreMemoryStrategyConfigurationSelfManagedConfigurationTriggerConditionsMessageBasedTriggerList
	_jsii_.Get(
		j,
		"messageBasedTrigger",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreMemoryStrategyConfigurationSelfManagedConfigurationTriggerConditionsOutputReference) MessageBasedTriggerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"messageBasedTriggerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreMemoryStrategyConfigurationSelfManagedConfigurationTriggerConditionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreMemoryStrategyConfigurationSelfManagedConfigurationTriggerConditionsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreMemoryStrategyConfigurationSelfManagedConfigurationTriggerConditionsOutputReference) TimeBasedTrigger() BedrockagentcoreMemoryStrategyConfigurationSelfManagedConfigurationTriggerConditionsTimeBasedTriggerList {
	var returns BedrockagentcoreMemoryStrategyConfigurationSelfManagedConfigurationTriggerConditionsTimeBasedTriggerList
	_jsii_.Get(
		j,
		"timeBasedTrigger",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreMemoryStrategyConfigurationSelfManagedConfigurationTriggerConditionsOutputReference) TimeBasedTriggerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeBasedTriggerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreMemoryStrategyConfigurationSelfManagedConfigurationTriggerConditionsOutputReference) TokenBasedTrigger() BedrockagentcoreMemoryStrategyConfigurationSelfManagedConfigurationTriggerConditionsTokenBasedTriggerList {
	var returns BedrockagentcoreMemoryStrategyConfigurationSelfManagedConfigurationTriggerConditionsTokenBasedTriggerList
	_jsii_.Get(
		j,
		"tokenBasedTrigger",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreMemoryStrategyConfigurationSelfManagedConfigurationTriggerConditionsOutputReference) TokenBasedTriggerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tokenBasedTriggerInput",
		&returns,
	)
	return returns
}


func NewBedrockagentcoreMemoryStrategyConfigurationSelfManagedConfigurationTriggerConditionsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) BedrockagentcoreMemoryStrategyConfigurationSelfManagedConfigurationTriggerConditionsOutputReference {
	_init_.Initialize()

	if err := validateNewBedrockagentcoreMemoryStrategyConfigurationSelfManagedConfigurationTriggerConditionsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_BedrockagentcoreMemoryStrategyConfigurationSelfManagedConfigurationTriggerConditionsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-aws.bedrockagentcoreMemoryStrategy.BedrockagentcoreMemoryStrategyConfigurationSelfManagedConfigurationTriggerConditionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewBedrockagentcoreMemoryStrategyConfigurationSelfManagedConfigurationTriggerConditionsOutputReference_Override(b BedrockagentcoreMemoryStrategyConfigurationSelfManagedConfigurationTriggerConditionsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.bedrockagentcoreMemoryStrategy.BedrockagentcoreMemoryStrategyConfigurationSelfManagedConfigurationTriggerConditionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		b,
	)
}

func (j *jsiiProxy_BedrockagentcoreMemoryStrategyConfigurationSelfManagedConfigurationTriggerConditionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreMemoryStrategyConfigurationSelfManagedConfigurationTriggerConditionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreMemoryStrategyConfigurationSelfManagedConfigurationTriggerConditionsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreMemoryStrategyConfigurationSelfManagedConfigurationTriggerConditionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreMemoryStrategyConfigurationSelfManagedConfigurationTriggerConditionsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BedrockagentcoreMemoryStrategyConfigurationSelfManagedConfigurationTriggerConditionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreMemoryStrategyConfigurationSelfManagedConfigurationTriggerConditionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (b *jsiiProxy_BedrockagentcoreMemoryStrategyConfigurationSelfManagedConfigurationTriggerConditionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (b *jsiiProxy_BedrockagentcoreMemoryStrategyConfigurationSelfManagedConfigurationTriggerConditionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (b *jsiiProxy_BedrockagentcoreMemoryStrategyConfigurationSelfManagedConfigurationTriggerConditionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (b *jsiiProxy_BedrockagentcoreMemoryStrategyConfigurationSelfManagedConfigurationTriggerConditionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (b *jsiiProxy_BedrockagentcoreMemoryStrategyConfigurationSelfManagedConfigurationTriggerConditionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (b *jsiiProxy_BedrockagentcoreMemoryStrategyConfigurationSelfManagedConfigurationTriggerConditionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (b *jsiiProxy_BedrockagentcoreMemoryStrategyConfigurationSelfManagedConfigurationTriggerConditionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (b *jsiiProxy_BedrockagentcoreMemoryStrategyConfigurationSelfManagedConfigurationTriggerConditionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (b *jsiiProxy_BedrockagentcoreMemoryStrategyConfigurationSelfManagedConfigurationTriggerConditionsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreMemoryStrategyConfigurationSelfManagedConfigurationTriggerConditionsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (b *jsiiProxy_BedrockagentcoreMemoryStrategyConfigurationSelfManagedConfigurationTriggerConditionsOutputReference) PutMessageBasedTrigger(value interface{}) {
	if err := b.validatePutMessageBasedTriggerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putMessageBasedTrigger",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentcoreMemoryStrategyConfigurationSelfManagedConfigurationTriggerConditionsOutputReference) PutTimeBasedTrigger(value interface{}) {
	if err := b.validatePutTimeBasedTriggerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putTimeBasedTrigger",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentcoreMemoryStrategyConfigurationSelfManagedConfigurationTriggerConditionsOutputReference) PutTokenBasedTrigger(value interface{}) {
	if err := b.validatePutTokenBasedTriggerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putTokenBasedTrigger",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentcoreMemoryStrategyConfigurationSelfManagedConfigurationTriggerConditionsOutputReference) ResetMessageBasedTrigger() {
	_jsii_.InvokeVoid(
		b,
		"resetMessageBasedTrigger",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreMemoryStrategyConfigurationSelfManagedConfigurationTriggerConditionsOutputReference) ResetTimeBasedTrigger() {
	_jsii_.InvokeVoid(
		b,
		"resetTimeBasedTrigger",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreMemoryStrategyConfigurationSelfManagedConfigurationTriggerConditionsOutputReference) ResetTokenBasedTrigger() {
	_jsii_.InvokeVoid(
		b,
		"resetTokenBasedTrigger",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreMemoryStrategyConfigurationSelfManagedConfigurationTriggerConditionsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (b *jsiiProxy_BedrockagentcoreMemoryStrategyConfigurationSelfManagedConfigurationTriggerConditionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

