// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudfrontdistributiontenant

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudfrontDistributionTenantConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_distribution_tenant#distribution_id CloudfrontDistributionTenant#distribution_id}.
	DistributionId *string `field:"required" json:"distributionId" yaml:"distributionId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_distribution_tenant#name CloudfrontDistributionTenant#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_distribution_tenant#connection_group_id CloudfrontDistributionTenant#connection_group_id}.
	ConnectionGroupId *string `field:"optional" json:"connectionGroupId" yaml:"connectionGroupId"`
	// customizations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_distribution_tenant#customizations CloudfrontDistributionTenant#customizations}
	Customizations interface{} `field:"optional" json:"customizations" yaml:"customizations"`
	// domain block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_distribution_tenant#domain CloudfrontDistributionTenant#domain}
	Domain interface{} `field:"optional" json:"domain" yaml:"domain"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_distribution_tenant#enabled CloudfrontDistributionTenant#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// managed_certificate_request block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_distribution_tenant#managed_certificate_request CloudfrontDistributionTenant#managed_certificate_request}
	ManagedCertificateRequest interface{} `field:"optional" json:"managedCertificateRequest" yaml:"managedCertificateRequest"`
	// parameter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_distribution_tenant#parameter CloudfrontDistributionTenant#parameter}
	Parameter interface{} `field:"optional" json:"parameter" yaml:"parameter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_distribution_tenant#tags CloudfrontDistributionTenant#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_distribution_tenant#timeouts CloudfrontDistributionTenant#timeouts}
	Timeouts *CloudfrontDistributionTenantTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_distribution_tenant#wait_for_deployment CloudfrontDistributionTenant#wait_for_deployment}.
	WaitForDeployment interface{} `field:"optional" json:"waitForDeployment" yaml:"waitForDeployment"`
}

