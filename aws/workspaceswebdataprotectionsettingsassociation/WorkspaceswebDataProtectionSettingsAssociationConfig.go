// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package workspaceswebdataprotectionsettingsassociation

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type WorkspaceswebDataProtectionSettingsAssociationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/workspacesweb_data_protection_settings_association#data_protection_settings_arn WorkspaceswebDataProtectionSettingsAssociation#data_protection_settings_arn}.
	DataProtectionSettingsArn *string `field:"required" json:"dataProtectionSettingsArn" yaml:"dataProtectionSettingsArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/workspacesweb_data_protection_settings_association#portal_arn WorkspaceswebDataProtectionSettingsAssociation#portal_arn}.
	PortalArn *string `field:"required" json:"portalArn" yaml:"portalArn"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/workspacesweb_data_protection_settings_association#region WorkspaceswebDataProtectionSettingsAssociation#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
}

