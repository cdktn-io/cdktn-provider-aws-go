// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package ecsdaemon

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EcsDaemonDeploymentConfigurationList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (e *jsiiProxy_EcsDaemonDeploymentConfigurationList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_EcsDaemonDeploymentConfigurationList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_EcsDaemonDeploymentConfigurationList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_EcsDaemonDeploymentConfigurationList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_EcsDaemonDeploymentConfigurationList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_EcsDaemonDeploymentConfigurationList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEcsDaemonDeploymentConfigurationListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

