// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package odbautonomousdatabase

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type OdbAutonomousDatabaseConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Password for the ADMIN user.
	//
	// This value is stored in Terraform state. Use admin_password_wo with Terraform 1.11 or later to avoid storing the password in state.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.68.0/docs/resources/odb_autonomous_database#admin_password OdbAutonomousDatabase#admin_password}
	AdminPassword *string `field:"optional" json:"adminPassword" yaml:"adminPassword"`
	// admin_password_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.68.0/docs/resources/odb_autonomous_database#admin_password_source OdbAutonomousDatabase#admin_password_source}
	AdminPasswordSource interface{} `field:"optional" json:"adminPasswordSource" yaml:"adminPasswordSource"`
	// Password for the ADMIN user. This write-only value is never stored in Terraform state.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.68.0/docs/resources/odb_autonomous_database#admin_password_wo OdbAutonomousDatabase#admin_password_wo}
	AdminPasswordWo *string `field:"optional" json:"adminPasswordWo" yaml:"adminPasswordWo"`
	// Arbitrary version used to trigger an ADMIN password update.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.68.0/docs/resources/odb_autonomous_database#admin_password_wo_version OdbAutonomousDatabase#admin_password_wo_version}
	AdminPasswordWoVersion *float64 `field:"optional" json:"adminPasswordWoVersion" yaml:"adminPasswordWoVersion"`
	// IP addresses allowed to access the Autonomous Database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.68.0/docs/resources/odb_autonomous_database#allowlisted_ips OdbAutonomousDatabase#allowlisted_ips}
	AllowlistedIps *[]*string `field:"optional" json:"allowlistedIps" yaml:"allowlistedIps"`
	// Maintenance schedule type for the Autonomous Database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.68.0/docs/resources/odb_autonomous_database#autonomous_maintenance_schedule_type OdbAutonomousDatabase#autonomous_maintenance_schedule_type}
	AutonomousMaintenanceScheduleType *string `field:"optional" json:"autonomousMaintenanceScheduleType" yaml:"autonomousMaintenanceScheduleType"`
	// Frequency at which a refreshable clone is automatically refreshed, in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.68.0/docs/resources/odb_autonomous_database#auto_refresh_frequency_in_seconds OdbAutonomousDatabase#auto_refresh_frequency_in_seconds}
	AutoRefreshFrequencyInSeconds *float64 `field:"optional" json:"autoRefreshFrequencyInSeconds" yaml:"autoRefreshFrequencyInSeconds"`
	// Time lag between a refreshable clone and its source, in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.68.0/docs/resources/odb_autonomous_database#auto_refresh_point_lag_in_seconds OdbAutonomousDatabase#auto_refresh_point_lag_in_seconds}
	AutoRefreshPointLagInSeconds *float64 `field:"optional" json:"autoRefreshPointLagInSeconds" yaml:"autoRefreshPointLagInSeconds"`
	// Retention period for automatic backups, in days.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.68.0/docs/resources/odb_autonomous_database#backup_retention_period_in_days OdbAutonomousDatabase#backup_retention_period_in_days}
	BackupRetentionPeriodInDays *float64 `field:"optional" json:"backupRetentionPeriodInDays" yaml:"backupRetentionPeriodInDays"`
	// Maximum compute capacity under the bring-your-own-license model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.68.0/docs/resources/odb_autonomous_database#byol_compute_count_limit OdbAutonomousDatabase#byol_compute_count_limit}
	ByolComputeCountLimit *float64 `field:"optional" json:"byolComputeCountLimit" yaml:"byolComputeCountLimit"`
	// Character set of the Autonomous Database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.68.0/docs/resources/odb_autonomous_database#character_set OdbAutonomousDatabase#character_set}
	CharacterSet *string `field:"optional" json:"characterSet" yaml:"characterSet"`
	// Compute capacity in ECPUs or OCPUs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.68.0/docs/resources/odb_autonomous_database#compute_count OdbAutonomousDatabase#compute_count}
	ComputeCount *float64 `field:"optional" json:"computeCount" yaml:"computeCount"`
	// Number of CPU cores allocated to the Autonomous Database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.68.0/docs/resources/odb_autonomous_database#cpu_core_count OdbAutonomousDatabase#cpu_core_count}
	CpuCoreCount *float64 `field:"optional" json:"cpuCoreCount" yaml:"cpuCoreCount"`
	// customer_contacts_to_send_to_oci block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.68.0/docs/resources/odb_autonomous_database#customer_contacts_to_send_to_oci OdbAutonomousDatabase#customer_contacts_to_send_to_oci}
	CustomerContactsToSendToOci interface{} `field:"optional" json:"customerContactsToSendToOci" yaml:"customerContactsToSendToOci"`
	// Oracle Database edition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.68.0/docs/resources/odb_autonomous_database#database_edition OdbAutonomousDatabase#database_edition}
	DatabaseEdition *string `field:"optional" json:"databaseEdition" yaml:"databaseEdition"`
	// Data volume size in GB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.68.0/docs/resources/odb_autonomous_database#data_storage_size_in_gbs OdbAutonomousDatabase#data_storage_size_in_gbs}
	DataStorageSizeInGbs *float64 `field:"optional" json:"dataStorageSizeInGbs" yaml:"dataStorageSizeInGbs"`
	// Data volume size in TB.
	//
	// Configured values must be whole numbers; computed values may be fractional when storage is configured in GB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.68.0/docs/resources/odb_autonomous_database#data_storage_size_in_tbs OdbAutonomousDatabase#data_storage_size_in_tbs}
	DataStorageSizeInTbs *float64 `field:"optional" json:"dataStorageSizeInTbs" yaml:"dataStorageSizeInTbs"`
	// Name of the Autonomous Database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.68.0/docs/resources/odb_autonomous_database#db_name OdbAutonomousDatabase#db_name}
	DbName *string `field:"optional" json:"dbName" yaml:"dbName"`
	// db_tools_details block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.68.0/docs/resources/odb_autonomous_database#db_tools_details OdbAutonomousDatabase#db_tools_details}
	DbToolsDetails interface{} `field:"optional" json:"dbToolsDetails" yaml:"dbToolsDetails"`
	// Oracle Database software version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.68.0/docs/resources/odb_autonomous_database#db_version OdbAutonomousDatabase#db_version}
	DbVersion *string `field:"optional" json:"dbVersion" yaml:"dbVersion"`
	// Intended database workload.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.68.0/docs/resources/odb_autonomous_database#db_workload OdbAutonomousDatabase#db_workload}
	DbWorkload *string `field:"optional" json:"dbWorkload" yaml:"dbWorkload"`
	// User-friendly name for the Autonomous Database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.68.0/docs/resources/odb_autonomous_database#display_name OdbAutonomousDatabase#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Encryption key provider. Configurable values are ORACLE_MANAGED and AWS_KMS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.68.0/docs/resources/odb_autonomous_database#encryption_key_provider OdbAutonomousDatabase#encryption_key_provider}
	EncryptionKeyProvider *string `field:"optional" json:"encryptionKeyProvider" yaml:"encryptionKeyProvider"`
	// Whether automatic compute scaling is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.68.0/docs/resources/odb_autonomous_database#is_auto_scaling_enabled OdbAutonomousDatabase#is_auto_scaling_enabled}
	IsAutoScalingEnabled interface{} `field:"optional" json:"isAutoScalingEnabled" yaml:"isAutoScalingEnabled"`
	// Whether automatic storage scaling is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.68.0/docs/resources/odb_autonomous_database#is_auto_scaling_for_storage_enabled OdbAutonomousDatabase#is_auto_scaling_for_storage_enabled}
	IsAutoScalingForStorageEnabled interface{} `field:"optional" json:"isAutoScalingForStorageEnabled" yaml:"isAutoScalingForStorageEnabled"`
	// Whether the backup retention period is locked.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.68.0/docs/resources/odb_autonomous_database#is_backup_retention_locked OdbAutonomousDatabase#is_backup_retention_locked}
	IsBackupRetentionLocked interface{} `field:"optional" json:"isBackupRetentionLocked" yaml:"isBackupRetentionLocked"`
	// Whether local Oracle Data Guard is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.68.0/docs/resources/odb_autonomous_database#is_local_data_guard_enabled OdbAutonomousDatabase#is_local_data_guard_enabled}
	IsLocalDataGuardEnabled interface{} `field:"optional" json:"isLocalDataGuardEnabled" yaml:"isLocalDataGuardEnabled"`
	// Whether mutual TLS authentication is required.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.68.0/docs/resources/odb_autonomous_database#is_mtls_connection_required OdbAutonomousDatabase#is_mtls_connection_required}
	IsMtlsConnectionRequired interface{} `field:"optional" json:"isMtlsConnectionRequired" yaml:"isMtlsConnectionRequired"`
	// Whether the Autonomous Database is a refreshable clone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.68.0/docs/resources/odb_autonomous_database#is_refreshable_clone OdbAutonomousDatabase#is_refreshable_clone}
	IsRefreshableClone interface{} `field:"optional" json:"isRefreshableClone" yaml:"isRefreshableClone"`
	// ARN of the AWS KMS key used to encrypt the Autonomous Database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.68.0/docs/resources/odb_autonomous_database#kms_key_id OdbAutonomousDatabase#kms_key_id}
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Oracle license model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.68.0/docs/resources/odb_autonomous_database#license_model OdbAutonomousDatabase#license_model}
	LicenseModel *string `field:"optional" json:"licenseModel" yaml:"licenseModel"`
	// Maximum data-loss limit for automatic local Data Guard failover, in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.68.0/docs/resources/odb_autonomous_database#local_adg_auto_failover_max_data_loss_limit OdbAutonomousDatabase#local_adg_auto_failover_max_data_loss_limit}
	LocalAdgAutoFailoverMaxDataLossLimit *float64 `field:"optional" json:"localAdgAutoFailoverMaxDataLossLimit" yaml:"localAdgAutoFailoverMaxDataLossLimit"`
	// long_term_backup_schedule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.68.0/docs/resources/odb_autonomous_database#long_term_backup_schedule OdbAutonomousDatabase#long_term_backup_schedule}
	LongTermBackupSchedule interface{} `field:"optional" json:"longTermBackupSchedule" yaml:"longTermBackupSchedule"`
	// National character set of the Autonomous Database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.68.0/docs/resources/odb_autonomous_database#ncharacter_set OdbAutonomousDatabase#ncharacter_set}
	NcharacterSet *string `field:"optional" json:"ncharacterSet" yaml:"ncharacterSet"`
	// ID of the associated ODB network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.68.0/docs/resources/odb_autonomous_database#odb_network_id OdbAutonomousDatabase#odb_network_id}
	OdbNetworkId *string `field:"optional" json:"odbNetworkId" yaml:"odbNetworkId"`
	// Open mode of the Autonomous Database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.68.0/docs/resources/odb_autonomous_database#open_mode OdbAutonomousDatabase#open_mode}
	OpenMode *string `field:"optional" json:"openMode" yaml:"openMode"`
	// Permission level of the Autonomous Database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.68.0/docs/resources/odb_autonomous_database#permission_level OdbAutonomousDatabase#permission_level}
	PermissionLevel *string `field:"optional" json:"permissionLevel" yaml:"permissionLevel"`
	// Private endpoint IP address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.68.0/docs/resources/odb_autonomous_database#private_endpoint_ip OdbAutonomousDatabase#private_endpoint_ip}
	PrivateEndpointIp *string `field:"optional" json:"privateEndpointIp" yaml:"privateEndpointIp"`
	// Private endpoint label.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.68.0/docs/resources/odb_autonomous_database#private_endpoint_label OdbAutonomousDatabase#private_endpoint_label}
	PrivateEndpointLabel *string `field:"optional" json:"privateEndpointLabel" yaml:"privateEndpointLabel"`
	// Refresh mode of a refreshable clone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.68.0/docs/resources/odb_autonomous_database#refreshable_mode OdbAutonomousDatabase#refreshable_mode}
	RefreshableMode *string `field:"optional" json:"refreshableMode" yaml:"refreshableMode"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.68.0/docs/resources/odb_autonomous_database#region OdbAutonomousDatabase#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// ID of the resource-pool leader Autonomous Database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.68.0/docs/resources/odb_autonomous_database#resource_pool_leader_id OdbAutonomousDatabase#resource_pool_leader_id}
	ResourcePoolLeaderId *string `field:"optional" json:"resourcePoolLeaderId" yaml:"resourcePoolLeaderId"`
	// resource_pool_summary block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.68.0/docs/resources/odb_autonomous_database#resource_pool_summary OdbAutonomousDatabase#resource_pool_summary}
	ResourcePoolSummary interface{} `field:"optional" json:"resourcePoolSummary" yaml:"resourcePoolSummary"`
	// scheduled_operations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.68.0/docs/resources/odb_autonomous_database#scheduled_operations OdbAutonomousDatabase#scheduled_operations}
	ScheduledOperations interface{} `field:"optional" json:"scheduledOperations" yaml:"scheduledOperations"`
	// Source from which to create the Autonomous Database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.68.0/docs/resources/odb_autonomous_database#source OdbAutonomousDatabase#source}
	Source *string `field:"optional" json:"source" yaml:"source"`
	// source_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.68.0/docs/resources/odb_autonomous_database#source_configuration OdbAutonomousDatabase#source_configuration}
	SourceConfiguration interface{} `field:"optional" json:"sourceConfiguration" yaml:"sourceConfiguration"`
	// IP addresses allowed to access the standby Autonomous Database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.68.0/docs/resources/odb_autonomous_database#standby_allowlisted_ips OdbAutonomousDatabase#standby_allowlisted_ips}
	StandbyAllowlistedIps *[]*string `field:"optional" json:"standbyAllowlistedIps" yaml:"standbyAllowlistedIps"`
	// Source of the standby allowlisted IP addresses.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.68.0/docs/resources/odb_autonomous_database#standby_allowlisted_ips_source OdbAutonomousDatabase#standby_allowlisted_ips_source}
	StandbyAllowlistedIpsSource *string `field:"optional" json:"standbyAllowlistedIpsSource" yaml:"standbyAllowlistedIpsSource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.68.0/docs/resources/odb_autonomous_database#tags OdbAutonomousDatabase#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Date and time when automatic refresh begins.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.68.0/docs/resources/odb_autonomous_database#time_of_auto_refresh_start OdbAutonomousDatabase#time_of_auto_refresh_start}
	TimeOfAutoRefreshStart *string `field:"optional" json:"timeOfAutoRefreshStart" yaml:"timeOfAutoRefreshStart"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.68.0/docs/resources/odb_autonomous_database#timeouts OdbAutonomousDatabase#timeouts}
	Timeouts *OdbAutonomousDatabaseTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// transportable_tablespace block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.68.0/docs/resources/odb_autonomous_database#transportable_tablespace OdbAutonomousDatabase#transportable_tablespace}
	TransportableTablespace interface{} `field:"optional" json:"transportableTablespace" yaml:"transportableTablespace"`
}

