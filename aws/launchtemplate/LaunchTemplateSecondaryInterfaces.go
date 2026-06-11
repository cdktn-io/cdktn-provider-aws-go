// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package launchtemplate


type LaunchTemplateSecondaryInterfaces struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/launch_template#delete_on_termination LaunchTemplate#delete_on_termination}.
	DeleteOnTermination interface{} `field:"optional" json:"deleteOnTermination" yaml:"deleteOnTermination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/launch_template#device_index LaunchTemplate#device_index}.
	DeviceIndex *float64 `field:"optional" json:"deviceIndex" yaml:"deviceIndex"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/launch_template#interface_type LaunchTemplate#interface_type}.
	InterfaceType *string `field:"optional" json:"interfaceType" yaml:"interfaceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/launch_template#network_card_index LaunchTemplate#network_card_index}.
	NetworkCardIndex *float64 `field:"optional" json:"networkCardIndex" yaml:"networkCardIndex"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/launch_template#private_ip_address_count LaunchTemplate#private_ip_address_count}.
	PrivateIpAddressCount *float64 `field:"optional" json:"privateIpAddressCount" yaml:"privateIpAddressCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/launch_template#private_ip_addresses LaunchTemplate#private_ip_addresses}.
	PrivateIpAddresses *[]*string `field:"optional" json:"privateIpAddresses" yaml:"privateIpAddresses"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/launch_template#secondary_subnet_id LaunchTemplate#secondary_subnet_id}.
	SecondarySubnetId *string `field:"optional" json:"secondarySubnetId" yaml:"secondarySubnetId"`
}

