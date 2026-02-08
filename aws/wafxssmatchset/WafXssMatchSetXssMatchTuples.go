// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wafxssmatchset


type WafXssMatchSetXssMatchTuples struct {
	// field_to_match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.31.0/docs/resources/waf_xss_match_set#field_to_match WafXssMatchSet#field_to_match}
	FieldToMatch *WafXssMatchSetXssMatchTuplesFieldToMatch `field:"required" json:"fieldToMatch" yaml:"fieldToMatch"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.31.0/docs/resources/waf_xss_match_set#text_transformation WafXssMatchSet#text_transformation}.
	TextTransformation *string `field:"required" json:"textTransformation" yaml:"textTransformation"`
}

