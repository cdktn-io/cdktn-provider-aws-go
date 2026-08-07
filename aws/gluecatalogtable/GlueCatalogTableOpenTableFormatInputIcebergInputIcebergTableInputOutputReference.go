// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluecatalogtable

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/jsii"

	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/gluecatalogtable/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference interface {
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
	InternalValue() *GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInput
	SetInternalValue(val *GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInput)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	PartitionSpec() GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputPartitionSpecOutputReference
	PartitionSpecInput() *GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputPartitionSpec
	Properties() *map[string]*string
	SetProperties(val *map[string]*string)
	PropertiesInput() *map[string]*string
	Schema() GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaOutputReference
	SchemaInput() *GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchema
	SortOrder() GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSortOrderOutputReference
	SortOrderInput() *GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSortOrder
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
	PutPartitionSpec(value *GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputPartitionSpec)
	PutSchema(value *GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchema)
	PutSortOrder(value *GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSortOrder)
	ResetPartitionSpec()
	ResetProperties()
	ResetSortOrder()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference
type jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference) InternalValue() *GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInput {
	var returns *GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInput
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference) PartitionSpec() GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputPartitionSpecOutputReference {
	var returns GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputPartitionSpecOutputReference
	_jsii_.Get(
		j,
		"partitionSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference) PartitionSpecInput() *GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputPartitionSpec {
	var returns *GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputPartitionSpec
	_jsii_.Get(
		j,
		"partitionSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference) Properties() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"properties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference) PropertiesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"propertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference) Schema() GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaOutputReference {
	var returns GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaOutputReference
	_jsii_.Get(
		j,
		"schema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference) SchemaInput() *GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchema {
	var returns *GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchema
	_jsii_.Get(
		j,
		"schemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference) SortOrder() GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSortOrderOutputReference {
	var returns GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSortOrderOutputReference
	_jsii_.Get(
		j,
		"sortOrder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference) SortOrderInput() *GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSortOrder {
	var returns *GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSortOrder
	_jsii_.Get(
		j,
		"sortOrderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference {
	_init_.Initialize()

	if err := validateNewGlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-aws.glueCatalogTable.GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference_Override(g GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.glueCatalogTable.GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference)SetInternalValue(val *GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInput) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference)SetProperties(val *map[string]*string) {
	if err := j.validateSetPropertiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"properties",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference) PutPartitionSpec(value *GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputPartitionSpec) {
	if err := g.validatePutPartitionSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPartitionSpec",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference) PutSchema(value *GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSchema) {
	if err := g.validatePutSchemaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSchema",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference) PutSortOrder(value *GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputSortOrder) {
	if err := g.validatePutSortOrderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSortOrder",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference) ResetPartitionSpec() {
	_jsii_.InvokeVoid(
		g,
		"resetPartitionSpec",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference) ResetProperties() {
	_jsii_.InvokeVoid(
		g,
		"resetProperties",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference) ResetSortOrder() {
	_jsii_.InvokeVoid(
		g,
		"resetSortOrder",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GlueCatalogTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

