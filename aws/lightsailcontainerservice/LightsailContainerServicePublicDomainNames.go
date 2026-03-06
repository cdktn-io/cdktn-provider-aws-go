// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lightsailcontainerservice


type LightsailContainerServicePublicDomainNames struct {
	// certificate block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.35.1/docs/resources/lightsail_container_service#certificate LightsailContainerService#certificate}
	Certificate interface{} `field:"required" json:"certificate" yaml:"certificate"`
}

