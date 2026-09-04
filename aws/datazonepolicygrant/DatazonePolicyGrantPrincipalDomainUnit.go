// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datazonepolicygrant


type DatazonePolicyGrantPrincipalDomainUnit struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/datazone_policy_grant#domain_unit_designation DatazonePolicyGrant#domain_unit_designation}.
	DomainUnitDesignation *string `field:"required" json:"domainUnitDesignation" yaml:"domainUnitDesignation"`
	// all_domain_units_grant_filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/datazone_policy_grant#all_domain_units_grant_filter DatazonePolicyGrant#all_domain_units_grant_filter}
	AllDomainUnitsGrantFilter interface{} `field:"optional" json:"allDomainUnitsGrantFilter" yaml:"allDomainUnitsGrantFilter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/datazone_policy_grant#domain_unit_identifier DatazonePolicyGrant#domain_unit_identifier}.
	DomainUnitIdentifier *string `field:"optional" json:"domainUnitIdentifier" yaml:"domainUnitIdentifier"`
}

