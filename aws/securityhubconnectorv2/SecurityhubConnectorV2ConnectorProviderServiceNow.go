// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securityhubconnectorv2


type SecurityhubConnectorV2ConnectorProviderServiceNow struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/securityhub_connector_v2#instance_name SecurityhubConnectorV2#instance_name}.
	InstanceName *string `field:"required" json:"instanceName" yaml:"instanceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/securityhub_connector_v2#secret_arn SecurityhubConnectorV2#secret_arn}.
	SecretArn *string `field:"required" json:"secretArn" yaml:"secretArn"`
}

