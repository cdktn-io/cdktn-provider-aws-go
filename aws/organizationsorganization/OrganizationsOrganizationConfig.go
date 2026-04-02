// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package organizationsorganization

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type OrganizationsOrganizationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.39.0/docs/resources/organizations_organization#aws_service_access_principals OrganizationsOrganization#aws_service_access_principals}.
	AwsServiceAccessPrincipals *[]*string `field:"optional" json:"awsServiceAccessPrincipals" yaml:"awsServiceAccessPrincipals"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.39.0/docs/resources/organizations_organization#enabled_policy_types OrganizationsOrganization#enabled_policy_types}.
	EnabledPolicyTypes *[]*string `field:"optional" json:"enabledPolicyTypes" yaml:"enabledPolicyTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.39.0/docs/resources/organizations_organization#feature_set OrganizationsOrganization#feature_set}.
	FeatureSet *string `field:"optional" json:"featureSet" yaml:"featureSet"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.39.0/docs/resources/organizations_organization#id OrganizationsOrganization#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.39.0/docs/resources/organizations_organization#return_organization_only OrganizationsOrganization#return_organization_only}.
	ReturnOrganizationOnly interface{} `field:"optional" json:"returnOrganizationOnly" yaml:"returnOrganizationOnly"`
}

