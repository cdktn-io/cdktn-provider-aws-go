// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ephemeralawsstswebidentitytoken

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type EphemeralAwsStsWebIdentityTokenConfig struct {
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
	// The intended recipients of the token (populates the `aud` claim in the JWT).
	//
	// Must contain between 1 and 10 items.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/ephemeral-resources/sts_web_identity_token#audience EphemeralAwsStsWebIdentityToken#audience}
	Audience *[]*string `field:"required" json:"audience" yaml:"audience"`
	// The cryptographic algorithm to use for signing the JWT.
	//
	// Valid values are `RS256` (RSA with SHA-256) and `ES384` (ECDSA using P-384 curve with SHA-384).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/ephemeral-resources/sts_web_identity_token#signing_algorithm EphemeralAwsStsWebIdentityToken#signing_algorithm}
	SigningAlgorithm *string `field:"required" json:"signingAlgorithm" yaml:"signingAlgorithm"`
	// The duration, in seconds, for which the JWT will remain valid.
	//
	// Value can range from 60 to 3600 seconds. Default is 300 seconds (5 minutes).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/ephemeral-resources/sts_web_identity_token#duration_seconds EphemeralAwsStsWebIdentityToken#duration_seconds}
	DurationSeconds *float64 `field:"optional" json:"durationSeconds" yaml:"durationSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/ephemeral-resources/sts_web_identity_token#tags EphemeralAwsStsWebIdentityToken#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

