// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudwatchomniaccessgrant


type CloudwatchomniAccessGrantScopedAction struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.70.0/docs/resources/cloudwatchomni_access_grant#actions CloudwatchomniAccessGrant#actions}.
	Actions *[]*string `field:"required" json:"actions" yaml:"actions"`
	// resource block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.70.0/docs/resources/cloudwatchomni_access_grant#resource CloudwatchomniAccessGrant#resource}
	Resource interface{} `field:"optional" json:"resource" yaml:"resource"`
}

