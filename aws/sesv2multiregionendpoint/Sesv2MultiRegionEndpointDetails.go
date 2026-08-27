// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sesv2multiregionendpoint


type Sesv2MultiRegionEndpointDetails struct {
	// routes_details block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sesv2_multi_region_endpoint#routes_details Sesv2MultiRegionEndpoint#routes_details}
	RoutesDetails interface{} `field:"optional" json:"routesDetails" yaml:"routesDetails"`
}

