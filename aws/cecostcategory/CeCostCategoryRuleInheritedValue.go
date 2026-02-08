// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cecostcategory


type CeCostCategoryRuleInheritedValue struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.31.0/docs/resources/ce_cost_category#dimension_key CeCostCategory#dimension_key}.
	DimensionKey *string `field:"optional" json:"dimensionKey" yaml:"dimensionKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.31.0/docs/resources/ce_cost_category#dimension_name CeCostCategory#dimension_name}.
	DimensionName *string `field:"optional" json:"dimensionName" yaml:"dimensionName"`
}

