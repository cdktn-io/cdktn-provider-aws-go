// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package redshiftidcapplication


type RedshiftIdcApplicationAuthorizedTokenIssuer struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/redshift_idc_application#authorized_audiences_list RedshiftIdcApplication#authorized_audiences_list}.
	AuthorizedAudiencesList *[]*string `field:"optional" json:"authorizedAudiencesList" yaml:"authorizedAudiencesList"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/redshift_idc_application#trusted_token_issuer_arn RedshiftIdcApplication#trusted_token_issuer_arn}.
	TrustedTokenIssuerArn *string `field:"optional" json:"trustedTokenIssuerArn" yaml:"trustedTokenIssuerArn"`
}

