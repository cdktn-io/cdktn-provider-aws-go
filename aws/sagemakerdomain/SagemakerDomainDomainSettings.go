// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerdomain


type SagemakerDomainDomainSettings struct {
	// docker_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.39.0/docs/resources/sagemaker_domain#docker_settings SagemakerDomain#docker_settings}
	DockerSettings *SagemakerDomainDomainSettingsDockerSettings `field:"optional" json:"dockerSettings" yaml:"dockerSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.39.0/docs/resources/sagemaker_domain#execution_role_identity_config SagemakerDomain#execution_role_identity_config}.
	ExecutionRoleIdentityConfig *string `field:"optional" json:"executionRoleIdentityConfig" yaml:"executionRoleIdentityConfig"`
	// r_studio_server_pro_domain_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.39.0/docs/resources/sagemaker_domain#r_studio_server_pro_domain_settings SagemakerDomain#r_studio_server_pro_domain_settings}
	RStudioServerProDomainSettings *SagemakerDomainDomainSettingsRStudioServerProDomainSettings `field:"optional" json:"rStudioServerProDomainSettings" yaml:"rStudioServerProDomainSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.39.0/docs/resources/sagemaker_domain#security_group_ids SagemakerDomain#security_group_ids}.
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// trusted_identity_propagation_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.39.0/docs/resources/sagemaker_domain#trusted_identity_propagation_settings SagemakerDomain#trusted_identity_propagation_settings}
	TrustedIdentityPropagationSettings *SagemakerDomainDomainSettingsTrustedIdentityPropagationSettings `field:"optional" json:"trustedIdentityPropagationSettings" yaml:"trustedIdentityPropagationSettings"`
}

