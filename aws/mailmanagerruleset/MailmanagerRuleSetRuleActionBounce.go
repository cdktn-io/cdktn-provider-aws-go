// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mailmanagerruleset


type MailmanagerRuleSetRuleActionBounce struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_rule_set#diagnostic_message MailmanagerRuleSet#diagnostic_message}.
	DiagnosticMessage *string `field:"required" json:"diagnosticMessage" yaml:"diagnosticMessage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_rule_set#role_arn MailmanagerRuleSet#role_arn}.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_rule_set#sender MailmanagerRuleSet#sender}.
	Sender *string `field:"required" json:"sender" yaml:"sender"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_rule_set#smtp_reply_code MailmanagerRuleSet#smtp_reply_code}.
	SmtpReplyCode *string `field:"required" json:"smtpReplyCode" yaml:"smtpReplyCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_rule_set#status_code MailmanagerRuleSet#status_code}.
	StatusCode *string `field:"required" json:"statusCode" yaml:"statusCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_rule_set#action_failure_policy MailmanagerRuleSet#action_failure_policy}.
	ActionFailurePolicy *string `field:"optional" json:"actionFailurePolicy" yaml:"actionFailurePolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_rule_set#message MailmanagerRuleSet#message}.
	Message *string `field:"optional" json:"message" yaml:"message"`
}

