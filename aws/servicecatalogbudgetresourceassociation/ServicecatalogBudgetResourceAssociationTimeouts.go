// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package servicecatalogbudgetresourceassociation


type ServicecatalogBudgetResourceAssociationTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.31.0/docs/resources/servicecatalog_budget_resource_association#create ServicecatalogBudgetResourceAssociation#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.31.0/docs/resources/servicecatalog_budget_resource_association#delete ServicecatalogBudgetResourceAssociation#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.31.0/docs/resources/servicecatalog_budget_resource_association#read ServicecatalogBudgetResourceAssociation#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

