// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package budgetsbudget

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BudgetsBudgetConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.54.0/docs/resources/budgets_budget#budget_type BudgetsBudget#budget_type}.
	BudgetType *string `field:"required" json:"budgetType" yaml:"budgetType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.54.0/docs/resources/budgets_budget#time_unit BudgetsBudget#time_unit}.
	TimeUnit *string `field:"required" json:"timeUnit" yaml:"timeUnit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.54.0/docs/resources/budgets_budget#account_id BudgetsBudget#account_id}.
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// auto_adjust_data block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.54.0/docs/resources/budgets_budget#auto_adjust_data BudgetsBudget#auto_adjust_data}
	AutoAdjustData *BudgetsBudgetAutoAdjustData `field:"optional" json:"autoAdjustData" yaml:"autoAdjustData"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.54.0/docs/resources/budgets_budget#billing_view_arn BudgetsBudget#billing_view_arn}.
	BillingViewArn *string `field:"optional" json:"billingViewArn" yaml:"billingViewArn"`
	// cost_filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.54.0/docs/resources/budgets_budget#cost_filter BudgetsBudget#cost_filter}
	CostFilter interface{} `field:"optional" json:"costFilter" yaml:"costFilter"`
	// cost_types block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.54.0/docs/resources/budgets_budget#cost_types BudgetsBudget#cost_types}
	CostTypes *BudgetsBudgetCostTypes `field:"optional" json:"costTypes" yaml:"costTypes"`
	// filter_expression block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.54.0/docs/resources/budgets_budget#filter_expression BudgetsBudget#filter_expression}
	FilterExpression *BudgetsBudgetFilterExpression `field:"optional" json:"filterExpression" yaml:"filterExpression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.54.0/docs/resources/budgets_budget#id BudgetsBudget#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.54.0/docs/resources/budgets_budget#limit_amount BudgetsBudget#limit_amount}.
	LimitAmount *string `field:"optional" json:"limitAmount" yaml:"limitAmount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.54.0/docs/resources/budgets_budget#limit_unit BudgetsBudget#limit_unit}.
	LimitUnit *string `field:"optional" json:"limitUnit" yaml:"limitUnit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.54.0/docs/resources/budgets_budget#metrics BudgetsBudget#metrics}.
	Metrics *[]*string `field:"optional" json:"metrics" yaml:"metrics"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.54.0/docs/resources/budgets_budget#name BudgetsBudget#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.54.0/docs/resources/budgets_budget#name_prefix BudgetsBudget#name_prefix}.
	NamePrefix *string `field:"optional" json:"namePrefix" yaml:"namePrefix"`
	// notification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.54.0/docs/resources/budgets_budget#notification BudgetsBudget#notification}
	Notification interface{} `field:"optional" json:"notification" yaml:"notification"`
	// planned_limit block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.54.0/docs/resources/budgets_budget#planned_limit BudgetsBudget#planned_limit}
	PlannedLimit interface{} `field:"optional" json:"plannedLimit" yaml:"plannedLimit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.54.0/docs/resources/budgets_budget#tags BudgetsBudget#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.54.0/docs/resources/budgets_budget#tags_all BudgetsBudget#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.54.0/docs/resources/budgets_budget#time_period_end BudgetsBudget#time_period_end}.
	TimePeriodEnd *string `field:"optional" json:"timePeriodEnd" yaml:"timePeriodEnd"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.54.0/docs/resources/budgets_budget#time_period_start BudgetsBudget#time_period_start}.
	TimePeriodStart *string `field:"optional" json:"timePeriodStart" yaml:"timePeriodStart"`
}

