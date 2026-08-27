// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package spotinstancerequest


type SpotInstanceRequestSecondaryNetworkInterface struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#network_card_index SpotInstanceRequest#network_card_index}.
	NetworkCardIndex *float64 `field:"required" json:"networkCardIndex" yaml:"networkCardIndex"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#secondary_subnet_id SpotInstanceRequest#secondary_subnet_id}.
	SecondarySubnetId *string `field:"required" json:"secondarySubnetId" yaml:"secondarySubnetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#delete_on_termination SpotInstanceRequest#delete_on_termination}.
	DeleteOnTermination interface{} `field:"optional" json:"deleteOnTermination" yaml:"deleteOnTermination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#device_index SpotInstanceRequest#device_index}.
	DeviceIndex *float64 `field:"optional" json:"deviceIndex" yaml:"deviceIndex"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#interface_type SpotInstanceRequest#interface_type}.
	InterfaceType *string `field:"optional" json:"interfaceType" yaml:"interfaceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#private_ip_address_count SpotInstanceRequest#private_ip_address_count}.
	PrivateIpAddressCount *float64 `field:"optional" json:"privateIpAddressCount" yaml:"privateIpAddressCount"`
}

