// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package bedrockinferenceprofile

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BedrockInferenceProfileModelsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (b *jsiiProxy_BedrockInferenceProfileModelsList) validateGetParameters(index *float64) error {
	return nil
}

func (b *jsiiProxy_BedrockInferenceProfileModelsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_BedrockInferenceProfileModelsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_BedrockInferenceProfileModelsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_BedrockInferenceProfileModelsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewBedrockInferenceProfileModelsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

