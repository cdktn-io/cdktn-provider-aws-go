// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudfrontdistribution


type CloudfrontDistributionViewerMtlsConfigTrustStoreConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_distribution#trust_store_id CloudfrontDistribution#trust_store_id}.
	TrustStoreId *string `field:"required" json:"trustStoreId" yaml:"trustStoreId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_distribution#advertise_trust_store_ca_names CloudfrontDistribution#advertise_trust_store_ca_names}.
	AdvertiseTrustStoreCaNames interface{} `field:"optional" json:"advertiseTrustStoreCaNames" yaml:"advertiseTrustStoreCaNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_distribution#ignore_certificate_expiry CloudfrontDistribution#ignore_certificate_expiry}.
	IgnoreCertificateExpiry interface{} `field:"optional" json:"ignoreCertificateExpiry" yaml:"ignoreCertificateExpiry"`
}

