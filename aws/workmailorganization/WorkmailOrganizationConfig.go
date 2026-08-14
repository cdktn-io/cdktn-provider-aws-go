// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package workmailorganization

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type WorkmailOrganizationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/workmail_organization#organization_alias WorkmailOrganization#organization_alias}.
	OrganizationAlias *string `field:"required" json:"organizationAlias" yaml:"organizationAlias"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/workmail_organization#delete_directory WorkmailOrganization#delete_directory}.
	DeleteDirectory interface{} `field:"optional" json:"deleteDirectory" yaml:"deleteDirectory"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/workmail_organization#delete_identity_center_application WorkmailOrganization#delete_identity_center_application}.
	DeleteIdentityCenterApplication interface{} `field:"optional" json:"deleteIdentityCenterApplication" yaml:"deleteIdentityCenterApplication"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/workmail_organization#directory_id WorkmailOrganization#directory_id}.
	DirectoryId *string `field:"optional" json:"directoryId" yaml:"directoryId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/workmail_organization#interoperability_enabled WorkmailOrganization#interoperability_enabled}.
	InteroperabilityEnabled interface{} `field:"optional" json:"interoperabilityEnabled" yaml:"interoperabilityEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/workmail_organization#kms_key_arn WorkmailOrganization#kms_key_arn}.
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/workmail_organization#region WorkmailOrganization#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/workmail_organization#tags WorkmailOrganization#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/workmail_organization#timeouts WorkmailOrganization#timeouts}
	Timeouts *WorkmailOrganizationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

