// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawssavingsplansofferings

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAwsSavingsplansOfferingsConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.45.0/docs/data-sources/savingsplans_offerings#currencies DataAwsSavingsplansOfferings#currencies}.
	Currencies *[]*string `field:"optional" json:"currencies" yaml:"currencies"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.45.0/docs/data-sources/savingsplans_offerings#descriptions DataAwsSavingsplansOfferings#descriptions}.
	Descriptions *[]*string `field:"optional" json:"descriptions" yaml:"descriptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.45.0/docs/data-sources/savingsplans_offerings#durations DataAwsSavingsplansOfferings#durations}.
	Durations *[]*float64 `field:"optional" json:"durations" yaml:"durations"`
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.45.0/docs/data-sources/savingsplans_offerings#filter DataAwsSavingsplansOfferings#filter}
	Filter interface{} `field:"optional" json:"filter" yaml:"filter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.45.0/docs/data-sources/savingsplans_offerings#offering_ids DataAwsSavingsplansOfferings#offering_ids}.
	OfferingIds *[]*string `field:"optional" json:"offeringIds" yaml:"offeringIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.45.0/docs/data-sources/savingsplans_offerings#operations DataAwsSavingsplansOfferings#operations}.
	Operations *[]*string `field:"optional" json:"operations" yaml:"operations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.45.0/docs/data-sources/savingsplans_offerings#payment_options DataAwsSavingsplansOfferings#payment_options}.
	PaymentOptions *[]*string `field:"optional" json:"paymentOptions" yaml:"paymentOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.45.0/docs/data-sources/savingsplans_offerings#plan_types DataAwsSavingsplansOfferings#plan_types}.
	PlanTypes *[]*string `field:"optional" json:"planTypes" yaml:"planTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.45.0/docs/data-sources/savingsplans_offerings#product_type DataAwsSavingsplansOfferings#product_type}.
	ProductType *string `field:"optional" json:"productType" yaml:"productType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.45.0/docs/data-sources/savingsplans_offerings#service_codes DataAwsSavingsplansOfferings#service_codes}.
	ServiceCodes *[]*string `field:"optional" json:"serviceCodes" yaml:"serviceCodes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.45.0/docs/data-sources/savingsplans_offerings#usage_types DataAwsSavingsplansOfferings#usage_types}.
	UsageTypes *[]*string `field:"optional" json:"usageTypes" yaml:"usageTypes"`
}

