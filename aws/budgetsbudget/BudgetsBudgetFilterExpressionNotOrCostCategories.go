// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package budgetsbudget


type BudgetsBudgetFilterExpressionNotOrCostCategories struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.34.0/docs/resources/budgets_budget#key BudgetsBudget#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.34.0/docs/resources/budgets_budget#match_options BudgetsBudget#match_options}.
	MatchOptions *[]*string `field:"optional" json:"matchOptions" yaml:"matchOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.34.0/docs/resources/budgets_budget#values BudgetsBudget#values}.
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

