// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudfronttruststore

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudfrontTrustStoreConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_trust_store#name CloudfrontTrustStore#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// ca_certificates_bundle_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_trust_store#ca_certificates_bundle_source CloudfrontTrustStore#ca_certificates_bundle_source}
	CaCertificatesBundleSource interface{} `field:"optional" json:"caCertificatesBundleSource" yaml:"caCertificatesBundleSource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_trust_store#tags CloudfrontTrustStore#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_trust_store#timeouts CloudfrontTrustStore#timeouts}
	Timeouts *CloudfrontTrustStoreTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

