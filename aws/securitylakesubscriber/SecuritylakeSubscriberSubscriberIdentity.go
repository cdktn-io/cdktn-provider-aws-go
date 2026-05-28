// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securitylakesubscriber


type SecuritylakeSubscriberSubscriberIdentity struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.47.0/docs/resources/securitylake_subscriber#external_id SecuritylakeSubscriber#external_id}.
	ExternalId *string `field:"required" json:"externalId" yaml:"externalId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.47.0/docs/resources/securitylake_subscriber#principal SecuritylakeSubscriber#principal}.
	Principal *string `field:"required" json:"principal" yaml:"principal"`
}

