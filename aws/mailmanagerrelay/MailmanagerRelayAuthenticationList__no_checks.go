// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package mailmanagerrelay

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_MailmanagerRelayAuthenticationList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (m *jsiiProxy_MailmanagerRelayAuthenticationList) validateGetParameters(index *float64) error {
	return nil
}

func (m *jsiiProxy_MailmanagerRelayAuthenticationList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_MailmanagerRelayAuthenticationList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_MailmanagerRelayAuthenticationList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_MailmanagerRelayAuthenticationList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_MailmanagerRelayAuthenticationList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewMailmanagerRelayAuthenticationListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

