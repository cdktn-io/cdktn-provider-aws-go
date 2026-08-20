// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudwatchlogs3tableintegrationsource

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CloudwatchLogS3TableIntegrationSourceConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.61.0/docs/resources/cloudwatch_log_s3_table_integration_source#integration_arn CloudwatchLogS3TableIntegrationSource#integration_arn}.
	IntegrationArn *string `field:"required" json:"integrationArn" yaml:"integrationArn"`
	// data_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.61.0/docs/resources/cloudwatch_log_s3_table_integration_source#data_source CloudwatchLogS3TableIntegrationSource#data_source}
	DataSource interface{} `field:"optional" json:"dataSource" yaml:"dataSource"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.61.0/docs/resources/cloudwatch_log_s3_table_integration_source#region CloudwatchLogS3TableIntegrationSource#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.61.0/docs/resources/cloudwatch_log_s3_table_integration_source#timeouts CloudwatchLogS3TableIntegrationSource#timeouts}
	Timeouts *CloudwatchLogS3TableIntegrationSourceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

