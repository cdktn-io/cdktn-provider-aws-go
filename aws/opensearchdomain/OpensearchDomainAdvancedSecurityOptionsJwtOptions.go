// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package opensearchdomain


type OpensearchDomainAdvancedSecurityOptionsJwtOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain#enabled OpensearchDomain#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain#jwks_url OpensearchDomain#jwks_url}.
	JwksUrl *string `field:"optional" json:"jwksUrl" yaml:"jwksUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain#public_key OpensearchDomain#public_key}.
	PublicKey *string `field:"optional" json:"publicKey" yaml:"publicKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain#roles_key OpensearchDomain#roles_key}.
	RolesKey *string `field:"optional" json:"rolesKey" yaml:"rolesKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain#subject_key OpensearchDomain#subject_key}.
	SubjectKey *string `field:"optional" json:"subjectKey" yaml:"subjectKey"`
}

