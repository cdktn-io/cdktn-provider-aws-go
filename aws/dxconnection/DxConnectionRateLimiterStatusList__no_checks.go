// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package dxconnection

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DxConnectionRateLimiterStatusList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DxConnectionRateLimiterStatusList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DxConnectionRateLimiterStatusList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DxConnectionRateLimiterStatusList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DxConnectionRateLimiterStatusList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DxConnectionRateLimiterStatusList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDxConnectionRateLimiterStatusListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

