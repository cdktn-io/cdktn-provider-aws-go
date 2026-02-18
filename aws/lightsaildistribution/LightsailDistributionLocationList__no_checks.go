// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package lightsaildistribution

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LightsailDistributionLocationList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (l *jsiiProxy_LightsailDistributionLocationList) validateGetParameters(index *float64) error {
	return nil
}

func (l *jsiiProxy_LightsailDistributionLocationList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_LightsailDistributionLocationList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LightsailDistributionLocationList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_LightsailDistributionLocationList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewLightsailDistributionLocationListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

