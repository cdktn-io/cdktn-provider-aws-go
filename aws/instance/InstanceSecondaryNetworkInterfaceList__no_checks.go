// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package instance

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_InstanceSecondaryNetworkInterfaceList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (i *jsiiProxy_InstanceSecondaryNetworkInterfaceList) validateGetParameters(index *float64) error {
	return nil
}

func (i *jsiiProxy_InstanceSecondaryNetworkInterfaceList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_InstanceSecondaryNetworkInterfaceList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_InstanceSecondaryNetworkInterfaceList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_InstanceSecondaryNetworkInterfaceList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_InstanceSecondaryNetworkInterfaceList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewInstanceSecondaryNetworkInterfaceListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

