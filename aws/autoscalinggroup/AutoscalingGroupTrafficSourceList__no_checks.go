// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package autoscalinggroup

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AutoscalingGroupTrafficSourceList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AutoscalingGroupTrafficSourceList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AutoscalingGroupTrafficSourceList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AutoscalingGroupTrafficSourceList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AutoscalingGroupTrafficSourceList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AutoscalingGroupTrafficSourceList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AutoscalingGroupTrafficSourceList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAutoscalingGroupTrafficSourceListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

