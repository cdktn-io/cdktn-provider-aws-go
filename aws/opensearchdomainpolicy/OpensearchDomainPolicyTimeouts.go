// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package opensearchdomainpolicy


type OpensearchDomainPolicyTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/opensearch_domain_policy#delete OpensearchDomainPolicy#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/opensearch_domain_policy#update OpensearchDomainPolicy#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

