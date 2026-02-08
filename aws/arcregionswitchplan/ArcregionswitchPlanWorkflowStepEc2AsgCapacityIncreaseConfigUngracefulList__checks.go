// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build !no_runtime_type_checking

package arcregionswitchplan

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepEc2AsgCapacityIncreaseConfigUngracefulList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	if mapKeyAttributeName == nil {
		return fmt.Errorf("parameter mapKeyAttributeName is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepEc2AsgCapacityIncreaseConfigUngracefulList) validateGetParameters(index *float64) error {
	if index == nil {
		return fmt.Errorf("parameter index is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepEc2AsgCapacityIncreaseConfigUngracefulList) validateResolveParameters(context cdktf.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepEc2AsgCapacityIncreaseConfigUngracefulList) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*ArcregionswitchPlanWorkflowStepEc2AsgCapacityIncreaseConfigUngraceful:
		val := val.(*[]*ArcregionswitchPlanWorkflowStepEc2AsgCapacityIncreaseConfigUngraceful)
		for idx_97dfc6, v := range *val {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
				return err
			}
		}
	case []*ArcregionswitchPlanWorkflowStepEc2AsgCapacityIncreaseConfigUngraceful:
		val_ := val.([]*ArcregionswitchPlanWorkflowStepEc2AsgCapacityIncreaseConfigUngraceful)
		val := &val_
		for idx_97dfc6, v := range *val {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktf.IResolvable, *[]*ArcregionswitchPlanWorkflowStepEc2AsgCapacityIncreaseConfigUngraceful; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepEc2AsgCapacityIncreaseConfigUngracefulList) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepEc2AsgCapacityIncreaseConfigUngracefulList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepEc2AsgCapacityIncreaseConfigUngracefulList) validateSetWrapsSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewArcregionswitchPlanWorkflowStepEc2AsgCapacityIncreaseConfigUngracefulListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
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

