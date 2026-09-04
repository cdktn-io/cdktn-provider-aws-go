// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datazonepolicygrant


type DatazonePolicyGrantDetail struct {
	// add_to_project_member_pool block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/datazone_policy_grant#add_to_project_member_pool DatazonePolicyGrant#add_to_project_member_pool}
	AddToProjectMemberPool interface{} `field:"optional" json:"addToProjectMemberPool" yaml:"addToProjectMemberPool"`
	// create_asset_type block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/datazone_policy_grant#create_asset_type DatazonePolicyGrant#create_asset_type}
	CreateAssetType interface{} `field:"optional" json:"createAssetType" yaml:"createAssetType"`
	// create_domain_unit block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/datazone_policy_grant#create_domain_unit DatazonePolicyGrant#create_domain_unit}
	CreateDomainUnit interface{} `field:"optional" json:"createDomainUnit" yaml:"createDomainUnit"`
	// create_environment block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/datazone_policy_grant#create_environment DatazonePolicyGrant#create_environment}
	CreateEnvironment interface{} `field:"optional" json:"createEnvironment" yaml:"createEnvironment"`
	// create_environment_from_blueprint block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/datazone_policy_grant#create_environment_from_blueprint DatazonePolicyGrant#create_environment_from_blueprint}
	CreateEnvironmentFromBlueprint interface{} `field:"optional" json:"createEnvironmentFromBlueprint" yaml:"createEnvironmentFromBlueprint"`
	// create_environment_profile block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/datazone_policy_grant#create_environment_profile DatazonePolicyGrant#create_environment_profile}
	CreateEnvironmentProfile interface{} `field:"optional" json:"createEnvironmentProfile" yaml:"createEnvironmentProfile"`
	// create_form_type block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/datazone_policy_grant#create_form_type DatazonePolicyGrant#create_form_type}
	CreateFormType interface{} `field:"optional" json:"createFormType" yaml:"createFormType"`
	// create_glossary block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/datazone_policy_grant#create_glossary DatazonePolicyGrant#create_glossary}
	CreateGlossary interface{} `field:"optional" json:"createGlossary" yaml:"createGlossary"`
	// create_project block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/datazone_policy_grant#create_project DatazonePolicyGrant#create_project}
	CreateProject interface{} `field:"optional" json:"createProject" yaml:"createProject"`
	// create_project_from_project_profile block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/datazone_policy_grant#create_project_from_project_profile DatazonePolicyGrant#create_project_from_project_profile}
	CreateProjectFromProjectProfile interface{} `field:"optional" json:"createProjectFromProjectProfile" yaml:"createProjectFromProjectProfile"`
	// delegate_create_environment_profile block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/datazone_policy_grant#delegate_create_environment_profile DatazonePolicyGrant#delegate_create_environment_profile}
	DelegateCreateEnvironmentProfile interface{} `field:"optional" json:"delegateCreateEnvironmentProfile" yaml:"delegateCreateEnvironmentProfile"`
	// override_domain_unit_owners block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/datazone_policy_grant#override_domain_unit_owners DatazonePolicyGrant#override_domain_unit_owners}
	OverrideDomainUnitOwners interface{} `field:"optional" json:"overrideDomainUnitOwners" yaml:"overrideDomainUnitOwners"`
	// override_project_owners block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/datazone_policy_grant#override_project_owners DatazonePolicyGrant#override_project_owners}
	OverrideProjectOwners interface{} `field:"optional" json:"overrideProjectOwners" yaml:"overrideProjectOwners"`
	// use_asset_type block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/datazone_policy_grant#use_asset_type DatazonePolicyGrant#use_asset_type}
	UseAssetType interface{} `field:"optional" json:"useAssetType" yaml:"useAssetType"`
}

