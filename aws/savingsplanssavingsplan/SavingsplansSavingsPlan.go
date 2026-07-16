// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package savingsplanssavingsplan

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v24/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v24/savingsplanssavingsplan/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.55.0/docs/resources/savingsplans_savings_plan aws_savingsplans_savings_plan}.
type SavingsplansSavingsPlan interface {
	cdktn.TerraformResource
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	Commitment() *string
	SetCommitment(val *string)
	CommitmentInput() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
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
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	PurchaseTime() *string
	SetPurchaseTime(val *string)
	PurchaseTimeInput() *string
	// Experimental.
	RawOverrides() interface{}
	RecurringPaymentAmount() *string
	Region() *string
	ReturnableUntil() *string
	SavingsPlanArn() *string
	SavingsPlanId() *string
	SavingsPlanOfferingId() *string
	SetSavingsPlanOfferingId(val *string)
	SavingsPlanOfferingIdInput() *string
	SavingsPlanType() *string
	Start() *string
	State() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsAll() cdktn.StringMap
	TagsInput() *map[string]*string
	TermDurationInSeconds() *float64
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() SavingsplansSavingsPlanTimeoutsOutputReference
	TimeoutsInput() interface{}
	UpfrontPaymentAmount() *string
	SetUpfrontPaymentAmount(val *string)
	UpfrontPaymentAmountInput() *string
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
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
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktn.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutTimeouts(value *SavingsplansSavingsPlanTimeouts)
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPurchaseTime()
	ResetTags()
	ResetTimeouts()
	ResetUpfrontPaymentAmount()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
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

// The jsii proxy struct for SavingsplansSavingsPlan
type jsiiProxy_SavingsplansSavingsPlan struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_SavingsplansSavingsPlan) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SavingsplansSavingsPlan) Commitment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commitment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SavingsplansSavingsPlan) CommitmentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commitmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SavingsplansSavingsPlan) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SavingsplansSavingsPlan) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SavingsplansSavingsPlan) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SavingsplansSavingsPlan) Currency() *string {
	var returns *string
	_jsii_.Get(
		j,
		"currency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SavingsplansSavingsPlan) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SavingsplansSavingsPlan) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SavingsplansSavingsPlan) Ec2InstanceFamily() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ec2InstanceFamily",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SavingsplansSavingsPlan) End() *string {
	var returns *string
	_jsii_.Get(
		j,
		"end",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SavingsplansSavingsPlan) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SavingsplansSavingsPlan) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SavingsplansSavingsPlan) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SavingsplansSavingsPlan) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SavingsplansSavingsPlan) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SavingsplansSavingsPlan) OfferingId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"offeringId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SavingsplansSavingsPlan) PaymentOption() *string {
	var returns *string
	_jsii_.Get(
		j,
		"paymentOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SavingsplansSavingsPlan) ProductTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"productTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SavingsplansSavingsPlan) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SavingsplansSavingsPlan) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SavingsplansSavingsPlan) PurchaseTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"purchaseTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SavingsplansSavingsPlan) PurchaseTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"purchaseTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SavingsplansSavingsPlan) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SavingsplansSavingsPlan) RecurringPaymentAmount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recurringPaymentAmount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SavingsplansSavingsPlan) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SavingsplansSavingsPlan) ReturnableUntil() *string {
	var returns *string
	_jsii_.Get(
		j,
		"returnableUntil",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SavingsplansSavingsPlan) SavingsPlanArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"savingsPlanArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SavingsplansSavingsPlan) SavingsPlanId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"savingsPlanId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SavingsplansSavingsPlan) SavingsPlanOfferingId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"savingsPlanOfferingId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SavingsplansSavingsPlan) SavingsPlanOfferingIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"savingsPlanOfferingIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SavingsplansSavingsPlan) SavingsPlanType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"savingsPlanType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SavingsplansSavingsPlan) Start() *string {
	var returns *string
	_jsii_.Get(
		j,
		"start",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SavingsplansSavingsPlan) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SavingsplansSavingsPlan) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SavingsplansSavingsPlan) TagsAll() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SavingsplansSavingsPlan) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SavingsplansSavingsPlan) TermDurationInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"termDurationInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SavingsplansSavingsPlan) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SavingsplansSavingsPlan) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SavingsplansSavingsPlan) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SavingsplansSavingsPlan) Timeouts() SavingsplansSavingsPlanTimeoutsOutputReference {
	var returns SavingsplansSavingsPlanTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SavingsplansSavingsPlan) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SavingsplansSavingsPlan) UpfrontPaymentAmount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"upfrontPaymentAmount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SavingsplansSavingsPlan) UpfrontPaymentAmountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"upfrontPaymentAmountInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.55.0/docs/resources/savingsplans_savings_plan aws_savingsplans_savings_plan} Resource.
func NewSavingsplansSavingsPlan(scope constructs.Construct, id *string, config *SavingsplansSavingsPlanConfig) SavingsplansSavingsPlan {
	_init_.Initialize()

	if err := validateNewSavingsplansSavingsPlanParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_SavingsplansSavingsPlan{}

	_jsii_.Create(
		"@cdktn/provider-aws.savingsplansSavingsPlan.SavingsplansSavingsPlan",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.55.0/docs/resources/savingsplans_savings_plan aws_savingsplans_savings_plan} Resource.
func NewSavingsplansSavingsPlan_Override(s SavingsplansSavingsPlan, scope constructs.Construct, id *string, config *SavingsplansSavingsPlanConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.savingsplansSavingsPlan.SavingsplansSavingsPlan",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SavingsplansSavingsPlan)SetCommitment(val *string) {
	if err := j.validateSetCommitmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"commitment",
		val,
	)
}

func (j *jsiiProxy_SavingsplansSavingsPlan)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_SavingsplansSavingsPlan)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SavingsplansSavingsPlan)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SavingsplansSavingsPlan)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_SavingsplansSavingsPlan)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SavingsplansSavingsPlan)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SavingsplansSavingsPlan)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_SavingsplansSavingsPlan)SetPurchaseTime(val *string) {
	if err := j.validateSetPurchaseTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"purchaseTime",
		val,
	)
}

func (j *jsiiProxy_SavingsplansSavingsPlan)SetSavingsPlanOfferingId(val *string) {
	if err := j.validateSetSavingsPlanOfferingIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"savingsPlanOfferingId",
		val,
	)
}

func (j *jsiiProxy_SavingsplansSavingsPlan)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_SavingsplansSavingsPlan)SetUpfrontPaymentAmount(val *string) {
	if err := j.validateSetUpfrontPaymentAmountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"upfrontPaymentAmount",
		val,
	)
}

// Generates CDKTN code for importing a SavingsplansSavingsPlan resource upon running "cdktn plan <stack-name>".
func SavingsplansSavingsPlan_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateSavingsplansSavingsPlan_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-aws.savingsplansSavingsPlan.SavingsplansSavingsPlan",
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
func SavingsplansSavingsPlan_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSavingsplansSavingsPlan_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-aws.savingsplansSavingsPlan.SavingsplansSavingsPlan",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SavingsplansSavingsPlan_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSavingsplansSavingsPlan_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-aws.savingsplansSavingsPlan.SavingsplansSavingsPlan",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SavingsplansSavingsPlan_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSavingsplansSavingsPlan_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-aws.savingsplansSavingsPlan.SavingsplansSavingsPlan",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SavingsplansSavingsPlan_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-aws.savingsplansSavingsPlan.SavingsplansSavingsPlan",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_SavingsplansSavingsPlan) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_SavingsplansSavingsPlan) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_SavingsplansSavingsPlan) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SavingsplansSavingsPlan) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SavingsplansSavingsPlan) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SavingsplansSavingsPlan) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SavingsplansSavingsPlan) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SavingsplansSavingsPlan) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SavingsplansSavingsPlan) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SavingsplansSavingsPlan) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SavingsplansSavingsPlan) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SavingsplansSavingsPlan) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SavingsplansSavingsPlan) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_SavingsplansSavingsPlan) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SavingsplansSavingsPlan) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SavingsplansSavingsPlan) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_SavingsplansSavingsPlan) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SavingsplansSavingsPlan) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SavingsplansSavingsPlan) PutTimeouts(value *SavingsplansSavingsPlanTimeouts) {
	if err := s.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SavingsplansSavingsPlan) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SavingsplansSavingsPlan) ResetPurchaseTime() {
	_jsii_.InvokeVoid(
		s,
		"resetPurchaseTime",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SavingsplansSavingsPlan) ResetTags() {
	_jsii_.InvokeVoid(
		s,
		"resetTags",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SavingsplansSavingsPlan) ResetTimeouts() {
	_jsii_.InvokeVoid(
		s,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SavingsplansSavingsPlan) ResetUpfrontPaymentAmount() {
	_jsii_.InvokeVoid(
		s,
		"resetUpfrontPaymentAmount",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SavingsplansSavingsPlan) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SavingsplansSavingsPlan) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SavingsplansSavingsPlan) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SavingsplansSavingsPlan) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SavingsplansSavingsPlan) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SavingsplansSavingsPlan) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SavingsplansSavingsPlan) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		s,
		"with",
		args,
		&returns,
	)

	return returns
}

