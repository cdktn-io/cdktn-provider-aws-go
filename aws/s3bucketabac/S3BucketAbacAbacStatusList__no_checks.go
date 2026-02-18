// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package s3bucketabac

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_S3BucketAbacAbacStatusList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_S3BucketAbacAbacStatusList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_S3BucketAbacAbacStatusList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_S3BucketAbacAbacStatusList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_S3BucketAbacAbacStatusList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_S3BucketAbacAbacStatusList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_S3BucketAbacAbacStatusList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewS3BucketAbacAbacStatusListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

