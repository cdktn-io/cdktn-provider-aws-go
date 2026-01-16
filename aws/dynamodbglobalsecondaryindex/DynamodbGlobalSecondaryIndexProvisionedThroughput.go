// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dynamodbglobalsecondaryindex


type DynamodbGlobalSecondaryIndexProvisionedThroughput struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/dynamodb_global_secondary_index#read_capacity_units DynamodbGlobalSecondaryIndex#read_capacity_units}.
	ReadCapacityUnits *float64 `field:"optional" json:"readCapacityUnits" yaml:"readCapacityUnits"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/dynamodb_global_secondary_index#write_capacity_units DynamodbGlobalSecondaryIndex#write_capacity_units}.
	WriteCapacityUnits *float64 `field:"optional" json:"writeCapacityUnits" yaml:"writeCapacityUnits"`
}

