// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudfrontdistribution


type CloudfrontDistributionOriginCustomOriginConfigOriginMtlsConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#client_certificate_arn CloudfrontDistribution#client_certificate_arn}.
	ClientCertificateArn *string `field:"required" json:"clientCertificateArn" yaml:"clientCertificateArn"`
}

