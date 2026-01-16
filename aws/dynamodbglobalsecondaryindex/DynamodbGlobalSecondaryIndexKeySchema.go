// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dynamodbglobalsecondaryindex


type DynamodbGlobalSecondaryIndexKeySchema struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/dynamodb_global_secondary_index#attribute_name DynamodbGlobalSecondaryIndex#attribute_name}.
	AttributeName *string `field:"required" json:"attributeName" yaml:"attributeName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/dynamodb_global_secondary_index#attribute_type DynamodbGlobalSecondaryIndex#attribute_type}.
	AttributeType *string `field:"required" json:"attributeType" yaml:"attributeType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/dynamodb_global_secondary_index#key_type DynamodbGlobalSecondaryIndex#key_type}.
	KeyType *string `field:"required" json:"keyType" yaml:"keyType"`
}

