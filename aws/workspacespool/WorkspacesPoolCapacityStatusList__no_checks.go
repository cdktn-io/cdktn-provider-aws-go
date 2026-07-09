// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package workspacespool

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WorkspacesPoolCapacityStatusList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (w *jsiiProxy_WorkspacesPoolCapacityStatusList) validateGetParameters(index *float64) error {
	return nil
}

func (w *jsiiProxy_WorkspacesPoolCapacityStatusList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_WorkspacesPoolCapacityStatusList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_WorkspacesPoolCapacityStatusList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_WorkspacesPoolCapacityStatusList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewWorkspacesPoolCapacityStatusListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

