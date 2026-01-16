// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudfronttruststore


type CloudfrontTrustStoreCaCertificatesBundleSource struct {
	// ca_certificates_bundle_s3_location block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_trust_store#ca_certificates_bundle_s3_location CloudfrontTrustStore#ca_certificates_bundle_s3_location}
	CaCertificatesBundleS3Location interface{} `field:"optional" json:"caCertificatesBundleS3Location" yaml:"caCertificatesBundleS3Location"`
}

