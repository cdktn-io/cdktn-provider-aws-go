// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package opensearchdomain


type OpensearchDomainDeploymentStrategyOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.49.0/docs/resources/opensearch_domain#deployment_strategy OpensearchDomain#deployment_strategy}.
	DeploymentStrategy *string `field:"required" json:"deploymentStrategy" yaml:"deploymentStrategy"`
}

