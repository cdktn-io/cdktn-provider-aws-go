// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cecostcategory


type CeCostCategoryRuleRuleAndOr struct {
	// cost_category block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.31.0/docs/resources/ce_cost_category#cost_category CeCostCategory#cost_category}
	CostCategory *CeCostCategoryRuleRuleAndOrCostCategory `field:"optional" json:"costCategory" yaml:"costCategory"`
	// dimension block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.31.0/docs/resources/ce_cost_category#dimension CeCostCategory#dimension}
	Dimension *CeCostCategoryRuleRuleAndOrDimension `field:"optional" json:"dimension" yaml:"dimension"`
	// tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.31.0/docs/resources/ce_cost_category#tags CeCostCategory#tags}
	Tags *CeCostCategoryRuleRuleAndOrTags `field:"optional" json:"tags" yaml:"tags"`
}

