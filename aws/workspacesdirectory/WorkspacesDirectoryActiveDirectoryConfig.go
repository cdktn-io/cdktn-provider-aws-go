// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package workspacesdirectory


type WorkspacesDirectoryActiveDirectoryConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.40.0/docs/resources/workspaces_directory#domain_name WorkspacesDirectory#domain_name}.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.40.0/docs/resources/workspaces_directory#service_account_secret_arn WorkspacesDirectory#service_account_secret_arn}.
	ServiceAccountSecretArn *string `field:"required" json:"serviceAccountSecretArn" yaml:"serviceAccountSecretArn"`
}

