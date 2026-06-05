// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wafv2webaclrulegroupassociation


type Wafv2WebAclRuleGroupAssociationManagedRuleGroupManagedRuleGroupConfigsAwsManagedRulesAcfpRuleSet struct {
	// Path to the account creation endpoint on the protected website.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.49.0/docs/resources/wafv2_web_acl_rule_group_association#creation_path Wafv2WebAclRuleGroupAssociation#creation_path}
	CreationPath *string `field:"required" json:"creationPath" yaml:"creationPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.49.0/docs/resources/wafv2_web_acl_rule_group_association#registration_page_path Wafv2WebAclRuleGroupAssociation#registration_page_path}.
	RegistrationPagePath *string `field:"required" json:"registrationPagePath" yaml:"registrationPagePath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.49.0/docs/resources/wafv2_web_acl_rule_group_association#enable_regex_in_path Wafv2WebAclRuleGroupAssociation#enable_regex_in_path}.
	EnableRegexInPath interface{} `field:"optional" json:"enableRegexInPath" yaml:"enableRegexInPath"`
	// request_inspection block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.49.0/docs/resources/wafv2_web_acl_rule_group_association#request_inspection Wafv2WebAclRuleGroupAssociation#request_inspection}
	RequestInspection interface{} `field:"optional" json:"requestInspection" yaml:"requestInspection"`
	// response_inspection block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.49.0/docs/resources/wafv2_web_acl_rule_group_association#response_inspection Wafv2WebAclRuleGroupAssociation#response_inspection}
	ResponseInspection interface{} `field:"optional" json:"responseInspection" yaml:"responseInspection"`
}

