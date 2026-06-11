// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksightanalysis


type QuicksightAnalysisParametersIntegerParameters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/quicksight_analysis#name QuicksightAnalysis#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/quicksight_analysis#values QuicksightAnalysis#values}.
	Values *[]*float64 `field:"required" json:"values" yaml:"values"`
}

