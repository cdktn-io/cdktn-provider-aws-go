// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package glueconnection


type GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2ClientApplication struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.49.0/docs/resources/glue_connection#aws_managed_client_application_reference GlueConnection#aws_managed_client_application_reference}.
	AwsManagedClientApplicationReference *string `field:"optional" json:"awsManagedClientApplicationReference" yaml:"awsManagedClientApplicationReference"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.49.0/docs/resources/glue_connection#user_managed_client_application_client_id GlueConnection#user_managed_client_application_client_id}.
	UserManagedClientApplicationClientId *string `field:"optional" json:"userManagedClientApplicationClientId" yaml:"userManagedClientApplicationClientId"`
}

