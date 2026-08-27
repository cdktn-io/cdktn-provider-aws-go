// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build !no_runtime_type_checking

package observabilityadmincentralizationrulefororganization

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

func (o *jsiiProxy_ObservabilityadminCentralizationRuleForOrganizationRuleDestinationDestinationLogsConfigurationOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (o *jsiiProxy_ObservabilityadminCentralizationRuleForOrganizationRuleDestinationDestinationLogsConfigurationOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (o *jsiiProxy_ObservabilityadminCentralizationRuleForOrganizationRuleDestinationDestinationLogsConfigurationOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (o *jsiiProxy_ObservabilityadminCentralizationRuleForOrganizationRuleDestinationDestinationLogsConfigurationOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (o *jsiiProxy_ObservabilityadminCentralizationRuleForOrganizationRuleDestinationDestinationLogsConfigurationOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (o *jsiiProxy_ObservabilityadminCentralizationRuleForOrganizationRuleDestinationDestinationLogsConfigurationOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (o *jsiiProxy_ObservabilityadminCentralizationRuleForOrganizationRuleDestinationDestinationLogsConfigurationOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (o *jsiiProxy_ObservabilityadminCentralizationRuleForOrganizationRuleDestinationDestinationLogsConfigurationOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (o *jsiiProxy_ObservabilityadminCentralizationRuleForOrganizationRuleDestinationDestinationLogsConfigurationOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (o *jsiiProxy_ObservabilityadminCentralizationRuleForOrganizationRuleDestinationDestinationLogsConfigurationOutputReference) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (o *jsiiProxy_ObservabilityadminCentralizationRuleForOrganizationRuleDestinationDestinationLogsConfigurationOutputReference) validatePutBackupConfigurationParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*ObservabilityadminCentralizationRuleForOrganizationRuleDestinationDestinationLogsConfigurationBackupConfiguration:
		value := value.(*[]*ObservabilityadminCentralizationRuleForOrganizationRuleDestinationDestinationLogsConfigurationBackupConfiguration)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*ObservabilityadminCentralizationRuleForOrganizationRuleDestinationDestinationLogsConfigurationBackupConfiguration:
		value_ := value.([]*ObservabilityadminCentralizationRuleForOrganizationRuleDestinationDestinationLogsConfigurationBackupConfiguration)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*ObservabilityadminCentralizationRuleForOrganizationRuleDestinationDestinationLogsConfigurationBackupConfiguration; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (o *jsiiProxy_ObservabilityadminCentralizationRuleForOrganizationRuleDestinationDestinationLogsConfigurationOutputReference) validatePutLogGroupNameConfigurationParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*ObservabilityadminCentralizationRuleForOrganizationRuleDestinationDestinationLogsConfigurationLogGroupNameConfiguration:
		value := value.(*[]*ObservabilityadminCentralizationRuleForOrganizationRuleDestinationDestinationLogsConfigurationLogGroupNameConfiguration)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*ObservabilityadminCentralizationRuleForOrganizationRuleDestinationDestinationLogsConfigurationLogGroupNameConfiguration:
		value_ := value.([]*ObservabilityadminCentralizationRuleForOrganizationRuleDestinationDestinationLogsConfigurationLogGroupNameConfiguration)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*ObservabilityadminCentralizationRuleForOrganizationRuleDestinationDestinationLogsConfigurationLogGroupNameConfiguration; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (o *jsiiProxy_ObservabilityadminCentralizationRuleForOrganizationRuleDestinationDestinationLogsConfigurationOutputReference) validatePutLogsEncryptionConfigurationParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*ObservabilityadminCentralizationRuleForOrganizationRuleDestinationDestinationLogsConfigurationLogsEncryptionConfiguration:
		value := value.(*[]*ObservabilityadminCentralizationRuleForOrganizationRuleDestinationDestinationLogsConfigurationLogsEncryptionConfiguration)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*ObservabilityadminCentralizationRuleForOrganizationRuleDestinationDestinationLogsConfigurationLogsEncryptionConfiguration:
		value_ := value.([]*ObservabilityadminCentralizationRuleForOrganizationRuleDestinationDestinationLogsConfigurationLogsEncryptionConfiguration)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*ObservabilityadminCentralizationRuleForOrganizationRuleDestinationDestinationLogsConfigurationLogsEncryptionConfiguration; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (o *jsiiProxy_ObservabilityadminCentralizationRuleForOrganizationRuleDestinationDestinationLogsConfigurationOutputReference) validatePutTagPropagationConfigurationParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*ObservabilityadminCentralizationRuleForOrganizationRuleDestinationDestinationLogsConfigurationTagPropagationConfiguration:
		value := value.(*[]*ObservabilityadminCentralizationRuleForOrganizationRuleDestinationDestinationLogsConfigurationTagPropagationConfiguration)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*ObservabilityadminCentralizationRuleForOrganizationRuleDestinationDestinationLogsConfigurationTagPropagationConfiguration:
		value_ := value.([]*ObservabilityadminCentralizationRuleForOrganizationRuleDestinationDestinationLogsConfigurationTagPropagationConfiguration)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*ObservabilityadminCentralizationRuleForOrganizationRuleDestinationDestinationLogsConfigurationTagPropagationConfiguration; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (o *jsiiProxy_ObservabilityadminCentralizationRuleForOrganizationRuleDestinationDestinationLogsConfigurationOutputReference) validateResolveParameters(context cdktn.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_ObservabilityadminCentralizationRuleForOrganizationRuleDestinationDestinationLogsConfigurationOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
	switch val.(type) {
	case *string:
		// ok
	case string:
		// ok
	case *float64:
		// ok
	case float64:
		// ok
	case *int:
		// ok
	case int:
		// ok
	case *uint:
		// ok
	case uint:
		// ok
	case *int8:
		// ok
	case int8:
		// ok
	case *int16:
		// ok
	case int16:
		// ok
	case *int32:
		// ok
	case int32:
		// ok
	case *int64:
		// ok
	case int64:
		// ok
	case *uint8:
		// ok
	case uint8:
		// ok
	case *uint16:
		// ok
	case uint16:
		// ok
	case *uint32:
		// ok
	case uint32:
		// ok
	case *uint64:
		// ok
	case uint64:
		// ok
	default:
		return fmt.Errorf("parameter val must be one of the allowed types: *string, *float64; received %#v (a %T)", val, val)
	}

	return nil
}

func (j *jsiiProxy_ObservabilityadminCentralizationRuleForOrganizationRuleDestinationDestinationLogsConfigurationOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_ObservabilityadminCentralizationRuleForOrganizationRuleDestinationDestinationLogsConfigurationOutputReference) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktn.IResolvable:
		// ok
	case *ObservabilityadminCentralizationRuleForOrganizationRuleDestinationDestinationLogsConfiguration:
		val := val.(*ObservabilityadminCentralizationRuleForOrganizationRuleDestinationDestinationLogsConfiguration)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case ObservabilityadminCentralizationRuleForOrganizationRuleDestinationDestinationLogsConfiguration:
		val_ := val.(ObservabilityadminCentralizationRuleForOrganizationRuleDestinationDestinationLogsConfiguration)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktn.IResolvable, *ObservabilityadminCentralizationRuleForOrganizationRuleDestinationDestinationLogsConfiguration; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_ObservabilityadminCentralizationRuleForOrganizationRuleDestinationDestinationLogsConfigurationOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_ObservabilityadminCentralizationRuleForOrganizationRuleDestinationDestinationLogsConfigurationOutputReference) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewObservabilityadminCentralizationRuleForOrganizationRuleDestinationDestinationLogsConfigurationOutputReferenceParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	if complexObjectIndex == nil {
		return fmt.Errorf("parameter complexObjectIndex is required, but nil was provided")
	}

	if complexObjectIsFromSet == nil {
		return fmt.Errorf("parameter complexObjectIsFromSet is required, but nil was provided")
	}

	return nil
}

