// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package costoptimizationhubpreferences

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CostoptimizationhubPreferencesConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.40.0/docs/resources/costoptimizationhub_preferences#member_account_discount_visibility CostoptimizationhubPreferences#member_account_discount_visibility}.
	MemberAccountDiscountVisibility *string `field:"optional" json:"memberAccountDiscountVisibility" yaml:"memberAccountDiscountVisibility"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.40.0/docs/resources/costoptimizationhub_preferences#savings_estimation_mode CostoptimizationhubPreferences#savings_estimation_mode}.
	SavingsEstimationMode *string `field:"optional" json:"savingsEstimationMode" yaml:"savingsEstimationMode"`
}

