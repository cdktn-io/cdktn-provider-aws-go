// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vpclatticeresourceconfiguration


type VpclatticeResourceConfigurationResourceConfigurationDefinition struct {
	// arn_resource block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.71.0/docs/resources/vpclattice_resource_configuration#arn_resource VpclatticeResourceConfiguration#arn_resource}
	ArnResource interface{} `field:"optional" json:"arnResource" yaml:"arnResource"`
	// cidr_resource block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.71.0/docs/resources/vpclattice_resource_configuration#cidr_resource VpclatticeResourceConfiguration#cidr_resource}
	CidrResource interface{} `field:"optional" json:"cidrResource" yaml:"cidrResource"`
	// dns_resource block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.71.0/docs/resources/vpclattice_resource_configuration#dns_resource VpclatticeResourceConfiguration#dns_resource}
	DnsResource interface{} `field:"optional" json:"dnsResource" yaml:"dnsResource"`
	// ip_resource block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.71.0/docs/resources/vpclattice_resource_configuration#ip_resource VpclatticeResourceConfiguration#ip_resource}
	IpResource interface{} `field:"optional" json:"ipResource" yaml:"ipResource"`
}

