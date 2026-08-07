// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentdatasource

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/jsii"

	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/bedrockagentdatasource/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BedrockagentDataSourceVectorIngestionConfigurationParsingConfigurationBedrockDataAutomationConfigurationList interface {
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
	Get(index *float64) BedrockagentDataSourceVectorIngestionConfigurationParsingConfigurationBedrockDataAutomationConfigurationOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BedrockagentDataSourceVectorIngestionConfigurationParsingConfigurationBedrockDataAutomationConfigurationList
type jsiiProxy_BedrockagentDataSourceVectorIngestionConfigurationParsingConfigurationBedrockDataAutomationConfigurationList struct {
	internal.Type__cdktnComplexList
}

func (j *jsiiProxy_BedrockagentDataSourceVectorIngestionConfigurationParsingConfigurationBedrockDataAutomationConfigurationList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentDataSourceVectorIngestionConfigurationParsingConfigurationBedrockDataAutomationConfigurationList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentDataSourceVectorIngestionConfigurationParsingConfigurationBedrockDataAutomationConfigurationList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentDataSourceVectorIngestionConfigurationParsingConfigurationBedrockDataAutomationConfigurationList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentDataSourceVectorIngestionConfigurationParsingConfigurationBedrockDataAutomationConfigurationList) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentDataSourceVectorIngestionConfigurationParsingConfigurationBedrockDataAutomationConfigurationList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewBedrockagentDataSourceVectorIngestionConfigurationParsingConfigurationBedrockDataAutomationConfigurationList(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) BedrockagentDataSourceVectorIngestionConfigurationParsingConfigurationBedrockDataAutomationConfigurationList {
	_init_.Initialize()

	if err := validateNewBedrockagentDataSourceVectorIngestionConfigurationParsingConfigurationBedrockDataAutomationConfigurationListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_BedrockagentDataSourceVectorIngestionConfigurationParsingConfigurationBedrockDataAutomationConfigurationList{}

	_jsii_.Create(
		"@cdktn/provider-aws.bedrockagentDataSource.BedrockagentDataSourceVectorIngestionConfigurationParsingConfigurationBedrockDataAutomationConfigurationList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewBedrockagentDataSourceVectorIngestionConfigurationParsingConfigurationBedrockDataAutomationConfigurationList_Override(b BedrockagentDataSourceVectorIngestionConfigurationParsingConfigurationBedrockDataAutomationConfigurationList, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.bedrockagentDataSource.BedrockagentDataSourceVectorIngestionConfigurationParsingConfigurationBedrockDataAutomationConfigurationList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		b,
	)
}

func (j *jsiiProxy_BedrockagentDataSourceVectorIngestionConfigurationParsingConfigurationBedrockDataAutomationConfigurationList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BedrockagentDataSourceVectorIngestionConfigurationParsingConfigurationBedrockDataAutomationConfigurationList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BedrockagentDataSourceVectorIngestionConfigurationParsingConfigurationBedrockDataAutomationConfigurationList)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_BedrockagentDataSourceVectorIngestionConfigurationParsingConfigurationBedrockDataAutomationConfigurationList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (b *jsiiProxy_BedrockagentDataSourceVectorIngestionConfigurationParsingConfigurationBedrockDataAutomationConfigurationList) AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator {
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

func (b *jsiiProxy_BedrockagentDataSourceVectorIngestionConfigurationParsingConfigurationBedrockDataAutomationConfigurationList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentDataSourceVectorIngestionConfigurationParsingConfigurationBedrockDataAutomationConfigurationList) Get(index *float64) BedrockagentDataSourceVectorIngestionConfigurationParsingConfigurationBedrockDataAutomationConfigurationOutputReference {
	if err := b.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns BedrockagentDataSourceVectorIngestionConfigurationParsingConfigurationBedrockDataAutomationConfigurationOutputReference

	_jsii_.Invoke(
		b,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentDataSourceVectorIngestionConfigurationParsingConfigurationBedrockDataAutomationConfigurationList) Resolve(context cdktn.IResolveContext) interface{} {
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

func (b *jsiiProxy_BedrockagentDataSourceVectorIngestionConfigurationParsingConfigurationBedrockDataAutomationConfigurationList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

