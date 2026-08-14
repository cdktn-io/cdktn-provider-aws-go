// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package opensearchserverlesssecurityconfig


type OpensearchserverlessSecurityConfigIamIdentityCenterOptions struct {
	// Instance ARN.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/opensearchserverless_security_config#instance_arn OpensearchserverlessSecurityConfig#instance_arn}
	InstanceArn *string `field:"required" json:"instanceArn" yaml:"instanceArn"`
	// Group attribute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/opensearchserverless_security_config#group_attribute OpensearchserverlessSecurityConfig#group_attribute}
	GroupAttribute *string `field:"optional" json:"groupAttribute" yaml:"groupAttribute"`
	// User attribute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/opensearchserverless_security_config#user_attribute OpensearchserverlessSecurityConfig#user_attribute}
	UserAttribute *string `field:"optional" json:"userAttribute" yaml:"userAttribute"`
}

