// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package odbautonomousdatabase

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktn/provider-aws.odbAutonomousDatabase.OdbAutonomousDatabase",
		reflect.TypeOf((*OdbAutonomousDatabase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "actualUsedDataStorageSizeInTbs", GoGetter: "ActualUsedDataStorageSizeInTbs"},
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "adminPassword", GoGetter: "AdminPassword"},
			_jsii_.MemberProperty{JsiiProperty: "adminPasswordInput", GoGetter: "AdminPasswordInput"},
			_jsii_.MemberProperty{JsiiProperty: "adminPasswordSource", GoGetter: "AdminPasswordSource"},
			_jsii_.MemberProperty{JsiiProperty: "adminPasswordSourceInput", GoGetter: "AdminPasswordSourceInput"},
			_jsii_.MemberProperty{JsiiProperty: "adminPasswordWo", GoGetter: "AdminPasswordWo"},
			_jsii_.MemberProperty{JsiiProperty: "adminPasswordWoInput", GoGetter: "AdminPasswordWoInput"},
			_jsii_.MemberProperty{JsiiProperty: "adminPasswordWoVersion", GoGetter: "AdminPasswordWoVersion"},
			_jsii_.MemberProperty{JsiiProperty: "adminPasswordWoVersionInput", GoGetter: "AdminPasswordWoVersionInput"},
			_jsii_.MemberProperty{JsiiProperty: "allocatedStorageSizeInTbs", GoGetter: "AllocatedStorageSizeInTbs"},
			_jsii_.MemberProperty{JsiiProperty: "allowlistedIps", GoGetter: "AllowlistedIps"},
			_jsii_.MemberProperty{JsiiProperty: "allowlistedIpsInput", GoGetter: "AllowlistedIpsInput"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberProperty{JsiiProperty: "autonomousMaintenanceScheduleType", GoGetter: "AutonomousMaintenanceScheduleType"},
			_jsii_.MemberProperty{JsiiProperty: "autonomousMaintenanceScheduleTypeInput", GoGetter: "AutonomousMaintenanceScheduleTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "autoRefreshFrequencyInSeconds", GoGetter: "AutoRefreshFrequencyInSeconds"},
			_jsii_.MemberProperty{JsiiProperty: "autoRefreshFrequencyInSecondsInput", GoGetter: "AutoRefreshFrequencyInSecondsInput"},
			_jsii_.MemberProperty{JsiiProperty: "autoRefreshPointLagInSeconds", GoGetter: "AutoRefreshPointLagInSeconds"},
			_jsii_.MemberProperty{JsiiProperty: "autoRefreshPointLagInSecondsInput", GoGetter: "AutoRefreshPointLagInSecondsInput"},
			_jsii_.MemberProperty{JsiiProperty: "availabilityZone", GoGetter: "AvailabilityZone"},
			_jsii_.MemberProperty{JsiiProperty: "availabilityZoneId", GoGetter: "AvailabilityZoneId"},
			_jsii_.MemberProperty{JsiiProperty: "availableUpgradeVersions", GoGetter: "AvailableUpgradeVersions"},
			_jsii_.MemberProperty{JsiiProperty: "backupRetentionPeriodInDays", GoGetter: "BackupRetentionPeriodInDays"},
			_jsii_.MemberProperty{JsiiProperty: "backupRetentionPeriodInDaysInput", GoGetter: "BackupRetentionPeriodInDaysInput"},
			_jsii_.MemberProperty{JsiiProperty: "byolComputeCountLimit", GoGetter: "ByolComputeCountLimit"},
			_jsii_.MemberProperty{JsiiProperty: "byolComputeCountLimitInput", GoGetter: "ByolComputeCountLimitInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "characterSet", GoGetter: "CharacterSet"},
			_jsii_.MemberProperty{JsiiProperty: "characterSetInput", GoGetter: "CharacterSetInput"},
			_jsii_.MemberProperty{JsiiProperty: "computeCount", GoGetter: "ComputeCount"},
			_jsii_.MemberProperty{JsiiProperty: "computeCountInput", GoGetter: "ComputeCountInput"},
			_jsii_.MemberProperty{JsiiProperty: "computeModel", GoGetter: "ComputeModel"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "cpuCoreCount", GoGetter: "CpuCoreCount"},
			_jsii_.MemberProperty{JsiiProperty: "cpuCoreCountInput", GoGetter: "CpuCoreCountInput"},
			_jsii_.MemberProperty{JsiiProperty: "createdAt", GoGetter: "CreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "customerContactsToSendToOci", GoGetter: "CustomerContactsToSendToOci"},
			_jsii_.MemberProperty{JsiiProperty: "customerContactsToSendToOciInput", GoGetter: "CustomerContactsToSendToOciInput"},
			_jsii_.MemberProperty{JsiiProperty: "databaseEdition", GoGetter: "DatabaseEdition"},
			_jsii_.MemberProperty{JsiiProperty: "databaseEditionInput", GoGetter: "DatabaseEditionInput"},
			_jsii_.MemberProperty{JsiiProperty: "databaseType", GoGetter: "DatabaseType"},
			_jsii_.MemberProperty{JsiiProperty: "dataStorageSizeInGbs", GoGetter: "DataStorageSizeInGbs"},
			_jsii_.MemberProperty{JsiiProperty: "dataStorageSizeInGbsInput", GoGetter: "DataStorageSizeInGbsInput"},
			_jsii_.MemberProperty{JsiiProperty: "dataStorageSizeInTbs", GoGetter: "DataStorageSizeInTbs"},
			_jsii_.MemberProperty{JsiiProperty: "dataStorageSizeInTbsInput", GoGetter: "DataStorageSizeInTbsInput"},
			_jsii_.MemberProperty{JsiiProperty: "dbName", GoGetter: "DbName"},
			_jsii_.MemberProperty{JsiiProperty: "dbNameInput", GoGetter: "DbNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "dbToolsDetails", GoGetter: "DbToolsDetails"},
			_jsii_.MemberProperty{JsiiProperty: "dbToolsDetailsInput", GoGetter: "DbToolsDetailsInput"},
			_jsii_.MemberProperty{JsiiProperty: "dbVersion", GoGetter: "DbVersion"},
			_jsii_.MemberProperty{JsiiProperty: "dbVersionInput", GoGetter: "DbVersionInput"},
			_jsii_.MemberProperty{JsiiProperty: "dbWorkload", GoGetter: "DbWorkload"},
			_jsii_.MemberProperty{JsiiProperty: "dbWorkloadInput", GoGetter: "DbWorkloadInput"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "displayName", GoGetter: "DisplayName"},
			_jsii_.MemberProperty{JsiiProperty: "displayNameInput", GoGetter: "DisplayNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "encryptionKeyProvider", GoGetter: "EncryptionKeyProvider"},
			_jsii_.MemberProperty{JsiiProperty: "encryptionKeyProviderInput", GoGetter: "EncryptionKeyProviderInput"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "hasResourceMove", GoMethod: "HasResourceMove"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isAutoScalingEnabled", GoGetter: "IsAutoScalingEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "isAutoScalingEnabledInput", GoGetter: "IsAutoScalingEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "isAutoScalingForStorageEnabled", GoGetter: "IsAutoScalingForStorageEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "isAutoScalingForStorageEnabledInput", GoGetter: "IsAutoScalingForStorageEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "isBackupRetentionLocked", GoGetter: "IsBackupRetentionLocked"},
			_jsii_.MemberProperty{JsiiProperty: "isBackupRetentionLockedInput", GoGetter: "IsBackupRetentionLockedInput"},
			_jsii_.MemberProperty{JsiiProperty: "isLocalDataGuardEnabled", GoGetter: "IsLocalDataGuardEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "isLocalDataGuardEnabledInput", GoGetter: "IsLocalDataGuardEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "isMtlsConnectionRequired", GoGetter: "IsMtlsConnectionRequired"},
			_jsii_.MemberProperty{JsiiProperty: "isMtlsConnectionRequiredInput", GoGetter: "IsMtlsConnectionRequiredInput"},
			_jsii_.MemberProperty{JsiiProperty: "isRefreshableClone", GoGetter: "IsRefreshableClone"},
			_jsii_.MemberProperty{JsiiProperty: "isRefreshableCloneInput", GoGetter: "IsRefreshableCloneInput"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKeyId", GoGetter: "KmsKeyId"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKeyIdInput", GoGetter: "KmsKeyIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "licenseModel", GoGetter: "LicenseModel"},
			_jsii_.MemberProperty{JsiiProperty: "licenseModelInput", GoGetter: "LicenseModelInput"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "localAdgAutoFailoverMaxDataLossLimit", GoGetter: "LocalAdgAutoFailoverMaxDataLossLimit"},
			_jsii_.MemberProperty{JsiiProperty: "localAdgAutoFailoverMaxDataLossLimitInput", GoGetter: "LocalAdgAutoFailoverMaxDataLossLimitInput"},
			_jsii_.MemberProperty{JsiiProperty: "longTermBackupSchedule", GoGetter: "LongTermBackupSchedule"},
			_jsii_.MemberProperty{JsiiProperty: "longTermBackupScheduleInput", GoGetter: "LongTermBackupScheduleInput"},
			_jsii_.MemberMethod{JsiiMethod: "markWriteOnlyAttribute", GoMethod: "MarkWriteOnlyAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "ncharacterSet", GoGetter: "NcharacterSet"},
			_jsii_.MemberProperty{JsiiProperty: "ncharacterSetInput", GoGetter: "NcharacterSetInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "ocid", GoGetter: "Ocid"},
			_jsii_.MemberProperty{JsiiProperty: "ociResourceAnchorName", GoGetter: "OciResourceAnchorName"},
			_jsii_.MemberProperty{JsiiProperty: "ociUrl", GoGetter: "OciUrl"},
			_jsii_.MemberProperty{JsiiProperty: "odbNetworkArn", GoGetter: "OdbNetworkArn"},
			_jsii_.MemberProperty{JsiiProperty: "odbNetworkId", GoGetter: "OdbNetworkId"},
			_jsii_.MemberProperty{JsiiProperty: "odbNetworkIdInput", GoGetter: "OdbNetworkIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "openMode", GoGetter: "OpenMode"},
			_jsii_.MemberProperty{JsiiProperty: "openModeInput", GoGetter: "OpenModeInput"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "percentProgress", GoGetter: "PercentProgress"},
			_jsii_.MemberProperty{JsiiProperty: "permissionLevel", GoGetter: "PermissionLevel"},
			_jsii_.MemberProperty{JsiiProperty: "permissionLevelInput", GoGetter: "PermissionLevelInput"},
			_jsii_.MemberProperty{JsiiProperty: "privateEndpoint", GoGetter: "PrivateEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "privateEndpointIp", GoGetter: "PrivateEndpointIp"},
			_jsii_.MemberProperty{JsiiProperty: "privateEndpointIpInput", GoGetter: "PrivateEndpointIpInput"},
			_jsii_.MemberProperty{JsiiProperty: "privateEndpointLabel", GoGetter: "PrivateEndpointLabel"},
			_jsii_.MemberProperty{JsiiProperty: "privateEndpointLabelInput", GoGetter: "PrivateEndpointLabelInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putAdminPasswordSource", GoMethod: "PutAdminPasswordSource"},
			_jsii_.MemberMethod{JsiiMethod: "putCustomerContactsToSendToOci", GoMethod: "PutCustomerContactsToSendToOci"},
			_jsii_.MemberMethod{JsiiMethod: "putDbToolsDetails", GoMethod: "PutDbToolsDetails"},
			_jsii_.MemberMethod{JsiiMethod: "putLongTermBackupSchedule", GoMethod: "PutLongTermBackupSchedule"},
			_jsii_.MemberMethod{JsiiMethod: "putResourcePoolSummary", GoMethod: "PutResourcePoolSummary"},
			_jsii_.MemberMethod{JsiiMethod: "putScheduledOperations", GoMethod: "PutScheduledOperations"},
			_jsii_.MemberMethod{JsiiMethod: "putSourceConfiguration", GoMethod: "PutSourceConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "putTimeouts", GoMethod: "PutTimeouts"},
			_jsii_.MemberMethod{JsiiMethod: "putTransportableTablespace", GoMethod: "PutTransportableTablespace"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "refreshableMode", GoGetter: "RefreshableMode"},
			_jsii_.MemberProperty{JsiiProperty: "refreshableModeInput", GoGetter: "RefreshableModeInput"},
			_jsii_.MemberProperty{JsiiProperty: "region", GoGetter: "Region"},
			_jsii_.MemberProperty{JsiiProperty: "regionInput", GoGetter: "RegionInput"},
			_jsii_.MemberMethod{JsiiMethod: "registerProviderFeatureUsage", GoMethod: "RegisterProviderFeatureUsage"},
			_jsii_.MemberMethod{JsiiMethod: "resetAdminPassword", GoMethod: "ResetAdminPassword"},
			_jsii_.MemberMethod{JsiiMethod: "resetAdminPasswordSource", GoMethod: "ResetAdminPasswordSource"},
			_jsii_.MemberMethod{JsiiMethod: "resetAdminPasswordWo", GoMethod: "ResetAdminPasswordWo"},
			_jsii_.MemberMethod{JsiiMethod: "resetAdminPasswordWoVersion", GoMethod: "ResetAdminPasswordWoVersion"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowlistedIps", GoMethod: "ResetAllowlistedIps"},
			_jsii_.MemberMethod{JsiiMethod: "resetAutonomousMaintenanceScheduleType", GoMethod: "ResetAutonomousMaintenanceScheduleType"},
			_jsii_.MemberMethod{JsiiMethod: "resetAutoRefreshFrequencyInSeconds", GoMethod: "ResetAutoRefreshFrequencyInSeconds"},
			_jsii_.MemberMethod{JsiiMethod: "resetAutoRefreshPointLagInSeconds", GoMethod: "ResetAutoRefreshPointLagInSeconds"},
			_jsii_.MemberMethod{JsiiMethod: "resetBackupRetentionPeriodInDays", GoMethod: "ResetBackupRetentionPeriodInDays"},
			_jsii_.MemberMethod{JsiiMethod: "resetByolComputeCountLimit", GoMethod: "ResetByolComputeCountLimit"},
			_jsii_.MemberMethod{JsiiMethod: "resetCharacterSet", GoMethod: "ResetCharacterSet"},
			_jsii_.MemberMethod{JsiiMethod: "resetComputeCount", GoMethod: "ResetComputeCount"},
			_jsii_.MemberMethod{JsiiMethod: "resetCpuCoreCount", GoMethod: "ResetCpuCoreCount"},
			_jsii_.MemberMethod{JsiiMethod: "resetCustomerContactsToSendToOci", GoMethod: "ResetCustomerContactsToSendToOci"},
			_jsii_.MemberMethod{JsiiMethod: "resetDatabaseEdition", GoMethod: "ResetDatabaseEdition"},
			_jsii_.MemberMethod{JsiiMethod: "resetDataStorageSizeInGbs", GoMethod: "ResetDataStorageSizeInGbs"},
			_jsii_.MemberMethod{JsiiMethod: "resetDataStorageSizeInTbs", GoMethod: "ResetDataStorageSizeInTbs"},
			_jsii_.MemberMethod{JsiiMethod: "resetDbName", GoMethod: "ResetDbName"},
			_jsii_.MemberMethod{JsiiMethod: "resetDbToolsDetails", GoMethod: "ResetDbToolsDetails"},
			_jsii_.MemberMethod{JsiiMethod: "resetDbVersion", GoMethod: "ResetDbVersion"},
			_jsii_.MemberMethod{JsiiMethod: "resetDbWorkload", GoMethod: "ResetDbWorkload"},
			_jsii_.MemberMethod{JsiiMethod: "resetDisplayName", GoMethod: "ResetDisplayName"},
			_jsii_.MemberMethod{JsiiMethod: "resetEncryptionKeyProvider", GoMethod: "ResetEncryptionKeyProvider"},
			_jsii_.MemberMethod{JsiiMethod: "resetIsAutoScalingEnabled", GoMethod: "ResetIsAutoScalingEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetIsAutoScalingForStorageEnabled", GoMethod: "ResetIsAutoScalingForStorageEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetIsBackupRetentionLocked", GoMethod: "ResetIsBackupRetentionLocked"},
			_jsii_.MemberMethod{JsiiMethod: "resetIsLocalDataGuardEnabled", GoMethod: "ResetIsLocalDataGuardEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetIsMtlsConnectionRequired", GoMethod: "ResetIsMtlsConnectionRequired"},
			_jsii_.MemberMethod{JsiiMethod: "resetIsRefreshableClone", GoMethod: "ResetIsRefreshableClone"},
			_jsii_.MemberMethod{JsiiMethod: "resetKmsKeyId", GoMethod: "ResetKmsKeyId"},
			_jsii_.MemberMethod{JsiiMethod: "resetLicenseModel", GoMethod: "ResetLicenseModel"},
			_jsii_.MemberMethod{JsiiMethod: "resetLocalAdgAutoFailoverMaxDataLossLimit", GoMethod: "ResetLocalAdgAutoFailoverMaxDataLossLimit"},
			_jsii_.MemberMethod{JsiiMethod: "resetLongTermBackupSchedule", GoMethod: "ResetLongTermBackupSchedule"},
			_jsii_.MemberMethod{JsiiMethod: "resetNcharacterSet", GoMethod: "ResetNcharacterSet"},
			_jsii_.MemberMethod{JsiiMethod: "resetOdbNetworkId", GoMethod: "ResetOdbNetworkId"},
			_jsii_.MemberMethod{JsiiMethod: "resetOpenMode", GoMethod: "ResetOpenMode"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPermissionLevel", GoMethod: "ResetPermissionLevel"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrivateEndpointIp", GoMethod: "ResetPrivateEndpointIp"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrivateEndpointLabel", GoMethod: "ResetPrivateEndpointLabel"},
			_jsii_.MemberMethod{JsiiMethod: "resetRefreshableMode", GoMethod: "ResetRefreshableMode"},
			_jsii_.MemberMethod{JsiiMethod: "resetRegion", GoMethod: "ResetRegion"},
			_jsii_.MemberMethod{JsiiMethod: "resetResourcePoolLeaderId", GoMethod: "ResetResourcePoolLeaderId"},
			_jsii_.MemberMethod{JsiiMethod: "resetResourcePoolSummary", GoMethod: "ResetResourcePoolSummary"},
			_jsii_.MemberMethod{JsiiMethod: "resetScheduledOperations", GoMethod: "ResetScheduledOperations"},
			_jsii_.MemberMethod{JsiiMethod: "resetSource", GoMethod: "ResetSource"},
			_jsii_.MemberMethod{JsiiMethod: "resetSourceConfiguration", GoMethod: "ResetSourceConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetStandbyAllowlistedIps", GoMethod: "ResetStandbyAllowlistedIps"},
			_jsii_.MemberMethod{JsiiMethod: "resetStandbyAllowlistedIpsSource", GoMethod: "ResetStandbyAllowlistedIpsSource"},
			_jsii_.MemberMethod{JsiiMethod: "resetTags", GoMethod: "ResetTags"},
			_jsii_.MemberMethod{JsiiMethod: "resetTimeOfAutoRefreshStart", GoMethod: "ResetTimeOfAutoRefreshStart"},
			_jsii_.MemberMethod{JsiiMethod: "resetTimeouts", GoMethod: "ResetTimeouts"},
			_jsii_.MemberMethod{JsiiMethod: "resetTransportableTablespace", GoMethod: "ResetTransportableTablespace"},
			_jsii_.MemberProperty{JsiiProperty: "resourcePoolLeaderId", GoGetter: "ResourcePoolLeaderId"},
			_jsii_.MemberProperty{JsiiProperty: "resourcePoolLeaderIdInput", GoGetter: "ResourcePoolLeaderIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "resourcePoolSummary", GoGetter: "ResourcePoolSummary"},
			_jsii_.MemberProperty{JsiiProperty: "resourcePoolSummaryInput", GoGetter: "ResourcePoolSummaryInput"},
			_jsii_.MemberProperty{JsiiProperty: "scheduledOperations", GoGetter: "ScheduledOperations"},
			_jsii_.MemberProperty{JsiiProperty: "scheduledOperationsInput", GoGetter: "ScheduledOperationsInput"},
			_jsii_.MemberProperty{JsiiProperty: "serviceConsoleUrl", GoGetter: "ServiceConsoleUrl"},
			_jsii_.MemberProperty{JsiiProperty: "source", GoGetter: "Source"},
			_jsii_.MemberProperty{JsiiProperty: "sourceConfiguration", GoGetter: "SourceConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "sourceConfigurationInput", GoGetter: "SourceConfigurationInput"},
			_jsii_.MemberProperty{JsiiProperty: "sourceId", GoGetter: "SourceId"},
			_jsii_.MemberProperty{JsiiProperty: "sourceInput", GoGetter: "SourceInput"},
			_jsii_.MemberProperty{JsiiProperty: "sqlWebDeveloperUrl", GoGetter: "SqlWebDeveloperUrl"},
			_jsii_.MemberProperty{JsiiProperty: "standbyAllowlistedIps", GoGetter: "StandbyAllowlistedIps"},
			_jsii_.MemberProperty{JsiiProperty: "standbyAllowlistedIpsInput", GoGetter: "StandbyAllowlistedIpsInput"},
			_jsii_.MemberProperty{JsiiProperty: "standbyAllowlistedIpsSource", GoGetter: "StandbyAllowlistedIpsSource"},
			_jsii_.MemberProperty{JsiiProperty: "standbyAllowlistedIpsSourceInput", GoGetter: "StandbyAllowlistedIpsSourceInput"},
			_jsii_.MemberProperty{JsiiProperty: "status", GoGetter: "Status"},
			_jsii_.MemberProperty{JsiiProperty: "statusReason", GoGetter: "StatusReason"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsAll", GoGetter: "TagsAll"},
			_jsii_.MemberProperty{JsiiProperty: "tagsInput", GoGetter: "TagsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "timeOfAutoRefreshStart", GoGetter: "TimeOfAutoRefreshStart"},
			_jsii_.MemberProperty{JsiiProperty: "timeOfAutoRefreshStartInput", GoGetter: "TimeOfAutoRefreshStartInput"},
			_jsii_.MemberProperty{JsiiProperty: "timeouts", GoGetter: "Timeouts"},
			_jsii_.MemberProperty{JsiiProperty: "timeoutsInput", GoGetter: "TimeoutsInput"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "transportableTablespace", GoGetter: "TransportableTablespace"},
			_jsii_.MemberProperty{JsiiProperty: "transportableTablespaceInput", GoGetter: "TransportableTablespaceInput"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_OdbAutonomousDatabase{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-aws.odbAutonomousDatabase.OdbAutonomousDatabaseAdminPasswordSource",
		reflect.TypeOf((*OdbAutonomousDatabaseAdminPasswordSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-aws.odbAutonomousDatabase.OdbAutonomousDatabaseAdminPasswordSourceCustomerManagedAwsSecret",
		reflect.TypeOf((*OdbAutonomousDatabaseAdminPasswordSourceCustomerManagedAwsSecret)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-aws.odbAutonomousDatabase.OdbAutonomousDatabaseAdminPasswordSourceCustomerManagedAwsSecretList",
		reflect.TypeOf((*OdbAutonomousDatabaseAdminPasswordSourceCustomerManagedAwsSecretList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_OdbAutonomousDatabaseAdminPasswordSourceCustomerManagedAwsSecretList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-aws.odbAutonomousDatabase.OdbAutonomousDatabaseAdminPasswordSourceCustomerManagedAwsSecretOutputReference",
		reflect.TypeOf((*OdbAutonomousDatabaseAdminPasswordSourceCustomerManagedAwsSecretOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "externalIdType", GoGetter: "ExternalIdType"},
			_jsii_.MemberProperty{JsiiProperty: "externalIdTypeInput", GoGetter: "ExternalIdTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "iamRoleArn", GoGetter: "IamRoleArn"},
			_jsii_.MemberProperty{JsiiProperty: "iamRoleArnInput", GoGetter: "IamRoleArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "secretArn", GoGetter: "SecretArn"},
			_jsii_.MemberProperty{JsiiProperty: "secretArnInput", GoGetter: "SecretArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_OdbAutonomousDatabaseAdminPasswordSourceCustomerManagedAwsSecretOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-aws.odbAutonomousDatabase.OdbAutonomousDatabaseAdminPasswordSourceList",
		reflect.TypeOf((*OdbAutonomousDatabaseAdminPasswordSourceList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_OdbAutonomousDatabaseAdminPasswordSourceList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-aws.odbAutonomousDatabase.OdbAutonomousDatabaseAdminPasswordSourceOutputReference",
		reflect.TypeOf((*OdbAutonomousDatabaseAdminPasswordSourceOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "customerManagedAwsSecret", GoGetter: "CustomerManagedAwsSecret"},
			_jsii_.MemberProperty{JsiiProperty: "customerManagedAwsSecretInput", GoGetter: "CustomerManagedAwsSecretInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "putCustomerManagedAwsSecret", GoMethod: "PutCustomerManagedAwsSecret"},
			_jsii_.MemberMethod{JsiiMethod: "resetCustomerManagedAwsSecret", GoMethod: "ResetCustomerManagedAwsSecret"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_OdbAutonomousDatabaseAdminPasswordSourceOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-aws.odbAutonomousDatabase.OdbAutonomousDatabaseConfig",
		reflect.TypeOf((*OdbAutonomousDatabaseConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-aws.odbAutonomousDatabase.OdbAutonomousDatabaseCustomerContactsToSendToOci",
		reflect.TypeOf((*OdbAutonomousDatabaseCustomerContactsToSendToOci)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-aws.odbAutonomousDatabase.OdbAutonomousDatabaseCustomerContactsToSendToOciList",
		reflect.TypeOf((*OdbAutonomousDatabaseCustomerContactsToSendToOciList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_OdbAutonomousDatabaseCustomerContactsToSendToOciList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-aws.odbAutonomousDatabase.OdbAutonomousDatabaseCustomerContactsToSendToOciOutputReference",
		reflect.TypeOf((*OdbAutonomousDatabaseCustomerContactsToSendToOciOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "email", GoGetter: "Email"},
			_jsii_.MemberProperty{JsiiProperty: "emailInput", GoGetter: "EmailInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_OdbAutonomousDatabaseCustomerContactsToSendToOciOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-aws.odbAutonomousDatabase.OdbAutonomousDatabaseDbToolsDetails",
		reflect.TypeOf((*OdbAutonomousDatabaseDbToolsDetails)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-aws.odbAutonomousDatabase.OdbAutonomousDatabaseDbToolsDetailsList",
		reflect.TypeOf((*OdbAutonomousDatabaseDbToolsDetailsList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_OdbAutonomousDatabaseDbToolsDetailsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-aws.odbAutonomousDatabase.OdbAutonomousDatabaseDbToolsDetailsOutputReference",
		reflect.TypeOf((*OdbAutonomousDatabaseDbToolsDetailsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberProperty{JsiiProperty: "computeCount", GoGetter: "ComputeCount"},
			_jsii_.MemberProperty{JsiiProperty: "computeCountInput", GoGetter: "ComputeCountInput"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isEnabled", GoGetter: "IsEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "isEnabledInput", GoGetter: "IsEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "maxIdleTimeInMinutes", GoGetter: "MaxIdleTimeInMinutes"},
			_jsii_.MemberProperty{JsiiProperty: "maxIdleTimeInMinutesInput", GoGetter: "MaxIdleTimeInMinutesInput"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetComputeCount", GoMethod: "ResetComputeCount"},
			_jsii_.MemberMethod{JsiiMethod: "resetIsEnabled", GoMethod: "ResetIsEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxIdleTimeInMinutes", GoMethod: "ResetMaxIdleTimeInMinutes"},
			_jsii_.MemberMethod{JsiiMethod: "resetName", GoMethod: "ResetName"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_OdbAutonomousDatabaseDbToolsDetailsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-aws.odbAutonomousDatabase.OdbAutonomousDatabaseLongTermBackupSchedule",
		reflect.TypeOf((*OdbAutonomousDatabaseLongTermBackupSchedule)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-aws.odbAutonomousDatabase.OdbAutonomousDatabaseLongTermBackupScheduleList",
		reflect.TypeOf((*OdbAutonomousDatabaseLongTermBackupScheduleList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_OdbAutonomousDatabaseLongTermBackupScheduleList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-aws.odbAutonomousDatabase.OdbAutonomousDatabaseLongTermBackupScheduleOutputReference",
		reflect.TypeOf((*OdbAutonomousDatabaseLongTermBackupScheduleOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isDisabled", GoGetter: "IsDisabled"},
			_jsii_.MemberProperty{JsiiProperty: "isDisabledInput", GoGetter: "IsDisabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "repeatCadence", GoGetter: "RepeatCadence"},
			_jsii_.MemberProperty{JsiiProperty: "repeatCadenceInput", GoGetter: "RepeatCadenceInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetIsDisabled", GoMethod: "ResetIsDisabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetRepeatCadence", GoMethod: "ResetRepeatCadence"},
			_jsii_.MemberMethod{JsiiMethod: "resetRetentionPeriodInDays", GoMethod: "ResetRetentionPeriodInDays"},
			_jsii_.MemberMethod{JsiiMethod: "resetTimeOfBackup", GoMethod: "ResetTimeOfBackup"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "retentionPeriodInDays", GoGetter: "RetentionPeriodInDays"},
			_jsii_.MemberProperty{JsiiProperty: "retentionPeriodInDaysInput", GoGetter: "RetentionPeriodInDaysInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "timeOfBackup", GoGetter: "TimeOfBackup"},
			_jsii_.MemberProperty{JsiiProperty: "timeOfBackupInput", GoGetter: "TimeOfBackupInput"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_OdbAutonomousDatabaseLongTermBackupScheduleOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-aws.odbAutonomousDatabase.OdbAutonomousDatabaseResourcePoolSummary",
		reflect.TypeOf((*OdbAutonomousDatabaseResourcePoolSummary)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-aws.odbAutonomousDatabase.OdbAutonomousDatabaseResourcePoolSummaryList",
		reflect.TypeOf((*OdbAutonomousDatabaseResourcePoolSummaryList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_OdbAutonomousDatabaseResourcePoolSummaryList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-aws.odbAutonomousDatabase.OdbAutonomousDatabaseResourcePoolSummaryOutputReference",
		reflect.TypeOf((*OdbAutonomousDatabaseResourcePoolSummaryOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "availableComputeCapacity", GoGetter: "AvailableComputeCapacity"},
			_jsii_.MemberProperty{JsiiProperty: "availableStorageCapacityInTbs", GoGetter: "AvailableStorageCapacityInTbs"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isDisabled", GoGetter: "IsDisabled"},
			_jsii_.MemberProperty{JsiiProperty: "isDisabledInput", GoGetter: "IsDisabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "poolSize", GoGetter: "PoolSize"},
			_jsii_.MemberProperty{JsiiProperty: "poolSizeInput", GoGetter: "PoolSizeInput"},
			_jsii_.MemberProperty{JsiiProperty: "poolStorageSizeInTbs", GoGetter: "PoolStorageSizeInTbs"},
			_jsii_.MemberProperty{JsiiProperty: "poolStorageSizeInTbsInput", GoGetter: "PoolStorageSizeInTbsInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetIsDisabled", GoMethod: "ResetIsDisabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetPoolSize", GoMethod: "ResetPoolSize"},
			_jsii_.MemberMethod{JsiiMethod: "resetPoolStorageSizeInTbs", GoMethod: "ResetPoolStorageSizeInTbs"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "totalComputeCapacity", GoGetter: "TotalComputeCapacity"},
		},
		func() interface{} {
			j := jsiiProxy_OdbAutonomousDatabaseResourcePoolSummaryOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-aws.odbAutonomousDatabase.OdbAutonomousDatabaseScheduledOperations",
		reflect.TypeOf((*OdbAutonomousDatabaseScheduledOperations)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-aws.odbAutonomousDatabase.OdbAutonomousDatabaseScheduledOperationsList",
		reflect.TypeOf((*OdbAutonomousDatabaseScheduledOperationsList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_OdbAutonomousDatabaseScheduledOperationsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-aws.odbAutonomousDatabase.OdbAutonomousDatabaseScheduledOperationsOutputReference",
		reflect.TypeOf((*OdbAutonomousDatabaseScheduledOperationsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "dayOfWeek", GoGetter: "DayOfWeek"},
			_jsii_.MemberProperty{JsiiProperty: "dayOfWeekInput", GoGetter: "DayOfWeekInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetScheduledStartTime", GoMethod: "ResetScheduledStartTime"},
			_jsii_.MemberMethod{JsiiMethod: "resetScheduledStopTime", GoMethod: "ResetScheduledStopTime"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "scheduledStartTime", GoGetter: "ScheduledStartTime"},
			_jsii_.MemberProperty{JsiiProperty: "scheduledStartTimeInput", GoGetter: "ScheduledStartTimeInput"},
			_jsii_.MemberProperty{JsiiProperty: "scheduledStopTime", GoGetter: "ScheduledStopTime"},
			_jsii_.MemberProperty{JsiiProperty: "scheduledStopTimeInput", GoGetter: "ScheduledStopTimeInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_OdbAutonomousDatabaseScheduledOperationsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-aws.odbAutonomousDatabase.OdbAutonomousDatabaseSourceConfiguration",
		reflect.TypeOf((*OdbAutonomousDatabaseSourceConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-aws.odbAutonomousDatabase.OdbAutonomousDatabaseSourceConfigurationCloneToRefreshable",
		reflect.TypeOf((*OdbAutonomousDatabaseSourceConfigurationCloneToRefreshable)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-aws.odbAutonomousDatabase.OdbAutonomousDatabaseSourceConfigurationCloneToRefreshableList",
		reflect.TypeOf((*OdbAutonomousDatabaseSourceConfigurationCloneToRefreshableList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_OdbAutonomousDatabaseSourceConfigurationCloneToRefreshableList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-aws.odbAutonomousDatabase.OdbAutonomousDatabaseSourceConfigurationCloneToRefreshableOutputReference",
		reflect.TypeOf((*OdbAutonomousDatabaseSourceConfigurationCloneToRefreshableOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "autoRefreshFrequencyInSeconds", GoGetter: "AutoRefreshFrequencyInSeconds"},
			_jsii_.MemberProperty{JsiiProperty: "autoRefreshFrequencyInSecondsInput", GoGetter: "AutoRefreshFrequencyInSecondsInput"},
			_jsii_.MemberProperty{JsiiProperty: "autoRefreshPointLagInSeconds", GoGetter: "AutoRefreshPointLagInSeconds"},
			_jsii_.MemberProperty{JsiiProperty: "autoRefreshPointLagInSecondsInput", GoGetter: "AutoRefreshPointLagInSecondsInput"},
			_jsii_.MemberProperty{JsiiProperty: "cloneType", GoGetter: "CloneType"},
			_jsii_.MemberProperty{JsiiProperty: "cloneTypeInput", GoGetter: "CloneTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "openMode", GoGetter: "OpenMode"},
			_jsii_.MemberProperty{JsiiProperty: "openModeInput", GoGetter: "OpenModeInput"},
			_jsii_.MemberProperty{JsiiProperty: "refreshableMode", GoGetter: "RefreshableMode"},
			_jsii_.MemberProperty{JsiiProperty: "refreshableModeInput", GoGetter: "RefreshableModeInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAutoRefreshFrequencyInSeconds", GoMethod: "ResetAutoRefreshFrequencyInSeconds"},
			_jsii_.MemberMethod{JsiiMethod: "resetAutoRefreshPointLagInSeconds", GoMethod: "ResetAutoRefreshPointLagInSeconds"},
			_jsii_.MemberMethod{JsiiMethod: "resetCloneType", GoMethod: "ResetCloneType"},
			_jsii_.MemberMethod{JsiiMethod: "resetOpenMode", GoMethod: "ResetOpenMode"},
			_jsii_.MemberMethod{JsiiMethod: "resetRefreshableMode", GoMethod: "ResetRefreshableMode"},
			_jsii_.MemberMethod{JsiiMethod: "resetTimeOfAutoRefreshStart", GoMethod: "ResetTimeOfAutoRefreshStart"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "sourceAutonomousDatabaseId", GoGetter: "SourceAutonomousDatabaseId"},
			_jsii_.MemberProperty{JsiiProperty: "sourceAutonomousDatabaseIdInput", GoGetter: "SourceAutonomousDatabaseIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "timeOfAutoRefreshStart", GoGetter: "TimeOfAutoRefreshStart"},
			_jsii_.MemberProperty{JsiiProperty: "timeOfAutoRefreshStartInput", GoGetter: "TimeOfAutoRefreshStartInput"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_OdbAutonomousDatabaseSourceConfigurationCloneToRefreshableOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-aws.odbAutonomousDatabase.OdbAutonomousDatabaseSourceConfigurationCrossRegionDataGuard",
		reflect.TypeOf((*OdbAutonomousDatabaseSourceConfigurationCrossRegionDataGuard)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-aws.odbAutonomousDatabase.OdbAutonomousDatabaseSourceConfigurationCrossRegionDataGuardList",
		reflect.TypeOf((*OdbAutonomousDatabaseSourceConfigurationCrossRegionDataGuardList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_OdbAutonomousDatabaseSourceConfigurationCrossRegionDataGuardList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-aws.odbAutonomousDatabase.OdbAutonomousDatabaseSourceConfigurationCrossRegionDataGuardOutputReference",
		reflect.TypeOf((*OdbAutonomousDatabaseSourceConfigurationCrossRegionDataGuardOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "sourceAutonomousDatabaseArn", GoGetter: "SourceAutonomousDatabaseArn"},
			_jsii_.MemberProperty{JsiiProperty: "sourceAutonomousDatabaseArnInput", GoGetter: "SourceAutonomousDatabaseArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_OdbAutonomousDatabaseSourceConfigurationCrossRegionDataGuardOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-aws.odbAutonomousDatabase.OdbAutonomousDatabaseSourceConfigurationCrossRegionDisasterRecovery",
		reflect.TypeOf((*OdbAutonomousDatabaseSourceConfigurationCrossRegionDisasterRecovery)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-aws.odbAutonomousDatabase.OdbAutonomousDatabaseSourceConfigurationCrossRegionDisasterRecoveryList",
		reflect.TypeOf((*OdbAutonomousDatabaseSourceConfigurationCrossRegionDisasterRecoveryList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_OdbAutonomousDatabaseSourceConfigurationCrossRegionDisasterRecoveryList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-aws.odbAutonomousDatabase.OdbAutonomousDatabaseSourceConfigurationCrossRegionDisasterRecoveryOutputReference",
		reflect.TypeOf((*OdbAutonomousDatabaseSourceConfigurationCrossRegionDisasterRecoveryOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isReplicateAutomaticBackups", GoGetter: "IsReplicateAutomaticBackups"},
			_jsii_.MemberProperty{JsiiProperty: "isReplicateAutomaticBackupsInput", GoGetter: "IsReplicateAutomaticBackupsInput"},
			_jsii_.MemberProperty{JsiiProperty: "remoteDisasterRecoveryType", GoGetter: "RemoteDisasterRecoveryType"},
			_jsii_.MemberProperty{JsiiProperty: "remoteDisasterRecoveryTypeInput", GoGetter: "RemoteDisasterRecoveryTypeInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetIsReplicateAutomaticBackups", GoMethod: "ResetIsReplicateAutomaticBackups"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "sourceAutonomousDatabaseArn", GoGetter: "SourceAutonomousDatabaseArn"},
			_jsii_.MemberProperty{JsiiProperty: "sourceAutonomousDatabaseArnInput", GoGetter: "SourceAutonomousDatabaseArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_OdbAutonomousDatabaseSourceConfigurationCrossRegionDisasterRecoveryOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-aws.odbAutonomousDatabase.OdbAutonomousDatabaseSourceConfigurationDatabaseClone",
		reflect.TypeOf((*OdbAutonomousDatabaseSourceConfigurationDatabaseClone)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-aws.odbAutonomousDatabase.OdbAutonomousDatabaseSourceConfigurationDatabaseCloneList",
		reflect.TypeOf((*OdbAutonomousDatabaseSourceConfigurationDatabaseCloneList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_OdbAutonomousDatabaseSourceConfigurationDatabaseCloneList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-aws.odbAutonomousDatabase.OdbAutonomousDatabaseSourceConfigurationDatabaseCloneOutputReference",
		reflect.TypeOf((*OdbAutonomousDatabaseSourceConfigurationDatabaseCloneOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "cloneType", GoGetter: "CloneType"},
			_jsii_.MemberProperty{JsiiProperty: "cloneTypeInput", GoGetter: "CloneTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "sourceAutonomousDatabaseId", GoGetter: "SourceAutonomousDatabaseId"},
			_jsii_.MemberProperty{JsiiProperty: "sourceAutonomousDatabaseIdInput", GoGetter: "SourceAutonomousDatabaseIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_OdbAutonomousDatabaseSourceConfigurationDatabaseCloneOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-aws.odbAutonomousDatabase.OdbAutonomousDatabaseSourceConfigurationList",
		reflect.TypeOf((*OdbAutonomousDatabaseSourceConfigurationList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_OdbAutonomousDatabaseSourceConfigurationList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-aws.odbAutonomousDatabase.OdbAutonomousDatabaseSourceConfigurationOutputReference",
		reflect.TypeOf((*OdbAutonomousDatabaseSourceConfigurationOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "cloneToRefreshable", GoGetter: "CloneToRefreshable"},
			_jsii_.MemberProperty{JsiiProperty: "cloneToRefreshableInput", GoGetter: "CloneToRefreshableInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "crossRegionDataGuard", GoGetter: "CrossRegionDataGuard"},
			_jsii_.MemberProperty{JsiiProperty: "crossRegionDataGuardInput", GoGetter: "CrossRegionDataGuardInput"},
			_jsii_.MemberProperty{JsiiProperty: "crossRegionDisasterRecovery", GoGetter: "CrossRegionDisasterRecovery"},
			_jsii_.MemberProperty{JsiiProperty: "crossRegionDisasterRecoveryInput", GoGetter: "CrossRegionDisasterRecoveryInput"},
			_jsii_.MemberProperty{JsiiProperty: "databaseClone", GoGetter: "DatabaseClone"},
			_jsii_.MemberProperty{JsiiProperty: "databaseCloneInput", GoGetter: "DatabaseCloneInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "pointInTimeRestore", GoGetter: "PointInTimeRestore"},
			_jsii_.MemberProperty{JsiiProperty: "pointInTimeRestoreInput", GoGetter: "PointInTimeRestoreInput"},
			_jsii_.MemberMethod{JsiiMethod: "putCloneToRefreshable", GoMethod: "PutCloneToRefreshable"},
			_jsii_.MemberMethod{JsiiMethod: "putCrossRegionDataGuard", GoMethod: "PutCrossRegionDataGuard"},
			_jsii_.MemberMethod{JsiiMethod: "putCrossRegionDisasterRecovery", GoMethod: "PutCrossRegionDisasterRecovery"},
			_jsii_.MemberMethod{JsiiMethod: "putDatabaseClone", GoMethod: "PutDatabaseClone"},
			_jsii_.MemberMethod{JsiiMethod: "putPointInTimeRestore", GoMethod: "PutPointInTimeRestore"},
			_jsii_.MemberMethod{JsiiMethod: "putRestoreFromBackup", GoMethod: "PutRestoreFromBackup"},
			_jsii_.MemberMethod{JsiiMethod: "resetCloneToRefreshable", GoMethod: "ResetCloneToRefreshable"},
			_jsii_.MemberMethod{JsiiMethod: "resetCrossRegionDataGuard", GoMethod: "ResetCrossRegionDataGuard"},
			_jsii_.MemberMethod{JsiiMethod: "resetCrossRegionDisasterRecovery", GoMethod: "ResetCrossRegionDisasterRecovery"},
			_jsii_.MemberMethod{JsiiMethod: "resetDatabaseClone", GoMethod: "ResetDatabaseClone"},
			_jsii_.MemberMethod{JsiiMethod: "resetPointInTimeRestore", GoMethod: "ResetPointInTimeRestore"},
			_jsii_.MemberMethod{JsiiMethod: "resetRestoreFromBackup", GoMethod: "ResetRestoreFromBackup"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "restoreFromBackup", GoGetter: "RestoreFromBackup"},
			_jsii_.MemberProperty{JsiiProperty: "restoreFromBackupInput", GoGetter: "RestoreFromBackupInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_OdbAutonomousDatabaseSourceConfigurationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-aws.odbAutonomousDatabase.OdbAutonomousDatabaseSourceConfigurationPointInTimeRestore",
		reflect.TypeOf((*OdbAutonomousDatabaseSourceConfigurationPointInTimeRestore)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-aws.odbAutonomousDatabase.OdbAutonomousDatabaseSourceConfigurationPointInTimeRestoreList",
		reflect.TypeOf((*OdbAutonomousDatabaseSourceConfigurationPointInTimeRestoreList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_OdbAutonomousDatabaseSourceConfigurationPointInTimeRestoreList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-aws.odbAutonomousDatabase.OdbAutonomousDatabaseSourceConfigurationPointInTimeRestoreOutputReference",
		reflect.TypeOf((*OdbAutonomousDatabaseSourceConfigurationPointInTimeRestoreOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "cloneTableSpaceList", GoGetter: "CloneTableSpaceList"},
			_jsii_.MemberProperty{JsiiProperty: "cloneTableSpaceListInput", GoGetter: "CloneTableSpaceListInput"},
			_jsii_.MemberProperty{JsiiProperty: "cloneType", GoGetter: "CloneType"},
			_jsii_.MemberProperty{JsiiProperty: "cloneTypeInput", GoGetter: "CloneTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetCloneTableSpaceList", GoMethod: "ResetCloneTableSpaceList"},
			_jsii_.MemberMethod{JsiiMethod: "resetTimestamp", GoMethod: "ResetTimestamp"},
			_jsii_.MemberMethod{JsiiMethod: "resetUseLatestAvailableBackupTimestamp", GoMethod: "ResetUseLatestAvailableBackupTimestamp"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "sourceAutonomousDatabaseId", GoGetter: "SourceAutonomousDatabaseId"},
			_jsii_.MemberProperty{JsiiProperty: "sourceAutonomousDatabaseIdInput", GoGetter: "SourceAutonomousDatabaseIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "timestamp", GoGetter: "Timestamp"},
			_jsii_.MemberProperty{JsiiProperty: "timestampInput", GoGetter: "TimestampInput"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "useLatestAvailableBackupTimestamp", GoGetter: "UseLatestAvailableBackupTimestamp"},
			_jsii_.MemberProperty{JsiiProperty: "useLatestAvailableBackupTimestampInput", GoGetter: "UseLatestAvailableBackupTimestampInput"},
		},
		func() interface{} {
			j := jsiiProxy_OdbAutonomousDatabaseSourceConfigurationPointInTimeRestoreOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-aws.odbAutonomousDatabase.OdbAutonomousDatabaseSourceConfigurationRestoreFromBackup",
		reflect.TypeOf((*OdbAutonomousDatabaseSourceConfigurationRestoreFromBackup)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-aws.odbAutonomousDatabase.OdbAutonomousDatabaseSourceConfigurationRestoreFromBackupList",
		reflect.TypeOf((*OdbAutonomousDatabaseSourceConfigurationRestoreFromBackupList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_OdbAutonomousDatabaseSourceConfigurationRestoreFromBackupList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-aws.odbAutonomousDatabase.OdbAutonomousDatabaseSourceConfigurationRestoreFromBackupOutputReference",
		reflect.TypeOf((*OdbAutonomousDatabaseSourceConfigurationRestoreFromBackupOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "autonomousDatabaseBackupId", GoGetter: "AutonomousDatabaseBackupId"},
			_jsii_.MemberProperty{JsiiProperty: "autonomousDatabaseBackupIdInput", GoGetter: "AutonomousDatabaseBackupIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "cloneTableSpaceList", GoGetter: "CloneTableSpaceList"},
			_jsii_.MemberProperty{JsiiProperty: "cloneTableSpaceListInput", GoGetter: "CloneTableSpaceListInput"},
			_jsii_.MemberProperty{JsiiProperty: "cloneType", GoGetter: "CloneType"},
			_jsii_.MemberProperty{JsiiProperty: "cloneTypeInput", GoGetter: "CloneTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetCloneTableSpaceList", GoMethod: "ResetCloneTableSpaceList"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_OdbAutonomousDatabaseSourceConfigurationRestoreFromBackupOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-aws.odbAutonomousDatabase.OdbAutonomousDatabaseTimeouts",
		reflect.TypeOf((*OdbAutonomousDatabaseTimeouts)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-aws.odbAutonomousDatabase.OdbAutonomousDatabaseTimeoutsOutputReference",
		reflect.TypeOf((*OdbAutonomousDatabaseTimeoutsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "create", GoGetter: "Create"},
			_jsii_.MemberProperty{JsiiProperty: "createInput", GoGetter: "CreateInput"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "delete", GoGetter: "Delete"},
			_jsii_.MemberProperty{JsiiProperty: "deleteInput", GoGetter: "DeleteInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetCreate", GoMethod: "ResetCreate"},
			_jsii_.MemberMethod{JsiiMethod: "resetDelete", GoMethod: "ResetDelete"},
			_jsii_.MemberMethod{JsiiMethod: "resetUpdate", GoMethod: "ResetUpdate"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "update", GoGetter: "Update"},
			_jsii_.MemberProperty{JsiiProperty: "updateInput", GoGetter: "UpdateInput"},
		},
		func() interface{} {
			j := jsiiProxy_OdbAutonomousDatabaseTimeoutsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-aws.odbAutonomousDatabase.OdbAutonomousDatabaseTransportableTablespace",
		reflect.TypeOf((*OdbAutonomousDatabaseTransportableTablespace)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-aws.odbAutonomousDatabase.OdbAutonomousDatabaseTransportableTablespaceList",
		reflect.TypeOf((*OdbAutonomousDatabaseTransportableTablespaceList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_OdbAutonomousDatabaseTransportableTablespaceList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-aws.odbAutonomousDatabase.OdbAutonomousDatabaseTransportableTablespaceOutputReference",
		reflect.TypeOf((*OdbAutonomousDatabaseTransportableTablespaceOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetTtsBundleUrl", GoMethod: "ResetTtsBundleUrl"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "ttsBundleUrl", GoGetter: "TtsBundleUrl"},
			_jsii_.MemberProperty{JsiiProperty: "ttsBundleUrlInput", GoGetter: "TtsBundleUrlInput"},
		},
		func() interface{} {
			j := jsiiProxy_OdbAutonomousDatabaseTransportableTablespaceOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
}
