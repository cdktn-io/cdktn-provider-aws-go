// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package mskchannel

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_MskChannelEncryptionConfigurationList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (m *jsiiProxy_MskChannelEncryptionConfigurationList) validateGetParameters(index *float64) error {
	return nil
}

func (m *jsiiProxy_MskChannelEncryptionConfigurationList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_MskChannelEncryptionConfigurationList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_MskChannelEncryptionConfigurationList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_MskChannelEncryptionConfigurationList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_MskChannelEncryptionConfigurationList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewMskChannelEncryptionConfigurationListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

