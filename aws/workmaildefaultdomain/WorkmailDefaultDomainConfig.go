// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package workmaildefaultdomain

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type WorkmailDefaultDomainConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Mail domain name to set as the default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.47.0/docs/resources/workmail_default_domain#domain_name WorkmailDefaultDomain#domain_name}
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// Identifier of the WorkMail organization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.47.0/docs/resources/workmail_default_domain#organization_id WorkmailDefaultDomain#organization_id}
	OrganizationId *string `field:"required" json:"organizationId" yaml:"organizationId"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.47.0/docs/resources/workmail_default_domain#region WorkmailDefaultDomain#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
}

