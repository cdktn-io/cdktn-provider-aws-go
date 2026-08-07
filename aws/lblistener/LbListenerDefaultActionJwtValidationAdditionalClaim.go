// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lblistener


type LbListenerDefaultActionJwtValidationAdditionalClaim struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.58.0/docs/resources/lb_listener#format LbListener#format}.
	Format *string `field:"required" json:"format" yaml:"format"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.58.0/docs/resources/lb_listener#name LbListener#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.58.0/docs/resources/lb_listener#values LbListener#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

