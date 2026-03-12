// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wafv2webaclrulegroupassociation


type Wafv2WebAclRuleGroupAssociationVisibilityConfig struct {
	// Indicates whether the rule is available for use in the metrics for the web ACL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.36.0/docs/resources/wafv2_web_acl_rule_group_association#cloudwatch_metrics_enabled Wafv2WebAclRuleGroupAssociation#cloudwatch_metrics_enabled}
	CloudwatchMetricsEnabled interface{} `field:"required" json:"cloudwatchMetricsEnabled" yaml:"cloudwatchMetricsEnabled"`
	// A name for the metrics for this rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.36.0/docs/resources/wafv2_web_acl_rule_group_association#metric_name Wafv2WebAclRuleGroupAssociation#metric_name}
	MetricName *string `field:"required" json:"metricName" yaml:"metricName"`
	// Indicates whether to store a sampling of the web requests that match the rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.36.0/docs/resources/wafv2_web_acl_rule_group_association#sampled_requests_enabled Wafv2WebAclRuleGroupAssociation#sampled_requests_enabled}
	SampledRequestsEnabled interface{} `field:"required" json:"sampledRequestsEnabled" yaml:"sampledRequestsEnabled"`
}

