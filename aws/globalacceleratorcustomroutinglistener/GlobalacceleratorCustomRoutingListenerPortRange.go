// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package globalacceleratorcustomroutinglistener


type GlobalacceleratorCustomRoutingListenerPortRange struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/globalaccelerator_custom_routing_listener#from_port GlobalacceleratorCustomRoutingListener#from_port}.
	FromPort *float64 `field:"optional" json:"fromPort" yaml:"fromPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/globalaccelerator_custom_routing_listener#to_port GlobalacceleratorCustomRoutingListener#to_port}.
	ToPort *float64 `field:"optional" json:"toPort" yaml:"toPort"`
}

