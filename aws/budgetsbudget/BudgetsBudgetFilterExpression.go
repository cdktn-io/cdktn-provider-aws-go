// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package budgetsbudget


type BudgetsBudgetFilterExpression struct {
	// and block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.45.0/docs/resources/budgets_budget#and BudgetsBudget#and}
	And interface{} `field:"optional" json:"and" yaml:"and"`
	// cost_categories block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.45.0/docs/resources/budgets_budget#cost_categories BudgetsBudget#cost_categories}
	CostCategories *BudgetsBudgetFilterExpressionCostCategories `field:"optional" json:"costCategories" yaml:"costCategories"`
	// dimensions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.45.0/docs/resources/budgets_budget#dimensions BudgetsBudget#dimensions}
	Dimensions *BudgetsBudgetFilterExpressionDimensions `field:"optional" json:"dimensions" yaml:"dimensions"`
	// not block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.45.0/docs/resources/budgets_budget#not BudgetsBudget#not}
	Not *BudgetsBudgetFilterExpressionNot `field:"optional" json:"not" yaml:"not"`
	// or block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.45.0/docs/resources/budgets_budget#or BudgetsBudget#or}
	Or interface{} `field:"optional" json:"or" yaml:"or"`
	// tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.45.0/docs/resources/budgets_budget#tags BudgetsBudget#tags}
	Tags *BudgetsBudgetFilterExpressionTags `field:"optional" json:"tags" yaml:"tags"`
}

