// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerlabelingjob


type SagemakerLabelingJobHumanTaskConfigPublicWorkforceTaskPriceAmountInUsd struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.54.0/docs/resources/sagemaker_labeling_job#cents SagemakerLabelingJob#cents}.
	Cents *float64 `field:"optional" json:"cents" yaml:"cents"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.54.0/docs/resources/sagemaker_labeling_job#dollars SagemakerLabelingJob#dollars}.
	Dollars *float64 `field:"optional" json:"dollars" yaml:"dollars"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.54.0/docs/resources/sagemaker_labeling_job#tenth_fractions_of_a_cent SagemakerLabelingJob#tenth_fractions_of_a_cent}.
	TenthFractionsOfACent *float64 `field:"optional" json:"tenthFractionsOfACent" yaml:"tenthFractionsOfACent"`
}

