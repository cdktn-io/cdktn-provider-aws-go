// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package acmcertificate


type AcmCertificateOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.32.1/docs/resources/acm_certificate#certificate_transparency_logging_preference AcmCertificate#certificate_transparency_logging_preference}.
	CertificateTransparencyLoggingPreference *string `field:"optional" json:"certificateTransparencyLoggingPreference" yaml:"certificateTransparencyLoggingPreference"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.32.1/docs/resources/acm_certificate#export AcmCertificate#export}.
	Export *string `field:"optional" json:"export" yaml:"export"`
}

