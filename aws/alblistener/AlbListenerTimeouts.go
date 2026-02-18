// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package alblistener


type AlbListenerTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.32.1/docs/resources/alb_listener#create AlbListener#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.32.1/docs/resources/alb_listener#update AlbListener#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

