// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lexv2modelsintent

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v24/jsii"

	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v24/lexv2modelsintent/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationKendraConfigurationOutputReference interface {
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
	ExactResponse() interface{}
	SetExactResponse(val interface{})
	ExactResponseInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	KendraIndex() *string
	SetKendraIndex(val *string)
	KendraIndexInput() *string
	QueryFilterString() *string
	SetQueryFilterString(val *string)
	QueryFilterStringEnabled() interface{}
	SetQueryFilterStringEnabled(val interface{})
	QueryFilterStringEnabledInput() interface{}
	QueryFilterStringInput() *string
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
	ResetExactResponse()
	ResetQueryFilterString()
	ResetQueryFilterStringEnabled()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationKendraConfigurationOutputReference
type jsiiProxy_Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationKendraConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationKendraConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationKendraConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationKendraConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationKendraConfigurationOutputReference) ExactResponse() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"exactResponse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationKendraConfigurationOutputReference) ExactResponseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"exactResponseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationKendraConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationKendraConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationKendraConfigurationOutputReference) KendraIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kendraIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationKendraConfigurationOutputReference) KendraIndexInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kendraIndexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationKendraConfigurationOutputReference) QueryFilterString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryFilterString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationKendraConfigurationOutputReference) QueryFilterStringEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"queryFilterStringEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationKendraConfigurationOutputReference) QueryFilterStringEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"queryFilterStringEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationKendraConfigurationOutputReference) QueryFilterStringInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryFilterStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationKendraConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationKendraConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewLexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationKendraConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationKendraConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewLexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationKendraConfigurationOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationKendraConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-aws.lexv2ModelsIntent.Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationKendraConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewLexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationKendraConfigurationOutputReference_Override(l Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationKendraConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.lexv2ModelsIntent.Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationKendraConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		l,
	)
}

func (j *jsiiProxy_Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationKendraConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationKendraConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationKendraConfigurationOutputReference)SetExactResponse(val interface{}) {
	if err := j.validateSetExactResponseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exactResponse",
		val,
	)
}

func (j *jsiiProxy_Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationKendraConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationKendraConfigurationOutputReference)SetKendraIndex(val *string) {
	if err := j.validateSetKendraIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kendraIndex",
		val,
	)
}

func (j *jsiiProxy_Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationKendraConfigurationOutputReference)SetQueryFilterString(val *string) {
	if err := j.validateSetQueryFilterStringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queryFilterString",
		val,
	)
}

func (j *jsiiProxy_Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationKendraConfigurationOutputReference)SetQueryFilterStringEnabled(val interface{}) {
	if err := j.validateSetQueryFilterStringEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queryFilterStringEnabled",
		val,
	)
}

func (j *jsiiProxy_Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationKendraConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationKendraConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (l *jsiiProxy_Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationKendraConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationKendraConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := l.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationKendraConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := l.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationKendraConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := l.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationKendraConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := l.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationKendraConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := l.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationKendraConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := l.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationKendraConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := l.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationKendraConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := l.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationKendraConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := l.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationKendraConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationKendraConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationKendraConfigurationOutputReference) ResetExactResponse() {
	_jsii_.InvokeVoid(
		l,
		"resetExactResponse",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationKendraConfigurationOutputReference) ResetQueryFilterString() {
	_jsii_.InvokeVoid(
		l,
		"resetQueryFilterString",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationKendraConfigurationOutputReference) ResetQueryFilterStringEnabled() {
	_jsii_.InvokeVoid(
		l,
		"resetQueryFilterStringEnabled",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationKendraConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := l.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		l,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationKendraConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

