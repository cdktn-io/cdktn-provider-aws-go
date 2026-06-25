// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2networkinsightsaccessscope


type Ec2NetworkInsightsAccessScopeExcludePaths struct {
	// destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/ec2_network_insights_access_scope#destination Ec2NetworkInsightsAccessScope#destination}
	Destination interface{} `field:"optional" json:"destination" yaml:"destination"`
	// source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/ec2_network_insights_access_scope#source Ec2NetworkInsightsAccessScope#source}
	Source interface{} `field:"optional" json:"source" yaml:"source"`
	// through_resources block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/ec2_network_insights_access_scope#through_resources Ec2NetworkInsightsAccessScope#through_resources}
	ThroughResources interface{} `field:"optional" json:"throughResources" yaml:"throughResources"`
}

