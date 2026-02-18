// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package m2application

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_M2ApplicationDefinitionList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (m *jsiiProxy_M2ApplicationDefinitionList) validateGetParameters(index *float64) error {
	return nil
}

func (m *jsiiProxy_M2ApplicationDefinitionList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_M2ApplicationDefinitionList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_M2ApplicationDefinitionList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_M2ApplicationDefinitionList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_M2ApplicationDefinitionList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewM2ApplicationDefinitionListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

