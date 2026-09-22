// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dmsdataprovider


type DmsDataProviderSettings struct {
	// doc_db_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.66.0/docs/resources/dms_data_provider#doc_db_settings DmsDataProvider#doc_db_settings}
	DocDbSettings interface{} `field:"optional" json:"docDbSettings" yaml:"docDbSettings"`
	// ibm_db2_luw_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.66.0/docs/resources/dms_data_provider#ibm_db2_luw_settings DmsDataProvider#ibm_db2_luw_settings}
	IbmDb2LuwSettings interface{} `field:"optional" json:"ibmDb2LuwSettings" yaml:"ibmDb2LuwSettings"`
	// ibm_db2_zos_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.66.0/docs/resources/dms_data_provider#ibm_db2_zos_settings DmsDataProvider#ibm_db2_zos_settings}
	IbmDb2ZosSettings interface{} `field:"optional" json:"ibmDb2ZosSettings" yaml:"ibmDb2ZosSettings"`
	// maria_db_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.66.0/docs/resources/dms_data_provider#maria_db_settings DmsDataProvider#maria_db_settings}
	MariaDbSettings interface{} `field:"optional" json:"mariaDbSettings" yaml:"mariaDbSettings"`
	// microsoft_sql_server_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.66.0/docs/resources/dms_data_provider#microsoft_sql_server_settings DmsDataProvider#microsoft_sql_server_settings}
	MicrosoftSqlServerSettings interface{} `field:"optional" json:"microsoftSqlServerSettings" yaml:"microsoftSqlServerSettings"`
	// mongo_db_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.66.0/docs/resources/dms_data_provider#mongo_db_settings DmsDataProvider#mongo_db_settings}
	MongoDbSettings interface{} `field:"optional" json:"mongoDbSettings" yaml:"mongoDbSettings"`
	// mysql_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.66.0/docs/resources/dms_data_provider#mysql_settings DmsDataProvider#mysql_settings}
	MysqlSettings interface{} `field:"optional" json:"mysqlSettings" yaml:"mysqlSettings"`
	// oracle_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.66.0/docs/resources/dms_data_provider#oracle_settings DmsDataProvider#oracle_settings}
	OracleSettings interface{} `field:"optional" json:"oracleSettings" yaml:"oracleSettings"`
	// postgresql_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.66.0/docs/resources/dms_data_provider#postgresql_settings DmsDataProvider#postgresql_settings}
	PostgresqlSettings interface{} `field:"optional" json:"postgresqlSettings" yaml:"postgresqlSettings"`
	// redshift_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.66.0/docs/resources/dms_data_provider#redshift_settings DmsDataProvider#redshift_settings}
	RedshiftSettings interface{} `field:"optional" json:"redshiftSettings" yaml:"redshiftSettings"`
	// sybase_ase_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.66.0/docs/resources/dms_data_provider#sybase_ase_settings DmsDataProvider#sybase_ase_settings}
	SybaseAseSettings interface{} `field:"optional" json:"sybaseAseSettings" yaml:"sybaseAseSettings"`
}

