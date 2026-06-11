// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ivschatroom


type IvschatRoomMessageReviewHandler struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/ivschat_room#fallback_result IvschatRoom#fallback_result}.
	FallbackResult *string `field:"optional" json:"fallbackResult" yaml:"fallbackResult"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.50.0/docs/resources/ivschat_room#uri IvschatRoom#uri}.
	Uri *string `field:"optional" json:"uri" yaml:"uri"`
}

