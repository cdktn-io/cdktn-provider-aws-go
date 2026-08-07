// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build !no_runtime_type_checking

package prometheusanomalydetector

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

func (p *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromBelowOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromBelowOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromBelowOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromBelowOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromBelowOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromBelowOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromBelowOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromBelowOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromBelowOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromBelowOutputReference) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromBelowOutputReference) validateResolveParameters(context cdktn.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromBelowOutputReference) validateSetAmountParameters(val *float64) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromBelowOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
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

func (j *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromBelowOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromBelowOutputReference) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktn.IResolvable:
		// ok
	case *PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromBelow:
		val := val.(*PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromBelow)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromBelow:
		val_ := val.(PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromBelow)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktn.IResolvable, *PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromBelow; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromBelowOutputReference) validateSetRatioParameters(val *float64) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromBelowOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromBelowOutputReference) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewPrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromBelowOutputReferenceParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) error {
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

