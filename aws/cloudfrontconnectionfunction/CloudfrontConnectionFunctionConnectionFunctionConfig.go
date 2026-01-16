// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudfrontconnectionfunction


type CloudfrontConnectionFunctionConnectionFunctionConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_connection_function#comment CloudfrontConnectionFunction#comment}.
	Comment *string `field:"required" json:"comment" yaml:"comment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_connection_function#runtime CloudfrontConnectionFunction#runtime}.
	Runtime *string `field:"required" json:"runtime" yaml:"runtime"`
	// key_value_store_association block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_connection_function#key_value_store_association CloudfrontConnectionFunction#key_value_store_association}
	KeyValueStoreAssociation interface{} `field:"optional" json:"keyValueStoreAssociation" yaml:"keyValueStoreAssociation"`
}

