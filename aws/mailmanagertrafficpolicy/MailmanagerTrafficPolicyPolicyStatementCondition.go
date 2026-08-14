// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mailmanagertrafficpolicy


type MailmanagerTrafficPolicyPolicyStatementCondition struct {
	// boolean_expression block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/mailmanager_traffic_policy#boolean_expression MailmanagerTrafficPolicy#boolean_expression}
	BooleanExpression interface{} `field:"optional" json:"booleanExpression" yaml:"booleanExpression"`
	// ip_expression block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/mailmanager_traffic_policy#ip_expression MailmanagerTrafficPolicy#ip_expression}
	IpExpression interface{} `field:"optional" json:"ipExpression" yaml:"ipExpression"`
	// ipv6_expression block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/mailmanager_traffic_policy#ipv6_expression MailmanagerTrafficPolicy#ipv6_expression}
	Ipv6Expression interface{} `field:"optional" json:"ipv6Expression" yaml:"ipv6Expression"`
	// string_expression block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/mailmanager_traffic_policy#string_expression MailmanagerTrafficPolicy#string_expression}
	StringExpression interface{} `field:"optional" json:"stringExpression" yaml:"stringExpression"`
	// tls_expression block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/resources/mailmanager_traffic_policy#tls_expression MailmanagerTrafficPolicy#tls_expression}
	TlsExpression interface{} `field:"optional" json:"tlsExpression" yaml:"tlsExpression"`
}

