// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build !no_runtime_type_checking

package dataawscognitouserpool

import (
	"fmt"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

func (d *jsiiProxy_DataAwsCognitoUserPoolUserPoolAddOnsAdvancedSecurityAdditionalFlowsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	if mapKeyAttributeName == nil {
		return fmt.Errorf("parameter mapKeyAttributeName is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DataAwsCognitoUserPoolUserPoolAddOnsAdvancedSecurityAdditionalFlowsList) validateGetParameters(index *float64) error {
	if index == nil {
		return fmt.Errorf("parameter index is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DataAwsCognitoUserPoolUserPoolAddOnsAdvancedSecurityAdditionalFlowsList) validateResolveParameters(context cdktn.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DataAwsCognitoUserPoolUserPoolAddOnsAdvancedSecurityAdditionalFlowsList) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DataAwsCognitoUserPoolUserPoolAddOnsAdvancedSecurityAdditionalFlowsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DataAwsCognitoUserPoolUserPoolAddOnsAdvancedSecurityAdditionalFlowsList) validateSetWrapsSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewDataAwsCognitoUserPoolUserPoolAddOnsAdvancedSecurityAdditionalFlowsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	if wrapsSet == nil {
		return fmt.Errorf("parameter wrapsSet is required, but nil was provided")
	}

	return nil
}

