// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudwatcheventconnection


type CloudwatchEventConnectionAuthParameters struct {
	// api_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.47.0/docs/resources/cloudwatch_event_connection#api_key CloudwatchEventConnection#api_key}
	ApiKey *CloudwatchEventConnectionAuthParametersApiKey `field:"optional" json:"apiKey" yaml:"apiKey"`
	// basic block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.47.0/docs/resources/cloudwatch_event_connection#basic CloudwatchEventConnection#basic}
	Basic *CloudwatchEventConnectionAuthParametersBasic `field:"optional" json:"basic" yaml:"basic"`
	// connectivity_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.47.0/docs/resources/cloudwatch_event_connection#connectivity_parameters CloudwatchEventConnection#connectivity_parameters}
	ConnectivityParameters *CloudwatchEventConnectionAuthParametersConnectivityParameters `field:"optional" json:"connectivityParameters" yaml:"connectivityParameters"`
	// invocation_http_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.47.0/docs/resources/cloudwatch_event_connection#invocation_http_parameters CloudwatchEventConnection#invocation_http_parameters}
	InvocationHttpParameters *CloudwatchEventConnectionAuthParametersInvocationHttpParameters `field:"optional" json:"invocationHttpParameters" yaml:"invocationHttpParameters"`
	// oauth block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.47.0/docs/resources/cloudwatch_event_connection#oauth CloudwatchEventConnection#oauth}
	Oauth *CloudwatchEventConnectionAuthParametersOauth `field:"optional" json:"oauth" yaml:"oauth"`
}

