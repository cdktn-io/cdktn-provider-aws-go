// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package flowlog

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FlowLogTagFieldSpecificationList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (f *jsiiProxy_FlowLogTagFieldSpecificationList) validateGetParameters(index *float64) error {
	return nil
}

func (f *jsiiProxy_FlowLogTagFieldSpecificationList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_FlowLogTagFieldSpecificationList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_FlowLogTagFieldSpecificationList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_FlowLogTagFieldSpecificationList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_FlowLogTagFieldSpecificationList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewFlowLogTagFieldSpecificationListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

