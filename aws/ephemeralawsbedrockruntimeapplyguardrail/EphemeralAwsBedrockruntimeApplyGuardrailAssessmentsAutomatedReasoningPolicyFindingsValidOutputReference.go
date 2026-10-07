// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ephemeralawsbedrockruntimeapplyguardrail

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/jsii"

	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/ephemeralawsbedrockruntimeapplyguardrail/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type EphemeralAwsBedrockruntimeApplyGuardrailAssessmentsAutomatedReasoningPolicyFindingsValidOutputReference interface {
	cdktn.ComplexObject
	ClaimsTrueScenario() EphemeralAwsBedrockruntimeApplyGuardrailAssessmentsAutomatedReasoningPolicyFindingsValidClaimsTrueScenarioList
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *EphemeralAwsBedrockruntimeApplyGuardrailAssessmentsAutomatedReasoningPolicyFindingsValid
	SetInternalValue(val *EphemeralAwsBedrockruntimeApplyGuardrailAssessmentsAutomatedReasoningPolicyFindingsValid)
	LogicWarning() EphemeralAwsBedrockruntimeApplyGuardrailAssessmentsAutomatedReasoningPolicyFindingsValidLogicWarningList
	SupportingRules() EphemeralAwsBedrockruntimeApplyGuardrailAssessmentsAutomatedReasoningPolicyFindingsValidSupportingRulesList
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Translation() EphemeralAwsBedrockruntimeApplyGuardrailAssessmentsAutomatedReasoningPolicyFindingsValidTranslationList
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktn.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EphemeralAwsBedrockruntimeApplyGuardrailAssessmentsAutomatedReasoningPolicyFindingsValidOutputReference
type jsiiProxy_EphemeralAwsBedrockruntimeApplyGuardrailAssessmentsAutomatedReasoningPolicyFindingsValidOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_EphemeralAwsBedrockruntimeApplyGuardrailAssessmentsAutomatedReasoningPolicyFindingsValidOutputReference) ClaimsTrueScenario() EphemeralAwsBedrockruntimeApplyGuardrailAssessmentsAutomatedReasoningPolicyFindingsValidClaimsTrueScenarioList {
	var returns EphemeralAwsBedrockruntimeApplyGuardrailAssessmentsAutomatedReasoningPolicyFindingsValidClaimsTrueScenarioList
	_jsii_.Get(
		j,
		"claimsTrueScenario",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAwsBedrockruntimeApplyGuardrailAssessmentsAutomatedReasoningPolicyFindingsValidOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAwsBedrockruntimeApplyGuardrailAssessmentsAutomatedReasoningPolicyFindingsValidOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAwsBedrockruntimeApplyGuardrailAssessmentsAutomatedReasoningPolicyFindingsValidOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAwsBedrockruntimeApplyGuardrailAssessmentsAutomatedReasoningPolicyFindingsValidOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAwsBedrockruntimeApplyGuardrailAssessmentsAutomatedReasoningPolicyFindingsValidOutputReference) InternalValue() *EphemeralAwsBedrockruntimeApplyGuardrailAssessmentsAutomatedReasoningPolicyFindingsValid {
	var returns *EphemeralAwsBedrockruntimeApplyGuardrailAssessmentsAutomatedReasoningPolicyFindingsValid
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAwsBedrockruntimeApplyGuardrailAssessmentsAutomatedReasoningPolicyFindingsValidOutputReference) LogicWarning() EphemeralAwsBedrockruntimeApplyGuardrailAssessmentsAutomatedReasoningPolicyFindingsValidLogicWarningList {
	var returns EphemeralAwsBedrockruntimeApplyGuardrailAssessmentsAutomatedReasoningPolicyFindingsValidLogicWarningList
	_jsii_.Get(
		j,
		"logicWarning",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAwsBedrockruntimeApplyGuardrailAssessmentsAutomatedReasoningPolicyFindingsValidOutputReference) SupportingRules() EphemeralAwsBedrockruntimeApplyGuardrailAssessmentsAutomatedReasoningPolicyFindingsValidSupportingRulesList {
	var returns EphemeralAwsBedrockruntimeApplyGuardrailAssessmentsAutomatedReasoningPolicyFindingsValidSupportingRulesList
	_jsii_.Get(
		j,
		"supportingRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAwsBedrockruntimeApplyGuardrailAssessmentsAutomatedReasoningPolicyFindingsValidOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAwsBedrockruntimeApplyGuardrailAssessmentsAutomatedReasoningPolicyFindingsValidOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAwsBedrockruntimeApplyGuardrailAssessmentsAutomatedReasoningPolicyFindingsValidOutputReference) Translation() EphemeralAwsBedrockruntimeApplyGuardrailAssessmentsAutomatedReasoningPolicyFindingsValidTranslationList {
	var returns EphemeralAwsBedrockruntimeApplyGuardrailAssessmentsAutomatedReasoningPolicyFindingsValidTranslationList
	_jsii_.Get(
		j,
		"translation",
		&returns,
	)
	return returns
}


func NewEphemeralAwsBedrockruntimeApplyGuardrailAssessmentsAutomatedReasoningPolicyFindingsValidOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) EphemeralAwsBedrockruntimeApplyGuardrailAssessmentsAutomatedReasoningPolicyFindingsValidOutputReference {
	_init_.Initialize()

	if err := validateNewEphemeralAwsBedrockruntimeApplyGuardrailAssessmentsAutomatedReasoningPolicyFindingsValidOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_EphemeralAwsBedrockruntimeApplyGuardrailAssessmentsAutomatedReasoningPolicyFindingsValidOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-aws.ephemeralAwsBedrockruntimeApplyGuardrail.EphemeralAwsBedrockruntimeApplyGuardrailAssessmentsAutomatedReasoningPolicyFindingsValidOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewEphemeralAwsBedrockruntimeApplyGuardrailAssessmentsAutomatedReasoningPolicyFindingsValidOutputReference_Override(e EphemeralAwsBedrockruntimeApplyGuardrailAssessmentsAutomatedReasoningPolicyFindingsValidOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.ephemeralAwsBedrockruntimeApplyGuardrail.EphemeralAwsBedrockruntimeApplyGuardrailAssessmentsAutomatedReasoningPolicyFindingsValidOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_EphemeralAwsBedrockruntimeApplyGuardrailAssessmentsAutomatedReasoningPolicyFindingsValidOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EphemeralAwsBedrockruntimeApplyGuardrailAssessmentsAutomatedReasoningPolicyFindingsValidOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EphemeralAwsBedrockruntimeApplyGuardrailAssessmentsAutomatedReasoningPolicyFindingsValidOutputReference)SetInternalValue(val *EphemeralAwsBedrockruntimeApplyGuardrailAssessmentsAutomatedReasoningPolicyFindingsValid) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EphemeralAwsBedrockruntimeApplyGuardrailAssessmentsAutomatedReasoningPolicyFindingsValidOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EphemeralAwsBedrockruntimeApplyGuardrailAssessmentsAutomatedReasoningPolicyFindingsValidOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_EphemeralAwsBedrockruntimeApplyGuardrailAssessmentsAutomatedReasoningPolicyFindingsValidOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralAwsBedrockruntimeApplyGuardrailAssessmentsAutomatedReasoningPolicyFindingsValidOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralAwsBedrockruntimeApplyGuardrailAssessmentsAutomatedReasoningPolicyFindingsValidOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralAwsBedrockruntimeApplyGuardrailAssessmentsAutomatedReasoningPolicyFindingsValidOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralAwsBedrockruntimeApplyGuardrailAssessmentsAutomatedReasoningPolicyFindingsValidOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralAwsBedrockruntimeApplyGuardrailAssessmentsAutomatedReasoningPolicyFindingsValidOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralAwsBedrockruntimeApplyGuardrailAssessmentsAutomatedReasoningPolicyFindingsValidOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralAwsBedrockruntimeApplyGuardrailAssessmentsAutomatedReasoningPolicyFindingsValidOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralAwsBedrockruntimeApplyGuardrailAssessmentsAutomatedReasoningPolicyFindingsValidOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralAwsBedrockruntimeApplyGuardrailAssessmentsAutomatedReasoningPolicyFindingsValidOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralAwsBedrockruntimeApplyGuardrailAssessmentsAutomatedReasoningPolicyFindingsValidOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralAwsBedrockruntimeApplyGuardrailAssessmentsAutomatedReasoningPolicyFindingsValidOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralAwsBedrockruntimeApplyGuardrailAssessmentsAutomatedReasoningPolicyFindingsValidOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := e.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralAwsBedrockruntimeApplyGuardrailAssessmentsAutomatedReasoningPolicyFindingsValidOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

