// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package codedeploydeploymentgroup


type CodedeployDeploymentGroupEc2TagSet struct {
	// ec2_tag_filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/codedeploy_deployment_group#ec2_tag_filter CodedeployDeploymentGroup#ec2_tag_filter}
	Ec2TagFilter interface{} `field:"optional" json:"ec2TagFilter" yaml:"ec2TagFilter"`
}

