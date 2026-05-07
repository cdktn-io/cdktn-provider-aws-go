// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package budgetsbudget

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v24/jsii"

	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v24/budgetsbudget/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BudgetsBudgetFilterExpressionOrOutputReference interface {
	cdktn.ComplexObject
	And() BudgetsBudgetFilterExpressionOrAndList
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
	CostCategories() BudgetsBudgetFilterExpressionOrCostCategoriesOutputReference
	CostCategoriesInput() *BudgetsBudgetFilterExpressionOrCostCategories
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Dimensions() BudgetsBudgetFilterExpressionOrDimensionsOutputReference
	DimensionsInput() *BudgetsBudgetFilterExpressionOrDimensions
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Not() BudgetsBudgetFilterExpressionOrNotOutputReference
	NotInput() *BudgetsBudgetFilterExpressionOrNot
	Or() BudgetsBudgetFilterExpressionOrOrList
	OrInput() interface{}
	Tags() BudgetsBudgetFilterExpressionOrTagsOutputReference
	TagsInput() *BudgetsBudgetFilterExpressionOrTags
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
	PutCostCategories(value *BudgetsBudgetFilterExpressionOrCostCategories)
	PutDimensions(value *BudgetsBudgetFilterExpressionOrDimensions)
	PutNot(value *BudgetsBudgetFilterExpressionOrNot)
	PutOr(value interface{})
	PutTags(value *BudgetsBudgetFilterExpressionOrTags)
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

// The jsii proxy struct for BudgetsBudgetFilterExpressionOrOutputReference
type jsiiProxy_BudgetsBudgetFilterExpressionOrOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_BudgetsBudgetFilterExpressionOrOutputReference) And() BudgetsBudgetFilterExpressionOrAndList {
	var returns BudgetsBudgetFilterExpressionOrAndList
	_jsii_.Get(
		j,
		"and",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetFilterExpressionOrOutputReference) AndInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"andInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetFilterExpressionOrOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetFilterExpressionOrOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetFilterExpressionOrOutputReference) CostCategories() BudgetsBudgetFilterExpressionOrCostCategoriesOutputReference {
	var returns BudgetsBudgetFilterExpressionOrCostCategoriesOutputReference
	_jsii_.Get(
		j,
		"costCategories",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetFilterExpressionOrOutputReference) CostCategoriesInput() *BudgetsBudgetFilterExpressionOrCostCategories {
	var returns *BudgetsBudgetFilterExpressionOrCostCategories
	_jsii_.Get(
		j,
		"costCategoriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetFilterExpressionOrOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetFilterExpressionOrOutputReference) Dimensions() BudgetsBudgetFilterExpressionOrDimensionsOutputReference {
	var returns BudgetsBudgetFilterExpressionOrDimensionsOutputReference
	_jsii_.Get(
		j,
		"dimensions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetFilterExpressionOrOutputReference) DimensionsInput() *BudgetsBudgetFilterExpressionOrDimensions {
	var returns *BudgetsBudgetFilterExpressionOrDimensions
	_jsii_.Get(
		j,
		"dimensionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetFilterExpressionOrOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetFilterExpressionOrOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetFilterExpressionOrOutputReference) Not() BudgetsBudgetFilterExpressionOrNotOutputReference {
	var returns BudgetsBudgetFilterExpressionOrNotOutputReference
	_jsii_.Get(
		j,
		"not",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetFilterExpressionOrOutputReference) NotInput() *BudgetsBudgetFilterExpressionOrNot {
	var returns *BudgetsBudgetFilterExpressionOrNot
	_jsii_.Get(
		j,
		"notInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetFilterExpressionOrOutputReference) Or() BudgetsBudgetFilterExpressionOrOrList {
	var returns BudgetsBudgetFilterExpressionOrOrList
	_jsii_.Get(
		j,
		"or",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetFilterExpressionOrOutputReference) OrInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"orInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetFilterExpressionOrOutputReference) Tags() BudgetsBudgetFilterExpressionOrTagsOutputReference {
	var returns BudgetsBudgetFilterExpressionOrTagsOutputReference
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetFilterExpressionOrOutputReference) TagsInput() *BudgetsBudgetFilterExpressionOrTags {
	var returns *BudgetsBudgetFilterExpressionOrTags
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetFilterExpressionOrOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetFilterExpressionOrOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBudgetsBudgetFilterExpressionOrOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) BudgetsBudgetFilterExpressionOrOutputReference {
	_init_.Initialize()

	if err := validateNewBudgetsBudgetFilterExpressionOrOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_BudgetsBudgetFilterExpressionOrOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-aws.budgetsBudget.BudgetsBudgetFilterExpressionOrOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewBudgetsBudgetFilterExpressionOrOutputReference_Override(b BudgetsBudgetFilterExpressionOrOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.budgetsBudget.BudgetsBudgetFilterExpressionOrOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		b,
	)
}

func (j *jsiiProxy_BudgetsBudgetFilterExpressionOrOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudgetFilterExpressionOrOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudgetFilterExpressionOrOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudgetFilterExpressionOrOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudgetFilterExpressionOrOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BudgetsBudgetFilterExpressionOrOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BudgetsBudgetFilterExpressionOrOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (b *jsiiProxy_BudgetsBudgetFilterExpressionOrOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (b *jsiiProxy_BudgetsBudgetFilterExpressionOrOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (b *jsiiProxy_BudgetsBudgetFilterExpressionOrOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (b *jsiiProxy_BudgetsBudgetFilterExpressionOrOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (b *jsiiProxy_BudgetsBudgetFilterExpressionOrOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (b *jsiiProxy_BudgetsBudgetFilterExpressionOrOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (b *jsiiProxy_BudgetsBudgetFilterExpressionOrOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (b *jsiiProxy_BudgetsBudgetFilterExpressionOrOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (b *jsiiProxy_BudgetsBudgetFilterExpressionOrOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BudgetsBudgetFilterExpressionOrOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (b *jsiiProxy_BudgetsBudgetFilterExpressionOrOutputReference) PutAnd(value interface{}) {
	if err := b.validatePutAndParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putAnd",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BudgetsBudgetFilterExpressionOrOutputReference) PutCostCategories(value *BudgetsBudgetFilterExpressionOrCostCategories) {
	if err := b.validatePutCostCategoriesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putCostCategories",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BudgetsBudgetFilterExpressionOrOutputReference) PutDimensions(value *BudgetsBudgetFilterExpressionOrDimensions) {
	if err := b.validatePutDimensionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putDimensions",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BudgetsBudgetFilterExpressionOrOutputReference) PutNot(value *BudgetsBudgetFilterExpressionOrNot) {
	if err := b.validatePutNotParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putNot",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BudgetsBudgetFilterExpressionOrOutputReference) PutOr(value interface{}) {
	if err := b.validatePutOrParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putOr",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BudgetsBudgetFilterExpressionOrOutputReference) PutTags(value *BudgetsBudgetFilterExpressionOrTags) {
	if err := b.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putTags",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BudgetsBudgetFilterExpressionOrOutputReference) ResetAnd() {
	_jsii_.InvokeVoid(
		b,
		"resetAnd",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BudgetsBudgetFilterExpressionOrOutputReference) ResetCostCategories() {
	_jsii_.InvokeVoid(
		b,
		"resetCostCategories",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BudgetsBudgetFilterExpressionOrOutputReference) ResetDimensions() {
	_jsii_.InvokeVoid(
		b,
		"resetDimensions",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BudgetsBudgetFilterExpressionOrOutputReference) ResetNot() {
	_jsii_.InvokeVoid(
		b,
		"resetNot",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BudgetsBudgetFilterExpressionOrOutputReference) ResetOr() {
	_jsii_.InvokeVoid(
		b,
		"resetOr",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BudgetsBudgetFilterExpressionOrOutputReference) ResetTags() {
	_jsii_.InvokeVoid(
		b,
		"resetTags",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BudgetsBudgetFilterExpressionOrOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (b *jsiiProxy_BudgetsBudgetFilterExpressionOrOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

