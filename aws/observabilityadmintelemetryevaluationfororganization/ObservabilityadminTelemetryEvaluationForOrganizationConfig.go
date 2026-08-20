// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilityadmintelemetryevaluationfororganization

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ObservabilityadminTelemetryEvaluationForOrganizationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.61.0/docs/resources/observabilityadmin_telemetry_evaluation_for_organization#all_regions ObservabilityadminTelemetryEvaluationForOrganization#all_regions}.
	AllRegions interface{} `field:"optional" json:"allRegions" yaml:"allRegions"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.61.0/docs/resources/observabilityadmin_telemetry_evaluation_for_organization#region ObservabilityadminTelemetryEvaluationForOrganization#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.61.0/docs/resources/observabilityadmin_telemetry_evaluation_for_organization#regions ObservabilityadminTelemetryEvaluationForOrganization#regions}.
	Regions *[]*string `field:"optional" json:"regions" yaml:"regions"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.61.0/docs/resources/observabilityadmin_telemetry_evaluation_for_organization#timeouts ObservabilityadminTelemetryEvaluationForOrganization#timeouts}
	Timeouts *ObservabilityadminTelemetryEvaluationForOrganizationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

