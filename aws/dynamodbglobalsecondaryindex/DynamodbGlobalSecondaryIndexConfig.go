// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dynamodbglobalsecondaryindex

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DynamodbGlobalSecondaryIndexConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/dynamodb_global_secondary_index#index_name DynamodbGlobalSecondaryIndex#index_name}.
	IndexName *string `field:"required" json:"indexName" yaml:"indexName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/dynamodb_global_secondary_index#table_name DynamodbGlobalSecondaryIndex#table_name}.
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
	// key_schema block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/dynamodb_global_secondary_index#key_schema DynamodbGlobalSecondaryIndex#key_schema}
	KeySchema interface{} `field:"optional" json:"keySchema" yaml:"keySchema"`
	// on_demand_throughput block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/dynamodb_global_secondary_index#on_demand_throughput DynamodbGlobalSecondaryIndex#on_demand_throughput}
	OnDemandThroughput interface{} `field:"optional" json:"onDemandThroughput" yaml:"onDemandThroughput"`
	// projection block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/dynamodb_global_secondary_index#projection DynamodbGlobalSecondaryIndex#projection}
	Projection interface{} `field:"optional" json:"projection" yaml:"projection"`
	// provisioned_throughput block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/dynamodb_global_secondary_index#provisioned_throughput DynamodbGlobalSecondaryIndex#provisioned_throughput}
	ProvisionedThroughput interface{} `field:"optional" json:"provisionedThroughput" yaml:"provisionedThroughput"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/dynamodb_global_secondary_index#region DynamodbGlobalSecondaryIndex#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/dynamodb_global_secondary_index#timeouts DynamodbGlobalSecondaryIndex#timeouts}
	Timeouts *DynamodbGlobalSecondaryIndexTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/dynamodb_global_secondary_index#warm_throughput DynamodbGlobalSecondaryIndex#warm_throughput}.
	WarmThroughput *DynamodbGlobalSecondaryIndexWarmThroughput `field:"optional" json:"warmThroughput" yaml:"warmThroughput"`
}

