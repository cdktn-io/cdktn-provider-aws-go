// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build !no_runtime_type_checking

package bedrockagentdatasource

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

func (b *jsiiProxy_BedrockagentDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationMediaExtractionConfigurationOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_BedrockagentDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationMediaExtractionConfigurationOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_BedrockagentDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationMediaExtractionConfigurationOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_BedrockagentDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationMediaExtractionConfigurationOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_BedrockagentDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationMediaExtractionConfigurationOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_BedrockagentDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationMediaExtractionConfigurationOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_BedrockagentDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationMediaExtractionConfigurationOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_BedrockagentDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationMediaExtractionConfigurationOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_BedrockagentDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationMediaExtractionConfigurationOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_BedrockagentDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationMediaExtractionConfigurationOutputReference) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_BedrockagentDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationMediaExtractionConfigurationOutputReference) validatePutAudioExtractionConfigurationParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*BedrockagentDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationMediaExtractionConfigurationAudioExtractionConfiguration:
		value := value.(*[]*BedrockagentDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationMediaExtractionConfigurationAudioExtractionConfiguration)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*BedrockagentDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationMediaExtractionConfigurationAudioExtractionConfiguration:
		value_ := value.([]*BedrockagentDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationMediaExtractionConfigurationAudioExtractionConfiguration)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*BedrockagentDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationMediaExtractionConfigurationAudioExtractionConfiguration; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (b *jsiiProxy_BedrockagentDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationMediaExtractionConfigurationOutputReference) validatePutImageExtractionConfigurationParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*BedrockagentDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationMediaExtractionConfigurationImageExtractionConfiguration:
		value := value.(*[]*BedrockagentDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationMediaExtractionConfigurationImageExtractionConfiguration)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*BedrockagentDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationMediaExtractionConfigurationImageExtractionConfiguration:
		value_ := value.([]*BedrockagentDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationMediaExtractionConfigurationImageExtractionConfiguration)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*BedrockagentDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationMediaExtractionConfigurationImageExtractionConfiguration; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (b *jsiiProxy_BedrockagentDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationMediaExtractionConfigurationOutputReference) validatePutVideoExtractionConfigurationParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*BedrockagentDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationMediaExtractionConfigurationVideoExtractionConfiguration:
		value := value.(*[]*BedrockagentDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationMediaExtractionConfigurationVideoExtractionConfiguration)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*BedrockagentDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationMediaExtractionConfigurationVideoExtractionConfiguration:
		value_ := value.([]*BedrockagentDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationMediaExtractionConfigurationVideoExtractionConfiguration)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*BedrockagentDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationMediaExtractionConfigurationVideoExtractionConfiguration; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (b *jsiiProxy_BedrockagentDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationMediaExtractionConfigurationOutputReference) validateResolveParameters(context cdktn.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_BedrockagentDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationMediaExtractionConfigurationOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
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

func (j *jsiiProxy_BedrockagentDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationMediaExtractionConfigurationOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_BedrockagentDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationMediaExtractionConfigurationOutputReference) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktn.IResolvable:
		// ok
	case *BedrockagentDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationMediaExtractionConfiguration:
		val := val.(*BedrockagentDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationMediaExtractionConfiguration)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case BedrockagentDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationMediaExtractionConfiguration:
		val_ := val.(BedrockagentDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationMediaExtractionConfiguration)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktn.IResolvable, *BedrockagentDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationMediaExtractionConfiguration; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_BedrockagentDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationMediaExtractionConfigurationOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_BedrockagentDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationMediaExtractionConfigurationOutputReference) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewBedrockagentDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationMediaExtractionConfigurationOutputReferenceParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) error {
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

