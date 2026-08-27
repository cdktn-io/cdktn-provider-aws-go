// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package budgetsbudget


type BudgetsBudgetFilterExpressionNot struct {
	// and block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget#and BudgetsBudget#and}
	And interface{} `field:"optional" json:"and" yaml:"and"`
	// cost_categories block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget#cost_categories BudgetsBudget#cost_categories}
	CostCategories *BudgetsBudgetFilterExpressionNotCostCategories `field:"optional" json:"costCategories" yaml:"costCategories"`
	// dimensions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget#dimensions BudgetsBudget#dimensions}
	Dimensions *BudgetsBudgetFilterExpressionNotDimensions `field:"optional" json:"dimensions" yaml:"dimensions"`
	// not block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget#not BudgetsBudget#not}
	Not *BudgetsBudgetFilterExpressionNotNot `field:"optional" json:"not" yaml:"not"`
	// or block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget#or BudgetsBudget#or}
	Or interface{} `field:"optional" json:"or" yaml:"or"`
	// tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget#tags BudgetsBudget#tags}
	Tags *BudgetsBudgetFilterExpressionNotTags `field:"optional" json:"tags" yaml:"tags"`
}

