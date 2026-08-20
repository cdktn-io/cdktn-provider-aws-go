// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mailmanagertrafficpolicy


type MailmanagerTrafficPolicyPolicyStatement struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.61.0/docs/resources/mailmanager_traffic_policy#action MailmanagerTrafficPolicy#action}.
	Action *string `field:"required" json:"action" yaml:"action"`
	// condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.61.0/docs/resources/mailmanager_traffic_policy#condition MailmanagerTrafficPolicy#condition}
	Condition interface{} `field:"optional" json:"condition" yaml:"condition"`
}

