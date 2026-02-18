// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package backupplan

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BackupPlanRuleScanActionList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (b *jsiiProxy_BackupPlanRuleScanActionList) validateGetParameters(index *float64) error {
	return nil
}

func (b *jsiiProxy_BackupPlanRuleScanActionList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_BackupPlanRuleScanActionList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_BackupPlanRuleScanActionList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_BackupPlanRuleScanActionList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_BackupPlanRuleScanActionList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewBackupPlanRuleScanActionListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

