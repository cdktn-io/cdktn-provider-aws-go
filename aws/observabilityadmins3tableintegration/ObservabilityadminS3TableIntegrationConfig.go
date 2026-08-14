// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilityadmins3tableintegration

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ObservabilityadminS3TableIntegrationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/observabilityadmin_s3_table_integration#role_arn ObservabilityadminS3TableIntegration#role_arn}.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// encryption block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/observabilityadmin_s3_table_integration#encryption ObservabilityadminS3TableIntegration#encryption}
	Encryption interface{} `field:"optional" json:"encryption" yaml:"encryption"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/observabilityadmin_s3_table_integration#region ObservabilityadminS3TableIntegration#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/observabilityadmin_s3_table_integration#tags ObservabilityadminS3TableIntegration#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/observabilityadmin_s3_table_integration#timeouts ObservabilityadminS3TableIntegration#timeouts}
	Timeouts *ObservabilityadminS3TableIntegrationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

