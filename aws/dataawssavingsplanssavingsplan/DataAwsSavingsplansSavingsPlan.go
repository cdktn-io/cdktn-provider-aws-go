// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawssavingsplanssavingsplan

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v23/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v23/dataawssavingsplanssavingsplan/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.35.1/docs/data-sources/savingsplans_savings_plan aws_savingsplans_savings_plan}.
type DataAwsSavingsplansSavingsPlan interface {
	cdktn.TerraformDataSource
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	Commitment() *string
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	Currency() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	Ec2InstanceFamily() *string
	End() *string
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
	OfferingId() *string
	PaymentOption() *string
	ProductTypes() *[]*string
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	PurchaseTime() *string
	// Experimental.
	RawOverrides() interface{}
	RecurringPaymentAmount() *string
	Region() *string
	ReturnableUntil() *string
	SavingsPlanArn() *string
	SavingsPlanId() *string
	SetSavingsPlanId(val *string)
	SavingsPlanIdInput() *string
	SavingsPlanOfferingId() *string
	SavingsPlanType() *string
	Start() *string
	State() *string
	Tags() cdktn.StringMap
	TermDurationInSeconds() *float64
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	UpfrontPaymentAmount() *string
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
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
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

// The jsii proxy struct for DataAwsSavingsplansSavingsPlan
type jsiiProxy_DataAwsSavingsplansSavingsPlan struct {
	internal.Type__cdktnTerraformDataSource
}

func (j *jsiiProxy_DataAwsSavingsplansSavingsPlan) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSavingsplansSavingsPlan) Commitment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commitment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSavingsplansSavingsPlan) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSavingsplansSavingsPlan) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSavingsplansSavingsPlan) Currency() *string {
	var returns *string
	_jsii_.Get(
		j,
		"currency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSavingsplansSavingsPlan) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSavingsplansSavingsPlan) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSavingsplansSavingsPlan) Ec2InstanceFamily() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ec2InstanceFamily",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSavingsplansSavingsPlan) End() *string {
	var returns *string
	_jsii_.Get(
		j,
		"end",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSavingsplansSavingsPlan) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSavingsplansSavingsPlan) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSavingsplansSavingsPlan) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSavingsplansSavingsPlan) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSavingsplansSavingsPlan) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSavingsplansSavingsPlan) OfferingId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"offeringId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSavingsplansSavingsPlan) PaymentOption() *string {
	var returns *string
	_jsii_.Get(
		j,
		"paymentOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSavingsplansSavingsPlan) ProductTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"productTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSavingsplansSavingsPlan) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSavingsplansSavingsPlan) PurchaseTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"purchaseTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSavingsplansSavingsPlan) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSavingsplansSavingsPlan) RecurringPaymentAmount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recurringPaymentAmount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSavingsplansSavingsPlan) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSavingsplansSavingsPlan) ReturnableUntil() *string {
	var returns *string
	_jsii_.Get(
		j,
		"returnableUntil",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSavingsplansSavingsPlan) SavingsPlanArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"savingsPlanArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSavingsplansSavingsPlan) SavingsPlanId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"savingsPlanId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSavingsplansSavingsPlan) SavingsPlanIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"savingsPlanIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSavingsplansSavingsPlan) SavingsPlanOfferingId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"savingsPlanOfferingId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSavingsplansSavingsPlan) SavingsPlanType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"savingsPlanType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSavingsplansSavingsPlan) Start() *string {
	var returns *string
	_jsii_.Get(
		j,
		"start",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSavingsplansSavingsPlan) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSavingsplansSavingsPlan) Tags() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSavingsplansSavingsPlan) TermDurationInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"termDurationInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSavingsplansSavingsPlan) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSavingsplansSavingsPlan) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSavingsplansSavingsPlan) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSavingsplansSavingsPlan) UpfrontPaymentAmount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"upfrontPaymentAmount",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.35.1/docs/data-sources/savingsplans_savings_plan aws_savingsplans_savings_plan} Data Source.
func NewDataAwsSavingsplansSavingsPlan(scope constructs.Construct, id *string, config *DataAwsSavingsplansSavingsPlanConfig) DataAwsSavingsplansSavingsPlan {
	_init_.Initialize()

	if err := validateNewDataAwsSavingsplansSavingsPlanParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsSavingsplansSavingsPlan{}

	_jsii_.Create(
		"@cdktn/provider-aws.dataAwsSavingsplansSavingsPlan.DataAwsSavingsplansSavingsPlan",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.35.1/docs/data-sources/savingsplans_savings_plan aws_savingsplans_savings_plan} Data Source.
func NewDataAwsSavingsplansSavingsPlan_Override(d DataAwsSavingsplansSavingsPlan, scope constructs.Construct, id *string, config *DataAwsSavingsplansSavingsPlanConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.dataAwsSavingsplansSavingsPlan.DataAwsSavingsplansSavingsPlan",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsSavingsplansSavingsPlan)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsSavingsplansSavingsPlan)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsSavingsplansSavingsPlan)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataAwsSavingsplansSavingsPlan)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsSavingsplansSavingsPlan)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataAwsSavingsplansSavingsPlan)SetSavingsPlanId(val *string) {
	if err := j.validateSetSavingsPlanIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"savingsPlanId",
		val,
	)
}

// Generates CDKTN code for importing a DataAwsSavingsplansSavingsPlan resource upon running "cdktn plan <stack-name>".
func DataAwsSavingsplansSavingsPlan_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateDataAwsSavingsplansSavingsPlan_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-aws.dataAwsSavingsplansSavingsPlan.DataAwsSavingsplansSavingsPlan",
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
func DataAwsSavingsplansSavingsPlan_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsSavingsplansSavingsPlan_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-aws.dataAwsSavingsplansSavingsPlan.DataAwsSavingsplansSavingsPlan",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAwsSavingsplansSavingsPlan_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsSavingsplansSavingsPlan_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-aws.dataAwsSavingsplansSavingsPlan.DataAwsSavingsplansSavingsPlan",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAwsSavingsplansSavingsPlan_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsSavingsplansSavingsPlan_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-aws.dataAwsSavingsplansSavingsPlan.DataAwsSavingsplansSavingsPlan",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsSavingsplansSavingsPlan_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-aws.dataAwsSavingsplansSavingsPlan.DataAwsSavingsplansSavingsPlan",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataAwsSavingsplansSavingsPlan) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataAwsSavingsplansSavingsPlan) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsSavingsplansSavingsPlan) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsSavingsplansSavingsPlan) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsSavingsplansSavingsPlan) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsSavingsplansSavingsPlan) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsSavingsplansSavingsPlan) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsSavingsplansSavingsPlan) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsSavingsplansSavingsPlan) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsSavingsplansSavingsPlan) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsSavingsplansSavingsPlan) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsSavingsplansSavingsPlan) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataAwsSavingsplansSavingsPlan) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsSavingsplansSavingsPlan) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsSavingsplansSavingsPlan) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsSavingsplansSavingsPlan) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsSavingsplansSavingsPlan) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsSavingsplansSavingsPlan) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsSavingsplansSavingsPlan) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsSavingsplansSavingsPlan) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

