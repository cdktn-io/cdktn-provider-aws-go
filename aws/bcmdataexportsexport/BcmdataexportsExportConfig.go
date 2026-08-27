// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bcmdataexportsexport

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BcmdataexportsExportConfig struct {
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
	// export block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bcmdataexports_export#export BcmdataexportsExport#export}
	Export interface{} `field:"optional" json:"export" yaml:"export"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bcmdataexports_export#tags BcmdataexportsExport#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bcmdataexports_export#timeouts BcmdataexportsExport#timeouts}
	Timeouts *BcmdataexportsExportTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

