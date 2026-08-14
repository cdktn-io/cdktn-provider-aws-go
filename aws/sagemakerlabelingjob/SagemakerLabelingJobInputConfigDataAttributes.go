// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerlabelingjob


type SagemakerLabelingJobInputConfigDataAttributes struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/sagemaker_labeling_job#content_classifiers SagemakerLabelingJob#content_classifiers}.
	ContentClassifiers *[]*string `field:"optional" json:"contentClassifiers" yaml:"contentClassifiers"`
}

