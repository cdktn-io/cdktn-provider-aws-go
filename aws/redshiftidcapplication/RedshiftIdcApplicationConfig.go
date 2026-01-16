// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package redshiftidcapplication

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RedshiftIdcApplicationConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/redshift_idc_application#iam_role_arn RedshiftIdcApplication#iam_role_arn}.
	IamRoleArn *string `field:"required" json:"iamRoleArn" yaml:"iamRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/redshift_idc_application#idc_display_name RedshiftIdcApplication#idc_display_name}.
	IdcDisplayName *string `field:"required" json:"idcDisplayName" yaml:"idcDisplayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/redshift_idc_application#idc_instance_arn RedshiftIdcApplication#idc_instance_arn}.
	IdcInstanceArn *string `field:"required" json:"idcInstanceArn" yaml:"idcInstanceArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/redshift_idc_application#redshift_idc_application_name RedshiftIdcApplication#redshift_idc_application_name}.
	RedshiftIdcApplicationName *string `field:"required" json:"redshiftIdcApplicationName" yaml:"redshiftIdcApplicationName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/redshift_idc_application#application_type RedshiftIdcApplication#application_type}.
	ApplicationType *string `field:"optional" json:"applicationType" yaml:"applicationType"`
	// authorized_token_issuer block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/redshift_idc_application#authorized_token_issuer RedshiftIdcApplication#authorized_token_issuer}
	AuthorizedTokenIssuer interface{} `field:"optional" json:"authorizedTokenIssuer" yaml:"authorizedTokenIssuer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/redshift_idc_application#identity_namespace RedshiftIdcApplication#identity_namespace}.
	IdentityNamespace *string `field:"optional" json:"identityNamespace" yaml:"identityNamespace"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/redshift_idc_application#region RedshiftIdcApplication#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// service_integration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/redshift_idc_application#service_integration RedshiftIdcApplication#service_integration}
	ServiceIntegration interface{} `field:"optional" json:"serviceIntegration" yaml:"serviceIntegration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/redshift_idc_application#tags RedshiftIdcApplication#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

