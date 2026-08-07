// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package redshiftnamespaceregistration

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type RedshiftNamespaceRegistrationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.58.0/docs/resources/redshift_namespace_registration#consumer_identifier RedshiftNamespaceRegistration#consumer_identifier}.
	ConsumerIdentifier *string `field:"required" json:"consumerIdentifier" yaml:"consumerIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.58.0/docs/resources/redshift_namespace_registration#namespace_type RedshiftNamespaceRegistration#namespace_type}.
	NamespaceType *string `field:"required" json:"namespaceType" yaml:"namespaceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.58.0/docs/resources/redshift_namespace_registration#provisioned_cluster_identifier RedshiftNamespaceRegistration#provisioned_cluster_identifier}.
	ProvisionedClusterIdentifier *string `field:"optional" json:"provisionedClusterIdentifier" yaml:"provisionedClusterIdentifier"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.58.0/docs/resources/redshift_namespace_registration#region RedshiftNamespaceRegistration#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.58.0/docs/resources/redshift_namespace_registration#serverless_namespace_identifier RedshiftNamespaceRegistration#serverless_namespace_identifier}.
	ServerlessNamespaceIdentifier *string `field:"optional" json:"serverlessNamespaceIdentifier" yaml:"serverlessNamespaceIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.58.0/docs/resources/redshift_namespace_registration#serverless_workgroup_identifier RedshiftNamespaceRegistration#serverless_workgroup_identifier}.
	ServerlessWorkgroupIdentifier *string `field:"optional" json:"serverlessWorkgroupIdentifier" yaml:"serverlessWorkgroupIdentifier"`
}

