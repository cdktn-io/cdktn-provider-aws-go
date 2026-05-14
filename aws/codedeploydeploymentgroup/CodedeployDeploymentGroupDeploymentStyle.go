// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package codedeploydeploymentgroup


type CodedeployDeploymentGroupDeploymentStyle struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.45.0/docs/resources/codedeploy_deployment_group#deployment_option CodedeployDeploymentGroup#deployment_option}.
	DeploymentOption *string `field:"optional" json:"deploymentOption" yaml:"deploymentOption"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.45.0/docs/resources/codedeploy_deployment_group#deployment_type CodedeployDeploymentGroup#deployment_type}.
	DeploymentType *string `field:"optional" json:"deploymentType" yaml:"deploymentType"`
}

