// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ephemeralawssecretsmanagerrandompassword

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type EphemeralAwsSecretsmanagerRandomPasswordConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/ephemeral-resources/secretsmanager_random_password#exclude_characters EphemeralAwsSecretsmanagerRandomPassword#exclude_characters}.
	ExcludeCharacters *string `field:"optional" json:"excludeCharacters" yaml:"excludeCharacters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/ephemeral-resources/secretsmanager_random_password#exclude_lowercase EphemeralAwsSecretsmanagerRandomPassword#exclude_lowercase}.
	ExcludeLowercase interface{} `field:"optional" json:"excludeLowercase" yaml:"excludeLowercase"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/ephemeral-resources/secretsmanager_random_password#exclude_numbers EphemeralAwsSecretsmanagerRandomPassword#exclude_numbers}.
	ExcludeNumbers interface{} `field:"optional" json:"excludeNumbers" yaml:"excludeNumbers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/ephemeral-resources/secretsmanager_random_password#exclude_punctuation EphemeralAwsSecretsmanagerRandomPassword#exclude_punctuation}.
	ExcludePunctuation interface{} `field:"optional" json:"excludePunctuation" yaml:"excludePunctuation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/ephemeral-resources/secretsmanager_random_password#exclude_uppercase EphemeralAwsSecretsmanagerRandomPassword#exclude_uppercase}.
	ExcludeUppercase interface{} `field:"optional" json:"excludeUppercase" yaml:"excludeUppercase"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/ephemeral-resources/secretsmanager_random_password#include_space EphemeralAwsSecretsmanagerRandomPassword#include_space}.
	IncludeSpace interface{} `field:"optional" json:"includeSpace" yaml:"includeSpace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/ephemeral-resources/secretsmanager_random_password#password_length EphemeralAwsSecretsmanagerRandomPassword#password_length}.
	PasswordLength *float64 `field:"optional" json:"passwordLength" yaml:"passwordLength"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/ephemeral-resources/secretsmanager_random_password#region EphemeralAwsSecretsmanagerRandomPassword#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/ephemeral-resources/secretsmanager_random_password#require_each_included_type EphemeralAwsSecretsmanagerRandomPassword#require_each_included_type}.
	RequireEachIncludedType interface{} `field:"optional" json:"requireEachIncludedType" yaml:"requireEachIncludedType"`
}

