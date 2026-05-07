// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vpclatticelistenerrule


type VpclatticeListenerRuleMatch struct {
	// http_match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.44.0/docs/resources/vpclattice_listener_rule#http_match VpclatticeListenerRule#http_match}
	HttpMatch *VpclatticeListenerRuleMatchHttpMatch `field:"required" json:"httpMatch" yaml:"httpMatch"`
}

