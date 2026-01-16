// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dynamodbglobalsecondaryindex


type DynamodbGlobalSecondaryIndexOnDemandThroughput struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/dynamodb_global_secondary_index#max_read_request_units DynamodbGlobalSecondaryIndex#max_read_request_units}.
	MaxReadRequestUnits *float64 `field:"optional" json:"maxReadRequestUnits" yaml:"maxReadRequestUnits"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/dynamodb_global_secondary_index#max_write_request_units DynamodbGlobalSecondaryIndex#max_write_request_units}.
	MaxWriteRequestUnits *float64 `field:"optional" json:"maxWriteRequestUnits" yaml:"maxWriteRequestUnits"`
}

