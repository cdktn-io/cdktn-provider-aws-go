// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package dataawsecsservice

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAwsEcsServiceLoadBalancerList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DataAwsEcsServiceLoadBalancerList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAwsEcsServiceLoadBalancerList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAwsEcsServiceLoadBalancerList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAwsEcsServiceLoadBalancerList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAwsEcsServiceLoadBalancerList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAwsEcsServiceLoadBalancerListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

