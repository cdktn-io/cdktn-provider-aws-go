// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package transferwebapp


type TransferWebAppEndpointDetails struct {
	// vpc block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.35.0/docs/resources/transfer_web_app#vpc TransferWebApp#vpc}
	Vpc interface{} `field:"optional" json:"vpc" yaml:"vpc"`
}

