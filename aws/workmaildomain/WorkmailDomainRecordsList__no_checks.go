// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package workmaildomain

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WorkmailDomainRecordsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (w *jsiiProxy_WorkmailDomainRecordsList) validateGetParameters(index *float64) error {
	return nil
}

func (w *jsiiProxy_WorkmailDomainRecordsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_WorkmailDomainRecordsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_WorkmailDomainRecordsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_WorkmailDomainRecordsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewWorkmailDomainRecordsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

