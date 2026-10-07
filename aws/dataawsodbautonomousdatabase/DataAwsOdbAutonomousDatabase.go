// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawsodbautonomousdatabase

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/dataawsodbautonomousdatabase/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.68.0/docs/data-sources/odb_autonomous_database aws_odb_autonomous_database}.
type DataAwsOdbAutonomousDatabase interface {
	cdktn.TerraformDataSource
	ActualUsedDataStorageSizeInTbs() *float64
	AdminPasswordSource() DataAwsOdbAutonomousDatabaseAdminPasswordSourceList
	AllocatedStorageSizeInTbs() *float64
	AllowlistedIps() *[]*string
	Arn() *string
	AutonomousMaintenanceScheduleType() *string
	AutoRefreshFrequencyInSeconds() *float64
	AutoRefreshPointLagInSeconds() *float64
	AvailabilityZone() *string
	AvailabilityZoneId() *string
	AvailableUpgradeVersions() *[]*string
	BackupRetentionPeriodInDays() *float64
	ByolComputeCountLimit() *float64
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	CharacterSet() *string
	ComputeCount() *float64
	ComputeModel() *string
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CpuCoreCount() *float64
	CreatedAt() *string
	CustomerContactsToSendToOci() DataAwsOdbAutonomousDatabaseCustomerContactsToSendToOciList
	DatabaseEdition() *string
	DatabaseType() *string
	DataStorageSizeInGbs() *float64
	DataStorageSizeInTbs() *float64
	DbName() *string
	DbToolsDetails() DataAwsOdbAutonomousDatabaseDbToolsDetailsList
	DbVersion() *string
	DbWorkload() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DisplayName() *string
	EncryptionKeyProvider() *string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	IsAutoScalingEnabled() cdktn.IResolvable
	IsAutoScalingForStorageEnabled() cdktn.IResolvable
	IsBackupRetentionLocked() cdktn.IResolvable
	IsLocalDataGuardEnabled() cdktn.IResolvable
	IsMtlsConnectionRequired() cdktn.IResolvable
	IsRefreshableClone() cdktn.IResolvable
	KmsKeyId() *string
	LicenseModel() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	LocalAdgAutoFailoverMaxDataLossLimit() *float64
	LongTermBackupSchedule() DataAwsOdbAutonomousDatabaseLongTermBackupScheduleList
	NcharacterSet() *string
	// The tree node.
	Node() constructs.Node
	Ocid() *string
	OciResourceAnchorName() *string
	OciUrl() *string
	OdbNetworkArn() *string
	OdbNetworkId() *string
	OpenMode() *string
	PercentProgress() *float64
	PermissionLevel() *string
	PrivateEndpoint() *string
	PrivateEndpointIp() *string
	PrivateEndpointLabel() *string
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	RefreshableMode() *string
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	ResourcePoolLeaderId() *string
	ResourcePoolSummary() DataAwsOdbAutonomousDatabaseResourcePoolSummaryList
	ScheduledOperations() DataAwsOdbAutonomousDatabaseScheduledOperationsList
	ServiceConsoleUrl() *string
	SourceId() *string
	SqlWebDeveloperUrl() *string
	StandbyAllowlistedIps() *[]*string
	StandbyAllowlistedIpsSource() *string
	Status() *string
	StatusReason() *string
	Tags() cdktn.StringMap
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TimeOfAutoRefreshStart() *string
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
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRegion()
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

// The jsii proxy struct for DataAwsOdbAutonomousDatabase
type jsiiProxy_DataAwsOdbAutonomousDatabase struct {
	internal.Type__cdktnTerraformDataSource
}

func (j *jsiiProxy_DataAwsOdbAutonomousDatabase) ActualUsedDataStorageSizeInTbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"actualUsedDataStorageSizeInTbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbAutonomousDatabase) AdminPasswordSource() DataAwsOdbAutonomousDatabaseAdminPasswordSourceList {
	var returns DataAwsOdbAutonomousDatabaseAdminPasswordSourceList
	_jsii_.Get(
		j,
		"adminPasswordSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbAutonomousDatabase) AllocatedStorageSizeInTbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"allocatedStorageSizeInTbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbAutonomousDatabase) AllowlistedIps() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowlistedIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbAutonomousDatabase) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbAutonomousDatabase) AutonomousMaintenanceScheduleType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autonomousMaintenanceScheduleType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbAutonomousDatabase) AutoRefreshFrequencyInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autoRefreshFrequencyInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbAutonomousDatabase) AutoRefreshPointLagInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autoRefreshPointLagInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbAutonomousDatabase) AvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbAutonomousDatabase) AvailabilityZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbAutonomousDatabase) AvailableUpgradeVersions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availableUpgradeVersions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbAutonomousDatabase) BackupRetentionPeriodInDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backupRetentionPeriodInDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbAutonomousDatabase) ByolComputeCountLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"byolComputeCountLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbAutonomousDatabase) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbAutonomousDatabase) CharacterSet() *string {
	var returns *string
	_jsii_.Get(
		j,
		"characterSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbAutonomousDatabase) ComputeCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"computeCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbAutonomousDatabase) ComputeModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbAutonomousDatabase) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbAutonomousDatabase) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbAutonomousDatabase) CpuCoreCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuCoreCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbAutonomousDatabase) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbAutonomousDatabase) CustomerContactsToSendToOci() DataAwsOdbAutonomousDatabaseCustomerContactsToSendToOciList {
	var returns DataAwsOdbAutonomousDatabaseCustomerContactsToSendToOciList
	_jsii_.Get(
		j,
		"customerContactsToSendToOci",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbAutonomousDatabase) DatabaseEdition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseEdition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbAutonomousDatabase) DatabaseType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbAutonomousDatabase) DataStorageSizeInGbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataStorageSizeInGbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbAutonomousDatabase) DataStorageSizeInTbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataStorageSizeInTbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbAutonomousDatabase) DbName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbAutonomousDatabase) DbToolsDetails() DataAwsOdbAutonomousDatabaseDbToolsDetailsList {
	var returns DataAwsOdbAutonomousDatabaseDbToolsDetailsList
	_jsii_.Get(
		j,
		"dbToolsDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbAutonomousDatabase) DbVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbAutonomousDatabase) DbWorkload() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbWorkload",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbAutonomousDatabase) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbAutonomousDatabase) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbAutonomousDatabase) EncryptionKeyProvider() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionKeyProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbAutonomousDatabase) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbAutonomousDatabase) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbAutonomousDatabase) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbAutonomousDatabase) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbAutonomousDatabase) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbAutonomousDatabase) IsAutoScalingEnabled() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"isAutoScalingEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbAutonomousDatabase) IsAutoScalingForStorageEnabled() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"isAutoScalingForStorageEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbAutonomousDatabase) IsBackupRetentionLocked() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"isBackupRetentionLocked",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbAutonomousDatabase) IsLocalDataGuardEnabled() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"isLocalDataGuardEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbAutonomousDatabase) IsMtlsConnectionRequired() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"isMtlsConnectionRequired",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbAutonomousDatabase) IsRefreshableClone() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"isRefreshableClone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbAutonomousDatabase) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbAutonomousDatabase) LicenseModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbAutonomousDatabase) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbAutonomousDatabase) LocalAdgAutoFailoverMaxDataLossLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"localAdgAutoFailoverMaxDataLossLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbAutonomousDatabase) LongTermBackupSchedule() DataAwsOdbAutonomousDatabaseLongTermBackupScheduleList {
	var returns DataAwsOdbAutonomousDatabaseLongTermBackupScheduleList
	_jsii_.Get(
		j,
		"longTermBackupSchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbAutonomousDatabase) NcharacterSet() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ncharacterSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbAutonomousDatabase) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbAutonomousDatabase) Ocid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ocid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbAutonomousDatabase) OciResourceAnchorName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ociResourceAnchorName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbAutonomousDatabase) OciUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ociUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbAutonomousDatabase) OdbNetworkArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"odbNetworkArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbAutonomousDatabase) OdbNetworkId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"odbNetworkId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbAutonomousDatabase) OpenMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"openMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbAutonomousDatabase) PercentProgress() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"percentProgress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbAutonomousDatabase) PermissionLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"permissionLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbAutonomousDatabase) PrivateEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbAutonomousDatabase) PrivateEndpointIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateEndpointIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbAutonomousDatabase) PrivateEndpointLabel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateEndpointLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbAutonomousDatabase) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbAutonomousDatabase) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbAutonomousDatabase) RefreshableMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"refreshableMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbAutonomousDatabase) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbAutonomousDatabase) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbAutonomousDatabase) ResourcePoolLeaderId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourcePoolLeaderId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbAutonomousDatabase) ResourcePoolSummary() DataAwsOdbAutonomousDatabaseResourcePoolSummaryList {
	var returns DataAwsOdbAutonomousDatabaseResourcePoolSummaryList
	_jsii_.Get(
		j,
		"resourcePoolSummary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbAutonomousDatabase) ScheduledOperations() DataAwsOdbAutonomousDatabaseScheduledOperationsList {
	var returns DataAwsOdbAutonomousDatabaseScheduledOperationsList
	_jsii_.Get(
		j,
		"scheduledOperations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbAutonomousDatabase) ServiceConsoleUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceConsoleUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbAutonomousDatabase) SourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbAutonomousDatabase) SqlWebDeveloperUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqlWebDeveloperUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbAutonomousDatabase) StandbyAllowlistedIps() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"standbyAllowlistedIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbAutonomousDatabase) StandbyAllowlistedIpsSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"standbyAllowlistedIpsSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbAutonomousDatabase) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbAutonomousDatabase) StatusReason() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusReason",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbAutonomousDatabase) Tags() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbAutonomousDatabase) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbAutonomousDatabase) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbAutonomousDatabase) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbAutonomousDatabase) TimeOfAutoRefreshStart() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeOfAutoRefreshStart",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.68.0/docs/data-sources/odb_autonomous_database aws_odb_autonomous_database} Data Source.
func NewDataAwsOdbAutonomousDatabase(scope constructs.Construct, id *string, config *DataAwsOdbAutonomousDatabaseConfig) DataAwsOdbAutonomousDatabase {
	_init_.Initialize()

	if err := validateNewDataAwsOdbAutonomousDatabaseParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsOdbAutonomousDatabase{}

	_jsii_.Create(
		"@cdktn/provider-aws.dataAwsOdbAutonomousDatabase.DataAwsOdbAutonomousDatabase",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.68.0/docs/data-sources/odb_autonomous_database aws_odb_autonomous_database} Data Source.
func NewDataAwsOdbAutonomousDatabase_Override(d DataAwsOdbAutonomousDatabase, scope constructs.Construct, id *string, config *DataAwsOdbAutonomousDatabaseConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.dataAwsOdbAutonomousDatabase.DataAwsOdbAutonomousDatabase",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsOdbAutonomousDatabase)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsOdbAutonomousDatabase)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsOdbAutonomousDatabase)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataAwsOdbAutonomousDatabase)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataAwsOdbAutonomousDatabase)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsOdbAutonomousDatabase)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataAwsOdbAutonomousDatabase)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

// Generates CDKTN code for importing a DataAwsOdbAutonomousDatabase resource upon running "cdktn plan <stack-name>".
func DataAwsOdbAutonomousDatabase_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateDataAwsOdbAutonomousDatabase_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-aws.dataAwsOdbAutonomousDatabase.DataAwsOdbAutonomousDatabase",
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
func DataAwsOdbAutonomousDatabase_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsOdbAutonomousDatabase_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-aws.dataAwsOdbAutonomousDatabase.DataAwsOdbAutonomousDatabase",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAwsOdbAutonomousDatabase_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsOdbAutonomousDatabase_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-aws.dataAwsOdbAutonomousDatabase.DataAwsOdbAutonomousDatabase",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAwsOdbAutonomousDatabase_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsOdbAutonomousDatabase_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-aws.dataAwsOdbAutonomousDatabase.DataAwsOdbAutonomousDatabase",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsOdbAutonomousDatabase_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-aws.dataAwsOdbAutonomousDatabase.DataAwsOdbAutonomousDatabase",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataAwsOdbAutonomousDatabase) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataAwsOdbAutonomousDatabase) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsOdbAutonomousDatabase) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsOdbAutonomousDatabase) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsOdbAutonomousDatabase) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsOdbAutonomousDatabase) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsOdbAutonomousDatabase) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsOdbAutonomousDatabase) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsOdbAutonomousDatabase) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsOdbAutonomousDatabase) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsOdbAutonomousDatabase) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsOdbAutonomousDatabase) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataAwsOdbAutonomousDatabase) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := d.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (d *jsiiProxy_DataAwsOdbAutonomousDatabase) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsOdbAutonomousDatabase) ResetRegion() {
	_jsii_.InvokeVoid(
		d,
		"resetRegion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsOdbAutonomousDatabase) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsOdbAutonomousDatabase) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsOdbAutonomousDatabase) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsOdbAutonomousDatabase) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsOdbAutonomousDatabase) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsOdbAutonomousDatabase) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsOdbAutonomousDatabase) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

