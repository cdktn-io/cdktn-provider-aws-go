// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package uxcaccountcustomizations

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type UxcAccountCustomizationsConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.58.0/docs/resources/uxc_account_customizations#account_color UxcAccountCustomizations#account_color}.
	AccountColor *string `field:"optional" json:"accountColor" yaml:"accountColor"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.58.0/docs/resources/uxc_account_customizations#visible_regions UxcAccountCustomizations#visible_regions}.
	VisibleRegions *[]*string `field:"optional" json:"visibleRegions" yaml:"visibleRegions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.58.0/docs/resources/uxc_account_customizations#visible_services UxcAccountCustomizations#visible_services}.
	VisibleServices *[]*string `field:"optional" json:"visibleServices" yaml:"visibleServices"`
}

