// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package guarddutymemberdetectorfeature

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GuarddutyMemberDetectorFeatureConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.49.0/docs/resources/guardduty_member_detector_feature#account_id GuarddutyMemberDetectorFeature#account_id}.
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.49.0/docs/resources/guardduty_member_detector_feature#detector_id GuarddutyMemberDetectorFeature#detector_id}.
	DetectorId *string `field:"required" json:"detectorId" yaml:"detectorId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.49.0/docs/resources/guardduty_member_detector_feature#name GuarddutyMemberDetectorFeature#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.49.0/docs/resources/guardduty_member_detector_feature#status GuarddutyMemberDetectorFeature#status}.
	Status *string `field:"required" json:"status" yaml:"status"`
	// additional_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.49.0/docs/resources/guardduty_member_detector_feature#additional_configuration GuarddutyMemberDetectorFeature#additional_configuration}
	AdditionalConfiguration interface{} `field:"optional" json:"additionalConfiguration" yaml:"additionalConfiguration"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.49.0/docs/resources/guardduty_member_detector_feature#region GuarddutyMemberDetectorFeature#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
}

