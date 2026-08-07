// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build !no_runtime_type_checking

package prometheusanomalydetector

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

func (p *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestOutputReference) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestOutputReference) validatePutIgnoreNearExpectedFromAboveParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromAbove:
		value := value.(*[]*PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromAbove)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromAbove:
		value_ := value.([]*PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromAbove)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromAbove; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (p *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestOutputReference) validatePutIgnoreNearExpectedFromBelowParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromBelow:
		value := value.(*[]*PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromBelow)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromBelow:
		value_ := value.([]*PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromBelow)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromBelow; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (p *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestOutputReference) validateResolveParameters(context cdktn.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
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

func (j *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestOutputReference) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktn.IResolvable:
		// ok
	case *PrometheusAnomalyDetectorConfigurationRandomCutForest:
		val := val.(*PrometheusAnomalyDetectorConfigurationRandomCutForest)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case PrometheusAnomalyDetectorConfigurationRandomCutForest:
		val_ := val.(PrometheusAnomalyDetectorConfigurationRandomCutForest)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktn.IResolvable, *PrometheusAnomalyDetectorConfigurationRandomCutForest; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestOutputReference) validateSetQueryParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestOutputReference) validateSetSampleSizeParameters(val *float64) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestOutputReference) validateSetShingleSizeParameters(val *float64) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestOutputReference) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewPrometheusAnomalyDetectorConfigurationRandomCutForestOutputReferenceParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) error {
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

