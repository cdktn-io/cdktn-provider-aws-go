// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudfrontmultitenantdistribution


type CloudfrontMultitenantDistributionOriginCustomOriginConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_multitenant_distribution#http_port CloudfrontMultitenantDistribution#http_port}.
	HttpPort *float64 `field:"required" json:"httpPort" yaml:"httpPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_multitenant_distribution#https_port CloudfrontMultitenantDistribution#https_port}.
	HttpsPort *float64 `field:"required" json:"httpsPort" yaml:"httpsPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_multitenant_distribution#origin_protocol_policy CloudfrontMultitenantDistribution#origin_protocol_policy}.
	OriginProtocolPolicy *string `field:"required" json:"originProtocolPolicy" yaml:"originProtocolPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_multitenant_distribution#origin_ssl_protocols CloudfrontMultitenantDistribution#origin_ssl_protocols}.
	OriginSslProtocols *[]*string `field:"required" json:"originSslProtocols" yaml:"originSslProtocols"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_multitenant_distribution#ip_address_type CloudfrontMultitenantDistribution#ip_address_type}.
	IpAddressType *string `field:"optional" json:"ipAddressType" yaml:"ipAddressType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_multitenant_distribution#origin_keepalive_timeout CloudfrontMultitenantDistribution#origin_keepalive_timeout}.
	OriginKeepaliveTimeout *float64 `field:"optional" json:"originKeepaliveTimeout" yaml:"originKeepaliveTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_multitenant_distribution#origin_read_timeout CloudfrontMultitenantDistribution#origin_read_timeout}.
	OriginReadTimeout *float64 `field:"optional" json:"originReadTimeout" yaml:"originReadTimeout"`
}

