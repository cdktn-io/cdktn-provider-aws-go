// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build !no_runtime_type_checking

package sagemakerlabelingjob

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func (s *jsiiProxy_SagemakerLabelingJobHumanTaskConfigPublicWorkforceTaskPriceAmountInUsdOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_SagemakerLabelingJobHumanTaskConfigPublicWorkforceTaskPriceAmountInUsdOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_SagemakerLabelingJobHumanTaskConfigPublicWorkforceTaskPriceAmountInUsdOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_SagemakerLabelingJobHumanTaskConfigPublicWorkforceTaskPriceAmountInUsdOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_SagemakerLabelingJobHumanTaskConfigPublicWorkforceTaskPriceAmountInUsdOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_SagemakerLabelingJobHumanTaskConfigPublicWorkforceTaskPriceAmountInUsdOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_SagemakerLabelingJobHumanTaskConfigPublicWorkforceTaskPriceAmountInUsdOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_SagemakerLabelingJobHumanTaskConfigPublicWorkforceTaskPriceAmountInUsdOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_SagemakerLabelingJobHumanTaskConfigPublicWorkforceTaskPriceAmountInUsdOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_SagemakerLabelingJobHumanTaskConfigPublicWorkforceTaskPriceAmountInUsdOutputReference) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_SagemakerLabelingJobHumanTaskConfigPublicWorkforceTaskPriceAmountInUsdOutputReference) validateResolveParameters(context cdktf.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_SagemakerLabelingJobHumanTaskConfigPublicWorkforceTaskPriceAmountInUsdOutputReference) validateSetCentsParameters(val *float64) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_SagemakerLabelingJobHumanTaskConfigPublicWorkforceTaskPriceAmountInUsdOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
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

func (j *jsiiProxy_SagemakerLabelingJobHumanTaskConfigPublicWorkforceTaskPriceAmountInUsdOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_SagemakerLabelingJobHumanTaskConfigPublicWorkforceTaskPriceAmountInUsdOutputReference) validateSetDollarsParameters(val *float64) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_SagemakerLabelingJobHumanTaskConfigPublicWorkforceTaskPriceAmountInUsdOutputReference) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktf.IResolvable:
		// ok
	case *SagemakerLabelingJobHumanTaskConfigPublicWorkforceTaskPriceAmountInUsd:
		val := val.(*SagemakerLabelingJobHumanTaskConfigPublicWorkforceTaskPriceAmountInUsd)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case SagemakerLabelingJobHumanTaskConfigPublicWorkforceTaskPriceAmountInUsd:
		val_ := val.(SagemakerLabelingJobHumanTaskConfigPublicWorkforceTaskPriceAmountInUsd)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktf.IResolvable, *SagemakerLabelingJobHumanTaskConfigPublicWorkforceTaskPriceAmountInUsd; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_SagemakerLabelingJobHumanTaskConfigPublicWorkforceTaskPriceAmountInUsdOutputReference) validateSetTenthFractionsOfACentParameters(val *float64) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_SagemakerLabelingJobHumanTaskConfigPublicWorkforceTaskPriceAmountInUsdOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_SagemakerLabelingJobHumanTaskConfigPublicWorkforceTaskPriceAmountInUsdOutputReference) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewSagemakerLabelingJobHumanTaskConfigPublicWorkforceTaskPriceAmountInUsdOutputReferenceParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) error {
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

