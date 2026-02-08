// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build !no_runtime_type_checking

package arcregionswitchplan

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigOutputReference) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigOutputReference) validatePutAsgParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*ArcregionswitchPlanWorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigAsg:
		value := value.(*[]*ArcregionswitchPlanWorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigAsg)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*ArcregionswitchPlanWorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigAsg:
		value_ := value.([]*ArcregionswitchPlanWorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigAsg)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*ArcregionswitchPlanWorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigAsg; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigOutputReference) validatePutUngracefulParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*ArcregionswitchPlanWorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigUngraceful:
		value := value.(*[]*ArcregionswitchPlanWorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigUngraceful)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*ArcregionswitchPlanWorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigUngraceful:
		value_ := value.([]*ArcregionswitchPlanWorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigUngraceful)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*ArcregionswitchPlanWorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigUngraceful; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigOutputReference) validateResolveParameters(context cdktf.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigOutputReference) validateSetCapacityMonitoringApproachParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
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

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigOutputReference) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktf.IResolvable:
		// ok
	case *ArcregionswitchPlanWorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfig:
		val := val.(*ArcregionswitchPlanWorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfig)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case ArcregionswitchPlanWorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfig:
		val_ := val.(ArcregionswitchPlanWorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfig)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktf.IResolvable, *ArcregionswitchPlanWorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfig; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigOutputReference) validateSetTargetPercentParameters(val *float64) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigOutputReference) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigOutputReference) validateSetTimeoutMinutesParameters(val *float64) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewArcregionswitchPlanWorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigOutputReferenceParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) error {
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

