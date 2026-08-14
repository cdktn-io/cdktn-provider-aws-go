// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mailmanagerruleset


type MailmanagerRuleSetRuleActionInvokeLambda struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/mailmanager_rule_set#function_arn MailmanagerRuleSet#function_arn}.
	FunctionArn *string `field:"required" json:"functionArn" yaml:"functionArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/mailmanager_rule_set#invocation_type MailmanagerRuleSet#invocation_type}.
	InvocationType *string `field:"required" json:"invocationType" yaml:"invocationType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/mailmanager_rule_set#role_arn MailmanagerRuleSet#role_arn}.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/mailmanager_rule_set#action_failure_policy MailmanagerRuleSet#action_failure_policy}.
	ActionFailurePolicy *string `field:"optional" json:"actionFailurePolicy" yaml:"actionFailurePolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/mailmanager_rule_set#retry_time_minutes MailmanagerRuleSet#retry_time_minutes}.
	RetryTimeMinutes *float64 `field:"optional" json:"retryTimeMinutes" yaml:"retryTimeMinutes"`
}

