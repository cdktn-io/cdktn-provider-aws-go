// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package odbcloudautonomousvmcluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v24/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v24/odbcloudautonomousvmcluster/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.46.0/docs/resources/odb_cloud_autonomous_vm_cluster aws_odb_cloud_autonomous_vm_cluster}.
type OdbCloudAutonomousVmCluster interface {
	cdktn.TerraformResource
	Arn() *string
	AutonomousDataStoragePercentage() *float64
	AutonomousDataStorageSizeInTbs() *float64
	SetAutonomousDataStorageSizeInTbs(val *float64)
	AutonomousDataStorageSizeInTbsInput() *float64
	AvailableAutonomousDataStorageSizeInTbs() *float64
	AvailableContainerDatabases() *float64
	AvailableCpus() *float64
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	CloudExadataInfrastructureArn() *string
	SetCloudExadataInfrastructureArn(val *string)
	CloudExadataInfrastructureArnInput() *string
	CloudExadataInfrastructureId() *string
	SetCloudExadataInfrastructureId(val *string)
	CloudExadataInfrastructureIdInput() *string
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
	CpuCoreCountPerNode() *float64
	SetCpuCoreCountPerNode(val *float64)
	CpuCoreCountPerNodeInput() *float64
	CpuPercentage() *float64
	CreatedAt() *string
	DataStorageSizeInGbs() *float64
	DataStorageSizeInTbs() *float64
	DbServers() *[]*string
	SetDbServers(val *[]*string)
	DbServersInput() *[]*string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	Domain() *string
	ExadataStorageInTbsLowestScaledValue() *float64
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Hostname() *string
	Id() *string
	IsMtlsEnabledVmCluster() interface{}
	SetIsMtlsEnabledVmCluster(val interface{})
	IsMtlsEnabledVmClusterInput() interface{}
	LicenseModel() *string
	SetLicenseModel(val *string)
	LicenseModelInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	MaintenanceWindow() OdbCloudAutonomousVmClusterMaintenanceWindowList
	MaintenanceWindowInput() interface{}
	MaxAcdsLowestScaledValue() *float64
	MemoryPerOracleComputeUnitInGbs() *float64
	SetMemoryPerOracleComputeUnitInGbs(val *float64)
	MemoryPerOracleComputeUnitInGbsInput() *float64
	MemorySizeInGbs() *float64
	// The tree node.
	Node() constructs.Node
	NodeCount() *float64
	NonProvisionableAutonomousContainerDatabases() *float64
	Ocid() *string
	OciResourceAnchorName() *string
	OciUrl() *string
	OdbNetworkArn() *string
	SetOdbNetworkArn(val *string)
	OdbNetworkArnInput() *string
	OdbNetworkId() *string
	SetOdbNetworkId(val *string)
	OdbNetworkIdInput() *string
	OdbNodeStorageSizeInGbs() *float64
	PercentProgress() *float64
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	ProvisionableAutonomousContainerDatabases() *float64
	ProvisionedAutonomousContainerDatabases() *float64
	ProvisionedCpus() *float64
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	ReclaimableCpus() *float64
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	ReservedCpus() *float64
	ScanListenerPortNonTls() *float64
	SetScanListenerPortNonTls(val *float64)
	ScanListenerPortNonTlsInput() *float64
	ScanListenerPortTls() *float64
	SetScanListenerPortTls(val *float64)
	ScanListenerPortTlsInput() *float64
	Shape() *string
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
	TimeDatabaseSslCertificateExpires() *string
	TimeOrdsCertificateExpires() *string
	Timeouts() OdbCloudAutonomousVmClusterTimeoutsOutputReference
	TimeoutsInput() interface{}
	TimeZone() *string
	SetTimeZone(val *string)
	TimeZoneInput() *string
	TotalContainerDatabases() *float64
	SetTotalContainerDatabases(val *float64)
	TotalContainerDatabasesInput() *float64
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
	PutMaintenanceWindow(value interface{})
	PutTimeouts(value *OdbCloudAutonomousVmClusterTimeouts)
	ResetCloudExadataInfrastructureArn()
	ResetCloudExadataInfrastructureId()
	ResetDescription()
	ResetIsMtlsEnabledVmCluster()
	ResetLicenseModel()
	ResetMaintenanceWindow()
	ResetOdbNetworkArn()
	ResetOdbNetworkId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRegion()
	ResetTags()
	ResetTimeouts()
	ResetTimeZone()
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

// The jsii proxy struct for OdbCloudAutonomousVmCluster
type jsiiProxy_OdbCloudAutonomousVmCluster struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster) AutonomousDataStoragePercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autonomousDataStoragePercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster) AutonomousDataStorageSizeInTbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autonomousDataStorageSizeInTbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster) AutonomousDataStorageSizeInTbsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autonomousDataStorageSizeInTbsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster) AvailableAutonomousDataStorageSizeInTbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"availableAutonomousDataStorageSizeInTbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster) AvailableContainerDatabases() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"availableContainerDatabases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster) AvailableCpus() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"availableCpus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster) CloudExadataInfrastructureArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudExadataInfrastructureArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster) CloudExadataInfrastructureArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudExadataInfrastructureArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster) CloudExadataInfrastructureId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudExadataInfrastructureId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster) CloudExadataInfrastructureIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudExadataInfrastructureIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster) ComputeModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster) CpuCoreCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuCoreCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster) CpuCoreCountPerNode() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuCoreCountPerNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster) CpuCoreCountPerNodeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuCoreCountPerNodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster) CpuPercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuPercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster) DataStorageSizeInGbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataStorageSizeInGbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster) DataStorageSizeInTbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataStorageSizeInTbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster) DbServers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dbServers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster) DbServersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dbServersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster) Domain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster) ExadataStorageInTbsLowestScaledValue() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"exadataStorageInTbsLowestScaledValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster) Hostname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster) IsMtlsEnabledVmCluster() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isMtlsEnabledVmCluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster) IsMtlsEnabledVmClusterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isMtlsEnabledVmClusterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster) LicenseModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster) LicenseModelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseModelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster) MaintenanceWindow() OdbCloudAutonomousVmClusterMaintenanceWindowList {
	var returns OdbCloudAutonomousVmClusterMaintenanceWindowList
	_jsii_.Get(
		j,
		"maintenanceWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster) MaintenanceWindowInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"maintenanceWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster) MaxAcdsLowestScaledValue() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAcdsLowestScaledValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster) MemoryPerOracleComputeUnitInGbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryPerOracleComputeUnitInGbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster) MemoryPerOracleComputeUnitInGbsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryPerOracleComputeUnitInGbsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster) MemorySizeInGbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memorySizeInGbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster) NodeCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nodeCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster) NonProvisionableAutonomousContainerDatabases() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nonProvisionableAutonomousContainerDatabases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster) Ocid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ocid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster) OciResourceAnchorName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ociResourceAnchorName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster) OciUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ociUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster) OdbNetworkArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"odbNetworkArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster) OdbNetworkArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"odbNetworkArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster) OdbNetworkId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"odbNetworkId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster) OdbNetworkIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"odbNetworkIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster) OdbNodeStorageSizeInGbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"odbNodeStorageSizeInGbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster) PercentProgress() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"percentProgress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster) ProvisionableAutonomousContainerDatabases() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"provisionableAutonomousContainerDatabases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster) ProvisionedAutonomousContainerDatabases() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"provisionedAutonomousContainerDatabases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster) ProvisionedCpus() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"provisionedCpus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster) ReclaimableCpus() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"reclaimableCpus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster) ReservedCpus() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"reservedCpus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster) ScanListenerPortNonTls() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scanListenerPortNonTls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster) ScanListenerPortNonTlsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scanListenerPortNonTlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster) ScanListenerPortTls() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scanListenerPortTls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster) ScanListenerPortTlsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scanListenerPortTlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster) Shape() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shape",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster) StatusReason() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusReason",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster) TagsAll() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster) TimeDatabaseSslCertificateExpires() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeDatabaseSslCertificateExpires",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster) TimeOrdsCertificateExpires() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeOrdsCertificateExpires",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster) Timeouts() OdbCloudAutonomousVmClusterTimeoutsOutputReference {
	var returns OdbCloudAutonomousVmClusterTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster) TimeZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster) TimeZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster) TotalContainerDatabases() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"totalContainerDatabases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster) TotalContainerDatabasesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"totalContainerDatabasesInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.46.0/docs/resources/odb_cloud_autonomous_vm_cluster aws_odb_cloud_autonomous_vm_cluster} Resource.
func NewOdbCloudAutonomousVmCluster(scope constructs.Construct, id *string, config *OdbCloudAutonomousVmClusterConfig) OdbCloudAutonomousVmCluster {
	_init_.Initialize()

	if err := validateNewOdbCloudAutonomousVmClusterParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_OdbCloudAutonomousVmCluster{}

	_jsii_.Create(
		"@cdktn/provider-aws.odbCloudAutonomousVmCluster.OdbCloudAutonomousVmCluster",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.46.0/docs/resources/odb_cloud_autonomous_vm_cluster aws_odb_cloud_autonomous_vm_cluster} Resource.
func NewOdbCloudAutonomousVmCluster_Override(o OdbCloudAutonomousVmCluster, scope constructs.Construct, id *string, config *OdbCloudAutonomousVmClusterConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.odbCloudAutonomousVmCluster.OdbCloudAutonomousVmCluster",
		[]interface{}{scope, id, config},
		o,
	)
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster)SetAutonomousDataStorageSizeInTbs(val *float64) {
	if err := j.validateSetAutonomousDataStorageSizeInTbsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autonomousDataStorageSizeInTbs",
		val,
	)
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster)SetCloudExadataInfrastructureArn(val *string) {
	if err := j.validateSetCloudExadataInfrastructureArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cloudExadataInfrastructureArn",
		val,
	)
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster)SetCloudExadataInfrastructureId(val *string) {
	if err := j.validateSetCloudExadataInfrastructureIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cloudExadataInfrastructureId",
		val,
	)
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster)SetCpuCoreCountPerNode(val *float64) {
	if err := j.validateSetCpuCoreCountPerNodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpuCoreCountPerNode",
		val,
	)
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster)SetDbServers(val *[]*string) {
	if err := j.validateSetDbServersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbServers",
		val,
	)
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster)SetIsMtlsEnabledVmCluster(val interface{}) {
	if err := j.validateSetIsMtlsEnabledVmClusterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isMtlsEnabledVmCluster",
		val,
	)
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster)SetLicenseModel(val *string) {
	if err := j.validateSetLicenseModelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"licenseModel",
		val,
	)
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster)SetMemoryPerOracleComputeUnitInGbs(val *float64) {
	if err := j.validateSetMemoryPerOracleComputeUnitInGbsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memoryPerOracleComputeUnitInGbs",
		val,
	)
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster)SetOdbNetworkArn(val *string) {
	if err := j.validateSetOdbNetworkArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"odbNetworkArn",
		val,
	)
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster)SetOdbNetworkId(val *string) {
	if err := j.validateSetOdbNetworkIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"odbNetworkId",
		val,
	)
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster)SetScanListenerPortNonTls(val *float64) {
	if err := j.validateSetScanListenerPortNonTlsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scanListenerPortNonTls",
		val,
	)
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster)SetScanListenerPortTls(val *float64) {
	if err := j.validateSetScanListenerPortTlsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scanListenerPortTls",
		val,
	)
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster)SetTimeZone(val *string) {
	if err := j.validateSetTimeZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeZone",
		val,
	)
}

func (j *jsiiProxy_OdbCloudAutonomousVmCluster)SetTotalContainerDatabases(val *float64) {
	if err := j.validateSetTotalContainerDatabasesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"totalContainerDatabases",
		val,
	)
}

// Generates CDKTN code for importing a OdbCloudAutonomousVmCluster resource upon running "cdktn plan <stack-name>".
func OdbCloudAutonomousVmCluster_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateOdbCloudAutonomousVmCluster_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-aws.odbCloudAutonomousVmCluster.OdbCloudAutonomousVmCluster",
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
func OdbCloudAutonomousVmCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOdbCloudAutonomousVmCluster_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-aws.odbCloudAutonomousVmCluster.OdbCloudAutonomousVmCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func OdbCloudAutonomousVmCluster_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOdbCloudAutonomousVmCluster_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-aws.odbCloudAutonomousVmCluster.OdbCloudAutonomousVmCluster",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func OdbCloudAutonomousVmCluster_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOdbCloudAutonomousVmCluster_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-aws.odbCloudAutonomousVmCluster.OdbCloudAutonomousVmCluster",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func OdbCloudAutonomousVmCluster_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-aws.odbCloudAutonomousVmCluster.OdbCloudAutonomousVmCluster",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (o *jsiiProxy_OdbCloudAutonomousVmCluster) AddMoveTarget(moveTarget *string) {
	if err := o.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (o *jsiiProxy_OdbCloudAutonomousVmCluster) AddOverride(path *string, value interface{}) {
	if err := o.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (o *jsiiProxy_OdbCloudAutonomousVmCluster) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_OdbCloudAutonomousVmCluster) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (o *jsiiProxy_OdbCloudAutonomousVmCluster) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_OdbCloudAutonomousVmCluster) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_OdbCloudAutonomousVmCluster) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_OdbCloudAutonomousVmCluster) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_OdbCloudAutonomousVmCluster) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_OdbCloudAutonomousVmCluster) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_OdbCloudAutonomousVmCluster) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_OdbCloudAutonomousVmCluster) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OdbCloudAutonomousVmCluster) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := o.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (o *jsiiProxy_OdbCloudAutonomousVmCluster) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (o *jsiiProxy_OdbCloudAutonomousVmCluster) MoveFromId(id *string) {
	if err := o.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"moveFromId",
		[]interface{}{id},
	)
}

func (o *jsiiProxy_OdbCloudAutonomousVmCluster) MoveTo(moveTarget *string, index interface{}) {
	if err := o.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (o *jsiiProxy_OdbCloudAutonomousVmCluster) MoveToId(id *string) {
	if err := o.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"moveToId",
		[]interface{}{id},
	)
}

func (o *jsiiProxy_OdbCloudAutonomousVmCluster) OverrideLogicalId(newLogicalId *string) {
	if err := o.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (o *jsiiProxy_OdbCloudAutonomousVmCluster) PutMaintenanceWindow(value interface{}) {
	if err := o.validatePutMaintenanceWindowParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putMaintenanceWindow",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OdbCloudAutonomousVmCluster) PutTimeouts(value *OdbCloudAutonomousVmClusterTimeouts) {
	if err := o.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OdbCloudAutonomousVmCluster) ResetCloudExadataInfrastructureArn() {
	_jsii_.InvokeVoid(
		o,
		"resetCloudExadataInfrastructureArn",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbCloudAutonomousVmCluster) ResetCloudExadataInfrastructureId() {
	_jsii_.InvokeVoid(
		o,
		"resetCloudExadataInfrastructureId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbCloudAutonomousVmCluster) ResetDescription() {
	_jsii_.InvokeVoid(
		o,
		"resetDescription",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbCloudAutonomousVmCluster) ResetIsMtlsEnabledVmCluster() {
	_jsii_.InvokeVoid(
		o,
		"resetIsMtlsEnabledVmCluster",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbCloudAutonomousVmCluster) ResetLicenseModel() {
	_jsii_.InvokeVoid(
		o,
		"resetLicenseModel",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbCloudAutonomousVmCluster) ResetMaintenanceWindow() {
	_jsii_.InvokeVoid(
		o,
		"resetMaintenanceWindow",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbCloudAutonomousVmCluster) ResetOdbNetworkArn() {
	_jsii_.InvokeVoid(
		o,
		"resetOdbNetworkArn",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbCloudAutonomousVmCluster) ResetOdbNetworkId() {
	_jsii_.InvokeVoid(
		o,
		"resetOdbNetworkId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbCloudAutonomousVmCluster) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		o,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbCloudAutonomousVmCluster) ResetRegion() {
	_jsii_.InvokeVoid(
		o,
		"resetRegion",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbCloudAutonomousVmCluster) ResetTags() {
	_jsii_.InvokeVoid(
		o,
		"resetTags",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbCloudAutonomousVmCluster) ResetTimeouts() {
	_jsii_.InvokeVoid(
		o,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbCloudAutonomousVmCluster) ResetTimeZone() {
	_jsii_.InvokeVoid(
		o,
		"resetTimeZone",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbCloudAutonomousVmCluster) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OdbCloudAutonomousVmCluster) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OdbCloudAutonomousVmCluster) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OdbCloudAutonomousVmCluster) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OdbCloudAutonomousVmCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OdbCloudAutonomousVmCluster) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OdbCloudAutonomousVmCluster) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

