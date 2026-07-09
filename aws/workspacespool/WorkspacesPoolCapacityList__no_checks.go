// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package workspacespool

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WorkspacesPoolCapacityList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (w *jsiiProxy_WorkspacesPoolCapacityList) validateGetParameters(index *float64) error {
	return nil
}

func (w *jsiiProxy_WorkspacesPoolCapacityList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_WorkspacesPoolCapacityList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_WorkspacesPoolCapacityList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_WorkspacesPoolCapacityList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_WorkspacesPoolCapacityList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewWorkspacesPoolCapacityListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

