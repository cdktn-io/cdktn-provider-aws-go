// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package alblistenerrule


type AlbListenerRuleTransformUrlRewriteConfig struct {
	// rewrite block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.35.1/docs/resources/alb_listener_rule#rewrite AlbListenerRule#rewrite}
	Rewrite *AlbListenerRuleTransformUrlRewriteConfigRewrite `field:"optional" json:"rewrite" yaml:"rewrite"`
}

