// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package guarddutymemberdetectorfeature


type GuarddutyMemberDetectorFeatureAdditionalConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.39.0/docs/resources/guardduty_member_detector_feature#name GuarddutyMemberDetectorFeature#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.39.0/docs/resources/guardduty_member_detector_feature#status GuarddutyMemberDetectorFeature#status}.
	Status *string `field:"required" json:"status" yaml:"status"`
}

