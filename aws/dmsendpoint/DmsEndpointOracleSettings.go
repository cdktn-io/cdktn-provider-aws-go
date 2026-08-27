// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dmsendpoint


type DmsEndpointOracleSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#access_alternate_directly DmsEndpoint#access_alternate_directly}.
	AccessAlternateDirectly interface{} `field:"optional" json:"accessAlternateDirectly" yaml:"accessAlternateDirectly"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#additional_archived_log_dest_id DmsEndpoint#additional_archived_log_dest_id}.
	AdditionalArchivedLogDestId *float64 `field:"optional" json:"additionalArchivedLogDestId" yaml:"additionalArchivedLogDestId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#add_supplemental_logging DmsEndpoint#add_supplemental_logging}.
	AddSupplementalLogging interface{} `field:"optional" json:"addSupplementalLogging" yaml:"addSupplementalLogging"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#allow_selected_nested_tables DmsEndpoint#allow_selected_nested_tables}.
	AllowSelectedNestedTables interface{} `field:"optional" json:"allowSelectedNestedTables" yaml:"allowSelectedNestedTables"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#archived_log_dest_id DmsEndpoint#archived_log_dest_id}.
	ArchivedLogDestId *float64 `field:"optional" json:"archivedLogDestId" yaml:"archivedLogDestId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#archived_logs_only DmsEndpoint#archived_logs_only}.
	ArchivedLogsOnly interface{} `field:"optional" json:"archivedLogsOnly" yaml:"archivedLogsOnly"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#asm_password DmsEndpoint#asm_password}.
	AsmPassword *string `field:"optional" json:"asmPassword" yaml:"asmPassword"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#asm_server DmsEndpoint#asm_server}.
	AsmServer *string `field:"optional" json:"asmServer" yaml:"asmServer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#asm_user DmsEndpoint#asm_user}.
	AsmUser *string `field:"optional" json:"asmUser" yaml:"asmUser"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#authentication_method DmsEndpoint#authentication_method}.
	AuthenticationMethod *string `field:"optional" json:"authenticationMethod" yaml:"authenticationMethod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#char_length_semantics DmsEndpoint#char_length_semantics}.
	CharLengthSemantics *string `field:"optional" json:"charLengthSemantics" yaml:"charLengthSemantics"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#convert_timestamp_with_zone_to_utc DmsEndpoint#convert_timestamp_with_zone_to_utc}.
	ConvertTimestampWithZoneToUtc interface{} `field:"optional" json:"convertTimestampWithZoneToUtc" yaml:"convertTimestampWithZoneToUtc"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#direct_path_no_log DmsEndpoint#direct_path_no_log}.
	DirectPathNoLog interface{} `field:"optional" json:"directPathNoLog" yaml:"directPathNoLog"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#direct_path_parallel_load DmsEndpoint#direct_path_parallel_load}.
	DirectPathParallelLoad interface{} `field:"optional" json:"directPathParallelLoad" yaml:"directPathParallelLoad"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#enable_homogenous_tablespace DmsEndpoint#enable_homogenous_tablespace}.
	EnableHomogenousTablespace interface{} `field:"optional" json:"enableHomogenousTablespace" yaml:"enableHomogenousTablespace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#extra_archived_log_dest_ids DmsEndpoint#extra_archived_log_dest_ids}.
	ExtraArchivedLogDestIds *[]*float64 `field:"optional" json:"extraArchivedLogDestIds" yaml:"extraArchivedLogDestIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#fail_task_on_lob_truncation DmsEndpoint#fail_task_on_lob_truncation}.
	FailTaskOnLobTruncation interface{} `field:"optional" json:"failTaskOnLobTruncation" yaml:"failTaskOnLobTruncation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#number_datatype_scale DmsEndpoint#number_datatype_scale}.
	NumberDatatypeScale *float64 `field:"optional" json:"numberDatatypeScale" yaml:"numberDatatypeScale"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#open_transaction_window DmsEndpoint#open_transaction_window}.
	OpenTransactionWindow *float64 `field:"optional" json:"openTransactionWindow" yaml:"openTransactionWindow"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#oracle_path_prefix DmsEndpoint#oracle_path_prefix}.
	OraclePathPrefix *string `field:"optional" json:"oraclePathPrefix" yaml:"oraclePathPrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#parallel_asm_read_threads DmsEndpoint#parallel_asm_read_threads}.
	ParallelAsmReadThreads *float64 `field:"optional" json:"parallelAsmReadThreads" yaml:"parallelAsmReadThreads"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#read_ahead_blocks DmsEndpoint#read_ahead_blocks}.
	ReadAheadBlocks *float64 `field:"optional" json:"readAheadBlocks" yaml:"readAheadBlocks"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#read_table_space_name DmsEndpoint#read_table_space_name}.
	ReadTableSpaceName interface{} `field:"optional" json:"readTableSpaceName" yaml:"readTableSpaceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#replace_path_prefix DmsEndpoint#replace_path_prefix}.
	ReplacePathPrefix interface{} `field:"optional" json:"replacePathPrefix" yaml:"replacePathPrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#retry_interval DmsEndpoint#retry_interval}.
	RetryInterval *float64 `field:"optional" json:"retryInterval" yaml:"retryInterval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#secrets_manager_oracle_asm_access_role_arn DmsEndpoint#secrets_manager_oracle_asm_access_role_arn}.
	SecretsManagerOracleAsmAccessRoleArn *string `field:"optional" json:"secretsManagerOracleAsmAccessRoleArn" yaml:"secretsManagerOracleAsmAccessRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#secrets_manager_oracle_asm_secret_id DmsEndpoint#secrets_manager_oracle_asm_secret_id}.
	SecretsManagerOracleAsmSecretId *string `field:"optional" json:"secretsManagerOracleAsmSecretId" yaml:"secretsManagerOracleAsmSecretId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#security_db_encryption DmsEndpoint#security_db_encryption}.
	SecurityDbEncryption *string `field:"optional" json:"securityDbEncryption" yaml:"securityDbEncryption"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#security_db_encryption_name DmsEndpoint#security_db_encryption_name}.
	SecurityDbEncryptionName *string `field:"optional" json:"securityDbEncryptionName" yaml:"securityDbEncryptionName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#spatial_data_option_to_geo_json_function_name DmsEndpoint#spatial_data_option_to_geo_json_function_name}.
	SpatialDataOptionToGeoJsonFunctionName *string `field:"optional" json:"spatialDataOptionToGeoJsonFunctionName" yaml:"spatialDataOptionToGeoJsonFunctionName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#standby_delay_time DmsEndpoint#standby_delay_time}.
	StandbyDelayTime *float64 `field:"optional" json:"standbyDelayTime" yaml:"standbyDelayTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#trim_space_in_char DmsEndpoint#trim_space_in_char}.
	TrimSpaceInChar interface{} `field:"optional" json:"trimSpaceInChar" yaml:"trimSpaceInChar"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#use_alternate_folder_for_online DmsEndpoint#use_alternate_folder_for_online}.
	UseAlternateFolderForOnline interface{} `field:"optional" json:"useAlternateFolderForOnline" yaml:"useAlternateFolderForOnline"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#use_bfile DmsEndpoint#use_bfile}.
	UseBfile interface{} `field:"optional" json:"useBfile" yaml:"useBfile"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#use_direct_path_full_load DmsEndpoint#use_direct_path_full_load}.
	UseDirectPathFullLoad interface{} `field:"optional" json:"useDirectPathFullLoad" yaml:"useDirectPathFullLoad"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#use_logminer_reader DmsEndpoint#use_logminer_reader}.
	UseLogminerReader interface{} `field:"optional" json:"useLogminerReader" yaml:"useLogminerReader"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#use_path_prefix DmsEndpoint#use_path_prefix}.
	UsePathPrefix *string `field:"optional" json:"usePathPrefix" yaml:"usePathPrefix"`
}

