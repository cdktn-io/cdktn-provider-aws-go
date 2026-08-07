// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package budgetsbudget


type BudgetsBudgetFilterExpressionAndOr struct {
	// cost_categories block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.58.0/docs/resources/budgets_budget#cost_categories BudgetsBudget#cost_categories}
	CostCategories *BudgetsBudgetFilterExpressionAndOrCostCategories `field:"optional" json:"costCategories" yaml:"costCategories"`
	// dimensions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.58.0/docs/resources/budgets_budget#dimensions BudgetsBudget#dimensions}
	Dimensions *BudgetsBudgetFilterExpressionAndOrDimensions `field:"optional" json:"dimensions" yaml:"dimensions"`
	// tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.58.0/docs/resources/budgets_budget#tags BudgetsBudget#tags}
	Tags *BudgetsBudgetFilterExpressionAndOrTags `field:"optional" json:"tags" yaml:"tags"`
}

