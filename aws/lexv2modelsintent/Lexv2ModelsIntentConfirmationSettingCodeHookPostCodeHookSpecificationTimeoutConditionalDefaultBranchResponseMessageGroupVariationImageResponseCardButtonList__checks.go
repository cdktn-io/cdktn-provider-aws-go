// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build !no_runtime_type_checking

package lexv2modelsintent

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationTimeoutConditionalDefaultBranchResponseMessageGroupVariationImageResponseCardButtonList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	if mapKeyAttributeName == nil {
		return fmt.Errorf("parameter mapKeyAttributeName is required, but nil was provided")
	}

	return nil
}

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationTimeoutConditionalDefaultBranchResponseMessageGroupVariationImageResponseCardButtonList) validateGetParameters(index *float64) error {
	if index == nil {
		return fmt.Errorf("parameter index is required, but nil was provided")
	}

	return nil
}

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationTimeoutConditionalDefaultBranchResponseMessageGroupVariationImageResponseCardButtonList) validateResolveParameters(context cdktn.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationTimeoutConditionalDefaultBranchResponseMessageGroupVariationImageResponseCardButtonList) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationTimeoutConditionalDefaultBranchResponseMessageGroupVariationImageResponseCardButton:
		val := val.(*[]*Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationTimeoutConditionalDefaultBranchResponseMessageGroupVariationImageResponseCardButton)
		for idx_97dfc6, v := range *val {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
				return err
			}
		}
	case []*Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationTimeoutConditionalDefaultBranchResponseMessageGroupVariationImageResponseCardButton:
		val_ := val.([]*Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationTimeoutConditionalDefaultBranchResponseMessageGroupVariationImageResponseCardButton)
		val := &val_
		for idx_97dfc6, v := range *val {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktn.IResolvable, *[]*Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationTimeoutConditionalDefaultBranchResponseMessageGroupVariationImageResponseCardButton; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationTimeoutConditionalDefaultBranchResponseMessageGroupVariationImageResponseCardButtonList) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationTimeoutConditionalDefaultBranchResponseMessageGroupVariationImageResponseCardButtonList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationTimeoutConditionalDefaultBranchResponseMessageGroupVariationImageResponseCardButtonList) validateSetWrapsSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewLexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationTimeoutConditionalDefaultBranchResponseMessageGroupVariationImageResponseCardButtonListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
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

