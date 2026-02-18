// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build !no_runtime_type_checking

package dataawstimestreamwritetable

import (
	"fmt"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

func (d *jsiiProxy_DataAwsTimestreamwriteTableMagneticStoreWritePropertiesMagneticStoreRejectedDataLocationS3ConfigurationList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	if mapKeyAttributeName == nil {
		return fmt.Errorf("parameter mapKeyAttributeName is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DataAwsTimestreamwriteTableMagneticStoreWritePropertiesMagneticStoreRejectedDataLocationS3ConfigurationList) validateGetParameters(index *float64) error {
	if index == nil {
		return fmt.Errorf("parameter index is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DataAwsTimestreamwriteTableMagneticStoreWritePropertiesMagneticStoreRejectedDataLocationS3ConfigurationList) validateResolveParameters(context cdktn.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DataAwsTimestreamwriteTableMagneticStoreWritePropertiesMagneticStoreRejectedDataLocationS3ConfigurationList) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DataAwsTimestreamwriteTableMagneticStoreWritePropertiesMagneticStoreRejectedDataLocationS3ConfigurationList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DataAwsTimestreamwriteTableMagneticStoreWritePropertiesMagneticStoreRejectedDataLocationS3ConfigurationList) validateSetWrapsSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewDataAwsTimestreamwriteTableMagneticStoreWritePropertiesMagneticStoreRejectedDataLocationS3ConfigurationListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	if wrapsSet == nil {
		return fmt.Errorf("parameter wrapsSet is required, but nil was provided")
	}

	return nil
}

