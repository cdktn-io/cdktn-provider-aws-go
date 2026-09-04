// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilityadmincentralizationrulefororganization


type ObservabilityadminCentralizationRuleForOrganizationRuleDestinationDestinationLogsConfigurationTagPropagationConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/observabilityadmin_centralization_rule_for_organization#destination_role_arn ObservabilityadminCentralizationRuleForOrganization#destination_role_arn}.
	DestinationRoleArn *string `field:"required" json:"destinationRoleArn" yaml:"destinationRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/resources/observabilityadmin_centralization_rule_for_organization#tag_conflict_resolution_strategy ObservabilityadminCentralizationRuleForOrganization#tag_conflict_resolution_strategy}.
	TagConflictResolutionStrategy *string `field:"optional" json:"tagConflictResolutionStrategy" yaml:"tagConflictResolutionStrategy"`
}

