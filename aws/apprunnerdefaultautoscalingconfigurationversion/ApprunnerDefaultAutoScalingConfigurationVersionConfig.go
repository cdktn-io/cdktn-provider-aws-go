// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apprunnerdefaultautoscalingconfigurationversion

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ApprunnerDefaultAutoScalingConfigurationVersionConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/apprunner_default_auto_scaling_configuration_version#auto_scaling_configuration_arn ApprunnerDefaultAutoScalingConfigurationVersion#auto_scaling_configuration_arn}.
	AutoScalingConfigurationArn *string `field:"required" json:"autoScalingConfigurationArn" yaml:"autoScalingConfigurationArn"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/apprunner_default_auto_scaling_configuration_version#region ApprunnerDefaultAutoScalingConfigurationVersion#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
}

