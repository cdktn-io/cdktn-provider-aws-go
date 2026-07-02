// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package arczonalshiftzonalautoshiftconfiguration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v24/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v24/arczonalshiftzonalautoshiftconfiguration/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.53.0/docs/resources/arczonalshift_zonal_autoshift_configuration aws_arczonalshift_zonal_autoshift_configuration}.
type ArczonalshiftZonalAutoshiftConfiguration interface {
	cdktn.TerraformResource
	AllowedWindows() *[]*string
	SetAllowedWindows(val *[]*string)
	AllowedWindowsInput() *[]*string
	BlockedDates() *[]*string
	SetBlockedDates(val *[]*string)
	BlockedDatesInput() *[]*string
	BlockedWindows() *[]*string
	SetBlockedWindows(val *[]*string)
	BlockedWindowsInput() *[]*string
	BlockingAlarms() ArczonalshiftZonalAutoshiftConfigurationBlockingAlarmsList
	BlockingAlarmsInput() interface{}
	// Experimental.
	CdktfStack() cdktn.TerraformStack
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
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
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
	OutcomeAlarms() ArczonalshiftZonalAutoshiftConfigurationOutcomeAlarmsList
	OutcomeAlarmsInput() interface{}
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	ResourceArn() *string
	SetResourceArn(val *string)
	ResourceArnInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	ZonalAutoshiftStatus() *string
	SetZonalAutoshiftStatus(val *string)
	ZonalAutoshiftStatusInput() *string
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
	PutBlockingAlarms(value interface{})
	PutOutcomeAlarms(value interface{})
	ResetAllowedWindows()
	ResetBlockedDates()
	ResetBlockedWindows()
	ResetBlockingAlarms()
	ResetOutcomeAlarms()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRegion()
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

// The jsii proxy struct for ArczonalshiftZonalAutoshiftConfiguration
type jsiiProxy_ArczonalshiftZonalAutoshiftConfiguration struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_ArczonalshiftZonalAutoshiftConfiguration) AllowedWindows() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedWindows",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArczonalshiftZonalAutoshiftConfiguration) AllowedWindowsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedWindowsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArczonalshiftZonalAutoshiftConfiguration) BlockedDates() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"blockedDates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArczonalshiftZonalAutoshiftConfiguration) BlockedDatesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"blockedDatesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArczonalshiftZonalAutoshiftConfiguration) BlockedWindows() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"blockedWindows",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArczonalshiftZonalAutoshiftConfiguration) BlockedWindowsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"blockedWindowsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArczonalshiftZonalAutoshiftConfiguration) BlockingAlarms() ArczonalshiftZonalAutoshiftConfigurationBlockingAlarmsList {
	var returns ArczonalshiftZonalAutoshiftConfigurationBlockingAlarmsList
	_jsii_.Get(
		j,
		"blockingAlarms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArczonalshiftZonalAutoshiftConfiguration) BlockingAlarmsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"blockingAlarmsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArczonalshiftZonalAutoshiftConfiguration) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArczonalshiftZonalAutoshiftConfiguration) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArczonalshiftZonalAutoshiftConfiguration) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArczonalshiftZonalAutoshiftConfiguration) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArczonalshiftZonalAutoshiftConfiguration) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArczonalshiftZonalAutoshiftConfiguration) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArczonalshiftZonalAutoshiftConfiguration) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArczonalshiftZonalAutoshiftConfiguration) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArczonalshiftZonalAutoshiftConfiguration) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArczonalshiftZonalAutoshiftConfiguration) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArczonalshiftZonalAutoshiftConfiguration) OutcomeAlarms() ArczonalshiftZonalAutoshiftConfigurationOutcomeAlarmsList {
	var returns ArczonalshiftZonalAutoshiftConfigurationOutcomeAlarmsList
	_jsii_.Get(
		j,
		"outcomeAlarms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArczonalshiftZonalAutoshiftConfiguration) OutcomeAlarmsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"outcomeAlarmsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArczonalshiftZonalAutoshiftConfiguration) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArczonalshiftZonalAutoshiftConfiguration) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArczonalshiftZonalAutoshiftConfiguration) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArczonalshiftZonalAutoshiftConfiguration) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArczonalshiftZonalAutoshiftConfiguration) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArczonalshiftZonalAutoshiftConfiguration) ResourceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArczonalshiftZonalAutoshiftConfiguration) ResourceArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArczonalshiftZonalAutoshiftConfiguration) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArczonalshiftZonalAutoshiftConfiguration) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArczonalshiftZonalAutoshiftConfiguration) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArczonalshiftZonalAutoshiftConfiguration) ZonalAutoshiftStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zonalAutoshiftStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArczonalshiftZonalAutoshiftConfiguration) ZonalAutoshiftStatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zonalAutoshiftStatusInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.53.0/docs/resources/arczonalshift_zonal_autoshift_configuration aws_arczonalshift_zonal_autoshift_configuration} Resource.
func NewArczonalshiftZonalAutoshiftConfiguration(scope constructs.Construct, id *string, config *ArczonalshiftZonalAutoshiftConfigurationConfig) ArczonalshiftZonalAutoshiftConfiguration {
	_init_.Initialize()

	if err := validateNewArczonalshiftZonalAutoshiftConfigurationParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ArczonalshiftZonalAutoshiftConfiguration{}

	_jsii_.Create(
		"@cdktn/provider-aws.arczonalshiftZonalAutoshiftConfiguration.ArczonalshiftZonalAutoshiftConfiguration",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.53.0/docs/resources/arczonalshift_zonal_autoshift_configuration aws_arczonalshift_zonal_autoshift_configuration} Resource.
func NewArczonalshiftZonalAutoshiftConfiguration_Override(a ArczonalshiftZonalAutoshiftConfiguration, scope constructs.Construct, id *string, config *ArczonalshiftZonalAutoshiftConfigurationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.arczonalshiftZonalAutoshiftConfiguration.ArczonalshiftZonalAutoshiftConfiguration",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_ArczonalshiftZonalAutoshiftConfiguration)SetAllowedWindows(val *[]*string) {
	if err := j.validateSetAllowedWindowsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedWindows",
		val,
	)
}

func (j *jsiiProxy_ArczonalshiftZonalAutoshiftConfiguration)SetBlockedDates(val *[]*string) {
	if err := j.validateSetBlockedDatesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"blockedDates",
		val,
	)
}

func (j *jsiiProxy_ArczonalshiftZonalAutoshiftConfiguration)SetBlockedWindows(val *[]*string) {
	if err := j.validateSetBlockedWindowsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"blockedWindows",
		val,
	)
}

func (j *jsiiProxy_ArczonalshiftZonalAutoshiftConfiguration)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ArczonalshiftZonalAutoshiftConfiguration)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ArczonalshiftZonalAutoshiftConfiguration)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ArczonalshiftZonalAutoshiftConfiguration)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ArczonalshiftZonalAutoshiftConfiguration)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ArczonalshiftZonalAutoshiftConfiguration)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ArczonalshiftZonalAutoshiftConfiguration)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ArczonalshiftZonalAutoshiftConfiguration)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_ArczonalshiftZonalAutoshiftConfiguration)SetResourceArn(val *string) {
	if err := j.validateSetResourceArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceArn",
		val,
	)
}

func (j *jsiiProxy_ArczonalshiftZonalAutoshiftConfiguration)SetZonalAutoshiftStatus(val *string) {
	if err := j.validateSetZonalAutoshiftStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zonalAutoshiftStatus",
		val,
	)
}

// Generates CDKTN code for importing a ArczonalshiftZonalAutoshiftConfiguration resource upon running "cdktn plan <stack-name>".
func ArczonalshiftZonalAutoshiftConfiguration_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateArczonalshiftZonalAutoshiftConfiguration_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-aws.arczonalshiftZonalAutoshiftConfiguration.ArczonalshiftZonalAutoshiftConfiguration",
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
func ArczonalshiftZonalAutoshiftConfiguration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateArczonalshiftZonalAutoshiftConfiguration_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-aws.arczonalshiftZonalAutoshiftConfiguration.ArczonalshiftZonalAutoshiftConfiguration",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ArczonalshiftZonalAutoshiftConfiguration_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateArczonalshiftZonalAutoshiftConfiguration_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-aws.arczonalshiftZonalAutoshiftConfiguration.ArczonalshiftZonalAutoshiftConfiguration",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ArczonalshiftZonalAutoshiftConfiguration_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateArczonalshiftZonalAutoshiftConfiguration_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-aws.arczonalshiftZonalAutoshiftConfiguration.ArczonalshiftZonalAutoshiftConfiguration",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ArczonalshiftZonalAutoshiftConfiguration_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-aws.arczonalshiftZonalAutoshiftConfiguration.ArczonalshiftZonalAutoshiftConfiguration",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_ArczonalshiftZonalAutoshiftConfiguration) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_ArczonalshiftZonalAutoshiftConfiguration) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_ArczonalshiftZonalAutoshiftConfiguration) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArczonalshiftZonalAutoshiftConfiguration) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArczonalshiftZonalAutoshiftConfiguration) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArczonalshiftZonalAutoshiftConfiguration) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArczonalshiftZonalAutoshiftConfiguration) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArczonalshiftZonalAutoshiftConfiguration) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArczonalshiftZonalAutoshiftConfiguration) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArczonalshiftZonalAutoshiftConfiguration) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArczonalshiftZonalAutoshiftConfiguration) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArczonalshiftZonalAutoshiftConfiguration) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArczonalshiftZonalAutoshiftConfiguration) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_ArczonalshiftZonalAutoshiftConfiguration) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArczonalshiftZonalAutoshiftConfiguration) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_ArczonalshiftZonalAutoshiftConfiguration) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_ArczonalshiftZonalAutoshiftConfiguration) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_ArczonalshiftZonalAutoshiftConfiguration) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_ArczonalshiftZonalAutoshiftConfiguration) PutBlockingAlarms(value interface{}) {
	if err := a.validatePutBlockingAlarmsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putBlockingAlarms",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ArczonalshiftZonalAutoshiftConfiguration) PutOutcomeAlarms(value interface{}) {
	if err := a.validatePutOutcomeAlarmsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putOutcomeAlarms",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ArczonalshiftZonalAutoshiftConfiguration) ResetAllowedWindows() {
	_jsii_.InvokeVoid(
		a,
		"resetAllowedWindows",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArczonalshiftZonalAutoshiftConfiguration) ResetBlockedDates() {
	_jsii_.InvokeVoid(
		a,
		"resetBlockedDates",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArczonalshiftZonalAutoshiftConfiguration) ResetBlockedWindows() {
	_jsii_.InvokeVoid(
		a,
		"resetBlockedWindows",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArczonalshiftZonalAutoshiftConfiguration) ResetBlockingAlarms() {
	_jsii_.InvokeVoid(
		a,
		"resetBlockingAlarms",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArczonalshiftZonalAutoshiftConfiguration) ResetOutcomeAlarms() {
	_jsii_.InvokeVoid(
		a,
		"resetOutcomeAlarms",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArczonalshiftZonalAutoshiftConfiguration) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArczonalshiftZonalAutoshiftConfiguration) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArczonalshiftZonalAutoshiftConfiguration) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArczonalshiftZonalAutoshiftConfiguration) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArczonalshiftZonalAutoshiftConfiguration) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArczonalshiftZonalAutoshiftConfiguration) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArczonalshiftZonalAutoshiftConfiguration) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArczonalshiftZonalAutoshiftConfiguration) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArczonalshiftZonalAutoshiftConfiguration) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		a,
		"with",
		args,
		&returns,
	)

	return returns
}

