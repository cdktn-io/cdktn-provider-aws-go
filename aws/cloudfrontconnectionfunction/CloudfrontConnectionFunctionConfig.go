// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudfrontconnectionfunction

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudfrontConnectionFunctionConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_connection_function#connection_function_code CloudfrontConnectionFunction#connection_function_code}.
	ConnectionFunctionCode *string `field:"required" json:"connectionFunctionCode" yaml:"connectionFunctionCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_connection_function#name CloudfrontConnectionFunction#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// connection_function_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_connection_function#connection_function_config CloudfrontConnectionFunction#connection_function_config}
	ConnectionFunctionConfig interface{} `field:"optional" json:"connectionFunctionConfig" yaml:"connectionFunctionConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_connection_function#publish CloudfrontConnectionFunction#publish}.
	Publish interface{} `field:"optional" json:"publish" yaml:"publish"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_connection_function#tags CloudfrontConnectionFunction#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

