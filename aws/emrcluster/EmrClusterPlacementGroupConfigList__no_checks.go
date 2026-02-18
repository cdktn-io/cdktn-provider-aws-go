// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package emrcluster

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EmrClusterPlacementGroupConfigList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (e *jsiiProxy_EmrClusterPlacementGroupConfigList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_EmrClusterPlacementGroupConfigList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_EmrClusterPlacementGroupConfigList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_EmrClusterPlacementGroupConfigList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_EmrClusterPlacementGroupConfigList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_EmrClusterPlacementGroupConfigList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEmrClusterPlacementGroupConfigListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

