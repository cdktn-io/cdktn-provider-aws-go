// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpcendpoint


type VpcEndpointDnsOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/vpc_endpoint#dns_record_ip_type VpcEndpoint#dns_record_ip_type}.
	DnsRecordIpType *string `field:"optional" json:"dnsRecordIpType" yaml:"dnsRecordIpType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/vpc_endpoint#private_dns_only_for_inbound_resolver_endpoint VpcEndpoint#private_dns_only_for_inbound_resolver_endpoint}.
	PrivateDnsOnlyForInboundResolverEndpoint interface{} `field:"optional" json:"privateDnsOnlyForInboundResolverEndpoint" yaml:"privateDnsOnlyForInboundResolverEndpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/vpc_endpoint#private_dns_preference VpcEndpoint#private_dns_preference}.
	PrivateDnsPreference *string `field:"optional" json:"privateDnsPreference" yaml:"privateDnsPreference"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/vpc_endpoint#private_dns_specified_domains VpcEndpoint#private_dns_specified_domains}.
	PrivateDnsSpecifiedDomains *[]*string `field:"optional" json:"privateDnsSpecifiedDomains" yaml:"privateDnsSpecifiedDomains"`
}

