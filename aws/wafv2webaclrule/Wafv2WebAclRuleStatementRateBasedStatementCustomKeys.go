// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wafv2webaclrule


type Wafv2WebAclRuleStatementRateBasedStatementCustomKeys struct {
	// asn block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.55.0/docs/resources/wafv2_web_acl_rule#asn Wafv2WebAclRuleA#asn}
	Asn interface{} `field:"optional" json:"asn" yaml:"asn"`
	// cookie block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.55.0/docs/resources/wafv2_web_acl_rule#cookie Wafv2WebAclRuleA#cookie}
	Cookie interface{} `field:"optional" json:"cookie" yaml:"cookie"`
	// forwarded_ip block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.55.0/docs/resources/wafv2_web_acl_rule#forwarded_ip Wafv2WebAclRuleA#forwarded_ip}
	ForwardedIp interface{} `field:"optional" json:"forwardedIp" yaml:"forwardedIp"`
	// header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.55.0/docs/resources/wafv2_web_acl_rule#header Wafv2WebAclRuleA#header}
	Header interface{} `field:"optional" json:"header" yaml:"header"`
	// http_method block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.55.0/docs/resources/wafv2_web_acl_rule#http_method Wafv2WebAclRuleA#http_method}
	HttpMethod interface{} `field:"optional" json:"httpMethod" yaml:"httpMethod"`
	// ip block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.55.0/docs/resources/wafv2_web_acl_rule#ip Wafv2WebAclRuleA#ip}
	Ip interface{} `field:"optional" json:"ip" yaml:"ip"`
	// ja3_fingerprint block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.55.0/docs/resources/wafv2_web_acl_rule#ja3_fingerprint Wafv2WebAclRuleA#ja3_fingerprint}
	Ja3Fingerprint interface{} `field:"optional" json:"ja3Fingerprint" yaml:"ja3Fingerprint"`
	// ja4_fingerprint block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.55.0/docs/resources/wafv2_web_acl_rule#ja4_fingerprint Wafv2WebAclRuleA#ja4_fingerprint}
	Ja4Fingerprint interface{} `field:"optional" json:"ja4Fingerprint" yaml:"ja4Fingerprint"`
	// label_namespace block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.55.0/docs/resources/wafv2_web_acl_rule#label_namespace Wafv2WebAclRuleA#label_namespace}
	LabelNamespace interface{} `field:"optional" json:"labelNamespace" yaml:"labelNamespace"`
	// query_argument block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.55.0/docs/resources/wafv2_web_acl_rule#query_argument Wafv2WebAclRuleA#query_argument}
	QueryArgument interface{} `field:"optional" json:"queryArgument" yaml:"queryArgument"`
	// query_string block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.55.0/docs/resources/wafv2_web_acl_rule#query_string Wafv2WebAclRuleA#query_string}
	QueryString interface{} `field:"optional" json:"queryString" yaml:"queryString"`
	// uri_path block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.55.0/docs/resources/wafv2_web_acl_rule#uri_path Wafv2WebAclRuleA#uri_path}
	UriPath interface{} `field:"optional" json:"uriPath" yaml:"uriPath"`
}

