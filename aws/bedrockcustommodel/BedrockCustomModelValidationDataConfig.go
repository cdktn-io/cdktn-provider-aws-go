// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockcustommodel


type BedrockCustomModelValidationDataConfig struct {
	// validator block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.35.0/docs/resources/bedrock_custom_model#validator BedrockCustomModel#validator}
	Validator interface{} `field:"optional" json:"validator" yaml:"validator"`
}

