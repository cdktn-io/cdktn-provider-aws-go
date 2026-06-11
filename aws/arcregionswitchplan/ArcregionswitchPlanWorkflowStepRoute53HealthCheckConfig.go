// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package arcregionswitchplan


type ArcregionswitchPlanWorkflowStepRoute53HealthCheckConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/arcregionswitch_plan#hosted_zone_id ArcregionswitchPlan#hosted_zone_id}.
	HostedZoneId *string `field:"required" json:"hostedZoneId" yaml:"hostedZoneId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/arcregionswitch_plan#record_name ArcregionswitchPlan#record_name}.
	RecordName *string `field:"required" json:"recordName" yaml:"recordName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/arcregionswitch_plan#cross_account_role ArcregionswitchPlan#cross_account_role}.
	CrossAccountRole *string `field:"optional" json:"crossAccountRole" yaml:"crossAccountRole"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/arcregionswitch_plan#external_id ArcregionswitchPlan#external_id}.
	ExternalId *string `field:"optional" json:"externalId" yaml:"externalId"`
	// record_set block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/arcregionswitch_plan#record_set ArcregionswitchPlan#record_set}
	RecordSet interface{} `field:"optional" json:"recordSet" yaml:"recordSet"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/arcregionswitch_plan#timeout_minutes ArcregionswitchPlan#timeout_minutes}.
	TimeoutMinutes *float64 `field:"optional" json:"timeoutMinutes" yaml:"timeoutMinutes"`
}

