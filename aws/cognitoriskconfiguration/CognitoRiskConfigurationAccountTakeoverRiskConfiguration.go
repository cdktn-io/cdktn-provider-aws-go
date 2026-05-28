// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cognitoriskconfiguration


type CognitoRiskConfigurationAccountTakeoverRiskConfiguration struct {
	// actions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.47.0/docs/resources/cognito_risk_configuration#actions CognitoRiskConfiguration#actions}
	Actions *CognitoRiskConfigurationAccountTakeoverRiskConfigurationActions `field:"required" json:"actions" yaml:"actions"`
	// notify_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.47.0/docs/resources/cognito_risk_configuration#notify_configuration CognitoRiskConfiguration#notify_configuration}
	NotifyConfiguration *CognitoRiskConfigurationAccountTakeoverRiskConfigurationNotifyConfiguration `field:"optional" json:"notifyConfiguration" yaml:"notifyConfiguration"`
}

