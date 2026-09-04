// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datazonepolicygrant


type DatazonePolicyGrantPrincipalProject struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/datazone_policy_grant#project_designation DatazonePolicyGrant#project_designation}.
	ProjectDesignation *string `field:"required" json:"projectDesignation" yaml:"projectDesignation"`
	// domain_unit_filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/datazone_policy_grant#domain_unit_filter DatazonePolicyGrant#domain_unit_filter}
	DomainUnitFilter interface{} `field:"optional" json:"domainUnitFilter" yaml:"domainUnitFilter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/datazone_policy_grant#project_identifier DatazonePolicyGrant#project_identifier}.
	ProjectIdentifier *string `field:"optional" json:"projectIdentifier" yaml:"projectIdentifier"`
}

