// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package odbautonomousdatabase


type OdbAutonomousDatabaseResourcePoolSummary struct {
	// Whether the resource pool is disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.68.0/docs/resources/odb_autonomous_database#is_disabled OdbAutonomousDatabase#is_disabled}
	IsDisabled interface{} `field:"optional" json:"isDisabled" yaml:"isDisabled"`
	// Number of Autonomous Databases the resource pool can contain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.68.0/docs/resources/odb_autonomous_database#pool_size OdbAutonomousDatabase#pool_size}
	PoolSize *float64 `field:"optional" json:"poolSize" yaml:"poolSize"`
	// Total storage size of the resource pool, in TB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.68.0/docs/resources/odb_autonomous_database#pool_storage_size_in_tbs OdbAutonomousDatabase#pool_storage_size_in_tbs}
	PoolStorageSizeInTbs *float64 `field:"optional" json:"poolStorageSizeInTbs" yaml:"poolStorageSizeInTbs"`
}

