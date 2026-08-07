// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mailmanagertrafficpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/jsii"

	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/mailmanagertrafficpolicy/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MailmanagerTrafficPolicyPolicyStatementConditionBooleanExpressionEvaluateAnalysisOutputReference interface {
	cdktn.ComplexObject
	Analyzer() *string
	SetAnalyzer(val *string)
	AnalyzerInput() *string
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	ResultField() *string
	SetResultField(val *string)
	ResultFieldInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
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

// The jsii proxy struct for MailmanagerTrafficPolicyPolicyStatementConditionBooleanExpressionEvaluateAnalysisOutputReference
type jsiiProxy_MailmanagerTrafficPolicyPolicyStatementConditionBooleanExpressionEvaluateAnalysisOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_MailmanagerTrafficPolicyPolicyStatementConditionBooleanExpressionEvaluateAnalysisOutputReference) Analyzer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"analyzer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MailmanagerTrafficPolicyPolicyStatementConditionBooleanExpressionEvaluateAnalysisOutputReference) AnalyzerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"analyzerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MailmanagerTrafficPolicyPolicyStatementConditionBooleanExpressionEvaluateAnalysisOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MailmanagerTrafficPolicyPolicyStatementConditionBooleanExpressionEvaluateAnalysisOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MailmanagerTrafficPolicyPolicyStatementConditionBooleanExpressionEvaluateAnalysisOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MailmanagerTrafficPolicyPolicyStatementConditionBooleanExpressionEvaluateAnalysisOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MailmanagerTrafficPolicyPolicyStatementConditionBooleanExpressionEvaluateAnalysisOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MailmanagerTrafficPolicyPolicyStatementConditionBooleanExpressionEvaluateAnalysisOutputReference) ResultField() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resultField",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MailmanagerTrafficPolicyPolicyStatementConditionBooleanExpressionEvaluateAnalysisOutputReference) ResultFieldInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resultFieldInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MailmanagerTrafficPolicyPolicyStatementConditionBooleanExpressionEvaluateAnalysisOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MailmanagerTrafficPolicyPolicyStatementConditionBooleanExpressionEvaluateAnalysisOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMailmanagerTrafficPolicyPolicyStatementConditionBooleanExpressionEvaluateAnalysisOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) MailmanagerTrafficPolicyPolicyStatementConditionBooleanExpressionEvaluateAnalysisOutputReference {
	_init_.Initialize()

	if err := validateNewMailmanagerTrafficPolicyPolicyStatementConditionBooleanExpressionEvaluateAnalysisOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_MailmanagerTrafficPolicyPolicyStatementConditionBooleanExpressionEvaluateAnalysisOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-aws.mailmanagerTrafficPolicy.MailmanagerTrafficPolicyPolicyStatementConditionBooleanExpressionEvaluateAnalysisOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewMailmanagerTrafficPolicyPolicyStatementConditionBooleanExpressionEvaluateAnalysisOutputReference_Override(m MailmanagerTrafficPolicyPolicyStatementConditionBooleanExpressionEvaluateAnalysisOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.mailmanagerTrafficPolicy.MailmanagerTrafficPolicyPolicyStatementConditionBooleanExpressionEvaluateAnalysisOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		m,
	)
}

func (j *jsiiProxy_MailmanagerTrafficPolicyPolicyStatementConditionBooleanExpressionEvaluateAnalysisOutputReference)SetAnalyzer(val *string) {
	if err := j.validateSetAnalyzerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"analyzer",
		val,
	)
}

func (j *jsiiProxy_MailmanagerTrafficPolicyPolicyStatementConditionBooleanExpressionEvaluateAnalysisOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MailmanagerTrafficPolicyPolicyStatementConditionBooleanExpressionEvaluateAnalysisOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MailmanagerTrafficPolicyPolicyStatementConditionBooleanExpressionEvaluateAnalysisOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MailmanagerTrafficPolicyPolicyStatementConditionBooleanExpressionEvaluateAnalysisOutputReference)SetResultField(val *string) {
	if err := j.validateSetResultFieldParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resultField",
		val,
	)
}

func (j *jsiiProxy_MailmanagerTrafficPolicyPolicyStatementConditionBooleanExpressionEvaluateAnalysisOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MailmanagerTrafficPolicyPolicyStatementConditionBooleanExpressionEvaluateAnalysisOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MailmanagerTrafficPolicyPolicyStatementConditionBooleanExpressionEvaluateAnalysisOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MailmanagerTrafficPolicyPolicyStatementConditionBooleanExpressionEvaluateAnalysisOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MailmanagerTrafficPolicyPolicyStatementConditionBooleanExpressionEvaluateAnalysisOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MailmanagerTrafficPolicyPolicyStatementConditionBooleanExpressionEvaluateAnalysisOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MailmanagerTrafficPolicyPolicyStatementConditionBooleanExpressionEvaluateAnalysisOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MailmanagerTrafficPolicyPolicyStatementConditionBooleanExpressionEvaluateAnalysisOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MailmanagerTrafficPolicyPolicyStatementConditionBooleanExpressionEvaluateAnalysisOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MailmanagerTrafficPolicyPolicyStatementConditionBooleanExpressionEvaluateAnalysisOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MailmanagerTrafficPolicyPolicyStatementConditionBooleanExpressionEvaluateAnalysisOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MailmanagerTrafficPolicyPolicyStatementConditionBooleanExpressionEvaluateAnalysisOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MailmanagerTrafficPolicyPolicyStatementConditionBooleanExpressionEvaluateAnalysisOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MailmanagerTrafficPolicyPolicyStatementConditionBooleanExpressionEvaluateAnalysisOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MailmanagerTrafficPolicyPolicyStatementConditionBooleanExpressionEvaluateAnalysisOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := m.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MailmanagerTrafficPolicyPolicyStatementConditionBooleanExpressionEvaluateAnalysisOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

