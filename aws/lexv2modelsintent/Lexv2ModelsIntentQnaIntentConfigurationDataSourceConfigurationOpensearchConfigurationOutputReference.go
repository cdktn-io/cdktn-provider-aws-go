// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lexv2modelsintent

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/jsii"

	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/lexv2modelsintent/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationOpensearchConfigurationOutputReference interface {
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
	DomainEndpoint() *string
	SetDomainEndpoint(val *string)
	DomainEndpointInput() *string
	ExactResponse() interface{}
	SetExactResponse(val interface{})
	ExactResponseFields() Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationOpensearchConfigurationExactResponseFieldsList
	ExactResponseFieldsInput() interface{}
	ExactResponseInput() interface{}
	// Experimental.
	Fqn() *string
	IncludeFields() *[]*string
	SetIncludeFields(val *[]*string)
	IncludeFieldsInput() *[]*string
	IndexName() *string
	SetIndexName(val *string)
	IndexNameInput() *string
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
	PutExactResponseFields(value interface{})
	ResetExactResponse()
	ResetExactResponseFields()
	ResetIncludeFields()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationOpensearchConfigurationOutputReference
type jsiiProxy_Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationOpensearchConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationOpensearchConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationOpensearchConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationOpensearchConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationOpensearchConfigurationOutputReference) DomainEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationOpensearchConfigurationOutputReference) DomainEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationOpensearchConfigurationOutputReference) ExactResponse() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"exactResponse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationOpensearchConfigurationOutputReference) ExactResponseFields() Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationOpensearchConfigurationExactResponseFieldsList {
	var returns Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationOpensearchConfigurationExactResponseFieldsList
	_jsii_.Get(
		j,
		"exactResponseFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationOpensearchConfigurationOutputReference) ExactResponseFieldsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"exactResponseFieldsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationOpensearchConfigurationOutputReference) ExactResponseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"exactResponseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationOpensearchConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationOpensearchConfigurationOutputReference) IncludeFields() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includeFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationOpensearchConfigurationOutputReference) IncludeFieldsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includeFieldsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationOpensearchConfigurationOutputReference) IndexName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"indexName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationOpensearchConfigurationOutputReference) IndexNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"indexNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationOpensearchConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationOpensearchConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationOpensearchConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewLexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationOpensearchConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationOpensearchConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewLexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationOpensearchConfigurationOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationOpensearchConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-aws.lexv2ModelsIntent.Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationOpensearchConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewLexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationOpensearchConfigurationOutputReference_Override(l Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationOpensearchConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.lexv2ModelsIntent.Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationOpensearchConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		l,
	)
}

func (j *jsiiProxy_Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationOpensearchConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationOpensearchConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationOpensearchConfigurationOutputReference)SetDomainEndpoint(val *string) {
	if err := j.validateSetDomainEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainEndpoint",
		val,
	)
}

func (j *jsiiProxy_Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationOpensearchConfigurationOutputReference)SetExactResponse(val interface{}) {
	if err := j.validateSetExactResponseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exactResponse",
		val,
	)
}

func (j *jsiiProxy_Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationOpensearchConfigurationOutputReference)SetIncludeFields(val *[]*string) {
	if err := j.validateSetIncludeFieldsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeFields",
		val,
	)
}

func (j *jsiiProxy_Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationOpensearchConfigurationOutputReference)SetIndexName(val *string) {
	if err := j.validateSetIndexNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"indexName",
		val,
	)
}

func (j *jsiiProxy_Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationOpensearchConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationOpensearchConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationOpensearchConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (l *jsiiProxy_Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationOpensearchConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationOpensearchConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (l *jsiiProxy_Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationOpensearchConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (l *jsiiProxy_Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationOpensearchConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (l *jsiiProxy_Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationOpensearchConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (l *jsiiProxy_Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationOpensearchConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (l *jsiiProxy_Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationOpensearchConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (l *jsiiProxy_Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationOpensearchConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (l *jsiiProxy_Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationOpensearchConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (l *jsiiProxy_Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationOpensearchConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (l *jsiiProxy_Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationOpensearchConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationOpensearchConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (l *jsiiProxy_Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationOpensearchConfigurationOutputReference) PutExactResponseFields(value interface{}) {
	if err := l.validatePutExactResponseFieldsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putExactResponseFields",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationOpensearchConfigurationOutputReference) ResetExactResponse() {
	_jsii_.InvokeVoid(
		l,
		"resetExactResponse",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationOpensearchConfigurationOutputReference) ResetExactResponseFields() {
	_jsii_.InvokeVoid(
		l,
		"resetExactResponseFields",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationOpensearchConfigurationOutputReference) ResetIncludeFields() {
	_jsii_.InvokeVoid(
		l,
		"resetIncludeFields",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationOpensearchConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (l *jsiiProxy_Lexv2ModelsIntentQnaIntentConfigurationDataSourceConfigurationOpensearchConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

