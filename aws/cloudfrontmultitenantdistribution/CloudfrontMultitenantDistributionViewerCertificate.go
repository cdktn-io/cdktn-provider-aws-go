// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudfrontmultitenantdistribution


type CloudfrontMultitenantDistributionViewerCertificate struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_multitenant_distribution#acm_certificate_arn CloudfrontMultitenantDistribution#acm_certificate_arn}.
	AcmCertificateArn *string `field:"optional" json:"acmCertificateArn" yaml:"acmCertificateArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_multitenant_distribution#cloudfront_default_certificate CloudfrontMultitenantDistribution#cloudfront_default_certificate}.
	CloudfrontDefaultCertificate interface{} `field:"optional" json:"cloudfrontDefaultCertificate" yaml:"cloudfrontDefaultCertificate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_multitenant_distribution#minimum_protocol_version CloudfrontMultitenantDistribution#minimum_protocol_version}.
	MinimumProtocolVersion *string `field:"optional" json:"minimumProtocolVersion" yaml:"minimumProtocolVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_multitenant_distribution#ssl_support_method CloudfrontMultitenantDistribution#ssl_support_method}.
	SslSupportMethod *string `field:"optional" json:"sslSupportMethod" yaml:"sslSupportMethod"`
}

