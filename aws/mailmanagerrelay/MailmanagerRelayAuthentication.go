// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mailmanagerrelay


type MailmanagerRelayAuthentication struct {
	// no_authentication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/mailmanager_relay#no_authentication MailmanagerRelay#no_authentication}
	NoAuthentication interface{} `field:"optional" json:"noAuthentication" yaml:"noAuthentication"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/mailmanager_relay#secret_arn MailmanagerRelay#secret_arn}.
	SecretArn *string `field:"optional" json:"secretArn" yaml:"secretArn"`
}

