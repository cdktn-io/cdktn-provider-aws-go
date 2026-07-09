// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pinpointsmsvoicev2pool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v24/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v24/pinpointsmsvoicev2pool/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.54.0/docs/resources/pinpointsmsvoicev2_pool aws_pinpointsmsvoicev2_pool}.
type Pinpointsmsvoicev2Pool interface {
	cdktn.TerraformResource
	Arn() *string
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
	DeletionProtectionEnabled() interface{}
	SetDeletionProtectionEnabled(val interface{})
	DeletionProtectionEnabledInput() interface{}
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
	Id() *string
	IsoCountryCode() *string
	SetIsoCountryCode(val *string)
	IsoCountryCodeInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	MessageType() *string
	SetMessageType(val *string)
	MessageTypeInput() *string
	// The tree node.
	Node() constructs.Node
	OptOutListName() *string
	SetOptOutListName(val *string)
	OptOutListNameInput() *string
	OriginationIdentities() *[]*string
	SetOriginationIdentities(val *[]*string)
	OriginationIdentitiesInput() *[]*string
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
	SelfManagedOptOutsEnabled() interface{}
	SetSelfManagedOptOutsEnabled(val interface{})
	SelfManagedOptOutsEnabledInput() interface{}
	SharedRoutesEnabled() interface{}
	SetSharedRoutesEnabled(val interface{})
	SharedRoutesEnabledInput() interface{}
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsAll() cdktn.StringMap
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() Pinpointsmsvoicev2PoolTimeoutsOutputReference
	TimeoutsInput() interface{}
	TwoWayChannelArn() *string
	SetTwoWayChannelArn(val *string)
	TwoWayChannelArnInput() *string
	TwoWayChannelRole() *string
	SetTwoWayChannelRole(val *string)
	TwoWayChannelRoleInput() *string
	TwoWayEnabled() interface{}
	SetTwoWayEnabled(val interface{})
	TwoWayEnabledInput() interface{}
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
	PutTimeouts(value *Pinpointsmsvoicev2PoolTimeouts)
	ResetDeletionProtectionEnabled()
	ResetIsoCountryCode()
	ResetOptOutListName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRegion()
	ResetSelfManagedOptOutsEnabled()
	ResetSharedRoutesEnabled()
	ResetTags()
	ResetTimeouts()
	ResetTwoWayChannelArn()
	ResetTwoWayChannelRole()
	ResetTwoWayEnabled()
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

// The jsii proxy struct for Pinpointsmsvoicev2Pool
type jsiiProxy_Pinpointsmsvoicev2Pool struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_Pinpointsmsvoicev2Pool) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pinpointsmsvoicev2Pool) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pinpointsmsvoicev2Pool) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pinpointsmsvoicev2Pool) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pinpointsmsvoicev2Pool) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pinpointsmsvoicev2Pool) DeletionProtectionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtectionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pinpointsmsvoicev2Pool) DeletionProtectionEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtectionEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pinpointsmsvoicev2Pool) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pinpointsmsvoicev2Pool) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pinpointsmsvoicev2Pool) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pinpointsmsvoicev2Pool) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pinpointsmsvoicev2Pool) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pinpointsmsvoicev2Pool) IsoCountryCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"isoCountryCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pinpointsmsvoicev2Pool) IsoCountryCodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"isoCountryCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pinpointsmsvoicev2Pool) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pinpointsmsvoicev2Pool) MessageType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"messageType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pinpointsmsvoicev2Pool) MessageTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"messageTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pinpointsmsvoicev2Pool) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pinpointsmsvoicev2Pool) OptOutListName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"optOutListName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pinpointsmsvoicev2Pool) OptOutListNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"optOutListNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pinpointsmsvoicev2Pool) OriginationIdentities() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"originationIdentities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pinpointsmsvoicev2Pool) OriginationIdentitiesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"originationIdentitiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pinpointsmsvoicev2Pool) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pinpointsmsvoicev2Pool) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pinpointsmsvoicev2Pool) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pinpointsmsvoicev2Pool) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pinpointsmsvoicev2Pool) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pinpointsmsvoicev2Pool) SelfManagedOptOutsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"selfManagedOptOutsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pinpointsmsvoicev2Pool) SelfManagedOptOutsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"selfManagedOptOutsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pinpointsmsvoicev2Pool) SharedRoutesEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sharedRoutesEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pinpointsmsvoicev2Pool) SharedRoutesEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sharedRoutesEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pinpointsmsvoicev2Pool) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pinpointsmsvoicev2Pool) TagsAll() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pinpointsmsvoicev2Pool) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pinpointsmsvoicev2Pool) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pinpointsmsvoicev2Pool) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pinpointsmsvoicev2Pool) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pinpointsmsvoicev2Pool) Timeouts() Pinpointsmsvoicev2PoolTimeoutsOutputReference {
	var returns Pinpointsmsvoicev2PoolTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pinpointsmsvoicev2Pool) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pinpointsmsvoicev2Pool) TwoWayChannelArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"twoWayChannelArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pinpointsmsvoicev2Pool) TwoWayChannelArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"twoWayChannelArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pinpointsmsvoicev2Pool) TwoWayChannelRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"twoWayChannelRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pinpointsmsvoicev2Pool) TwoWayChannelRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"twoWayChannelRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pinpointsmsvoicev2Pool) TwoWayEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"twoWayEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pinpointsmsvoicev2Pool) TwoWayEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"twoWayEnabledInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.54.0/docs/resources/pinpointsmsvoicev2_pool aws_pinpointsmsvoicev2_pool} Resource.
func NewPinpointsmsvoicev2Pool(scope constructs.Construct, id *string, config *Pinpointsmsvoicev2PoolConfig) Pinpointsmsvoicev2Pool {
	_init_.Initialize()

	if err := validateNewPinpointsmsvoicev2PoolParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Pinpointsmsvoicev2Pool{}

	_jsii_.Create(
		"@cdktn/provider-aws.pinpointsmsvoicev2Pool.Pinpointsmsvoicev2Pool",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.54.0/docs/resources/pinpointsmsvoicev2_pool aws_pinpointsmsvoicev2_pool} Resource.
func NewPinpointsmsvoicev2Pool_Override(p Pinpointsmsvoicev2Pool, scope constructs.Construct, id *string, config *Pinpointsmsvoicev2PoolConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.pinpointsmsvoicev2Pool.Pinpointsmsvoicev2Pool",
		[]interface{}{scope, id, config},
		p,
	)
}

func (j *jsiiProxy_Pinpointsmsvoicev2Pool)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Pinpointsmsvoicev2Pool)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Pinpointsmsvoicev2Pool)SetDeletionProtectionEnabled(val interface{}) {
	if err := j.validateSetDeletionProtectionEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deletionProtectionEnabled",
		val,
	)
}

func (j *jsiiProxy_Pinpointsmsvoicev2Pool)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Pinpointsmsvoicev2Pool)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Pinpointsmsvoicev2Pool)SetIsoCountryCode(val *string) {
	if err := j.validateSetIsoCountryCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isoCountryCode",
		val,
	)
}

func (j *jsiiProxy_Pinpointsmsvoicev2Pool)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Pinpointsmsvoicev2Pool)SetMessageType(val *string) {
	if err := j.validateSetMessageTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"messageType",
		val,
	)
}

func (j *jsiiProxy_Pinpointsmsvoicev2Pool)SetOptOutListName(val *string) {
	if err := j.validateSetOptOutListNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"optOutListName",
		val,
	)
}

func (j *jsiiProxy_Pinpointsmsvoicev2Pool)SetOriginationIdentities(val *[]*string) {
	if err := j.validateSetOriginationIdentitiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"originationIdentities",
		val,
	)
}

func (j *jsiiProxy_Pinpointsmsvoicev2Pool)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Pinpointsmsvoicev2Pool)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Pinpointsmsvoicev2Pool)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_Pinpointsmsvoicev2Pool)SetSelfManagedOptOutsEnabled(val interface{}) {
	if err := j.validateSetSelfManagedOptOutsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"selfManagedOptOutsEnabled",
		val,
	)
}

func (j *jsiiProxy_Pinpointsmsvoicev2Pool)SetSharedRoutesEnabled(val interface{}) {
	if err := j.validateSetSharedRoutesEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sharedRoutesEnabled",
		val,
	)
}

func (j *jsiiProxy_Pinpointsmsvoicev2Pool)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_Pinpointsmsvoicev2Pool)SetTwoWayChannelArn(val *string) {
	if err := j.validateSetTwoWayChannelArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"twoWayChannelArn",
		val,
	)
}

func (j *jsiiProxy_Pinpointsmsvoicev2Pool)SetTwoWayChannelRole(val *string) {
	if err := j.validateSetTwoWayChannelRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"twoWayChannelRole",
		val,
	)
}

func (j *jsiiProxy_Pinpointsmsvoicev2Pool)SetTwoWayEnabled(val interface{}) {
	if err := j.validateSetTwoWayEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"twoWayEnabled",
		val,
	)
}

// Generates CDKTN code for importing a Pinpointsmsvoicev2Pool resource upon running "cdktn plan <stack-name>".
func Pinpointsmsvoicev2Pool_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validatePinpointsmsvoicev2Pool_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-aws.pinpointsmsvoicev2Pool.Pinpointsmsvoicev2Pool",
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
func Pinpointsmsvoicev2Pool_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePinpointsmsvoicev2Pool_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-aws.pinpointsmsvoicev2Pool.Pinpointsmsvoicev2Pool",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Pinpointsmsvoicev2Pool_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePinpointsmsvoicev2Pool_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-aws.pinpointsmsvoicev2Pool.Pinpointsmsvoicev2Pool",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Pinpointsmsvoicev2Pool_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePinpointsmsvoicev2Pool_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-aws.pinpointsmsvoicev2Pool.Pinpointsmsvoicev2Pool",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Pinpointsmsvoicev2Pool_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-aws.pinpointsmsvoicev2Pool.Pinpointsmsvoicev2Pool",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (p *jsiiProxy_Pinpointsmsvoicev2Pool) AddMoveTarget(moveTarget *string) {
	if err := p.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (p *jsiiProxy_Pinpointsmsvoicev2Pool) AddOverride(path *string, value interface{}) {
	if err := p.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (p *jsiiProxy_Pinpointsmsvoicev2Pool) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Pinpointsmsvoicev2Pool) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Pinpointsmsvoicev2Pool) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Pinpointsmsvoicev2Pool) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Pinpointsmsvoicev2Pool) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Pinpointsmsvoicev2Pool) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Pinpointsmsvoicev2Pool) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Pinpointsmsvoicev2Pool) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Pinpointsmsvoicev2Pool) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Pinpointsmsvoicev2Pool) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Pinpointsmsvoicev2Pool) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := p.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (p *jsiiProxy_Pinpointsmsvoicev2Pool) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Pinpointsmsvoicev2Pool) MoveFromId(id *string) {
	if err := p.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveFromId",
		[]interface{}{id},
	)
}

func (p *jsiiProxy_Pinpointsmsvoicev2Pool) MoveTo(moveTarget *string, index interface{}) {
	if err := p.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (p *jsiiProxy_Pinpointsmsvoicev2Pool) MoveToId(id *string) {
	if err := p.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveToId",
		[]interface{}{id},
	)
}

func (p *jsiiProxy_Pinpointsmsvoicev2Pool) OverrideLogicalId(newLogicalId *string) {
	if err := p.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (p *jsiiProxy_Pinpointsmsvoicev2Pool) PutTimeouts(value *Pinpointsmsvoicev2PoolTimeouts) {
	if err := p.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_Pinpointsmsvoicev2Pool) ResetDeletionProtectionEnabled() {
	_jsii_.InvokeVoid(
		p,
		"resetDeletionProtectionEnabled",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Pinpointsmsvoicev2Pool) ResetIsoCountryCode() {
	_jsii_.InvokeVoid(
		p,
		"resetIsoCountryCode",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Pinpointsmsvoicev2Pool) ResetOptOutListName() {
	_jsii_.InvokeVoid(
		p,
		"resetOptOutListName",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Pinpointsmsvoicev2Pool) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		p,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Pinpointsmsvoicev2Pool) ResetRegion() {
	_jsii_.InvokeVoid(
		p,
		"resetRegion",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Pinpointsmsvoicev2Pool) ResetSelfManagedOptOutsEnabled() {
	_jsii_.InvokeVoid(
		p,
		"resetSelfManagedOptOutsEnabled",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Pinpointsmsvoicev2Pool) ResetSharedRoutesEnabled() {
	_jsii_.InvokeVoid(
		p,
		"resetSharedRoutesEnabled",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Pinpointsmsvoicev2Pool) ResetTags() {
	_jsii_.InvokeVoid(
		p,
		"resetTags",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Pinpointsmsvoicev2Pool) ResetTimeouts() {
	_jsii_.InvokeVoid(
		p,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Pinpointsmsvoicev2Pool) ResetTwoWayChannelArn() {
	_jsii_.InvokeVoid(
		p,
		"resetTwoWayChannelArn",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Pinpointsmsvoicev2Pool) ResetTwoWayChannelRole() {
	_jsii_.InvokeVoid(
		p,
		"resetTwoWayChannelRole",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Pinpointsmsvoicev2Pool) ResetTwoWayEnabled() {
	_jsii_.InvokeVoid(
		p,
		"resetTwoWayEnabled",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Pinpointsmsvoicev2Pool) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Pinpointsmsvoicev2Pool) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Pinpointsmsvoicev2Pool) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Pinpointsmsvoicev2Pool) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Pinpointsmsvoicev2Pool) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Pinpointsmsvoicev2Pool) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Pinpointsmsvoicev2Pool) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		p,
		"with",
		args,
		&returns,
	)

	return returns
}

