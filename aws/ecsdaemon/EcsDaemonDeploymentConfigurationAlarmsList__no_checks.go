// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package ecsdaemon

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EcsDaemonDeploymentConfigurationAlarmsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (e *jsiiProxy_EcsDaemonDeploymentConfigurationAlarmsList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_EcsDaemonDeploymentConfigurationAlarmsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_EcsDaemonDeploymentConfigurationAlarmsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_EcsDaemonDeploymentConfigurationAlarmsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_EcsDaemonDeploymentConfigurationAlarmsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_EcsDaemonDeploymentConfigurationAlarmsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEcsDaemonDeploymentConfigurationAlarmsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

