// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mailmanagerruleset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/jsii"

	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/mailmanagerruleset/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MailmanagerRuleSetRuleUnlessOutputReference interface {
	cdktn.ComplexObject
	BooleanExpression() MailmanagerRuleSetRuleUnlessBooleanExpressionList
	BooleanExpressionInput() interface{}
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
	DmarcExpression() MailmanagerRuleSetRuleUnlessDmarcExpressionList
	DmarcExpressionInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IpExpression() MailmanagerRuleSetRuleUnlessIpExpressionList
	IpExpressionInput() interface{}
	NumberExpression() MailmanagerRuleSetRuleUnlessNumberExpressionList
	NumberExpressionInput() interface{}
	StringExpression() MailmanagerRuleSetRuleUnlessStringExpressionList
	StringExpressionInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	VerdictExpression() MailmanagerRuleSetRuleUnlessVerdictExpressionList
	VerdictExpressionInput() interface{}
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
	PutBooleanExpression(value interface{})
	PutDmarcExpression(value interface{})
	PutIpExpression(value interface{})
	PutNumberExpression(value interface{})
	PutStringExpression(value interface{})
	PutVerdictExpression(value interface{})
	ResetBooleanExpression()
	ResetDmarcExpression()
	ResetIpExpression()
	ResetNumberExpression()
	ResetStringExpression()
	ResetVerdictExpression()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MailmanagerRuleSetRuleUnlessOutputReference
type jsiiProxy_MailmanagerRuleSetRuleUnlessOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_MailmanagerRuleSetRuleUnlessOutputReference) BooleanExpression() MailmanagerRuleSetRuleUnlessBooleanExpressionList {
	var returns MailmanagerRuleSetRuleUnlessBooleanExpressionList
	_jsii_.Get(
		j,
		"booleanExpression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MailmanagerRuleSetRuleUnlessOutputReference) BooleanExpressionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"booleanExpressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MailmanagerRuleSetRuleUnlessOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MailmanagerRuleSetRuleUnlessOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MailmanagerRuleSetRuleUnlessOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MailmanagerRuleSetRuleUnlessOutputReference) DmarcExpression() MailmanagerRuleSetRuleUnlessDmarcExpressionList {
	var returns MailmanagerRuleSetRuleUnlessDmarcExpressionList
	_jsii_.Get(
		j,
		"dmarcExpression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MailmanagerRuleSetRuleUnlessOutputReference) DmarcExpressionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dmarcExpressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MailmanagerRuleSetRuleUnlessOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MailmanagerRuleSetRuleUnlessOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MailmanagerRuleSetRuleUnlessOutputReference) IpExpression() MailmanagerRuleSetRuleUnlessIpExpressionList {
	var returns MailmanagerRuleSetRuleUnlessIpExpressionList
	_jsii_.Get(
		j,
		"ipExpression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MailmanagerRuleSetRuleUnlessOutputReference) IpExpressionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipExpressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MailmanagerRuleSetRuleUnlessOutputReference) NumberExpression() MailmanagerRuleSetRuleUnlessNumberExpressionList {
	var returns MailmanagerRuleSetRuleUnlessNumberExpressionList
	_jsii_.Get(
		j,
		"numberExpression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MailmanagerRuleSetRuleUnlessOutputReference) NumberExpressionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"numberExpressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MailmanagerRuleSetRuleUnlessOutputReference) StringExpression() MailmanagerRuleSetRuleUnlessStringExpressionList {
	var returns MailmanagerRuleSetRuleUnlessStringExpressionList
	_jsii_.Get(
		j,
		"stringExpression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MailmanagerRuleSetRuleUnlessOutputReference) StringExpressionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stringExpressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MailmanagerRuleSetRuleUnlessOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MailmanagerRuleSetRuleUnlessOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MailmanagerRuleSetRuleUnlessOutputReference) VerdictExpression() MailmanagerRuleSetRuleUnlessVerdictExpressionList {
	var returns MailmanagerRuleSetRuleUnlessVerdictExpressionList
	_jsii_.Get(
		j,
		"verdictExpression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MailmanagerRuleSetRuleUnlessOutputReference) VerdictExpressionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"verdictExpressionInput",
		&returns,
	)
	return returns
}


func NewMailmanagerRuleSetRuleUnlessOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) MailmanagerRuleSetRuleUnlessOutputReference {
	_init_.Initialize()

	if err := validateNewMailmanagerRuleSetRuleUnlessOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_MailmanagerRuleSetRuleUnlessOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-aws.mailmanagerRuleSet.MailmanagerRuleSetRuleUnlessOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewMailmanagerRuleSetRuleUnlessOutputReference_Override(m MailmanagerRuleSetRuleUnlessOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.mailmanagerRuleSet.MailmanagerRuleSetRuleUnlessOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		m,
	)
}

func (j *jsiiProxy_MailmanagerRuleSetRuleUnlessOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MailmanagerRuleSetRuleUnlessOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MailmanagerRuleSetRuleUnlessOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MailmanagerRuleSetRuleUnlessOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MailmanagerRuleSetRuleUnlessOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MailmanagerRuleSetRuleUnlessOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MailmanagerRuleSetRuleUnlessOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MailmanagerRuleSetRuleUnlessOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MailmanagerRuleSetRuleUnlessOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MailmanagerRuleSetRuleUnlessOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MailmanagerRuleSetRuleUnlessOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MailmanagerRuleSetRuleUnlessOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MailmanagerRuleSetRuleUnlessOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MailmanagerRuleSetRuleUnlessOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MailmanagerRuleSetRuleUnlessOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MailmanagerRuleSetRuleUnlessOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MailmanagerRuleSetRuleUnlessOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MailmanagerRuleSetRuleUnlessOutputReference) PutBooleanExpression(value interface{}) {
	if err := m.validatePutBooleanExpressionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putBooleanExpression",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MailmanagerRuleSetRuleUnlessOutputReference) PutDmarcExpression(value interface{}) {
	if err := m.validatePutDmarcExpressionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putDmarcExpression",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MailmanagerRuleSetRuleUnlessOutputReference) PutIpExpression(value interface{}) {
	if err := m.validatePutIpExpressionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putIpExpression",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MailmanagerRuleSetRuleUnlessOutputReference) PutNumberExpression(value interface{}) {
	if err := m.validatePutNumberExpressionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putNumberExpression",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MailmanagerRuleSetRuleUnlessOutputReference) PutStringExpression(value interface{}) {
	if err := m.validatePutStringExpressionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putStringExpression",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MailmanagerRuleSetRuleUnlessOutputReference) PutVerdictExpression(value interface{}) {
	if err := m.validatePutVerdictExpressionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putVerdictExpression",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MailmanagerRuleSetRuleUnlessOutputReference) ResetBooleanExpression() {
	_jsii_.InvokeVoid(
		m,
		"resetBooleanExpression",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MailmanagerRuleSetRuleUnlessOutputReference) ResetDmarcExpression() {
	_jsii_.InvokeVoid(
		m,
		"resetDmarcExpression",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MailmanagerRuleSetRuleUnlessOutputReference) ResetIpExpression() {
	_jsii_.InvokeVoid(
		m,
		"resetIpExpression",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MailmanagerRuleSetRuleUnlessOutputReference) ResetNumberExpression() {
	_jsii_.InvokeVoid(
		m,
		"resetNumberExpression",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MailmanagerRuleSetRuleUnlessOutputReference) ResetStringExpression() {
	_jsii_.InvokeVoid(
		m,
		"resetStringExpression",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MailmanagerRuleSetRuleUnlessOutputReference) ResetVerdictExpression() {
	_jsii_.InvokeVoid(
		m,
		"resetVerdictExpression",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MailmanagerRuleSetRuleUnlessOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := m.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MailmanagerRuleSetRuleUnlessOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

