// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lambdamicrovmsmicrovm


type LambdamicrovmsMicrovmLogging struct {
	// cloudwatch block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/lambdamicrovms_microvm#cloudwatch LambdamicrovmsMicrovm#cloudwatch}
	Cloudwatch interface{} `field:"optional" json:"cloudwatch" yaml:"cloudwatch"`
	// disabled block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/lambdamicrovms_microvm#disabled LambdamicrovmsMicrovm#disabled}
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
}

