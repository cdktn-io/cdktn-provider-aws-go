// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package odbautonomousdatabase

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/odbautonomousdatabase/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.68.0/docs/resources/odb_autonomous_database aws_odb_autonomous_database}.
type OdbAutonomousDatabase interface {
	cdktn.TerraformResource
	ActualUsedDataStorageSizeInTbs() *float64
	AdminPassword() *string
	SetAdminPassword(val *string)
	AdminPasswordInput() *string
	AdminPasswordSource() OdbAutonomousDatabaseAdminPasswordSourceList
	AdminPasswordSourceInput() interface{}
	// Deprecated: Write-only: the provider never returns this value; reading it always yields null by protocol contract. The getter remains for compatibility and will be removed in a future prebuilt-provider major.
	AdminPasswordWo() *string
	// Deprecated: Write-only: the provider never returns this value; reading it always yields null by protocol contract. The getter remains for compatibility and will be removed in a future prebuilt-provider major.
	SetAdminPasswordWo(val *string)
	AdminPasswordWoInput() *string
	AdminPasswordWoVersion() *float64
	SetAdminPasswordWoVersion(val *float64)
	AdminPasswordWoVersionInput() *float64
	AllocatedStorageSizeInTbs() *float64
	AllowlistedIps() *[]*string
	SetAllowlistedIps(val *[]*string)
	AllowlistedIpsInput() *[]*string
	Arn() *string
	AutonomousMaintenanceScheduleType() *string
	SetAutonomousMaintenanceScheduleType(val *string)
	AutonomousMaintenanceScheduleTypeInput() *string
	AutoRefreshFrequencyInSeconds() *float64
	SetAutoRefreshFrequencyInSeconds(val *float64)
	AutoRefreshFrequencyInSecondsInput() *float64
	AutoRefreshPointLagInSeconds() *float64
	SetAutoRefreshPointLagInSeconds(val *float64)
	AutoRefreshPointLagInSecondsInput() *float64
	AvailabilityZone() *string
	AvailabilityZoneId() *string
	AvailableUpgradeVersions() *[]*string
	BackupRetentionPeriodInDays() *float64
	SetBackupRetentionPeriodInDays(val *float64)
	BackupRetentionPeriodInDaysInput() *float64
	ByolComputeCountLimit() *float64
	SetByolComputeCountLimit(val *float64)
	ByolComputeCountLimitInput() *float64
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	CharacterSet() *string
	SetCharacterSet(val *string)
	CharacterSetInput() *string
	ComputeCount() *float64
	SetComputeCount(val *float64)
	ComputeCountInput() *float64
	ComputeModel() *string
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
	CpuCoreCount() *float64
	SetCpuCoreCount(val *float64)
	CpuCoreCountInput() *float64
	CreatedAt() *string
	CustomerContactsToSendToOci() OdbAutonomousDatabaseCustomerContactsToSendToOciList
	CustomerContactsToSendToOciInput() interface{}
	DatabaseEdition() *string
	SetDatabaseEdition(val *string)
	DatabaseEditionInput() *string
	DatabaseType() *string
	DataStorageSizeInGbs() *float64
	SetDataStorageSizeInGbs(val *float64)
	DataStorageSizeInGbsInput() *float64
	DataStorageSizeInTbs() *float64
	SetDataStorageSizeInTbs(val *float64)
	DataStorageSizeInTbsInput() *float64
	DbName() *string
	SetDbName(val *string)
	DbNameInput() *string
	DbToolsDetails() OdbAutonomousDatabaseDbToolsDetailsList
	DbToolsDetailsInput() interface{}
	DbVersion() *string
	SetDbVersion(val *string)
	DbVersionInput() *string
	DbWorkload() *string
	SetDbWorkload(val *string)
	DbWorkloadInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	EncryptionKeyProvider() *string
	SetEncryptionKeyProvider(val *string)
	EncryptionKeyProviderInput() *string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	IsAutoScalingEnabled() interface{}
	SetIsAutoScalingEnabled(val interface{})
	IsAutoScalingEnabledInput() interface{}
	IsAutoScalingForStorageEnabled() interface{}
	SetIsAutoScalingForStorageEnabled(val interface{})
	IsAutoScalingForStorageEnabledInput() interface{}
	IsBackupRetentionLocked() interface{}
	SetIsBackupRetentionLocked(val interface{})
	IsBackupRetentionLockedInput() interface{}
	IsLocalDataGuardEnabled() interface{}
	SetIsLocalDataGuardEnabled(val interface{})
	IsLocalDataGuardEnabledInput() interface{}
	IsMtlsConnectionRequired() interface{}
	SetIsMtlsConnectionRequired(val interface{})
	IsMtlsConnectionRequiredInput() interface{}
	IsRefreshableClone() interface{}
	SetIsRefreshableClone(val interface{})
	IsRefreshableCloneInput() interface{}
	KmsKeyId() *string
	SetKmsKeyId(val *string)
	KmsKeyIdInput() *string
	LicenseModel() *string
	SetLicenseModel(val *string)
	LicenseModelInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	LocalAdgAutoFailoverMaxDataLossLimit() *float64
	SetLocalAdgAutoFailoverMaxDataLossLimit(val *float64)
	LocalAdgAutoFailoverMaxDataLossLimitInput() *float64
	LongTermBackupSchedule() OdbAutonomousDatabaseLongTermBackupScheduleList
	LongTermBackupScheduleInput() interface{}
	NcharacterSet() *string
	SetNcharacterSet(val *string)
	NcharacterSetInput() *string
	// The tree node.
	Node() constructs.Node
	Ocid() *string
	OciResourceAnchorName() *string
	OciUrl() *string
	OdbNetworkArn() *string
	OdbNetworkId() *string
	SetOdbNetworkId(val *string)
	OdbNetworkIdInput() *string
	OpenMode() *string
	SetOpenMode(val *string)
	OpenModeInput() *string
	PercentProgress() *float64
	PermissionLevel() *string
	SetPermissionLevel(val *string)
	PermissionLevelInput() *string
	PrivateEndpoint() *string
	PrivateEndpointIp() *string
	SetPrivateEndpointIp(val *string)
	PrivateEndpointIpInput() *string
	PrivateEndpointLabel() *string
	SetPrivateEndpointLabel(val *string)
	PrivateEndpointLabelInput() *string
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
	RefreshableMode() *string
	SetRefreshableMode(val *string)
	RefreshableModeInput() *string
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	ResourcePoolLeaderId() *string
	SetResourcePoolLeaderId(val *string)
	ResourcePoolLeaderIdInput() *string
	ResourcePoolSummary() OdbAutonomousDatabaseResourcePoolSummaryList
	ResourcePoolSummaryInput() interface{}
	ScheduledOperations() OdbAutonomousDatabaseScheduledOperationsList
	ScheduledOperationsInput() interface{}
	ServiceConsoleUrl() *string
	Source() *string
	SetSource(val *string)
	SourceConfiguration() OdbAutonomousDatabaseSourceConfigurationList
	SourceConfigurationInput() interface{}
	SourceId() *string
	SourceInput() *string
	SqlWebDeveloperUrl() *string
	StandbyAllowlistedIps() *[]*string
	SetStandbyAllowlistedIps(val *[]*string)
	StandbyAllowlistedIpsInput() *[]*string
	StandbyAllowlistedIpsSource() *string
	SetStandbyAllowlistedIpsSource(val *string)
	StandbyAllowlistedIpsSourceInput() *string
	Status() *string
	StatusReason() *string
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
	TimeOfAutoRefreshStart() *string
	SetTimeOfAutoRefreshStart(val *string)
	TimeOfAutoRefreshStartInput() *string
	Timeouts() OdbAutonomousDatabaseTimeoutsOutputReference
	TimeoutsInput() interface{}
	TransportableTablespace() OdbAutonomousDatabaseTransportableTablespaceList
	TransportableTablespaceInput() interface{}
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
	PutAdminPasswordSource(value interface{})
	PutCustomerContactsToSendToOci(value interface{})
	PutDbToolsDetails(value interface{})
	PutLongTermBackupSchedule(value interface{})
	PutResourcePoolSummary(value interface{})
	PutScheduledOperations(value interface{})
	PutSourceConfiguration(value interface{})
	PutTimeouts(value *OdbAutonomousDatabaseTimeouts)
	PutTransportableTablespace(value interface{})
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
	ResetAdminPassword()
	ResetAdminPasswordSource()
	ResetAdminPasswordWo()
	ResetAdminPasswordWoVersion()
	ResetAllowlistedIps()
	ResetAutonomousMaintenanceScheduleType()
	ResetAutoRefreshFrequencyInSeconds()
	ResetAutoRefreshPointLagInSeconds()
	ResetBackupRetentionPeriodInDays()
	ResetByolComputeCountLimit()
	ResetCharacterSet()
	ResetComputeCount()
	ResetCpuCoreCount()
	ResetCustomerContactsToSendToOci()
	ResetDatabaseEdition()
	ResetDataStorageSizeInGbs()
	ResetDataStorageSizeInTbs()
	ResetDbName()
	ResetDbToolsDetails()
	ResetDbVersion()
	ResetDbWorkload()
	ResetDisplayName()
	ResetEncryptionKeyProvider()
	ResetIsAutoScalingEnabled()
	ResetIsAutoScalingForStorageEnabled()
	ResetIsBackupRetentionLocked()
	ResetIsLocalDataGuardEnabled()
	ResetIsMtlsConnectionRequired()
	ResetIsRefreshableClone()
	ResetKmsKeyId()
	ResetLicenseModel()
	ResetLocalAdgAutoFailoverMaxDataLossLimit()
	ResetLongTermBackupSchedule()
	ResetNcharacterSet()
	ResetOdbNetworkId()
	ResetOpenMode()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPermissionLevel()
	ResetPrivateEndpointIp()
	ResetPrivateEndpointLabel()
	ResetRefreshableMode()
	ResetRegion()
	ResetResourcePoolLeaderId()
	ResetResourcePoolSummary()
	ResetScheduledOperations()
	ResetSource()
	ResetSourceConfiguration()
	ResetStandbyAllowlistedIps()
	ResetStandbyAllowlistedIpsSource()
	ResetTags()
	ResetTimeOfAutoRefreshStart()
	ResetTimeouts()
	ResetTransportableTablespace()
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

// The jsii proxy struct for OdbAutonomousDatabase
type jsiiProxy_OdbAutonomousDatabase struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_OdbAutonomousDatabase) ActualUsedDataStorageSizeInTbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"actualUsedDataStorageSizeInTbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) AdminPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) AdminPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) AdminPasswordSource() OdbAutonomousDatabaseAdminPasswordSourceList {
	var returns OdbAutonomousDatabaseAdminPasswordSourceList
	_jsii_.Get(
		j,
		"adminPasswordSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) AdminPasswordSourceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"adminPasswordSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) AdminPasswordWo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminPasswordWo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) AdminPasswordWoInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminPasswordWoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) AdminPasswordWoVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"adminPasswordWoVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) AdminPasswordWoVersionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"adminPasswordWoVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) AllocatedStorageSizeInTbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"allocatedStorageSizeInTbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) AllowlistedIps() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowlistedIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) AllowlistedIpsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowlistedIpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) AutonomousMaintenanceScheduleType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autonomousMaintenanceScheduleType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) AutonomousMaintenanceScheduleTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autonomousMaintenanceScheduleTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) AutoRefreshFrequencyInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autoRefreshFrequencyInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) AutoRefreshFrequencyInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autoRefreshFrequencyInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) AutoRefreshPointLagInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autoRefreshPointLagInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) AutoRefreshPointLagInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autoRefreshPointLagInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) AvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) AvailabilityZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) AvailableUpgradeVersions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availableUpgradeVersions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) BackupRetentionPeriodInDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backupRetentionPeriodInDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) BackupRetentionPeriodInDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backupRetentionPeriodInDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) ByolComputeCountLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"byolComputeCountLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) ByolComputeCountLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"byolComputeCountLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) CharacterSet() *string {
	var returns *string
	_jsii_.Get(
		j,
		"characterSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) CharacterSetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"characterSetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) ComputeCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"computeCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) ComputeCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"computeCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) ComputeModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) CpuCoreCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuCoreCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) CpuCoreCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuCoreCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) CustomerContactsToSendToOci() OdbAutonomousDatabaseCustomerContactsToSendToOciList {
	var returns OdbAutonomousDatabaseCustomerContactsToSendToOciList
	_jsii_.Get(
		j,
		"customerContactsToSendToOci",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) CustomerContactsToSendToOciInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customerContactsToSendToOciInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) DatabaseEdition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseEdition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) DatabaseEditionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseEditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) DatabaseType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) DataStorageSizeInGbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataStorageSizeInGbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) DataStorageSizeInGbsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataStorageSizeInGbsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) DataStorageSizeInTbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataStorageSizeInTbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) DataStorageSizeInTbsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataStorageSizeInTbsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) DbName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) DbNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) DbToolsDetails() OdbAutonomousDatabaseDbToolsDetailsList {
	var returns OdbAutonomousDatabaseDbToolsDetailsList
	_jsii_.Get(
		j,
		"dbToolsDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) DbToolsDetailsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dbToolsDetailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) DbVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) DbVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) DbWorkload() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbWorkload",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) DbWorkloadInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbWorkloadInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) EncryptionKeyProvider() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionKeyProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) EncryptionKeyProviderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionKeyProviderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) IsAutoScalingEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isAutoScalingEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) IsAutoScalingEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isAutoScalingEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) IsAutoScalingForStorageEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isAutoScalingForStorageEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) IsAutoScalingForStorageEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isAutoScalingForStorageEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) IsBackupRetentionLocked() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isBackupRetentionLocked",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) IsBackupRetentionLockedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isBackupRetentionLockedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) IsLocalDataGuardEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isLocalDataGuardEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) IsLocalDataGuardEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isLocalDataGuardEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) IsMtlsConnectionRequired() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isMtlsConnectionRequired",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) IsMtlsConnectionRequiredInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isMtlsConnectionRequiredInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) IsRefreshableClone() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isRefreshableClone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) IsRefreshableCloneInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isRefreshableCloneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) KmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) LicenseModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) LicenseModelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseModelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) LocalAdgAutoFailoverMaxDataLossLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"localAdgAutoFailoverMaxDataLossLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) LocalAdgAutoFailoverMaxDataLossLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"localAdgAutoFailoverMaxDataLossLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) LongTermBackupSchedule() OdbAutonomousDatabaseLongTermBackupScheduleList {
	var returns OdbAutonomousDatabaseLongTermBackupScheduleList
	_jsii_.Get(
		j,
		"longTermBackupSchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) LongTermBackupScheduleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"longTermBackupScheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) NcharacterSet() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ncharacterSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) NcharacterSetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ncharacterSetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) Ocid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ocid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) OciResourceAnchorName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ociResourceAnchorName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) OciUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ociUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) OdbNetworkArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"odbNetworkArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) OdbNetworkId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"odbNetworkId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) OdbNetworkIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"odbNetworkIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) OpenMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"openMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) OpenModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"openModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) PercentProgress() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"percentProgress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) PermissionLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"permissionLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) PermissionLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"permissionLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) PrivateEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) PrivateEndpointIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateEndpointIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) PrivateEndpointIpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateEndpointIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) PrivateEndpointLabel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateEndpointLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) PrivateEndpointLabelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateEndpointLabelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) RefreshableMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"refreshableMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) RefreshableModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"refreshableModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) ResourcePoolLeaderId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourcePoolLeaderId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) ResourcePoolLeaderIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourcePoolLeaderIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) ResourcePoolSummary() OdbAutonomousDatabaseResourcePoolSummaryList {
	var returns OdbAutonomousDatabaseResourcePoolSummaryList
	_jsii_.Get(
		j,
		"resourcePoolSummary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) ResourcePoolSummaryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourcePoolSummaryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) ScheduledOperations() OdbAutonomousDatabaseScheduledOperationsList {
	var returns OdbAutonomousDatabaseScheduledOperationsList
	_jsii_.Get(
		j,
		"scheduledOperations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) ScheduledOperationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scheduledOperationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) ServiceConsoleUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceConsoleUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) Source() *string {
	var returns *string
	_jsii_.Get(
		j,
		"source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) SourceConfiguration() OdbAutonomousDatabaseSourceConfigurationList {
	var returns OdbAutonomousDatabaseSourceConfigurationList
	_jsii_.Get(
		j,
		"sourceConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) SourceConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourceConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) SourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) SourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) SqlWebDeveloperUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqlWebDeveloperUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) StandbyAllowlistedIps() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"standbyAllowlistedIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) StandbyAllowlistedIpsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"standbyAllowlistedIpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) StandbyAllowlistedIpsSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"standbyAllowlistedIpsSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) StandbyAllowlistedIpsSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"standbyAllowlistedIpsSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) StatusReason() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusReason",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) TagsAll() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) TimeOfAutoRefreshStart() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeOfAutoRefreshStart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) TimeOfAutoRefreshStartInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeOfAutoRefreshStartInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) Timeouts() OdbAutonomousDatabaseTimeoutsOutputReference {
	var returns OdbAutonomousDatabaseTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) TransportableTablespace() OdbAutonomousDatabaseTransportableTablespaceList {
	var returns OdbAutonomousDatabaseTransportableTablespaceList
	_jsii_.Get(
		j,
		"transportableTablespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbAutonomousDatabase) TransportableTablespaceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"transportableTablespaceInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.68.0/docs/resources/odb_autonomous_database aws_odb_autonomous_database} Resource.
func NewOdbAutonomousDatabase(scope constructs.Construct, id *string, config *OdbAutonomousDatabaseConfig) OdbAutonomousDatabase {
	_init_.Initialize()

	if err := validateNewOdbAutonomousDatabaseParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_OdbAutonomousDatabase{}

	_jsii_.Create(
		"@cdktn/provider-aws.odbAutonomousDatabase.OdbAutonomousDatabase",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.68.0/docs/resources/odb_autonomous_database aws_odb_autonomous_database} Resource.
func NewOdbAutonomousDatabase_Override(o OdbAutonomousDatabase, scope constructs.Construct, id *string, config *OdbAutonomousDatabaseConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.odbAutonomousDatabase.OdbAutonomousDatabase",
		[]interface{}{scope, id, config},
		o,
	)
}

func (j *jsiiProxy_OdbAutonomousDatabase)SetAdminPassword(val *string) {
	if err := j.validateSetAdminPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adminPassword",
		val,
	)
}

func (j *jsiiProxy_OdbAutonomousDatabase)SetAdminPasswordWo(val *string) {
	if err := j.validateSetAdminPasswordWoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adminPasswordWo",
		val,
	)
}

func (j *jsiiProxy_OdbAutonomousDatabase)SetAdminPasswordWoVersion(val *float64) {
	if err := j.validateSetAdminPasswordWoVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adminPasswordWoVersion",
		val,
	)
}

func (j *jsiiProxy_OdbAutonomousDatabase)SetAllowlistedIps(val *[]*string) {
	if err := j.validateSetAllowlistedIpsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowlistedIps",
		val,
	)
}

func (j *jsiiProxy_OdbAutonomousDatabase)SetAutonomousMaintenanceScheduleType(val *string) {
	if err := j.validateSetAutonomousMaintenanceScheduleTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autonomousMaintenanceScheduleType",
		val,
	)
}

func (j *jsiiProxy_OdbAutonomousDatabase)SetAutoRefreshFrequencyInSeconds(val *float64) {
	if err := j.validateSetAutoRefreshFrequencyInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoRefreshFrequencyInSeconds",
		val,
	)
}

func (j *jsiiProxy_OdbAutonomousDatabase)SetAutoRefreshPointLagInSeconds(val *float64) {
	if err := j.validateSetAutoRefreshPointLagInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoRefreshPointLagInSeconds",
		val,
	)
}

func (j *jsiiProxy_OdbAutonomousDatabase)SetBackupRetentionPeriodInDays(val *float64) {
	if err := j.validateSetBackupRetentionPeriodInDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backupRetentionPeriodInDays",
		val,
	)
}

func (j *jsiiProxy_OdbAutonomousDatabase)SetByolComputeCountLimit(val *float64) {
	if err := j.validateSetByolComputeCountLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"byolComputeCountLimit",
		val,
	)
}

func (j *jsiiProxy_OdbAutonomousDatabase)SetCharacterSet(val *string) {
	if err := j.validateSetCharacterSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"characterSet",
		val,
	)
}

func (j *jsiiProxy_OdbAutonomousDatabase)SetComputeCount(val *float64) {
	if err := j.validateSetComputeCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"computeCount",
		val,
	)
}

func (j *jsiiProxy_OdbAutonomousDatabase)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_OdbAutonomousDatabase)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_OdbAutonomousDatabase)SetCpuCoreCount(val *float64) {
	if err := j.validateSetCpuCoreCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpuCoreCount",
		val,
	)
}

func (j *jsiiProxy_OdbAutonomousDatabase)SetDatabaseEdition(val *string) {
	if err := j.validateSetDatabaseEditionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"databaseEdition",
		val,
	)
}

func (j *jsiiProxy_OdbAutonomousDatabase)SetDataStorageSizeInGbs(val *float64) {
	if err := j.validateSetDataStorageSizeInGbsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataStorageSizeInGbs",
		val,
	)
}

func (j *jsiiProxy_OdbAutonomousDatabase)SetDataStorageSizeInTbs(val *float64) {
	if err := j.validateSetDataStorageSizeInTbsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataStorageSizeInTbs",
		val,
	)
}

func (j *jsiiProxy_OdbAutonomousDatabase)SetDbName(val *string) {
	if err := j.validateSetDbNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbName",
		val,
	)
}

func (j *jsiiProxy_OdbAutonomousDatabase)SetDbVersion(val *string) {
	if err := j.validateSetDbVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbVersion",
		val,
	)
}

func (j *jsiiProxy_OdbAutonomousDatabase)SetDbWorkload(val *string) {
	if err := j.validateSetDbWorkloadParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbWorkload",
		val,
	)
}

func (j *jsiiProxy_OdbAutonomousDatabase)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_OdbAutonomousDatabase)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_OdbAutonomousDatabase)SetEncryptionKeyProvider(val *string) {
	if err := j.validateSetEncryptionKeyProviderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encryptionKeyProvider",
		val,
	)
}

func (j *jsiiProxy_OdbAutonomousDatabase)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_OdbAutonomousDatabase)SetIsAutoScalingEnabled(val interface{}) {
	if err := j.validateSetIsAutoScalingEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isAutoScalingEnabled",
		val,
	)
}

func (j *jsiiProxy_OdbAutonomousDatabase)SetIsAutoScalingForStorageEnabled(val interface{}) {
	if err := j.validateSetIsAutoScalingForStorageEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isAutoScalingForStorageEnabled",
		val,
	)
}

func (j *jsiiProxy_OdbAutonomousDatabase)SetIsBackupRetentionLocked(val interface{}) {
	if err := j.validateSetIsBackupRetentionLockedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isBackupRetentionLocked",
		val,
	)
}

func (j *jsiiProxy_OdbAutonomousDatabase)SetIsLocalDataGuardEnabled(val interface{}) {
	if err := j.validateSetIsLocalDataGuardEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isLocalDataGuardEnabled",
		val,
	)
}

func (j *jsiiProxy_OdbAutonomousDatabase)SetIsMtlsConnectionRequired(val interface{}) {
	if err := j.validateSetIsMtlsConnectionRequiredParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isMtlsConnectionRequired",
		val,
	)
}

func (j *jsiiProxy_OdbAutonomousDatabase)SetIsRefreshableClone(val interface{}) {
	if err := j.validateSetIsRefreshableCloneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isRefreshableClone",
		val,
	)
}

func (j *jsiiProxy_OdbAutonomousDatabase)SetKmsKeyId(val *string) {
	if err := j.validateSetKmsKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

func (j *jsiiProxy_OdbAutonomousDatabase)SetLicenseModel(val *string) {
	if err := j.validateSetLicenseModelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"licenseModel",
		val,
	)
}

func (j *jsiiProxy_OdbAutonomousDatabase)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_OdbAutonomousDatabase)SetLocalAdgAutoFailoverMaxDataLossLimit(val *float64) {
	if err := j.validateSetLocalAdgAutoFailoverMaxDataLossLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localAdgAutoFailoverMaxDataLossLimit",
		val,
	)
}

func (j *jsiiProxy_OdbAutonomousDatabase)SetNcharacterSet(val *string) {
	if err := j.validateSetNcharacterSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ncharacterSet",
		val,
	)
}

func (j *jsiiProxy_OdbAutonomousDatabase)SetOdbNetworkId(val *string) {
	if err := j.validateSetOdbNetworkIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"odbNetworkId",
		val,
	)
}

func (j *jsiiProxy_OdbAutonomousDatabase)SetOpenMode(val *string) {
	if err := j.validateSetOpenModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"openMode",
		val,
	)
}

func (j *jsiiProxy_OdbAutonomousDatabase)SetPermissionLevel(val *string) {
	if err := j.validateSetPermissionLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"permissionLevel",
		val,
	)
}

func (j *jsiiProxy_OdbAutonomousDatabase)SetPrivateEndpointIp(val *string) {
	if err := j.validateSetPrivateEndpointIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateEndpointIp",
		val,
	)
}

func (j *jsiiProxy_OdbAutonomousDatabase)SetPrivateEndpointLabel(val *string) {
	if err := j.validateSetPrivateEndpointLabelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateEndpointLabel",
		val,
	)
}

func (j *jsiiProxy_OdbAutonomousDatabase)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_OdbAutonomousDatabase)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_OdbAutonomousDatabase)SetRefreshableMode(val *string) {
	if err := j.validateSetRefreshableModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"refreshableMode",
		val,
	)
}

func (j *jsiiProxy_OdbAutonomousDatabase)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_OdbAutonomousDatabase)SetResourcePoolLeaderId(val *string) {
	if err := j.validateSetResourcePoolLeaderIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourcePoolLeaderId",
		val,
	)
}

func (j *jsiiProxy_OdbAutonomousDatabase)SetSource(val *string) {
	if err := j.validateSetSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"source",
		val,
	)
}

func (j *jsiiProxy_OdbAutonomousDatabase)SetStandbyAllowlistedIps(val *[]*string) {
	if err := j.validateSetStandbyAllowlistedIpsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"standbyAllowlistedIps",
		val,
	)
}

func (j *jsiiProxy_OdbAutonomousDatabase)SetStandbyAllowlistedIpsSource(val *string) {
	if err := j.validateSetStandbyAllowlistedIpsSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"standbyAllowlistedIpsSource",
		val,
	)
}

func (j *jsiiProxy_OdbAutonomousDatabase)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_OdbAutonomousDatabase)SetTimeOfAutoRefreshStart(val *string) {
	if err := j.validateSetTimeOfAutoRefreshStartParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeOfAutoRefreshStart",
		val,
	)
}

// Generates CDKTN code for importing a OdbAutonomousDatabase resource upon running "cdktn plan <stack-name>".
func OdbAutonomousDatabase_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateOdbAutonomousDatabase_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-aws.odbAutonomousDatabase.OdbAutonomousDatabase",
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
func OdbAutonomousDatabase_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOdbAutonomousDatabase_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-aws.odbAutonomousDatabase.OdbAutonomousDatabase",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func OdbAutonomousDatabase_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOdbAutonomousDatabase_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-aws.odbAutonomousDatabase.OdbAutonomousDatabase",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func OdbAutonomousDatabase_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOdbAutonomousDatabase_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-aws.odbAutonomousDatabase.OdbAutonomousDatabase",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func OdbAutonomousDatabase_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-aws.odbAutonomousDatabase.OdbAutonomousDatabase",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (o *jsiiProxy_OdbAutonomousDatabase) AddMoveTarget(moveTarget *string) {
	if err := o.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (o *jsiiProxy_OdbAutonomousDatabase) AddOverride(path *string, value interface{}) {
	if err := o.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (o *jsiiProxy_OdbAutonomousDatabase) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := o.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OdbAutonomousDatabase) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := o.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OdbAutonomousDatabase) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := o.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OdbAutonomousDatabase) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := o.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OdbAutonomousDatabase) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := o.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OdbAutonomousDatabase) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := o.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OdbAutonomousDatabase) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := o.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OdbAutonomousDatabase) GetStringAttribute(terraformAttribute *string) *string {
	if err := o.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OdbAutonomousDatabase) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := o.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OdbAutonomousDatabase) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OdbAutonomousDatabase) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := o.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (o *jsiiProxy_OdbAutonomousDatabase) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := o.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OdbAutonomousDatabase) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := o.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		o,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OdbAutonomousDatabase) MoveFromId(id *string) {
	if err := o.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"moveFromId",
		[]interface{}{id},
	)
}

func (o *jsiiProxy_OdbAutonomousDatabase) MoveTo(moveTarget *string, index interface{}) {
	if err := o.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (o *jsiiProxy_OdbAutonomousDatabase) MoveToId(id *string) {
	if err := o.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"moveToId",
		[]interface{}{id},
	)
}

func (o *jsiiProxy_OdbAutonomousDatabase) OverrideLogicalId(newLogicalId *string) {
	if err := o.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (o *jsiiProxy_OdbAutonomousDatabase) PutAdminPasswordSource(value interface{}) {
	if err := o.validatePutAdminPasswordSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putAdminPasswordSource",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OdbAutonomousDatabase) PutCustomerContactsToSendToOci(value interface{}) {
	if err := o.validatePutCustomerContactsToSendToOciParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putCustomerContactsToSendToOci",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OdbAutonomousDatabase) PutDbToolsDetails(value interface{}) {
	if err := o.validatePutDbToolsDetailsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putDbToolsDetails",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OdbAutonomousDatabase) PutLongTermBackupSchedule(value interface{}) {
	if err := o.validatePutLongTermBackupScheduleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putLongTermBackupSchedule",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OdbAutonomousDatabase) PutResourcePoolSummary(value interface{}) {
	if err := o.validatePutResourcePoolSummaryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putResourcePoolSummary",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OdbAutonomousDatabase) PutScheduledOperations(value interface{}) {
	if err := o.validatePutScheduledOperationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putScheduledOperations",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OdbAutonomousDatabase) PutSourceConfiguration(value interface{}) {
	if err := o.validatePutSourceConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putSourceConfiguration",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OdbAutonomousDatabase) PutTimeouts(value *OdbAutonomousDatabaseTimeouts) {
	if err := o.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OdbAutonomousDatabase) PutTransportableTablespace(value interface{}) {
	if err := o.validatePutTransportableTablespaceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putTransportableTablespace",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OdbAutonomousDatabase) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := o.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (o *jsiiProxy_OdbAutonomousDatabase) ResetAdminPassword() {
	_jsii_.InvokeVoid(
		o,
		"resetAdminPassword",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbAutonomousDatabase) ResetAdminPasswordSource() {
	_jsii_.InvokeVoid(
		o,
		"resetAdminPasswordSource",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbAutonomousDatabase) ResetAdminPasswordWo() {
	_jsii_.InvokeVoid(
		o,
		"resetAdminPasswordWo",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbAutonomousDatabase) ResetAdminPasswordWoVersion() {
	_jsii_.InvokeVoid(
		o,
		"resetAdminPasswordWoVersion",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbAutonomousDatabase) ResetAllowlistedIps() {
	_jsii_.InvokeVoid(
		o,
		"resetAllowlistedIps",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbAutonomousDatabase) ResetAutonomousMaintenanceScheduleType() {
	_jsii_.InvokeVoid(
		o,
		"resetAutonomousMaintenanceScheduleType",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbAutonomousDatabase) ResetAutoRefreshFrequencyInSeconds() {
	_jsii_.InvokeVoid(
		o,
		"resetAutoRefreshFrequencyInSeconds",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbAutonomousDatabase) ResetAutoRefreshPointLagInSeconds() {
	_jsii_.InvokeVoid(
		o,
		"resetAutoRefreshPointLagInSeconds",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbAutonomousDatabase) ResetBackupRetentionPeriodInDays() {
	_jsii_.InvokeVoid(
		o,
		"resetBackupRetentionPeriodInDays",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbAutonomousDatabase) ResetByolComputeCountLimit() {
	_jsii_.InvokeVoid(
		o,
		"resetByolComputeCountLimit",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbAutonomousDatabase) ResetCharacterSet() {
	_jsii_.InvokeVoid(
		o,
		"resetCharacterSet",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbAutonomousDatabase) ResetComputeCount() {
	_jsii_.InvokeVoid(
		o,
		"resetComputeCount",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbAutonomousDatabase) ResetCpuCoreCount() {
	_jsii_.InvokeVoid(
		o,
		"resetCpuCoreCount",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbAutonomousDatabase) ResetCustomerContactsToSendToOci() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomerContactsToSendToOci",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbAutonomousDatabase) ResetDatabaseEdition() {
	_jsii_.InvokeVoid(
		o,
		"resetDatabaseEdition",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbAutonomousDatabase) ResetDataStorageSizeInGbs() {
	_jsii_.InvokeVoid(
		o,
		"resetDataStorageSizeInGbs",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbAutonomousDatabase) ResetDataStorageSizeInTbs() {
	_jsii_.InvokeVoid(
		o,
		"resetDataStorageSizeInTbs",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbAutonomousDatabase) ResetDbName() {
	_jsii_.InvokeVoid(
		o,
		"resetDbName",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbAutonomousDatabase) ResetDbToolsDetails() {
	_jsii_.InvokeVoid(
		o,
		"resetDbToolsDetails",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbAutonomousDatabase) ResetDbVersion() {
	_jsii_.InvokeVoid(
		o,
		"resetDbVersion",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbAutonomousDatabase) ResetDbWorkload() {
	_jsii_.InvokeVoid(
		o,
		"resetDbWorkload",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbAutonomousDatabase) ResetDisplayName() {
	_jsii_.InvokeVoid(
		o,
		"resetDisplayName",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbAutonomousDatabase) ResetEncryptionKeyProvider() {
	_jsii_.InvokeVoid(
		o,
		"resetEncryptionKeyProvider",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbAutonomousDatabase) ResetIsAutoScalingEnabled() {
	_jsii_.InvokeVoid(
		o,
		"resetIsAutoScalingEnabled",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbAutonomousDatabase) ResetIsAutoScalingForStorageEnabled() {
	_jsii_.InvokeVoid(
		o,
		"resetIsAutoScalingForStorageEnabled",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbAutonomousDatabase) ResetIsBackupRetentionLocked() {
	_jsii_.InvokeVoid(
		o,
		"resetIsBackupRetentionLocked",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbAutonomousDatabase) ResetIsLocalDataGuardEnabled() {
	_jsii_.InvokeVoid(
		o,
		"resetIsLocalDataGuardEnabled",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbAutonomousDatabase) ResetIsMtlsConnectionRequired() {
	_jsii_.InvokeVoid(
		o,
		"resetIsMtlsConnectionRequired",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbAutonomousDatabase) ResetIsRefreshableClone() {
	_jsii_.InvokeVoid(
		o,
		"resetIsRefreshableClone",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbAutonomousDatabase) ResetKmsKeyId() {
	_jsii_.InvokeVoid(
		o,
		"resetKmsKeyId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbAutonomousDatabase) ResetLicenseModel() {
	_jsii_.InvokeVoid(
		o,
		"resetLicenseModel",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbAutonomousDatabase) ResetLocalAdgAutoFailoverMaxDataLossLimit() {
	_jsii_.InvokeVoid(
		o,
		"resetLocalAdgAutoFailoverMaxDataLossLimit",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbAutonomousDatabase) ResetLongTermBackupSchedule() {
	_jsii_.InvokeVoid(
		o,
		"resetLongTermBackupSchedule",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbAutonomousDatabase) ResetNcharacterSet() {
	_jsii_.InvokeVoid(
		o,
		"resetNcharacterSet",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbAutonomousDatabase) ResetOdbNetworkId() {
	_jsii_.InvokeVoid(
		o,
		"resetOdbNetworkId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbAutonomousDatabase) ResetOpenMode() {
	_jsii_.InvokeVoid(
		o,
		"resetOpenMode",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbAutonomousDatabase) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		o,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbAutonomousDatabase) ResetPermissionLevel() {
	_jsii_.InvokeVoid(
		o,
		"resetPermissionLevel",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbAutonomousDatabase) ResetPrivateEndpointIp() {
	_jsii_.InvokeVoid(
		o,
		"resetPrivateEndpointIp",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbAutonomousDatabase) ResetPrivateEndpointLabel() {
	_jsii_.InvokeVoid(
		o,
		"resetPrivateEndpointLabel",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbAutonomousDatabase) ResetRefreshableMode() {
	_jsii_.InvokeVoid(
		o,
		"resetRefreshableMode",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbAutonomousDatabase) ResetRegion() {
	_jsii_.InvokeVoid(
		o,
		"resetRegion",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbAutonomousDatabase) ResetResourcePoolLeaderId() {
	_jsii_.InvokeVoid(
		o,
		"resetResourcePoolLeaderId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbAutonomousDatabase) ResetResourcePoolSummary() {
	_jsii_.InvokeVoid(
		o,
		"resetResourcePoolSummary",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbAutonomousDatabase) ResetScheduledOperations() {
	_jsii_.InvokeVoid(
		o,
		"resetScheduledOperations",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbAutonomousDatabase) ResetSource() {
	_jsii_.InvokeVoid(
		o,
		"resetSource",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbAutonomousDatabase) ResetSourceConfiguration() {
	_jsii_.InvokeVoid(
		o,
		"resetSourceConfiguration",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbAutonomousDatabase) ResetStandbyAllowlistedIps() {
	_jsii_.InvokeVoid(
		o,
		"resetStandbyAllowlistedIps",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbAutonomousDatabase) ResetStandbyAllowlistedIpsSource() {
	_jsii_.InvokeVoid(
		o,
		"resetStandbyAllowlistedIpsSource",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbAutonomousDatabase) ResetTags() {
	_jsii_.InvokeVoid(
		o,
		"resetTags",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbAutonomousDatabase) ResetTimeOfAutoRefreshStart() {
	_jsii_.InvokeVoid(
		o,
		"resetTimeOfAutoRefreshStart",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbAutonomousDatabase) ResetTimeouts() {
	_jsii_.InvokeVoid(
		o,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbAutonomousDatabase) ResetTransportableTablespace() {
	_jsii_.InvokeVoid(
		o,
		"resetTransportableTablespace",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbAutonomousDatabase) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OdbAutonomousDatabase) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OdbAutonomousDatabase) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OdbAutonomousDatabase) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OdbAutonomousDatabase) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OdbAutonomousDatabase) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OdbAutonomousDatabase) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		o,
		"with",
		args,
		&returns,
	)

	return returns
}

