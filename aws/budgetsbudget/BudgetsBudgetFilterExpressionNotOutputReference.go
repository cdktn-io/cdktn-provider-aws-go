// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package budgetsbudget

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/jsii"

	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/budgetsbudget/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BudgetsBudgetFilterExpressionNotOutputReference interface {
	cdktn.ComplexObject
	And() BudgetsBudgetFilterExpressionNotAndList
	AndInput() interface{}
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
	CostCategories() BudgetsBudgetFilterExpressionNotCostCategoriesOutputReference
	CostCategoriesInput() *BudgetsBudgetFilterExpressionNotCostCategories
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Dimensions() BudgetsBudgetFilterExpressionNotDimensionsOutputReference
	DimensionsInput() *BudgetsBudgetFilterExpressionNotDimensions
	// Experimental.
	Fqn() *string
	InternalValue() *BudgetsBudgetFilterExpressionNot
	SetInternalValue(val *BudgetsBudgetFilterExpressionNot)
	Not() BudgetsBudgetFilterExpressionNotNotOutputReference
	NotInput() *BudgetsBudgetFilterExpressionNotNot
	Or() BudgetsBudgetFilterExpressionNotOrList
	OrInput() interface{}
	Tags() BudgetsBudgetFilterExpressionNotTagsOutputReference
	TagsInput() *BudgetsBudgetFilterExpressionNotTags
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
	PutAnd(value interface{})
	PutCostCategories(value *BudgetsBudgetFilterExpressionNotCostCategories)
	PutDimensions(value *BudgetsBudgetFilterExpressionNotDimensions)
	PutNot(value *BudgetsBudgetFilterExpressionNotNot)
	PutOr(value interface{})
	PutTags(value *BudgetsBudgetFilterExpressionNotTags)
	ResetAnd()
	ResetCostCategories()
	ResetDimensions()
	ResetNot()
	ResetOr()
	ResetTags()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BudgetsBudgetFilterExpressionNotOutputReference
type jsiiProxy_BudgetsBudgetFilterExpressionNotOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_BudgetsBudgetFilterExpressionNotOutputReference) And() BudgetsBudgetFilterExpressionNotAndList {
	var returns BudgetsBudgetFilterExpressionNotAndList
	_jsii_.Get(
		j,
		"and",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetFilterExpressionNotOutputReference) AndInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"andInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetFilterExpressionNotOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetFilterExpressionNotOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetFilterExpressionNotOutputReference) CostCategories() BudgetsBudgetFilterExpressionNotCostCategoriesOutputReference {
	var returns BudgetsBudgetFilterExpressionNotCostCategoriesOutputReference
	_jsii_.Get(
		j,
		"costCategories",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetFilterExpressionNotOutputReference) CostCategoriesInput() *BudgetsBudgetFilterExpressionNotCostCategories {
	var returns *BudgetsBudgetFilterExpressionNotCostCategories
	_jsii_.Get(
		j,
		"costCategoriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetFilterExpressionNotOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetFilterExpressionNotOutputReference) Dimensions() BudgetsBudgetFilterExpressionNotDimensionsOutputReference {
	var returns BudgetsBudgetFilterExpressionNotDimensionsOutputReference
	_jsii_.Get(
		j,
		"dimensions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetFilterExpressionNotOutputReference) DimensionsInput() *BudgetsBudgetFilterExpressionNotDimensions {
	var returns *BudgetsBudgetFilterExpressionNotDimensions
	_jsii_.Get(
		j,
		"dimensionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetFilterExpressionNotOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetFilterExpressionNotOutputReference) InternalValue() *BudgetsBudgetFilterExpressionNot {
	var returns *BudgetsBudgetFilterExpressionNot
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetFilterExpressionNotOutputReference) Not() BudgetsBudgetFilterExpressionNotNotOutputReference {
	var returns BudgetsBudgetFilterExpressionNotNotOutputReference
	_jsii_.Get(
		j,
		"not",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetFilterExpressionNotOutputReference) NotInput() *BudgetsBudgetFilterExpressionNotNot {
	var returns *BudgetsBudgetFilterExpressionNotNot
	_jsii_.Get(
		j,
		"notInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetFilterExpressionNotOutputReference) Or() BudgetsBudgetFilterExpressionNotOrList {
	var returns BudgetsBudgetFilterExpressionNotOrList
	_jsii_.Get(
		j,
		"or",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetFilterExpressionNotOutputReference) OrInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"orInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetFilterExpressionNotOutputReference) Tags() BudgetsBudgetFilterExpressionNotTagsOutputReference {
	var returns BudgetsBudgetFilterExpressionNotTagsOutputReference
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetFilterExpressionNotOutputReference) TagsInput() *BudgetsBudgetFilterExpressionNotTags {
	var returns *BudgetsBudgetFilterExpressionNotTags
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetFilterExpressionNotOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetFilterExpressionNotOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBudgetsBudgetFilterExpressionNotOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) BudgetsBudgetFilterExpressionNotOutputReference {
	_init_.Initialize()

	if err := validateNewBudgetsBudgetFilterExpressionNotOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_BudgetsBudgetFilterExpressionNotOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-aws.budgetsBudget.BudgetsBudgetFilterExpressionNotOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBudgetsBudgetFilterExpressionNotOutputReference_Override(b BudgetsBudgetFilterExpressionNotOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.budgetsBudget.BudgetsBudgetFilterExpressionNotOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BudgetsBudgetFilterExpressionNotOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudgetFilterExpressionNotOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudgetFilterExpressionNotOutputReference)SetInternalValue(val *BudgetsBudgetFilterExpressionNot) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudgetFilterExpressionNotOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudgetFilterExpressionNotOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BudgetsBudgetFilterExpressionNotOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BudgetsBudgetFilterExpressionNotOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := b.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BudgetsBudgetFilterExpressionNotOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := b.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BudgetsBudgetFilterExpressionNotOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := b.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BudgetsBudgetFilterExpressionNotOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := b.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BudgetsBudgetFilterExpressionNotOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := b.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BudgetsBudgetFilterExpressionNotOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := b.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BudgetsBudgetFilterExpressionNotOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := b.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BudgetsBudgetFilterExpressionNotOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := b.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BudgetsBudgetFilterExpressionNotOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := b.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BudgetsBudgetFilterExpressionNotOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BudgetsBudgetFilterExpressionNotOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BudgetsBudgetFilterExpressionNotOutputReference) PutAnd(value interface{}) {
	if err := b.validatePutAndParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putAnd",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BudgetsBudgetFilterExpressionNotOutputReference) PutCostCategories(value *BudgetsBudgetFilterExpressionNotCostCategories) {
	if err := b.validatePutCostCategoriesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putCostCategories",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BudgetsBudgetFilterExpressionNotOutputReference) PutDimensions(value *BudgetsBudgetFilterExpressionNotDimensions) {
	if err := b.validatePutDimensionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putDimensions",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BudgetsBudgetFilterExpressionNotOutputReference) PutNot(value *BudgetsBudgetFilterExpressionNotNot) {
	if err := b.validatePutNotParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putNot",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BudgetsBudgetFilterExpressionNotOutputReference) PutOr(value interface{}) {
	if err := b.validatePutOrParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putOr",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BudgetsBudgetFilterExpressionNotOutputReference) PutTags(value *BudgetsBudgetFilterExpressionNotTags) {
	if err := b.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putTags",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BudgetsBudgetFilterExpressionNotOutputReference) ResetAnd() {
	_jsii_.InvokeVoid(
		b,
		"resetAnd",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BudgetsBudgetFilterExpressionNotOutputReference) ResetCostCategories() {
	_jsii_.InvokeVoid(
		b,
		"resetCostCategories",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BudgetsBudgetFilterExpressionNotOutputReference) ResetDimensions() {
	_jsii_.InvokeVoid(
		b,
		"resetDimensions",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BudgetsBudgetFilterExpressionNotOutputReference) ResetNot() {
	_jsii_.InvokeVoid(
		b,
		"resetNot",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BudgetsBudgetFilterExpressionNotOutputReference) ResetOr() {
	_jsii_.InvokeVoid(
		b,
		"resetOr",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BudgetsBudgetFilterExpressionNotOutputReference) ResetTags() {
	_jsii_.InvokeVoid(
		b,
		"resetTags",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BudgetsBudgetFilterExpressionNotOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (b *jsiiProxy_BudgetsBudgetFilterExpressionNotOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

