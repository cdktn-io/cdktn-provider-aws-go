// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package budgetsbudget

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v23/jsii"

	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v23/budgetsbudget/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BudgetsBudgetFilterExpressionAndOutputReference interface {
	cdktn.ComplexObject
	And() BudgetsBudgetFilterExpressionAndAndList
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
	CostCategories() BudgetsBudgetFilterExpressionAndCostCategoriesOutputReference
	CostCategoriesInput() *BudgetsBudgetFilterExpressionAndCostCategories
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Dimensions() BudgetsBudgetFilterExpressionAndDimensionsOutputReference
	DimensionsInput() *BudgetsBudgetFilterExpressionAndDimensions
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Not() BudgetsBudgetFilterExpressionAndNotOutputReference
	NotInput() *BudgetsBudgetFilterExpressionAndNot
	Or() BudgetsBudgetFilterExpressionAndOrList
	OrInput() interface{}
	Tags() BudgetsBudgetFilterExpressionAndTagsOutputReference
	TagsInput() *BudgetsBudgetFilterExpressionAndTags
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
	PutCostCategories(value *BudgetsBudgetFilterExpressionAndCostCategories)
	PutDimensions(value *BudgetsBudgetFilterExpressionAndDimensions)
	PutNot(value *BudgetsBudgetFilterExpressionAndNot)
	PutOr(value interface{})
	PutTags(value *BudgetsBudgetFilterExpressionAndTags)
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

// The jsii proxy struct for BudgetsBudgetFilterExpressionAndOutputReference
type jsiiProxy_BudgetsBudgetFilterExpressionAndOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_BudgetsBudgetFilterExpressionAndOutputReference) And() BudgetsBudgetFilterExpressionAndAndList {
	var returns BudgetsBudgetFilterExpressionAndAndList
	_jsii_.Get(
		j,
		"and",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetFilterExpressionAndOutputReference) AndInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"andInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetFilterExpressionAndOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetFilterExpressionAndOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetFilterExpressionAndOutputReference) CostCategories() BudgetsBudgetFilterExpressionAndCostCategoriesOutputReference {
	var returns BudgetsBudgetFilterExpressionAndCostCategoriesOutputReference
	_jsii_.Get(
		j,
		"costCategories",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetFilterExpressionAndOutputReference) CostCategoriesInput() *BudgetsBudgetFilterExpressionAndCostCategories {
	var returns *BudgetsBudgetFilterExpressionAndCostCategories
	_jsii_.Get(
		j,
		"costCategoriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetFilterExpressionAndOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetFilterExpressionAndOutputReference) Dimensions() BudgetsBudgetFilterExpressionAndDimensionsOutputReference {
	var returns BudgetsBudgetFilterExpressionAndDimensionsOutputReference
	_jsii_.Get(
		j,
		"dimensions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetFilterExpressionAndOutputReference) DimensionsInput() *BudgetsBudgetFilterExpressionAndDimensions {
	var returns *BudgetsBudgetFilterExpressionAndDimensions
	_jsii_.Get(
		j,
		"dimensionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetFilterExpressionAndOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetFilterExpressionAndOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetFilterExpressionAndOutputReference) Not() BudgetsBudgetFilterExpressionAndNotOutputReference {
	var returns BudgetsBudgetFilterExpressionAndNotOutputReference
	_jsii_.Get(
		j,
		"not",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetFilterExpressionAndOutputReference) NotInput() *BudgetsBudgetFilterExpressionAndNot {
	var returns *BudgetsBudgetFilterExpressionAndNot
	_jsii_.Get(
		j,
		"notInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetFilterExpressionAndOutputReference) Or() BudgetsBudgetFilterExpressionAndOrList {
	var returns BudgetsBudgetFilterExpressionAndOrList
	_jsii_.Get(
		j,
		"or",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetFilterExpressionAndOutputReference) OrInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"orInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetFilterExpressionAndOutputReference) Tags() BudgetsBudgetFilterExpressionAndTagsOutputReference {
	var returns BudgetsBudgetFilterExpressionAndTagsOutputReference
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetFilterExpressionAndOutputReference) TagsInput() *BudgetsBudgetFilterExpressionAndTags {
	var returns *BudgetsBudgetFilterExpressionAndTags
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetFilterExpressionAndOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetFilterExpressionAndOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBudgetsBudgetFilterExpressionAndOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) BudgetsBudgetFilterExpressionAndOutputReference {
	_init_.Initialize()

	if err := validateNewBudgetsBudgetFilterExpressionAndOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_BudgetsBudgetFilterExpressionAndOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-aws.budgetsBudget.BudgetsBudgetFilterExpressionAndOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewBudgetsBudgetFilterExpressionAndOutputReference_Override(b BudgetsBudgetFilterExpressionAndOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.budgetsBudget.BudgetsBudgetFilterExpressionAndOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		b,
	)
}

func (j *jsiiProxy_BudgetsBudgetFilterExpressionAndOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudgetFilterExpressionAndOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudgetFilterExpressionAndOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudgetFilterExpressionAndOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudgetFilterExpressionAndOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BudgetsBudgetFilterExpressionAndOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BudgetsBudgetFilterExpressionAndOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (b *jsiiProxy_BudgetsBudgetFilterExpressionAndOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (b *jsiiProxy_BudgetsBudgetFilterExpressionAndOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (b *jsiiProxy_BudgetsBudgetFilterExpressionAndOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (b *jsiiProxy_BudgetsBudgetFilterExpressionAndOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (b *jsiiProxy_BudgetsBudgetFilterExpressionAndOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (b *jsiiProxy_BudgetsBudgetFilterExpressionAndOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (b *jsiiProxy_BudgetsBudgetFilterExpressionAndOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (b *jsiiProxy_BudgetsBudgetFilterExpressionAndOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (b *jsiiProxy_BudgetsBudgetFilterExpressionAndOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BudgetsBudgetFilterExpressionAndOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (b *jsiiProxy_BudgetsBudgetFilterExpressionAndOutputReference) PutAnd(value interface{}) {
	if err := b.validatePutAndParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putAnd",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BudgetsBudgetFilterExpressionAndOutputReference) PutCostCategories(value *BudgetsBudgetFilterExpressionAndCostCategories) {
	if err := b.validatePutCostCategoriesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putCostCategories",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BudgetsBudgetFilterExpressionAndOutputReference) PutDimensions(value *BudgetsBudgetFilterExpressionAndDimensions) {
	if err := b.validatePutDimensionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putDimensions",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BudgetsBudgetFilterExpressionAndOutputReference) PutNot(value *BudgetsBudgetFilterExpressionAndNot) {
	if err := b.validatePutNotParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putNot",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BudgetsBudgetFilterExpressionAndOutputReference) PutOr(value interface{}) {
	if err := b.validatePutOrParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putOr",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BudgetsBudgetFilterExpressionAndOutputReference) PutTags(value *BudgetsBudgetFilterExpressionAndTags) {
	if err := b.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putTags",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BudgetsBudgetFilterExpressionAndOutputReference) ResetAnd() {
	_jsii_.InvokeVoid(
		b,
		"resetAnd",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BudgetsBudgetFilterExpressionAndOutputReference) ResetCostCategories() {
	_jsii_.InvokeVoid(
		b,
		"resetCostCategories",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BudgetsBudgetFilterExpressionAndOutputReference) ResetDimensions() {
	_jsii_.InvokeVoid(
		b,
		"resetDimensions",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BudgetsBudgetFilterExpressionAndOutputReference) ResetNot() {
	_jsii_.InvokeVoid(
		b,
		"resetNot",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BudgetsBudgetFilterExpressionAndOutputReference) ResetOr() {
	_jsii_.InvokeVoid(
		b,
		"resetOr",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BudgetsBudgetFilterExpressionAndOutputReference) ResetTags() {
	_jsii_.InvokeVoid(
		b,
		"resetTags",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BudgetsBudgetFilterExpressionAndOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (b *jsiiProxy_BudgetsBudgetFilterExpressionAndOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

