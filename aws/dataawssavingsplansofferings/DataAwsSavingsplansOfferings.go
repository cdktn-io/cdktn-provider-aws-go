// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawssavingsplansofferings

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v24/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v24/dataawssavingsplansofferings/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.56.0/docs/data-sources/savingsplans_offerings aws_savingsplans_offerings}.
type DataAwsSavingsplansOfferings interface {
	cdktn.TerraformDataSource
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	Currencies() *[]*string
	SetCurrencies(val *[]*string)
	CurrenciesInput() *[]*string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Descriptions() *[]*string
	SetDescriptions(val *[]*string)
	DescriptionsInput() *[]*string
	Durations() *[]*float64
	SetDurations(val *[]*float64)
	DurationsInput() *[]*float64
	Filter() DataAwsSavingsplansOfferingsFilterList
	FilterInput() interface{}
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	OfferingIds() *[]*string
	SetOfferingIds(val *[]*string)
	OfferingIdsInput() *[]*string
	Offerings() DataAwsSavingsplansOfferingsOfferingsList
	Operations() *[]*string
	SetOperations(val *[]*string)
	OperationsInput() *[]*string
	PaymentOptions() *[]*string
	SetPaymentOptions(val *[]*string)
	PaymentOptionsInput() *[]*string
	PlanTypes() *[]*string
	SetPlanTypes(val *[]*string)
	PlanTypesInput() *[]*string
	ProductType() *string
	SetProductType(val *string)
	ProductTypeInput() *string
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	ServiceCodes() *[]*string
	SetServiceCodes(val *[]*string)
	ServiceCodesInput() *[]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	UsageTypes() *[]*string
	SetUsageTypes(val *[]*string)
	UsageTypesInput() *[]*string
	// Experimental.
	AddOverride(path *string, value interface{})
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
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutFilter(value interface{})
	ResetCurrencies()
	ResetDescriptions()
	ResetDurations()
	ResetFilter()
	ResetOfferingIds()
	ResetOperations()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPaymentOptions()
	ResetPlanTypes()
	ResetProductType()
	ResetServiceCodes()
	ResetUsageTypes()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
	// Applies one or more mixins to this construct.
	//
	// Mixins are applied in order. The list of constructs is captured at the
	// start of the call, so constructs added by a mixin will not be visited.
	// Use multiple `with()` calls if subsequent mixins should apply to added
	// constructs.
	//
	// Returns: This construct for chaining.
	With(mixins ...constructs.IMixin) constructs.IConstruct
}

// The jsii proxy struct for DataAwsSavingsplansOfferings
type jsiiProxy_DataAwsSavingsplansOfferings struct {
	internal.Type__cdktnTerraformDataSource
}

func (j *jsiiProxy_DataAwsSavingsplansOfferings) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSavingsplansOfferings) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSavingsplansOfferings) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSavingsplansOfferings) Currencies() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"currencies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSavingsplansOfferings) CurrenciesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"currenciesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSavingsplansOfferings) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSavingsplansOfferings) Descriptions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"descriptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSavingsplansOfferings) DescriptionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"descriptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSavingsplansOfferings) Durations() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"durations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSavingsplansOfferings) DurationsInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"durationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSavingsplansOfferings) Filter() DataAwsSavingsplansOfferingsFilterList {
	var returns DataAwsSavingsplansOfferingsFilterList
	_jsii_.Get(
		j,
		"filter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSavingsplansOfferings) FilterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"filterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSavingsplansOfferings) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSavingsplansOfferings) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSavingsplansOfferings) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSavingsplansOfferings) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSavingsplansOfferings) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSavingsplansOfferings) OfferingIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"offeringIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSavingsplansOfferings) OfferingIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"offeringIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSavingsplansOfferings) Offerings() DataAwsSavingsplansOfferingsOfferingsList {
	var returns DataAwsSavingsplansOfferingsOfferingsList
	_jsii_.Get(
		j,
		"offerings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSavingsplansOfferings) Operations() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"operations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSavingsplansOfferings) OperationsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"operationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSavingsplansOfferings) PaymentOptions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"paymentOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSavingsplansOfferings) PaymentOptionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"paymentOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSavingsplansOfferings) PlanTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"planTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSavingsplansOfferings) PlanTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"planTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSavingsplansOfferings) ProductType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"productType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSavingsplansOfferings) ProductTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"productTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSavingsplansOfferings) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSavingsplansOfferings) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSavingsplansOfferings) ServiceCodes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"serviceCodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSavingsplansOfferings) ServiceCodesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"serviceCodesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSavingsplansOfferings) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSavingsplansOfferings) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSavingsplansOfferings) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSavingsplansOfferings) UsageTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"usageTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSavingsplansOfferings) UsageTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"usageTypesInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.56.0/docs/data-sources/savingsplans_offerings aws_savingsplans_offerings} Data Source.
func NewDataAwsSavingsplansOfferings(scope constructs.Construct, id *string, config *DataAwsSavingsplansOfferingsConfig) DataAwsSavingsplansOfferings {
	_init_.Initialize()

	if err := validateNewDataAwsSavingsplansOfferingsParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsSavingsplansOfferings{}

	_jsii_.Create(
		"@cdktn/provider-aws.dataAwsSavingsplansOfferings.DataAwsSavingsplansOfferings",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.56.0/docs/data-sources/savingsplans_offerings aws_savingsplans_offerings} Data Source.
func NewDataAwsSavingsplansOfferings_Override(d DataAwsSavingsplansOfferings, scope constructs.Construct, id *string, config *DataAwsSavingsplansOfferingsConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.dataAwsSavingsplansOfferings.DataAwsSavingsplansOfferings",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsSavingsplansOfferings)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsSavingsplansOfferings)SetCurrencies(val *[]*string) {
	if err := j.validateSetCurrenciesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"currencies",
		val,
	)
}

func (j *jsiiProxy_DataAwsSavingsplansOfferings)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsSavingsplansOfferings)SetDescriptions(val *[]*string) {
	if err := j.validateSetDescriptionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"descriptions",
		val,
	)
}

func (j *jsiiProxy_DataAwsSavingsplansOfferings)SetDurations(val *[]*float64) {
	if err := j.validateSetDurationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"durations",
		val,
	)
}

func (j *jsiiProxy_DataAwsSavingsplansOfferings)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataAwsSavingsplansOfferings)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsSavingsplansOfferings)SetOfferingIds(val *[]*string) {
	if err := j.validateSetOfferingIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"offeringIds",
		val,
	)
}

func (j *jsiiProxy_DataAwsSavingsplansOfferings)SetOperations(val *[]*string) {
	if err := j.validateSetOperationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"operations",
		val,
	)
}

func (j *jsiiProxy_DataAwsSavingsplansOfferings)SetPaymentOptions(val *[]*string) {
	if err := j.validateSetPaymentOptionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"paymentOptions",
		val,
	)
}

func (j *jsiiProxy_DataAwsSavingsplansOfferings)SetPlanTypes(val *[]*string) {
	if err := j.validateSetPlanTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"planTypes",
		val,
	)
}

func (j *jsiiProxy_DataAwsSavingsplansOfferings)SetProductType(val *string) {
	if err := j.validateSetProductTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"productType",
		val,
	)
}

func (j *jsiiProxy_DataAwsSavingsplansOfferings)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataAwsSavingsplansOfferings)SetServiceCodes(val *[]*string) {
	if err := j.validateSetServiceCodesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceCodes",
		val,
	)
}

func (j *jsiiProxy_DataAwsSavingsplansOfferings)SetUsageTypes(val *[]*string) {
	if err := j.validateSetUsageTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"usageTypes",
		val,
	)
}

// Generates CDKTN code for importing a DataAwsSavingsplansOfferings resource upon running "cdktn plan <stack-name>".
func DataAwsSavingsplansOfferings_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateDataAwsSavingsplansOfferings_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-aws.dataAwsSavingsplansOfferings.DataAwsSavingsplansOfferings",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func DataAwsSavingsplansOfferings_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsSavingsplansOfferings_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-aws.dataAwsSavingsplansOfferings.DataAwsSavingsplansOfferings",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAwsSavingsplansOfferings_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsSavingsplansOfferings_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-aws.dataAwsSavingsplansOfferings.DataAwsSavingsplansOfferings",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAwsSavingsplansOfferings_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsSavingsplansOfferings_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-aws.dataAwsSavingsplansOfferings.DataAwsSavingsplansOfferings",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsSavingsplansOfferings_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-aws.dataAwsSavingsplansOfferings.DataAwsSavingsplansOfferings",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataAwsSavingsplansOfferings) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataAwsSavingsplansOfferings) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsSavingsplansOfferings) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsSavingsplansOfferings) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsSavingsplansOfferings) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsSavingsplansOfferings) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsSavingsplansOfferings) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsSavingsplansOfferings) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsSavingsplansOfferings) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsSavingsplansOfferings) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsSavingsplansOfferings) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsSavingsplansOfferings) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataAwsSavingsplansOfferings) PutFilter(value interface{}) {
	if err := d.validatePutFilterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putFilter",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataAwsSavingsplansOfferings) ResetCurrencies() {
	_jsii_.InvokeVoid(
		d,
		"resetCurrencies",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsSavingsplansOfferings) ResetDescriptions() {
	_jsii_.InvokeVoid(
		d,
		"resetDescriptions",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsSavingsplansOfferings) ResetDurations() {
	_jsii_.InvokeVoid(
		d,
		"resetDurations",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsSavingsplansOfferings) ResetFilter() {
	_jsii_.InvokeVoid(
		d,
		"resetFilter",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsSavingsplansOfferings) ResetOfferingIds() {
	_jsii_.InvokeVoid(
		d,
		"resetOfferingIds",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsSavingsplansOfferings) ResetOperations() {
	_jsii_.InvokeVoid(
		d,
		"resetOperations",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsSavingsplansOfferings) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsSavingsplansOfferings) ResetPaymentOptions() {
	_jsii_.InvokeVoid(
		d,
		"resetPaymentOptions",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsSavingsplansOfferings) ResetPlanTypes() {
	_jsii_.InvokeVoid(
		d,
		"resetPlanTypes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsSavingsplansOfferings) ResetProductType() {
	_jsii_.InvokeVoid(
		d,
		"resetProductType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsSavingsplansOfferings) ResetServiceCodes() {
	_jsii_.InvokeVoid(
		d,
		"resetServiceCodes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsSavingsplansOfferings) ResetUsageTypes() {
	_jsii_.InvokeVoid(
		d,
		"resetUsageTypes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsSavingsplansOfferings) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsSavingsplansOfferings) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsSavingsplansOfferings) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsSavingsplansOfferings) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsSavingsplansOfferings) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsSavingsplansOfferings) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsSavingsplansOfferings) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		d,
		"with",
		args,
		&returns,
	)

	return returns
}

