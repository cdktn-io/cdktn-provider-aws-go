// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dynamodbglobalsecondaryindex


type DynamodbGlobalSecondaryIndexWarmThroughput struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/dynamodb_global_secondary_index#read_units_per_second DynamodbGlobalSecondaryIndex#read_units_per_second}.
	ReadUnitsPerSecond *float64 `field:"optional" json:"readUnitsPerSecond" yaml:"readUnitsPerSecond"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/dynamodb_global_secondary_index#write_units_per_second DynamodbGlobalSecondaryIndex#write_units_per_second}.
	WriteUnitsPerSecond *float64 `field:"optional" json:"writeUnitsPerSecond" yaml:"writeUnitsPerSecond"`
}

