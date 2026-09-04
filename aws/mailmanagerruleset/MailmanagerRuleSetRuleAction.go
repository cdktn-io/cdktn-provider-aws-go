// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mailmanagerruleset


type MailmanagerRuleSetRuleAction struct {
	// add_header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/mailmanager_rule_set#add_header MailmanagerRuleSet#add_header}
	AddHeader interface{} `field:"optional" json:"addHeader" yaml:"addHeader"`
	// archive block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/mailmanager_rule_set#archive MailmanagerRuleSet#archive}
	Archive interface{} `field:"optional" json:"archive" yaml:"archive"`
	// bounce block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/mailmanager_rule_set#bounce MailmanagerRuleSet#bounce}
	Bounce interface{} `field:"optional" json:"bounce" yaml:"bounce"`
	// deliver_to_mailbox block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/mailmanager_rule_set#deliver_to_mailbox MailmanagerRuleSet#deliver_to_mailbox}
	DeliverToMailbox interface{} `field:"optional" json:"deliverToMailbox" yaml:"deliverToMailbox"`
	// deliver_to_q_business block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/mailmanager_rule_set#deliver_to_q_business MailmanagerRuleSet#deliver_to_q_business}
	DeliverToQBusiness interface{} `field:"optional" json:"deliverToQBusiness" yaml:"deliverToQBusiness"`
	// drop block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/mailmanager_rule_set#drop MailmanagerRuleSet#drop}
	Drop interface{} `field:"optional" json:"drop" yaml:"drop"`
	// invoke_lambda block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/mailmanager_rule_set#invoke_lambda MailmanagerRuleSet#invoke_lambda}
	InvokeLambda interface{} `field:"optional" json:"invokeLambda" yaml:"invokeLambda"`
	// publish_to_sns block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/mailmanager_rule_set#publish_to_sns MailmanagerRuleSet#publish_to_sns}
	PublishToSns interface{} `field:"optional" json:"publishToSns" yaml:"publishToSns"`
	// relay block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/mailmanager_rule_set#relay MailmanagerRuleSet#relay}
	Relay interface{} `field:"optional" json:"relay" yaml:"relay"`
	// replace_recipient block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/mailmanager_rule_set#replace_recipient MailmanagerRuleSet#replace_recipient}
	ReplaceRecipient interface{} `field:"optional" json:"replaceRecipient" yaml:"replaceRecipient"`
	// send block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/mailmanager_rule_set#send MailmanagerRuleSet#send}
	Send interface{} `field:"optional" json:"send" yaml:"send"`
	// write_to_s3 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/mailmanager_rule_set#write_to_s3 MailmanagerRuleSet#write_to_s3}
	WriteToS3 interface{} `field:"optional" json:"writeToS3" yaml:"writeToS3"`
}

