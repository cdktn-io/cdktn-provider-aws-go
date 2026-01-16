// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package transferwebapp


type TransferWebAppEndpointDetails struct {
	// vpc block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/transfer_web_app#vpc TransferWebApp#vpc}
	Vpc interface{} `field:"optional" json:"vpc" yaml:"vpc"`
}

