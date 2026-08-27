// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerlabelingjob


type SagemakerLabelingJobInputConfigDataSourceS3DataSource struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_labeling_job#manifest_s3_uri SagemakerLabelingJob#manifest_s3_uri}.
	ManifestS3Uri *string `field:"required" json:"manifestS3Uri" yaml:"manifestS3Uri"`
}

