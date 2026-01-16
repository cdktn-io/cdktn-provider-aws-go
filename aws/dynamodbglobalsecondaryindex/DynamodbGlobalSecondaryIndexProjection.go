// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dynamodbglobalsecondaryindex


type DynamodbGlobalSecondaryIndexProjection struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/dynamodb_global_secondary_index#projection_type DynamodbGlobalSecondaryIndex#projection_type}.
	ProjectionType *string `field:"required" json:"projectionType" yaml:"projectionType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/dynamodb_global_secondary_index#non_key_attributes DynamodbGlobalSecondaryIndex#non_key_attributes}.
	NonKeyAttributes *[]*string `field:"optional" json:"nonKeyAttributes" yaml:"nonKeyAttributes"`
}

