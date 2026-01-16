// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudfrontdistribution


type CloudfrontDistributionViewerMtlsConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_distribution#mode CloudfrontDistribution#mode}.
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// trust_store_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_distribution#trust_store_config CloudfrontDistribution#trust_store_config}
	TrustStoreConfig *CloudfrontDistributionViewerMtlsConfigTrustStoreConfig `field:"optional" json:"trustStoreConfig" yaml:"trustStoreConfig"`
}

