// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mailmanagerruleset


type MailmanagerRuleSetRuleActionWriteToS3 struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.61.0/docs/resources/mailmanager_rule_set#role_arn MailmanagerRuleSet#role_arn}.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.61.0/docs/resources/mailmanager_rule_set#s3_bucket MailmanagerRuleSet#s3_bucket}.
	S3Bucket *string `field:"required" json:"s3Bucket" yaml:"s3Bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.61.0/docs/resources/mailmanager_rule_set#action_failure_policy MailmanagerRuleSet#action_failure_policy}.
	ActionFailurePolicy *string `field:"optional" json:"actionFailurePolicy" yaml:"actionFailurePolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.61.0/docs/resources/mailmanager_rule_set#s3_prefix MailmanagerRuleSet#s3_prefix}.
	S3Prefix *string `field:"optional" json:"s3Prefix" yaml:"s3Prefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.61.0/docs/resources/mailmanager_rule_set#s3_sse_kms_key_id MailmanagerRuleSet#s3_sse_kms_key_id}.
	S3SseKmsKeyId *string `field:"optional" json:"s3SseKmsKeyId" yaml:"s3SseKmsKeyId"`
}

