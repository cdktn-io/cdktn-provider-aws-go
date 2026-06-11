// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package shieldapplicationlayerautomaticresponse

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ShieldApplicationLayerAutomaticResponseConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/shield_application_layer_automatic_response#action ShieldApplicationLayerAutomaticResponse#action}.
	Action *string `field:"required" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/shield_application_layer_automatic_response#resource_arn ShieldApplicationLayerAutomaticResponse#resource_arn}.
	ResourceArn *string `field:"required" json:"resourceArn" yaml:"resourceArn"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/shield_application_layer_automatic_response#timeouts ShieldApplicationLayerAutomaticResponse#timeouts}
	Timeouts *ShieldApplicationLayerAutomaticResponseTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

