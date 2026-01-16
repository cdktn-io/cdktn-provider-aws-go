// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudfrontdistributiontenant


type CloudfrontDistributionTenantManagedCertificateRequest struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_distribution_tenant#certificate_transparency_logging_preference CloudfrontDistributionTenant#certificate_transparency_logging_preference}.
	CertificateTransparencyLoggingPreference *string `field:"optional" json:"certificateTransparencyLoggingPreference" yaml:"certificateTransparencyLoggingPreference"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_distribution_tenant#primary_domain_name CloudfrontDistributionTenant#primary_domain_name}.
	PrimaryDomainName *string `field:"optional" json:"primaryDomainName" yaml:"primaryDomainName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_distribution_tenant#validation_token_host CloudfrontDistributionTenant#validation_token_host}.
	ValidationTokenHost *string `field:"optional" json:"validationTokenHost" yaml:"validationTokenHost"`
}

