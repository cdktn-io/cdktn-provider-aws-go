// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudfrontmultitenantdistribution

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudfrontMultitenantDistributionConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_multitenant_distribution#comment CloudfrontMultitenantDistribution#comment}.
	Comment *string `field:"required" json:"comment" yaml:"comment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_multitenant_distribution#enabled CloudfrontMultitenantDistribution#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// active_trusted_key_groups block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_multitenant_distribution#active_trusted_key_groups CloudfrontMultitenantDistribution#active_trusted_key_groups}
	ActiveTrustedKeyGroups interface{} `field:"optional" json:"activeTrustedKeyGroups" yaml:"activeTrustedKeyGroups"`
	// cache_behavior block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_multitenant_distribution#cache_behavior CloudfrontMultitenantDistribution#cache_behavior}
	CacheBehavior interface{} `field:"optional" json:"cacheBehavior" yaml:"cacheBehavior"`
	// custom_error_response block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_multitenant_distribution#custom_error_response CloudfrontMultitenantDistribution#custom_error_response}
	CustomErrorResponse interface{} `field:"optional" json:"customErrorResponse" yaml:"customErrorResponse"`
	// default_cache_behavior block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_multitenant_distribution#default_cache_behavior CloudfrontMultitenantDistribution#default_cache_behavior}
	DefaultCacheBehavior interface{} `field:"optional" json:"defaultCacheBehavior" yaml:"defaultCacheBehavior"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_multitenant_distribution#default_root_object CloudfrontMultitenantDistribution#default_root_object}.
	DefaultRootObject *string `field:"optional" json:"defaultRootObject" yaml:"defaultRootObject"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_multitenant_distribution#http_version CloudfrontMultitenantDistribution#http_version}.
	HttpVersion *string `field:"optional" json:"httpVersion" yaml:"httpVersion"`
	// origin block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_multitenant_distribution#origin CloudfrontMultitenantDistribution#origin}
	Origin interface{} `field:"optional" json:"origin" yaml:"origin"`
	// origin_group block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_multitenant_distribution#origin_group CloudfrontMultitenantDistribution#origin_group}
	OriginGroup interface{} `field:"optional" json:"originGroup" yaml:"originGroup"`
	// restrictions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_multitenant_distribution#restrictions CloudfrontMultitenantDistribution#restrictions}
	Restrictions interface{} `field:"optional" json:"restrictions" yaml:"restrictions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_multitenant_distribution#tags CloudfrontMultitenantDistribution#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// tenant_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_multitenant_distribution#tenant_config CloudfrontMultitenantDistribution#tenant_config}
	TenantConfig interface{} `field:"optional" json:"tenantConfig" yaml:"tenantConfig"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_multitenant_distribution#timeouts CloudfrontMultitenantDistribution#timeouts}
	Timeouts *CloudfrontMultitenantDistributionTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// viewer_certificate block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_multitenant_distribution#viewer_certificate CloudfrontMultitenantDistribution#viewer_certificate}
	ViewerCertificate interface{} `field:"optional" json:"viewerCertificate" yaml:"viewerCertificate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_multitenant_distribution#web_acl_id CloudfrontMultitenantDistribution#web_acl_id}.
	WebAclId *string `field:"optional" json:"webAclId" yaml:"webAclId"`
}

