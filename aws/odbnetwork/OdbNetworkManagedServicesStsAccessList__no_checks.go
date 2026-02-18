// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package odbnetwork

// Building without runtime type checking enabled, so all the below just return nil

func (o *jsiiProxy_OdbNetworkManagedServicesStsAccessList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (o *jsiiProxy_OdbNetworkManagedServicesStsAccessList) validateGetParameters(index *float64) error {
	return nil
}

func (o *jsiiProxy_OdbNetworkManagedServicesStsAccessList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_OdbNetworkManagedServicesStsAccessList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_OdbNetworkManagedServicesStsAccessList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_OdbNetworkManagedServicesStsAccessList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewOdbNetworkManagedServicesStsAccessListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

