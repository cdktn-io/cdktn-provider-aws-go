// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appmeshvirtualgateway


type AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrustAcm struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#certificate_authority_arns AppmeshVirtualGateway#certificate_authority_arns}.
	CertificateAuthorityArns *[]*string `field:"required" json:"certificateAuthorityArns" yaml:"certificateAuthorityArns"`
}

