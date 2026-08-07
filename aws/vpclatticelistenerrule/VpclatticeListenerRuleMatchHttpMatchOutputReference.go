// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vpclatticelistenerrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/jsii"

	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/vpclatticelistenerrule/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type VpclatticeListenerRuleMatchHttpMatchOutputReference interface {
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
	HeaderMatches() VpclatticeListenerRuleMatchHttpMatchHeaderMatchesList
	HeaderMatchesInput() interface{}
	InternalValue() *VpclatticeListenerRuleMatchHttpMatch
	SetInternalValue(val *VpclatticeListenerRuleMatchHttpMatch)
	Method() *string
	SetMethod(val *string)
	MethodInput() *string
	PathMatch() VpclatticeListenerRuleMatchHttpMatchPathMatchOutputReference
	PathMatchInput() *VpclatticeListenerRuleMatchHttpMatchPathMatch
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
	PutHeaderMatches(value interface{})
	PutPathMatch(value *VpclatticeListenerRuleMatchHttpMatchPathMatch)
	ResetHeaderMatches()
	ResetMethod()
	ResetPathMatch()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VpclatticeListenerRuleMatchHttpMatchOutputReference
type jsiiProxy_VpclatticeListenerRuleMatchHttpMatchOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_VpclatticeListenerRuleMatchHttpMatchOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpclatticeListenerRuleMatchHttpMatchOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpclatticeListenerRuleMatchHttpMatchOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpclatticeListenerRuleMatchHttpMatchOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpclatticeListenerRuleMatchHttpMatchOutputReference) HeaderMatches() VpclatticeListenerRuleMatchHttpMatchHeaderMatchesList {
	var returns VpclatticeListenerRuleMatchHttpMatchHeaderMatchesList
	_jsii_.Get(
		j,
		"headerMatches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpclatticeListenerRuleMatchHttpMatchOutputReference) HeaderMatchesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"headerMatchesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpclatticeListenerRuleMatchHttpMatchOutputReference) InternalValue() *VpclatticeListenerRuleMatchHttpMatch {
	var returns *VpclatticeListenerRuleMatchHttpMatch
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpclatticeListenerRuleMatchHttpMatchOutputReference) Method() *string {
	var returns *string
	_jsii_.Get(
		j,
		"method",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpclatticeListenerRuleMatchHttpMatchOutputReference) MethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"methodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpclatticeListenerRuleMatchHttpMatchOutputReference) PathMatch() VpclatticeListenerRuleMatchHttpMatchPathMatchOutputReference {
	var returns VpclatticeListenerRuleMatchHttpMatchPathMatchOutputReference
	_jsii_.Get(
		j,
		"pathMatch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpclatticeListenerRuleMatchHttpMatchOutputReference) PathMatchInput() *VpclatticeListenerRuleMatchHttpMatchPathMatch {
	var returns *VpclatticeListenerRuleMatchHttpMatchPathMatch
	_jsii_.Get(
		j,
		"pathMatchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpclatticeListenerRuleMatchHttpMatchOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpclatticeListenerRuleMatchHttpMatchOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewVpclatticeListenerRuleMatchHttpMatchOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) VpclatticeListenerRuleMatchHttpMatchOutputReference {
	_init_.Initialize()

	if err := validateNewVpclatticeListenerRuleMatchHttpMatchOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_VpclatticeListenerRuleMatchHttpMatchOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-aws.vpclatticeListenerRule.VpclatticeListenerRuleMatchHttpMatchOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewVpclatticeListenerRuleMatchHttpMatchOutputReference_Override(v VpclatticeListenerRuleMatchHttpMatchOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.vpclatticeListenerRule.VpclatticeListenerRuleMatchHttpMatchOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_VpclatticeListenerRuleMatchHttpMatchOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VpclatticeListenerRuleMatchHttpMatchOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VpclatticeListenerRuleMatchHttpMatchOutputReference)SetInternalValue(val *VpclatticeListenerRuleMatchHttpMatch) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VpclatticeListenerRuleMatchHttpMatchOutputReference)SetMethod(val *string) {
	if err := j.validateSetMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"method",
		val,
	)
}

func (j *jsiiProxy_VpclatticeListenerRuleMatchHttpMatchOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VpclatticeListenerRuleMatchHttpMatchOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_VpclatticeListenerRuleMatchHttpMatchOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpclatticeListenerRuleMatchHttpMatchOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := v.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpclatticeListenerRuleMatchHttpMatchOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := v.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		v,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpclatticeListenerRuleMatchHttpMatchOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := v.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		v,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpclatticeListenerRuleMatchHttpMatchOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := v.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpclatticeListenerRuleMatchHttpMatchOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := v.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		v,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpclatticeListenerRuleMatchHttpMatchOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := v.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		v,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpclatticeListenerRuleMatchHttpMatchOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := v.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		v,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpclatticeListenerRuleMatchHttpMatchOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := v.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		v,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpclatticeListenerRuleMatchHttpMatchOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := v.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		v,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpclatticeListenerRuleMatchHttpMatchOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpclatticeListenerRuleMatchHttpMatchOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := v.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpclatticeListenerRuleMatchHttpMatchOutputReference) PutHeaderMatches(value interface{}) {
	if err := v.validatePutHeaderMatchesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putHeaderMatches",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VpclatticeListenerRuleMatchHttpMatchOutputReference) PutPathMatch(value *VpclatticeListenerRuleMatchHttpMatchPathMatch) {
	if err := v.validatePutPathMatchParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putPathMatch",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VpclatticeListenerRuleMatchHttpMatchOutputReference) ResetHeaderMatches() {
	_jsii_.InvokeVoid(
		v,
		"resetHeaderMatches",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpclatticeListenerRuleMatchHttpMatchOutputReference) ResetMethod() {
	_jsii_.InvokeVoid(
		v,
		"resetMethod",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpclatticeListenerRuleMatchHttpMatchOutputReference) ResetPathMatch() {
	_jsii_.InvokeVoid(
		v,
		"resetPathMatch",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpclatticeListenerRuleMatchHttpMatchOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := v.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		v,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpclatticeListenerRuleMatchHttpMatchOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

