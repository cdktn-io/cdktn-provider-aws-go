// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package codedeploydeploymentconfig


type CodedeployDeploymentConfigTrafficRoutingConfigTimeBasedLinear struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.47.0/docs/resources/codedeploy_deployment_config#interval CodedeployDeploymentConfig#interval}.
	Interval *float64 `field:"optional" json:"interval" yaml:"interval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.47.0/docs/resources/codedeploy_deployment_config#percentage CodedeployDeploymentConfig#percentage}.
	Percentage *float64 `field:"optional" json:"percentage" yaml:"percentage"`
}

