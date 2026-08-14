// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreharness


type BedrockagentcoreHarnessMemoryManagedMemoryConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/bedrockagentcore_harness#encryption_key_arn BedrockagentcoreHarness#encryption_key_arn}.
	EncryptionKeyArn *string `field:"optional" json:"encryptionKeyArn" yaml:"encryptionKeyArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/bedrockagentcore_harness#event_expiry_duration BedrockagentcoreHarness#event_expiry_duration}.
	EventExpiryDuration *float64 `field:"optional" json:"eventExpiryDuration" yaml:"eventExpiryDuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/bedrockagentcore_harness#strategies BedrockagentcoreHarness#strategies}.
	Strategies *[]*string `field:"optional" json:"strategies" yaml:"strategies"`
}

