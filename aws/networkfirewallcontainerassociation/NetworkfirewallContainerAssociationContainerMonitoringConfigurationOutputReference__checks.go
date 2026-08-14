// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build !no_runtime_type_checking

package networkfirewallcontainerassociation

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

func (n *jsiiProxy_NetworkfirewallContainerAssociationContainerMonitoringConfigurationOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (n *jsiiProxy_NetworkfirewallContainerAssociationContainerMonitoringConfigurationOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (n *jsiiProxy_NetworkfirewallContainerAssociationContainerMonitoringConfigurationOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (n *jsiiProxy_NetworkfirewallContainerAssociationContainerMonitoringConfigurationOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (n *jsiiProxy_NetworkfirewallContainerAssociationContainerMonitoringConfigurationOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (n *jsiiProxy_NetworkfirewallContainerAssociationContainerMonitoringConfigurationOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (n *jsiiProxy_NetworkfirewallContainerAssociationContainerMonitoringConfigurationOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (n *jsiiProxy_NetworkfirewallContainerAssociationContainerMonitoringConfigurationOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (n *jsiiProxy_NetworkfirewallContainerAssociationContainerMonitoringConfigurationOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (n *jsiiProxy_NetworkfirewallContainerAssociationContainerMonitoringConfigurationOutputReference) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (n *jsiiProxy_NetworkfirewallContainerAssociationContainerMonitoringConfigurationOutputReference) validatePutAttributeFilterParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*NetworkfirewallContainerAssociationContainerMonitoringConfigurationAttributeFilter:
		value := value.(*[]*NetworkfirewallContainerAssociationContainerMonitoringConfigurationAttributeFilter)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*NetworkfirewallContainerAssociationContainerMonitoringConfigurationAttributeFilter:
		value_ := value.([]*NetworkfirewallContainerAssociationContainerMonitoringConfigurationAttributeFilter)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*NetworkfirewallContainerAssociationContainerMonitoringConfigurationAttributeFilter; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (n *jsiiProxy_NetworkfirewallContainerAssociationContainerMonitoringConfigurationOutputReference) validateResolveParameters(context cdktn.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_NetworkfirewallContainerAssociationContainerMonitoringConfigurationOutputReference) validateSetClusterArnParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_NetworkfirewallContainerAssociationContainerMonitoringConfigurationOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
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

func (j *jsiiProxy_NetworkfirewallContainerAssociationContainerMonitoringConfigurationOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_NetworkfirewallContainerAssociationContainerMonitoringConfigurationOutputReference) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktn.IResolvable:
		// ok
	case *NetworkfirewallContainerAssociationContainerMonitoringConfiguration:
		val := val.(*NetworkfirewallContainerAssociationContainerMonitoringConfiguration)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case NetworkfirewallContainerAssociationContainerMonitoringConfiguration:
		val_ := val.(NetworkfirewallContainerAssociationContainerMonitoringConfiguration)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktn.IResolvable, *NetworkfirewallContainerAssociationContainerMonitoringConfiguration; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_NetworkfirewallContainerAssociationContainerMonitoringConfigurationOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_NetworkfirewallContainerAssociationContainerMonitoringConfigurationOutputReference) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewNetworkfirewallContainerAssociationContainerMonitoringConfigurationOutputReferenceParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) error {
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

