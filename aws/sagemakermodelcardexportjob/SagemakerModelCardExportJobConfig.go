// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakermodelcardexportjob

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SagemakerModelCardExportJobConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/sagemaker_model_card_export_job#model_card_export_job_name SagemakerModelCardExportJob#model_card_export_job_name}.
	ModelCardExportJobName *string `field:"required" json:"modelCardExportJobName" yaml:"modelCardExportJobName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/sagemaker_model_card_export_job#model_card_name SagemakerModelCardExportJob#model_card_name}.
	ModelCardName *string `field:"required" json:"modelCardName" yaml:"modelCardName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/sagemaker_model_card_export_job#model_card_version SagemakerModelCardExportJob#model_card_version}.
	ModelCardVersion *float64 `field:"optional" json:"modelCardVersion" yaml:"modelCardVersion"`
	// output_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/sagemaker_model_card_export_job#output_config SagemakerModelCardExportJob#output_config}
	OutputConfig interface{} `field:"optional" json:"outputConfig" yaml:"outputConfig"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/sagemaker_model_card_export_job#region SagemakerModelCardExportJob#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/sagemaker_model_card_export_job#timeouts SagemakerModelCardExportJob#timeouts}
	Timeouts *SagemakerModelCardExportJobTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

