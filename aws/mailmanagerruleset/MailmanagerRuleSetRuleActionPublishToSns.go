// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mailmanagerruleset


type MailmanagerRuleSetRuleActionPublishToSns struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_rule_set#role_arn MailmanagerRuleSet#role_arn}.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_rule_set#topic_arn MailmanagerRuleSet#topic_arn}.
	TopicArn *string `field:"required" json:"topicArn" yaml:"topicArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_rule_set#action_failure_policy MailmanagerRuleSet#action_failure_policy}.
	ActionFailurePolicy *string `field:"optional" json:"actionFailurePolicy" yaml:"actionFailurePolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_rule_set#encoding MailmanagerRuleSet#encoding}.
	Encoding *string `field:"optional" json:"encoding" yaml:"encoding"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_rule_set#payload_type MailmanagerRuleSet#payload_type}.
	PayloadType *string `field:"optional" json:"payloadType" yaml:"payloadType"`
}

