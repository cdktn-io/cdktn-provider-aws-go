// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build !no_runtime_type_checking

package sagemakerhyperparametertuningjob

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

func (s *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionHyperParameterRangesOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionHyperParameterRangesOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionHyperParameterRangesOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionHyperParameterRangesOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionHyperParameterRangesOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionHyperParameterRangesOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionHyperParameterRangesOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionHyperParameterRangesOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionHyperParameterRangesOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionHyperParameterRangesOutputReference) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionHyperParameterRangesOutputReference) validatePutAutoParametersParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SagemakerHyperParameterTuningJobTrainingJobDefinitionHyperParameterRangesAutoParameters:
		value := value.(*[]*SagemakerHyperParameterTuningJobTrainingJobDefinitionHyperParameterRangesAutoParameters)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SagemakerHyperParameterTuningJobTrainingJobDefinitionHyperParameterRangesAutoParameters:
		value_ := value.([]*SagemakerHyperParameterTuningJobTrainingJobDefinitionHyperParameterRangesAutoParameters)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SagemakerHyperParameterTuningJobTrainingJobDefinitionHyperParameterRangesAutoParameters; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionHyperParameterRangesOutputReference) validatePutCategoricalParameterRangesParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SagemakerHyperParameterTuningJobTrainingJobDefinitionHyperParameterRangesCategoricalParameterRanges:
		value := value.(*[]*SagemakerHyperParameterTuningJobTrainingJobDefinitionHyperParameterRangesCategoricalParameterRanges)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SagemakerHyperParameterTuningJobTrainingJobDefinitionHyperParameterRangesCategoricalParameterRanges:
		value_ := value.([]*SagemakerHyperParameterTuningJobTrainingJobDefinitionHyperParameterRangesCategoricalParameterRanges)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SagemakerHyperParameterTuningJobTrainingJobDefinitionHyperParameterRangesCategoricalParameterRanges; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionHyperParameterRangesOutputReference) validatePutContinuousParameterRangesParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SagemakerHyperParameterTuningJobTrainingJobDefinitionHyperParameterRangesContinuousParameterRanges:
		value := value.(*[]*SagemakerHyperParameterTuningJobTrainingJobDefinitionHyperParameterRangesContinuousParameterRanges)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SagemakerHyperParameterTuningJobTrainingJobDefinitionHyperParameterRangesContinuousParameterRanges:
		value_ := value.([]*SagemakerHyperParameterTuningJobTrainingJobDefinitionHyperParameterRangesContinuousParameterRanges)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SagemakerHyperParameterTuningJobTrainingJobDefinitionHyperParameterRangesContinuousParameterRanges; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionHyperParameterRangesOutputReference) validatePutIntegerParameterRangesParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SagemakerHyperParameterTuningJobTrainingJobDefinitionHyperParameterRangesIntegerParameterRanges:
		value := value.(*[]*SagemakerHyperParameterTuningJobTrainingJobDefinitionHyperParameterRangesIntegerParameterRanges)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SagemakerHyperParameterTuningJobTrainingJobDefinitionHyperParameterRangesIntegerParameterRanges:
		value_ := value.([]*SagemakerHyperParameterTuningJobTrainingJobDefinitionHyperParameterRangesIntegerParameterRanges)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SagemakerHyperParameterTuningJobTrainingJobDefinitionHyperParameterRangesIntegerParameterRanges; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionHyperParameterRangesOutputReference) validateResolveParameters(context cdktn.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionHyperParameterRangesOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
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

func (j *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionHyperParameterRangesOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionHyperParameterRangesOutputReference) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktn.IResolvable:
		// ok
	case *SagemakerHyperParameterTuningJobTrainingJobDefinitionHyperParameterRanges:
		val := val.(*SagemakerHyperParameterTuningJobTrainingJobDefinitionHyperParameterRanges)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case SagemakerHyperParameterTuningJobTrainingJobDefinitionHyperParameterRanges:
		val_ := val.(SagemakerHyperParameterTuningJobTrainingJobDefinitionHyperParameterRanges)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktn.IResolvable, *SagemakerHyperParameterTuningJobTrainingJobDefinitionHyperParameterRanges; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionHyperParameterRangesOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_SagemakerHyperParameterTuningJobTrainingJobDefinitionHyperParameterRangesOutputReference) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewSagemakerHyperParameterTuningJobTrainingJobDefinitionHyperParameterRangesOutputReferenceParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) error {
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

