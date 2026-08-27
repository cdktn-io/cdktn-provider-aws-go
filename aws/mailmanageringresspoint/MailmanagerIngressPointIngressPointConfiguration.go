// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mailmanageringresspoint


type MailmanagerIngressPointIngressPointConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_ingress_point#secret_arn MailmanagerIngressPoint#secret_arn}.
	SecretArn *string `field:"optional" json:"secretArn" yaml:"secretArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_ingress_point#smtp_password_wo MailmanagerIngressPoint#smtp_password_wo}.
	SmtpPasswordWo *string `field:"optional" json:"smtpPasswordWo" yaml:"smtpPasswordWo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_ingress_point#smtp_password_wo_version MailmanagerIngressPoint#smtp_password_wo_version}.
	SmtpPasswordWoVersion *float64 `field:"optional" json:"smtpPasswordWoVersion" yaml:"smtpPasswordWoVersion"`
	// tls_auth_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_ingress_point#tls_auth_configuration MailmanagerIngressPoint#tls_auth_configuration}
	TlsAuthConfiguration interface{} `field:"optional" json:"tlsAuthConfiguration" yaml:"tlsAuthConfiguration"`
}

