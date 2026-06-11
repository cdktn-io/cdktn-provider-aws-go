// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2networkinsightsaccessscope

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Ec2NetworkInsightsAccessScopeConfig struct {
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
	// exclude_paths block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/ec2_network_insights_access_scope#exclude_paths Ec2NetworkInsightsAccessScope#exclude_paths}
	ExcludePaths interface{} `field:"optional" json:"excludePaths" yaml:"excludePaths"`
	// match_paths block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/ec2_network_insights_access_scope#match_paths Ec2NetworkInsightsAccessScope#match_paths}
	MatchPaths interface{} `field:"optional" json:"matchPaths" yaml:"matchPaths"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/ec2_network_insights_access_scope#region Ec2NetworkInsightsAccessScope#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/ec2_network_insights_access_scope#tags Ec2NetworkInsightsAccessScope#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

