// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package rdsclustersnapshotcopy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/rdsclustersnapshotcopy/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/rds_cluster_snapshot_copy aws_rds_cluster_snapshot_copy}.
type RdsClusterSnapshotCopy interface {
	cdktn.TerraformResource
	AllocatedStorage() *float64
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	CopyTags() interface{}
	SetCopyTags(val interface{})
	CopyTagsInput() interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	DbClusterSnapshotArn() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DestinationRegion() *string
	SetDestinationRegion(val *string)
	DestinationRegionInput() *string
	Engine() *string
	EngineVersion() *string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	KmsKeyId() *string
	SetKmsKeyId(val *string)
	KmsKeyIdInput() *string
	LicenseModel() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	PresignedUrl() *string
	SetPresignedUrl(val *string)
	PresignedUrlInput() *string
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
	SharedAccounts() *[]*string
	SetSharedAccounts(val *[]*string)
	SharedAccountsInput() *[]*string
	SnapshotType() *string
	SourceDbClusterSnapshotIdentifier() *string
	SetSourceDbClusterSnapshotIdentifier(val *string)
	SourceDbClusterSnapshotIdentifierInput() *string
	StorageEncrypted() cdktn.IResolvable
	StorageType() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsAll() cdktn.StringMap
	TagsInput() *map[string]*string
	TargetDbClusterSnapshotIdentifier() *string
	SetTargetDbClusterSnapshotIdentifier(val *string)
	TargetDbClusterSnapshotIdentifierInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() RdsClusterSnapshotCopyTimeoutsOutputReference
	TimeoutsInput() interface{}
	VpcId() *string
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
	// Wraps a write-only attribute's already-mapped value so that `ProviderFeature.WRITE_ONLY_ATTRIBUTES` usage is registered at *resolve* time instead of at mutation time (setter/constructor). Called by generated bindings from `synthesizeAttributes()` and `synthesizeHclAttributes()`, e.g. `secret_key_wo: this.markWriteOnlyAttribute(cdktn.stringToTerraform(this._secretKeyWo))`; not intended to be called directly.
	//
	// `undefined` passes through completely unchanged, so the existing
	// undefined-filtering that omits unset attributes from synthesized
	// output (see `resolve()` in `tokens/private/resolve.ts`, and the
	// `value.value !== undefined` filter in generated
	// `synthesizeHclAttributes()`) keeps working untouched. `null` is also
	// passed through unchanged: it already renders as an explicit
	// null-out and must not arm the validation either.
	//
	// Any other value - including one that will itself resolve to nothing
	// (e.g. a `Lazy`/`IResolvable` producer with no value to contribute) -
	// is wrapped in a token whose `resolve()` defers to the real resolver
	// first and registers usage only if what comes back is not
	// `null`/`undefined`; the resolved value is then returned unchanged,
	// so what actually renders is untouched by this wrapper. A producer
	// that resolves to `undefined` therefore neither registers usage nor
	// leaves anything behind in the synthesized attribute - the omission
	// behaves exactly as if the attribute had never been set.
	//
	// Registration goes through `_registerResolveDiscoveredProviderFeatureUsage`
	// rather than `registerProviderFeatureUsage`: usage here is only known at
	// resolve time, and a given element can be resolved across many
	// synthesis passes over its lifetime (repeated `app.synth()` calls,
	// tests reusing a construct tree), so it must represent only the CURRENT
	// pass rather than accumulate forever. Every validation-enabled entry
	// point (`App.synth`; `Testing.synth`/`synthHcl` with validations;
	// `StackSynthesizer.synthesize`) runs a prepare step that deactivates any
	// stale registration and then resolves every element's `toTerraform()`
	// before that same entry point's validations run - see
	// `TerraformStack._runPreparingResolve` - so whatever this closure
	// (re-)registers during that prepare step is always visible to the
	// validation that reads it afterwards, and nothing left over from an
	// earlier pass leaks into the current one.
	// Experimental.
	MarkWriteOnlyAttribute(value interface{}) interface{}
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using its instance function.
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
	PutTimeouts(value *RdsClusterSnapshotCopyTimeouts)
	// Registers a synth-time validation that the project's declared targetVersions admit the given provider-protocol feature family.
	//
	// Called by generated provider bindings when a versioned feature is
	// structurally in use - the element's existence in the construct tree
	// already implies the feature is used, e.g. constructing a
	// `TerraformEphemeralResource` at all - so, unlike
	// `_registerResolveDiscoveredProviderFeatureUsage`, this registration is
	// never deactivated by `_resetResolveDiscoveredProviderFeatureUsage`. Not
	// intended to be called directly by user code. Lives on `TerraformElement`
	// (rather than `TerraformResource`) so it covers any element subclass
	// that needs it.
	// Experimental.
	RegisterProviderFeatureUsage(feature cdktn.ProviderFeature)
	ResetCopyTags()
	ResetDestinationRegion()
	ResetKmsKeyId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPresignedUrl()
	ResetRegion()
	ResetSharedAccounts()
	ResetTags()
	ResetTimeouts()
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

// The jsii proxy struct for RdsClusterSnapshotCopy
type jsiiProxy_RdsClusterSnapshotCopy struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_RdsClusterSnapshotCopy) AllocatedStorage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"allocatedStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsClusterSnapshotCopy) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsClusterSnapshotCopy) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsClusterSnapshotCopy) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsClusterSnapshotCopy) CopyTags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsClusterSnapshotCopy) CopyTagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsClusterSnapshotCopy) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsClusterSnapshotCopy) DbClusterSnapshotArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbClusterSnapshotArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsClusterSnapshotCopy) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsClusterSnapshotCopy) DestinationRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsClusterSnapshotCopy) DestinationRegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationRegionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsClusterSnapshotCopy) Engine() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsClusterSnapshotCopy) EngineVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsClusterSnapshotCopy) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsClusterSnapshotCopy) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsClusterSnapshotCopy) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsClusterSnapshotCopy) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsClusterSnapshotCopy) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsClusterSnapshotCopy) KmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsClusterSnapshotCopy) LicenseModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsClusterSnapshotCopy) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsClusterSnapshotCopy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsClusterSnapshotCopy) PresignedUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"presignedUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsClusterSnapshotCopy) PresignedUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"presignedUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsClusterSnapshotCopy) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsClusterSnapshotCopy) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsClusterSnapshotCopy) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsClusterSnapshotCopy) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsClusterSnapshotCopy) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsClusterSnapshotCopy) SharedAccounts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sharedAccounts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsClusterSnapshotCopy) SharedAccountsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sharedAccountsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsClusterSnapshotCopy) SnapshotType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsClusterSnapshotCopy) SourceDbClusterSnapshotIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceDbClusterSnapshotIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsClusterSnapshotCopy) SourceDbClusterSnapshotIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceDbClusterSnapshotIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsClusterSnapshotCopy) StorageEncrypted() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"storageEncrypted",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsClusterSnapshotCopy) StorageType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsClusterSnapshotCopy) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsClusterSnapshotCopy) TagsAll() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsClusterSnapshotCopy) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsClusterSnapshotCopy) TargetDbClusterSnapshotIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetDbClusterSnapshotIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsClusterSnapshotCopy) TargetDbClusterSnapshotIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetDbClusterSnapshotIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsClusterSnapshotCopy) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsClusterSnapshotCopy) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsClusterSnapshotCopy) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsClusterSnapshotCopy) Timeouts() RdsClusterSnapshotCopyTimeoutsOutputReference {
	var returns RdsClusterSnapshotCopyTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsClusterSnapshotCopy) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsClusterSnapshotCopy) VpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcId",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/rds_cluster_snapshot_copy aws_rds_cluster_snapshot_copy} Resource.
func NewRdsClusterSnapshotCopy(scope constructs.Construct, id *string, config *RdsClusterSnapshotCopyConfig) RdsClusterSnapshotCopy {
	_init_.Initialize()

	if err := validateNewRdsClusterSnapshotCopyParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_RdsClusterSnapshotCopy{}

	_jsii_.Create(
		"@cdktn/provider-aws.rdsClusterSnapshotCopy.RdsClusterSnapshotCopy",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/rds_cluster_snapshot_copy aws_rds_cluster_snapshot_copy} Resource.
func NewRdsClusterSnapshotCopy_Override(r RdsClusterSnapshotCopy, scope constructs.Construct, id *string, config *RdsClusterSnapshotCopyConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.rdsClusterSnapshotCopy.RdsClusterSnapshotCopy",
		[]interface{}{scope, id, config},
		r,
	)
}

func (j *jsiiProxy_RdsClusterSnapshotCopy)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_RdsClusterSnapshotCopy)SetCopyTags(val interface{}) {
	if err := j.validateSetCopyTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"copyTags",
		val,
	)
}

func (j *jsiiProxy_RdsClusterSnapshotCopy)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_RdsClusterSnapshotCopy)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_RdsClusterSnapshotCopy)SetDestinationRegion(val *string) {
	if err := j.validateSetDestinationRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationRegion",
		val,
	)
}

func (j *jsiiProxy_RdsClusterSnapshotCopy)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_RdsClusterSnapshotCopy)SetKmsKeyId(val *string) {
	if err := j.validateSetKmsKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

func (j *jsiiProxy_RdsClusterSnapshotCopy)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_RdsClusterSnapshotCopy)SetPresignedUrl(val *string) {
	if err := j.validateSetPresignedUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"presignedUrl",
		val,
	)
}

func (j *jsiiProxy_RdsClusterSnapshotCopy)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_RdsClusterSnapshotCopy)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_RdsClusterSnapshotCopy)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_RdsClusterSnapshotCopy)SetSharedAccounts(val *[]*string) {
	if err := j.validateSetSharedAccountsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sharedAccounts",
		val,
	)
}

func (j *jsiiProxy_RdsClusterSnapshotCopy)SetSourceDbClusterSnapshotIdentifier(val *string) {
	if err := j.validateSetSourceDbClusterSnapshotIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceDbClusterSnapshotIdentifier",
		val,
	)
}

func (j *jsiiProxy_RdsClusterSnapshotCopy)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_RdsClusterSnapshotCopy)SetTargetDbClusterSnapshotIdentifier(val *string) {
	if err := j.validateSetTargetDbClusterSnapshotIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetDbClusterSnapshotIdentifier",
		val,
	)
}

// Generates CDKTN code for importing a RdsClusterSnapshotCopy resource upon running "cdktn plan <stack-name>".
func RdsClusterSnapshotCopy_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateRdsClusterSnapshotCopy_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-aws.rdsClusterSnapshotCopy.RdsClusterSnapshotCopy",
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
func RdsClusterSnapshotCopy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRdsClusterSnapshotCopy_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-aws.rdsClusterSnapshotCopy.RdsClusterSnapshotCopy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func RdsClusterSnapshotCopy_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRdsClusterSnapshotCopy_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-aws.rdsClusterSnapshotCopy.RdsClusterSnapshotCopy",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func RdsClusterSnapshotCopy_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRdsClusterSnapshotCopy_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-aws.rdsClusterSnapshotCopy.RdsClusterSnapshotCopy",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func RdsClusterSnapshotCopy_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-aws.rdsClusterSnapshotCopy.RdsClusterSnapshotCopy",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (r *jsiiProxy_RdsClusterSnapshotCopy) AddMoveTarget(moveTarget *string) {
	if err := r.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (r *jsiiProxy_RdsClusterSnapshotCopy) AddOverride(path *string, value interface{}) {
	if err := r.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (r *jsiiProxy_RdsClusterSnapshotCopy) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := r.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsClusterSnapshotCopy) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := r.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsClusterSnapshotCopy) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := r.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		r,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsClusterSnapshotCopy) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := r.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsClusterSnapshotCopy) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := r.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsClusterSnapshotCopy) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := r.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		r,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsClusterSnapshotCopy) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := r.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		r,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsClusterSnapshotCopy) GetStringAttribute(terraformAttribute *string) *string {
	if err := r.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsClusterSnapshotCopy) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := r.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		r,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsClusterSnapshotCopy) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsClusterSnapshotCopy) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := r.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (r *jsiiProxy_RdsClusterSnapshotCopy) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := r.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsClusterSnapshotCopy) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := r.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		r,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsClusterSnapshotCopy) MoveFromId(id *string) {
	if err := r.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveFromId",
		[]interface{}{id},
	)
}

func (r *jsiiProxy_RdsClusterSnapshotCopy) MoveTo(moveTarget *string, index interface{}) {
	if err := r.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (r *jsiiProxy_RdsClusterSnapshotCopy) MoveToId(id *string) {
	if err := r.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveToId",
		[]interface{}{id},
	)
}

func (r *jsiiProxy_RdsClusterSnapshotCopy) OverrideLogicalId(newLogicalId *string) {
	if err := r.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (r *jsiiProxy_RdsClusterSnapshotCopy) PutTimeouts(value *RdsClusterSnapshotCopyTimeouts) {
	if err := r.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RdsClusterSnapshotCopy) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := r.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (r *jsiiProxy_RdsClusterSnapshotCopy) ResetCopyTags() {
	_jsii_.InvokeVoid(
		r,
		"resetCopyTags",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsClusterSnapshotCopy) ResetDestinationRegion() {
	_jsii_.InvokeVoid(
		r,
		"resetDestinationRegion",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsClusterSnapshotCopy) ResetKmsKeyId() {
	_jsii_.InvokeVoid(
		r,
		"resetKmsKeyId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsClusterSnapshotCopy) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		r,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsClusterSnapshotCopy) ResetPresignedUrl() {
	_jsii_.InvokeVoid(
		r,
		"resetPresignedUrl",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsClusterSnapshotCopy) ResetRegion() {
	_jsii_.InvokeVoid(
		r,
		"resetRegion",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsClusterSnapshotCopy) ResetSharedAccounts() {
	_jsii_.InvokeVoid(
		r,
		"resetSharedAccounts",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsClusterSnapshotCopy) ResetTags() {
	_jsii_.InvokeVoid(
		r,
		"resetTags",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsClusterSnapshotCopy) ResetTimeouts() {
	_jsii_.InvokeVoid(
		r,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsClusterSnapshotCopy) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsClusterSnapshotCopy) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsClusterSnapshotCopy) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsClusterSnapshotCopy) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsClusterSnapshotCopy) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsClusterSnapshotCopy) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsClusterSnapshotCopy) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		r,
		"with",
		args,
		&returns,
	)

	return returns
}

