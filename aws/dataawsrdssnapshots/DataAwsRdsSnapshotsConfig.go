// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawsrdssnapshots

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAwsRdsSnapshotsConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/rds_snapshots#db_instance_identifier DataAwsRdsSnapshots#db_instance_identifier}.
	DbInstanceIdentifier *string `field:"optional" json:"dbInstanceIdentifier" yaml:"dbInstanceIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/rds_snapshots#db_snapshot_identifier DataAwsRdsSnapshots#db_snapshot_identifier}.
	DbSnapshotIdentifier *string `field:"optional" json:"dbSnapshotIdentifier" yaml:"dbSnapshotIdentifier"`
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/rds_snapshots#filter DataAwsRdsSnapshots#filter}
	Filter interface{} `field:"optional" json:"filter" yaml:"filter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/rds_snapshots#include_public DataAwsRdsSnapshots#include_public}.
	IncludePublic interface{} `field:"optional" json:"includePublic" yaml:"includePublic"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/rds_snapshots#include_shared DataAwsRdsSnapshots#include_shared}.
	IncludeShared interface{} `field:"optional" json:"includeShared" yaml:"includeShared"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/rds_snapshots#region DataAwsRdsSnapshots#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/rds_snapshots#snapshot_type DataAwsRdsSnapshots#snapshot_type}.
	SnapshotType *string `field:"optional" json:"snapshotType" yaml:"snapshotType"`
}

