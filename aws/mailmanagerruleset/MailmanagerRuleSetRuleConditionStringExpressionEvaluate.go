// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mailmanagerruleset


type MailmanagerRuleSetRuleConditionStringExpressionEvaluate struct {
	// analysis block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/mailmanager_rule_set#analysis MailmanagerRuleSet#analysis}
	Analysis interface{} `field:"optional" json:"analysis" yaml:"analysis"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/mailmanager_rule_set#attribute MailmanagerRuleSet#attribute}.
	Attribute *string `field:"optional" json:"attribute" yaml:"attribute"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/mailmanager_rule_set#client_certificate_attribute MailmanagerRuleSet#client_certificate_attribute}.
	ClientCertificateAttribute *string `field:"optional" json:"clientCertificateAttribute" yaml:"clientCertificateAttribute"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/mailmanager_rule_set#mime_header_attribute MailmanagerRuleSet#mime_header_attribute}.
	MimeHeaderAttribute *string `field:"optional" json:"mimeHeaderAttribute" yaml:"mimeHeaderAttribute"`
}

