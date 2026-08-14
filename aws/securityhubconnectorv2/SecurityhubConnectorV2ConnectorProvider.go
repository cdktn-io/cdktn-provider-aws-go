// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securityhubconnectorv2


type SecurityhubConnectorV2ConnectorProvider struct {
	// jira_cloud block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/securityhub_connector_v2#jira_cloud SecurityhubConnectorV2#jira_cloud}
	JiraCloud interface{} `field:"optional" json:"jiraCloud" yaml:"jiraCloud"`
	// service_now block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/securityhub_connector_v2#service_now SecurityhubConnectorV2#service_now}
	ServiceNow interface{} `field:"optional" json:"serviceNow" yaml:"serviceNow"`
}

