// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluecatalogtable

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/jsii"

	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/gluecatalogtable/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsOutputReference interface {
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
	Doc() *string
	SetDoc(val *string)
	DocInput() *string
	// Experimental.
	Fqn() *string
	Id() *float64
	SetId(val *float64)
	IdInput() *float64
	InitialDefault() *string
	SetInitialDefault(val *string)
	InitialDefaultInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	Required() interface{}
	SetRequired(val interface{})
	RequiredInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeInput() *string
	WriteDefault() *string
	SetWriteDefault(val *string)
	WriteDefaultInput() *string
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
	ResetDoc()
	ResetInitialDefault()
	ResetWriteDefault()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsOutputReference
type jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsOutputReference) Doc() *string {
	var returns *string
	_jsii_.Get(
		j,
		"doc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsOutputReference) DocInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"docInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsOutputReference) Id() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsOutputReference) IdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsOutputReference) InitialDefault() *string {
	var returns *string
	_jsii_.Get(
		j,
		"initialDefault",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsOutputReference) InitialDefaultInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"initialDefaultInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsOutputReference) Required() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"required",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsOutputReference) RequiredInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requiredInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsOutputReference) WriteDefault() *string {
	var returns *string
	_jsii_.Get(
		j,
		"writeDefault",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsOutputReference) WriteDefaultInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"writeDefaultInput",
		&returns,
	)
	return returns
}


func NewGlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsOutputReference {
	_init_.Initialize()

	if err := validateNewGlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-aws.glueCatalogTable.GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsOutputReference_Override(g GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.glueCatalogTable.GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsOutputReference)SetDoc(val *string) {
	if err := j.validateSetDocParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"doc",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsOutputReference)SetId(val *float64) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsOutputReference)SetInitialDefault(val *string) {
	if err := j.validateSetInitialDefaultParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"initialDefault",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsOutputReference)SetRequired(val interface{}) {
	if err := j.validateSetRequiredParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"required",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsOutputReference)SetWriteDefault(val *string) {
	if err := j.validateSetWriteDefaultParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"writeDefault",
		val,
	)
}

func (g *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsOutputReference) ResetDoc() {
	_jsii_.InvokeVoid(
		g,
		"resetDoc",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsOutputReference) ResetInitialDefault() {
	_jsii_.InvokeVoid(
		g,
		"resetInitialDefault",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsOutputReference) ResetWriteDefault() {
	_jsii_.InvokeVoid(
		g,
		"resetWriteDefault",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := g.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

