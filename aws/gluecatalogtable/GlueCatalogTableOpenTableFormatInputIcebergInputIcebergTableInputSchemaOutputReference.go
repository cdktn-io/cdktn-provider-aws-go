// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluecatalogtable

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v24/jsii"

	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v24/gluecatalogtable/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaOutputReference interface {
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
	Fields() GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsList
	FieldsInput() interface{}
	// Experimental.
	Fqn() *string
	IdentifierFieldIds() *[]*float64
	SetIdentifierFieldIds(val *[]*float64)
	IdentifierFieldIdsInput() *[]*float64
	InternalValue() *GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchema
	SetInternalValue(val *GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchema)
	SchemaId() *float64
	SetSchemaId(val *float64)
	SchemaIdInput() *float64
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
	PutFields(value interface{})
	ResetIdentifierFieldIds()
	ResetSchemaId()
	ResetType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaOutputReference
type jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaOutputReference) Fields() GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsList {
	var returns GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsList
	_jsii_.Get(
		j,
		"fields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaOutputReference) FieldsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fieldsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaOutputReference) IdentifierFieldIds() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"identifierFieldIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaOutputReference) IdentifierFieldIdsInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"identifierFieldIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaOutputReference) InternalValue() *GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchema {
	var returns *GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchema
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaOutputReference) SchemaId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"schemaId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaOutputReference) SchemaIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"schemaIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewGlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaOutputReference {
	_init_.Initialize()

	if err := validateNewGlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-aws.glueCatalogTable.GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaOutputReference_Override(g GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.glueCatalogTable.GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaOutputReference)SetIdentifierFieldIds(val *[]*float64) {
	if err := j.validateSetIdentifierFieldIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"identifierFieldIds",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaOutputReference)SetInternalValue(val *GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchema) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaOutputReference)SetSchemaId(val *float64) {
	if err := j.validateSetSchemaIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schemaId",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (g *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaOutputReference) PutFields(value interface{}) {
	if err := g.validatePutFieldsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putFields",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaOutputReference) ResetIdentifierFieldIds() {
	_jsii_.InvokeVoid(
		g,
		"resetIdentifierFieldIds",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaOutputReference) ResetSchemaId() {
	_jsii_.InvokeVoid(
		g,
		"resetSchemaId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaOutputReference) ResetType() {
	_jsii_.InvokeVoid(
		g,
		"resetType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

