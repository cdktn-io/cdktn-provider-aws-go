// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mailmanagertrafficpolicy


type MailmanagerTrafficPolicyPolicyStatementConditionStringExpressionEvaluateAnalysis struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_traffic_policy#analyzer MailmanagerTrafficPolicy#analyzer}.
	Analyzer *string `field:"required" json:"analyzer" yaml:"analyzer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_traffic_policy#result_field MailmanagerTrafficPolicy#result_field}.
	ResultField *string `field:"required" json:"resultField" yaml:"resultField"`
}

