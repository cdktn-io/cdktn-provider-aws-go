// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerlabelingjob

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SagemakerLabelingJobConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/sagemaker_labeling_job#label_attribute_name SagemakerLabelingJob#label_attribute_name}.
	LabelAttributeName *string `field:"required" json:"labelAttributeName" yaml:"labelAttributeName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/sagemaker_labeling_job#labeling_job_name SagemakerLabelingJob#labeling_job_name}.
	LabelingJobName *string `field:"required" json:"labelingJobName" yaml:"labelingJobName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/sagemaker_labeling_job#role_arn SagemakerLabelingJob#role_arn}.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// human_task_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/sagemaker_labeling_job#human_task_config SagemakerLabelingJob#human_task_config}
	HumanTaskConfig interface{} `field:"optional" json:"humanTaskConfig" yaml:"humanTaskConfig"`
	// input_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/sagemaker_labeling_job#input_config SagemakerLabelingJob#input_config}
	InputConfig interface{} `field:"optional" json:"inputConfig" yaml:"inputConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/sagemaker_labeling_job#label_category_config_s3_uri SagemakerLabelingJob#label_category_config_s3_uri}.
	LabelCategoryConfigS3Uri *string `field:"optional" json:"labelCategoryConfigS3Uri" yaml:"labelCategoryConfigS3Uri"`
	// labeling_job_algorithms_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/sagemaker_labeling_job#labeling_job_algorithms_config SagemakerLabelingJob#labeling_job_algorithms_config}
	LabelingJobAlgorithmsConfig interface{} `field:"optional" json:"labelingJobAlgorithmsConfig" yaml:"labelingJobAlgorithmsConfig"`
	// output_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/sagemaker_labeling_job#output_config SagemakerLabelingJob#output_config}
	OutputConfig interface{} `field:"optional" json:"outputConfig" yaml:"outputConfig"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/sagemaker_labeling_job#region SagemakerLabelingJob#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/sagemaker_labeling_job#stopping_conditions SagemakerLabelingJob#stopping_conditions}.
	StoppingConditions interface{} `field:"optional" json:"stoppingConditions" yaml:"stoppingConditions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/sagemaker_labeling_job#tags SagemakerLabelingJob#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

