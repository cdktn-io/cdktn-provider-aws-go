// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package iotcacertificate

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IotCaCertificateValidityList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (i *jsiiProxy_IotCaCertificateValidityList) validateGetParameters(index *float64) error {
	return nil
}

func (i *jsiiProxy_IotCaCertificateValidityList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_IotCaCertificateValidityList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_IotCaCertificateValidityList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_IotCaCertificateValidityList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewIotCaCertificateValidityListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

