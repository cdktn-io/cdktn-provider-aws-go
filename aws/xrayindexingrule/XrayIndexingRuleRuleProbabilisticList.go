// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package xrayindexingrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/jsii"

	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/xrayindexingrule/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type XrayIndexingRuleRuleProbabilisticList interface {
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
	Get(index *float64) XrayIndexingRuleRuleProbabilisticOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for XrayIndexingRuleRuleProbabilisticList
type jsiiProxy_XrayIndexingRuleRuleProbabilisticList struct {
	internal.Type__cdktnComplexList
}

func (j *jsiiProxy_XrayIndexingRuleRuleProbabilisticList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XrayIndexingRuleRuleProbabilisticList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XrayIndexingRuleRuleProbabilisticList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XrayIndexingRuleRuleProbabilisticList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XrayIndexingRuleRuleProbabilisticList) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XrayIndexingRuleRuleProbabilisticList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewXrayIndexingRuleRuleProbabilisticList(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) XrayIndexingRuleRuleProbabilisticList {
	_init_.Initialize()

	if err := validateNewXrayIndexingRuleRuleProbabilisticListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_XrayIndexingRuleRuleProbabilisticList{}

	_jsii_.Create(
		"@cdktn/provider-aws.xrayIndexingRule.XrayIndexingRuleRuleProbabilisticList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewXrayIndexingRuleRuleProbabilisticList_Override(x XrayIndexingRuleRuleProbabilisticList, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.xrayIndexingRule.XrayIndexingRuleRuleProbabilisticList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		x,
	)
}

func (j *jsiiProxy_XrayIndexingRuleRuleProbabilisticList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_XrayIndexingRuleRuleProbabilisticList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_XrayIndexingRuleRuleProbabilisticList)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_XrayIndexingRuleRuleProbabilisticList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (x *jsiiProxy_XrayIndexingRuleRuleProbabilisticList) AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator {
	if err := x.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktn.DynamicListTerraformIterator

	_jsii_.Invoke(
		x,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (x *jsiiProxy_XrayIndexingRuleRuleProbabilisticList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		x,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (x *jsiiProxy_XrayIndexingRuleRuleProbabilisticList) Get(index *float64) XrayIndexingRuleRuleProbabilisticOutputReference {
	if err := x.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns XrayIndexingRuleRuleProbabilisticOutputReference

	_jsii_.Invoke(
		x,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (x *jsiiProxy_XrayIndexingRuleRuleProbabilisticList) Resolve(context cdktn.IResolveContext) interface{} {
	if err := x.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		x,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (x *jsiiProxy_XrayIndexingRuleRuleProbabilisticList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		x,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

