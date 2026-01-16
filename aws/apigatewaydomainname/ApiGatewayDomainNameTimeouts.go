// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigatewaydomainname


type ApiGatewayDomainNameTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/api_gateway_domain_name#create ApiGatewayDomainName#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/api_gateway_domain_name#update ApiGatewayDomainName#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

