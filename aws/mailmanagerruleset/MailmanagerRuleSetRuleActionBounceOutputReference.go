// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mailmanagerruleset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/jsii"

	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/mailmanagerruleset/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MailmanagerRuleSetRuleActionBounceOutputReference interface {
	cdktn.ComplexObject
	ActionFailurePolicy() *string
	SetActionFailurePolicy(val *string)
	ActionFailurePolicyInput() *string
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
	DiagnosticMessage() *string
	SetDiagnosticMessage(val *string)
	DiagnosticMessageInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Message() *string
	SetMessage(val *string)
	MessageInput() *string
	RoleArn() *string
	SetRoleArn(val *string)
	RoleArnInput() *string
	Sender() *string
	SetSender(val *string)
	SenderInput() *string
	SmtpReplyCode() *string
	SetSmtpReplyCode(val *string)
	SmtpReplyCodeInput() *string
	StatusCode() *string
	SetStatusCode(val *string)
	StatusCodeInput() *string
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
	ResetActionFailurePolicy()
	ResetMessage()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MailmanagerRuleSetRuleActionBounceOutputReference
type jsiiProxy_MailmanagerRuleSetRuleActionBounceOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_MailmanagerRuleSetRuleActionBounceOutputReference) ActionFailurePolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionFailurePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MailmanagerRuleSetRuleActionBounceOutputReference) ActionFailurePolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionFailurePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MailmanagerRuleSetRuleActionBounceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MailmanagerRuleSetRuleActionBounceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MailmanagerRuleSetRuleActionBounceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MailmanagerRuleSetRuleActionBounceOutputReference) DiagnosticMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diagnosticMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MailmanagerRuleSetRuleActionBounceOutputReference) DiagnosticMessageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diagnosticMessageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MailmanagerRuleSetRuleActionBounceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MailmanagerRuleSetRuleActionBounceOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MailmanagerRuleSetRuleActionBounceOutputReference) Message() *string {
	var returns *string
	_jsii_.Get(
		j,
		"message",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MailmanagerRuleSetRuleActionBounceOutputReference) MessageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"messageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MailmanagerRuleSetRuleActionBounceOutputReference) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MailmanagerRuleSetRuleActionBounceOutputReference) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MailmanagerRuleSetRuleActionBounceOutputReference) Sender() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sender",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MailmanagerRuleSetRuleActionBounceOutputReference) SenderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"senderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MailmanagerRuleSetRuleActionBounceOutputReference) SmtpReplyCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"smtpReplyCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MailmanagerRuleSetRuleActionBounceOutputReference) SmtpReplyCodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"smtpReplyCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MailmanagerRuleSetRuleActionBounceOutputReference) StatusCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MailmanagerRuleSetRuleActionBounceOutputReference) StatusCodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MailmanagerRuleSetRuleActionBounceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MailmanagerRuleSetRuleActionBounceOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMailmanagerRuleSetRuleActionBounceOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) MailmanagerRuleSetRuleActionBounceOutputReference {
	_init_.Initialize()

	if err := validateNewMailmanagerRuleSetRuleActionBounceOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_MailmanagerRuleSetRuleActionBounceOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-aws.mailmanagerRuleSet.MailmanagerRuleSetRuleActionBounceOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewMailmanagerRuleSetRuleActionBounceOutputReference_Override(m MailmanagerRuleSetRuleActionBounceOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.mailmanagerRuleSet.MailmanagerRuleSetRuleActionBounceOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		m,
	)
}

func (j *jsiiProxy_MailmanagerRuleSetRuleActionBounceOutputReference)SetActionFailurePolicy(val *string) {
	if err := j.validateSetActionFailurePolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"actionFailurePolicy",
		val,
	)
}

func (j *jsiiProxy_MailmanagerRuleSetRuleActionBounceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MailmanagerRuleSetRuleActionBounceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MailmanagerRuleSetRuleActionBounceOutputReference)SetDiagnosticMessage(val *string) {
	if err := j.validateSetDiagnosticMessageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diagnosticMessage",
		val,
	)
}

func (j *jsiiProxy_MailmanagerRuleSetRuleActionBounceOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MailmanagerRuleSetRuleActionBounceOutputReference)SetMessage(val *string) {
	if err := j.validateSetMessageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"message",
		val,
	)
}

func (j *jsiiProxy_MailmanagerRuleSetRuleActionBounceOutputReference)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_MailmanagerRuleSetRuleActionBounceOutputReference)SetSender(val *string) {
	if err := j.validateSetSenderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sender",
		val,
	)
}

func (j *jsiiProxy_MailmanagerRuleSetRuleActionBounceOutputReference)SetSmtpReplyCode(val *string) {
	if err := j.validateSetSmtpReplyCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"smtpReplyCode",
		val,
	)
}

func (j *jsiiProxy_MailmanagerRuleSetRuleActionBounceOutputReference)SetStatusCode(val *string) {
	if err := j.validateSetStatusCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"statusCode",
		val,
	)
}

func (j *jsiiProxy_MailmanagerRuleSetRuleActionBounceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MailmanagerRuleSetRuleActionBounceOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MailmanagerRuleSetRuleActionBounceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MailmanagerRuleSetRuleActionBounceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MailmanagerRuleSetRuleActionBounceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (m *jsiiProxy_MailmanagerRuleSetRuleActionBounceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MailmanagerRuleSetRuleActionBounceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MailmanagerRuleSetRuleActionBounceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MailmanagerRuleSetRuleActionBounceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MailmanagerRuleSetRuleActionBounceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MailmanagerRuleSetRuleActionBounceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MailmanagerRuleSetRuleActionBounceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MailmanagerRuleSetRuleActionBounceOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MailmanagerRuleSetRuleActionBounceOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (m *jsiiProxy_MailmanagerRuleSetRuleActionBounceOutputReference) ResetActionFailurePolicy() {
	_jsii_.InvokeVoid(
		m,
		"resetActionFailurePolicy",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MailmanagerRuleSetRuleActionBounceOutputReference) ResetMessage() {
	_jsii_.InvokeVoid(
		m,
		"resetMessage",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MailmanagerRuleSetRuleActionBounceOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (m *jsiiProxy_MailmanagerRuleSetRuleActionBounceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

