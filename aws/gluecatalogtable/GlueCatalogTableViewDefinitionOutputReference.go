// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluecatalogtable

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v24/jsii"

	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v24/gluecatalogtable/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GlueCatalogTableViewDefinitionOutputReference interface {
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
	Definer() *string
	SetDefiner(val *string)
	DefinerInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *GlueCatalogTableViewDefinition
	SetInternalValue(val *GlueCatalogTableViewDefinition)
	IsProtected() interface{}
	SetIsProtected(val interface{})
	IsProtectedInput() interface{}
	LastRefreshType() *string
	SetLastRefreshType(val *string)
	LastRefreshTypeInput() *string
	RefreshSeconds() *float64
	SetRefreshSeconds(val *float64)
	RefreshSecondsInput() *float64
	Representations() GlueCatalogTableViewDefinitionRepresentationsList
	RepresentationsInput() interface{}
	SubObjects() *[]*string
	SetSubObjects(val *[]*string)
	SubObjectsInput() *[]*string
	SubObjectVersionIds() *[]*float64
	SetSubObjectVersionIds(val *[]*float64)
	SubObjectVersionIdsInput() *[]*float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	ViewVersionId() *float64
	SetViewVersionId(val *float64)
	ViewVersionIdInput() *float64
	ViewVersionToken() *string
	SetViewVersionToken(val *string)
	ViewVersionTokenInput() *string
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
	PutRepresentations(value interface{})
	ResetDefiner()
	ResetIsProtected()
	ResetLastRefreshType()
	ResetRefreshSeconds()
	ResetRepresentations()
	ResetSubObjects()
	ResetSubObjectVersionIds()
	ResetViewVersionId()
	ResetViewVersionToken()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GlueCatalogTableViewDefinitionOutputReference
type jsiiProxy_GlueCatalogTableViewDefinitionOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GlueCatalogTableViewDefinitionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableViewDefinitionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableViewDefinitionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableViewDefinitionOutputReference) Definer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"definer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableViewDefinitionOutputReference) DefinerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"definerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableViewDefinitionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableViewDefinitionOutputReference) InternalValue() *GlueCatalogTableViewDefinition {
	var returns *GlueCatalogTableViewDefinition
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableViewDefinitionOutputReference) IsProtected() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isProtected",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableViewDefinitionOutputReference) IsProtectedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isProtectedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableViewDefinitionOutputReference) LastRefreshType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastRefreshType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableViewDefinitionOutputReference) LastRefreshTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastRefreshTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableViewDefinitionOutputReference) RefreshSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"refreshSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableViewDefinitionOutputReference) RefreshSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"refreshSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableViewDefinitionOutputReference) Representations() GlueCatalogTableViewDefinitionRepresentationsList {
	var returns GlueCatalogTableViewDefinitionRepresentationsList
	_jsii_.Get(
		j,
		"representations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableViewDefinitionOutputReference) RepresentationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"representationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableViewDefinitionOutputReference) SubObjects() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subObjects",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableViewDefinitionOutputReference) SubObjectsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subObjectsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableViewDefinitionOutputReference) SubObjectVersionIds() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"subObjectVersionIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableViewDefinitionOutputReference) SubObjectVersionIdsInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"subObjectVersionIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableViewDefinitionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableViewDefinitionOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableViewDefinitionOutputReference) ViewVersionId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"viewVersionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableViewDefinitionOutputReference) ViewVersionIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"viewVersionIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableViewDefinitionOutputReference) ViewVersionToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"viewVersionToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableViewDefinitionOutputReference) ViewVersionTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"viewVersionTokenInput",
		&returns,
	)
	return returns
}


func NewGlueCatalogTableViewDefinitionOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GlueCatalogTableViewDefinitionOutputReference {
	_init_.Initialize()

	if err := validateNewGlueCatalogTableViewDefinitionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GlueCatalogTableViewDefinitionOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-aws.glueCatalogTable.GlueCatalogTableViewDefinitionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGlueCatalogTableViewDefinitionOutputReference_Override(g GlueCatalogTableViewDefinitionOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.glueCatalogTable.GlueCatalogTableViewDefinitionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GlueCatalogTableViewDefinitionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTableViewDefinitionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTableViewDefinitionOutputReference)SetDefiner(val *string) {
	if err := j.validateSetDefinerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"definer",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTableViewDefinitionOutputReference)SetInternalValue(val *GlueCatalogTableViewDefinition) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTableViewDefinitionOutputReference)SetIsProtected(val interface{}) {
	if err := j.validateSetIsProtectedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isProtected",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTableViewDefinitionOutputReference)SetLastRefreshType(val *string) {
	if err := j.validateSetLastRefreshTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lastRefreshType",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTableViewDefinitionOutputReference)SetRefreshSeconds(val *float64) {
	if err := j.validateSetRefreshSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"refreshSeconds",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTableViewDefinitionOutputReference)SetSubObjects(val *[]*string) {
	if err := j.validateSetSubObjectsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subObjects",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTableViewDefinitionOutputReference)SetSubObjectVersionIds(val *[]*float64) {
	if err := j.validateSetSubObjectVersionIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subObjectVersionIds",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTableViewDefinitionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTableViewDefinitionOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTableViewDefinitionOutputReference)SetViewVersionId(val *float64) {
	if err := j.validateSetViewVersionIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"viewVersionId",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTableViewDefinitionOutputReference)SetViewVersionToken(val *string) {
	if err := j.validateSetViewVersionTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"viewVersionToken",
		val,
	)
}

func (g *jsiiProxy_GlueCatalogTableViewDefinitionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCatalogTableViewDefinitionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GlueCatalogTableViewDefinitionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GlueCatalogTableViewDefinitionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GlueCatalogTableViewDefinitionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GlueCatalogTableViewDefinitionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GlueCatalogTableViewDefinitionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GlueCatalogTableViewDefinitionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GlueCatalogTableViewDefinitionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GlueCatalogTableViewDefinitionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GlueCatalogTableViewDefinitionOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCatalogTableViewDefinitionOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GlueCatalogTableViewDefinitionOutputReference) PutRepresentations(value interface{}) {
	if err := g.validatePutRepresentationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRepresentations",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueCatalogTableViewDefinitionOutputReference) ResetDefiner() {
	_jsii_.InvokeVoid(
		g,
		"resetDefiner",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCatalogTableViewDefinitionOutputReference) ResetIsProtected() {
	_jsii_.InvokeVoid(
		g,
		"resetIsProtected",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCatalogTableViewDefinitionOutputReference) ResetLastRefreshType() {
	_jsii_.InvokeVoid(
		g,
		"resetLastRefreshType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCatalogTableViewDefinitionOutputReference) ResetRefreshSeconds() {
	_jsii_.InvokeVoid(
		g,
		"resetRefreshSeconds",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCatalogTableViewDefinitionOutputReference) ResetRepresentations() {
	_jsii_.InvokeVoid(
		g,
		"resetRepresentations",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCatalogTableViewDefinitionOutputReference) ResetSubObjects() {
	_jsii_.InvokeVoid(
		g,
		"resetSubObjects",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCatalogTableViewDefinitionOutputReference) ResetSubObjectVersionIds() {
	_jsii_.InvokeVoid(
		g,
		"resetSubObjectVersionIds",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCatalogTableViewDefinitionOutputReference) ResetViewVersionId() {
	_jsii_.InvokeVoid(
		g,
		"resetViewVersionId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCatalogTableViewDefinitionOutputReference) ResetViewVersionToken() {
	_jsii_.InvokeVoid(
		g,
		"resetViewVersionToken",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCatalogTableViewDefinitionOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GlueCatalogTableViewDefinitionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

