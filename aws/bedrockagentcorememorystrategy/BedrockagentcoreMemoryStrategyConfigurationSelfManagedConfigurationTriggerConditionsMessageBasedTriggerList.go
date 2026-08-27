// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcorememorystrategy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/jsii"

	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/bedrockagentcorememorystrategy/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BedrockagentcoreMemoryStrategyConfigurationSelfManagedConfigurationTriggerConditionsMessageBasedTriggerList interface {
	cdktn.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	WrapsSet() *bool
	// Experimental.
	SetWrapsSet(val *bool)
	// Creating an iterator for this complex list.
	//
	// The list will be converted into a map with the mapKeyAttributeName as the key.
	// Experimental.
	AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) BedrockagentcoreMemoryStrategyConfigurationSelfManagedConfigurationTriggerConditionsMessageBasedTriggerOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BedrockagentcoreMemoryStrategyConfigurationSelfManagedConfigurationTriggerConditionsMessageBasedTriggerList
type jsiiProxy_BedrockagentcoreMemoryStrategyConfigurationSelfManagedConfigurationTriggerConditionsMessageBasedTriggerList struct {
	internal.Type__cdktnComplexList
}

func (j *jsiiProxy_BedrockagentcoreMemoryStrategyConfigurationSelfManagedConfigurationTriggerConditionsMessageBasedTriggerList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreMemoryStrategyConfigurationSelfManagedConfigurationTriggerConditionsMessageBasedTriggerList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreMemoryStrategyConfigurationSelfManagedConfigurationTriggerConditionsMessageBasedTriggerList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreMemoryStrategyConfigurationSelfManagedConfigurationTriggerConditionsMessageBasedTriggerList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreMemoryStrategyConfigurationSelfManagedConfigurationTriggerConditionsMessageBasedTriggerList) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreMemoryStrategyConfigurationSelfManagedConfigurationTriggerConditionsMessageBasedTriggerList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewBedrockagentcoreMemoryStrategyConfigurationSelfManagedConfigurationTriggerConditionsMessageBasedTriggerList(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) BedrockagentcoreMemoryStrategyConfigurationSelfManagedConfigurationTriggerConditionsMessageBasedTriggerList {
	_init_.Initialize()

	if err := validateNewBedrockagentcoreMemoryStrategyConfigurationSelfManagedConfigurationTriggerConditionsMessageBasedTriggerListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_BedrockagentcoreMemoryStrategyConfigurationSelfManagedConfigurationTriggerConditionsMessageBasedTriggerList{}

	_jsii_.Create(
		"@cdktn/provider-aws.bedrockagentcoreMemoryStrategy.BedrockagentcoreMemoryStrategyConfigurationSelfManagedConfigurationTriggerConditionsMessageBasedTriggerList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewBedrockagentcoreMemoryStrategyConfigurationSelfManagedConfigurationTriggerConditionsMessageBasedTriggerList_Override(b BedrockagentcoreMemoryStrategyConfigurationSelfManagedConfigurationTriggerConditionsMessageBasedTriggerList, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.bedrockagentcoreMemoryStrategy.BedrockagentcoreMemoryStrategyConfigurationSelfManagedConfigurationTriggerConditionsMessageBasedTriggerList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		b,
	)
}

func (j *jsiiProxy_BedrockagentcoreMemoryStrategyConfigurationSelfManagedConfigurationTriggerConditionsMessageBasedTriggerList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreMemoryStrategyConfigurationSelfManagedConfigurationTriggerConditionsMessageBasedTriggerList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreMemoryStrategyConfigurationSelfManagedConfigurationTriggerConditionsMessageBasedTriggerList)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreMemoryStrategyConfigurationSelfManagedConfigurationTriggerConditionsMessageBasedTriggerList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (b *jsiiProxy_BedrockagentcoreMemoryStrategyConfigurationSelfManagedConfigurationTriggerConditionsMessageBasedTriggerList) AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator {
	if err := b.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktn.DynamicListTerraformIterator

	_jsii_.Invoke(
		b,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreMemoryStrategyConfigurationSelfManagedConfigurationTriggerConditionsMessageBasedTriggerList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreMemoryStrategyConfigurationSelfManagedConfigurationTriggerConditionsMessageBasedTriggerList) Get(index *float64) BedrockagentcoreMemoryStrategyConfigurationSelfManagedConfigurationTriggerConditionsMessageBasedTriggerOutputReference {
	if err := b.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns BedrockagentcoreMemoryStrategyConfigurationSelfManagedConfigurationTriggerConditionsMessageBasedTriggerOutputReference

	_jsii_.Invoke(
		b,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreMemoryStrategyConfigurationSelfManagedConfigurationTriggerConditionsMessageBasedTriggerList) Resolve(context cdktn.IResolveContext) interface{} {
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

func (b *jsiiProxy_BedrockagentcoreMemoryStrategyConfigurationSelfManagedConfigurationTriggerConditionsMessageBasedTriggerList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

