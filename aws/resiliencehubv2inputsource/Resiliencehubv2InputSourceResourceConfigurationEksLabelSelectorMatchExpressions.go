// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package resiliencehubv2inputsource


type Resiliencehubv2InputSourceResourceConfigurationEksLabelSelectorMatchExpressions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.70.0/docs/resources/resiliencehubv2_input_source#key Resiliencehubv2InputSource#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.70.0/docs/resources/resiliencehubv2_input_source#operator Resiliencehubv2InputSource#operator}.
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.70.0/docs/resources/resiliencehubv2_input_source#values Resiliencehubv2InputSource#values}.
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

