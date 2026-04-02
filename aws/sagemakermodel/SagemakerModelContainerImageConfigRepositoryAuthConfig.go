// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakermodel


type SagemakerModelContainerImageConfigRepositoryAuthConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.39.0/docs/resources/sagemaker_model#repository_credentials_provider_arn SagemakerModel#repository_credentials_provider_arn}.
	RepositoryCredentialsProviderArn *string `field:"required" json:"repositoryCredentialsProviderArn" yaml:"repositoryCredentialsProviderArn"`
}

