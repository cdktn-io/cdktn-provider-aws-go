// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudfrontdistributiontenant


type CloudfrontDistributionTenantParameter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_distribution_tenant#name CloudfrontDistributionTenant#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_distribution_tenant#value CloudfrontDistributionTenant#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

