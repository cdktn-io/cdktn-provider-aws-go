// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ephemeralawssecretsmanagersecretversion

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type EphemeralAwsSecretsmanagerSecretVersionConfig struct {
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformEphemeralResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.61.0/docs/ephemeral-resources/secretsmanager_secret_version#secret_id EphemeralAwsSecretsmanagerSecretVersion#secret_id}.
	SecretId *string `field:"required" json:"secretId" yaml:"secretId"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.61.0/docs/ephemeral-resources/secretsmanager_secret_version#region EphemeralAwsSecretsmanagerSecretVersion#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.61.0/docs/ephemeral-resources/secretsmanager_secret_version#version_id EphemeralAwsSecretsmanagerSecretVersion#version_id}.
	VersionId *string `field:"optional" json:"versionId" yaml:"versionId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.61.0/docs/ephemeral-resources/secretsmanager_secret_version#version_stage EphemeralAwsSecretsmanagerSecretVersion#version_stage}.
	VersionStage *string `field:"optional" json:"versionStage" yaml:"versionStage"`
}

