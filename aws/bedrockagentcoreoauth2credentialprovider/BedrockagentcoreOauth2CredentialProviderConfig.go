// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreoauth2credentialprovider

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BedrockagentcoreOauth2CredentialProviderConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.47.0/docs/resources/bedrockagentcore_oauth2_credential_provider#credential_provider_vendor BedrockagentcoreOauth2CredentialProvider#credential_provider_vendor}.
	CredentialProviderVendor *string `field:"required" json:"credentialProviderVendor" yaml:"credentialProviderVendor"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.47.0/docs/resources/bedrockagentcore_oauth2_credential_provider#name BedrockagentcoreOauth2CredentialProvider#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// oauth2_provider_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.47.0/docs/resources/bedrockagentcore_oauth2_credential_provider#oauth2_provider_config BedrockagentcoreOauth2CredentialProvider#oauth2_provider_config}
	Oauth2ProviderConfig interface{} `field:"optional" json:"oauth2ProviderConfig" yaml:"oauth2ProviderConfig"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.47.0/docs/resources/bedrockagentcore_oauth2_credential_provider#region BedrockagentcoreOauth2CredentialProvider#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.47.0/docs/resources/bedrockagentcore_oauth2_credential_provider#tags BedrockagentcoreOauth2CredentialProvider#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

