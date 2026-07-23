// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mskcluster


type MskClusterClientAuthenticationTls struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.56.0/docs/resources/msk_cluster#certificate_authority_arns MskCluster#certificate_authority_arns}.
	CertificateAuthorityArns *[]*string `field:"optional" json:"certificateAuthorityArns" yaml:"certificateAuthorityArns"`
}

