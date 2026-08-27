// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package shieldsubscription

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ShieldSubscriptionConfig struct {
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
	// Whether to automatically renew the subscription when it expires.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/shield_subscription#auto_renew ShieldSubscription#auto_renew}
	AutoRenew *string `field:"optional" json:"autoRenew" yaml:"autoRenew"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/shield_subscription#skip_destroy ShieldSubscription#skip_destroy}.
	SkipDestroy interface{} `field:"optional" json:"skipDestroy" yaml:"skipDestroy"`
}

