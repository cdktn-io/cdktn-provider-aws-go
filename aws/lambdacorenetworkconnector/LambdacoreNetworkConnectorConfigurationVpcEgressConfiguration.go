// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lambdacorenetworkconnector


type LambdacoreNetworkConnectorConfigurationVpcEgressConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.61.0/docs/resources/lambdacore_network_connector#associated_compute_resource_types LambdacoreNetworkConnector#associated_compute_resource_types}.
	AssociatedComputeResourceTypes *[]*string `field:"required" json:"associatedComputeResourceTypes" yaml:"associatedComputeResourceTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.61.0/docs/resources/lambdacore_network_connector#security_group_ids LambdacoreNetworkConnector#security_group_ids}.
	SecurityGroupIds *[]*string `field:"required" json:"securityGroupIds" yaml:"securityGroupIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.61.0/docs/resources/lambdacore_network_connector#subnet_ids LambdacoreNetworkConnector#subnet_ids}.
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.61.0/docs/resources/lambdacore_network_connector#network_protocol LambdacoreNetworkConnector#network_protocol}.
	NetworkProtocol *string `field:"optional" json:"networkProtocol" yaml:"networkProtocol"`
}

