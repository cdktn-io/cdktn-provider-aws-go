// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mailmanageringresspoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/jsii"

	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/mailmanageringresspoint/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MailmanagerIngressPointIngressPointConfigurationOutputReference interface {
	cdktn.ComplexObject
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
	SecretArn() *string
	SetSecretArn(val *string)
	SecretArnInput() *string
	// Deprecated: Write-only: the provider never returns this value; reading it always yields null by protocol contract. The getter remains for compatibility and will be removed in a future prebuilt-provider major.
	SmtpPasswordWo() *string
	// Deprecated: Write-only: the provider never returns this value; reading it always yields null by protocol contract. The getter remains for compatibility and will be removed in a future prebuilt-provider major.
	SetSmtpPasswordWo(val *string)
	SmtpPasswordWoInput() *string
	SmtpPasswordWoVersion() *float64
	SetSmtpPasswordWoVersion(val *float64)
	SmtpPasswordWoVersionInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TlsAuthConfiguration() MailmanagerIngressPointIngressPointConfigurationTlsAuthConfigurationList
	TlsAuthConfigurationInput() interface{}
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
	PutTlsAuthConfiguration(value interface{})
	ResetSecretArn()
	ResetSmtpPasswordWo()
	ResetSmtpPasswordWoVersion()
	ResetTlsAuthConfiguration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MailmanagerIngressPointIngressPointConfigurationOutputReference
type jsiiProxy_MailmanagerIngressPointIngressPointConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_MailmanagerIngressPointIngressPointConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MailmanagerIngressPointIngressPointConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MailmanagerIngressPointIngressPointConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MailmanagerIngressPointIngressPointConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MailmanagerIngressPointIngressPointConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MailmanagerIngressPointIngressPointConfigurationOutputReference) SecretArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MailmanagerIngressPointIngressPointConfigurationOutputReference) SecretArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MailmanagerIngressPointIngressPointConfigurationOutputReference) SmtpPasswordWo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"smtpPasswordWo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MailmanagerIngressPointIngressPointConfigurationOutputReference) SmtpPasswordWoInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"smtpPasswordWoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MailmanagerIngressPointIngressPointConfigurationOutputReference) SmtpPasswordWoVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"smtpPasswordWoVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MailmanagerIngressPointIngressPointConfigurationOutputReference) SmtpPasswordWoVersionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"smtpPasswordWoVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MailmanagerIngressPointIngressPointConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MailmanagerIngressPointIngressPointConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MailmanagerIngressPointIngressPointConfigurationOutputReference) TlsAuthConfiguration() MailmanagerIngressPointIngressPointConfigurationTlsAuthConfigurationList {
	var returns MailmanagerIngressPointIngressPointConfigurationTlsAuthConfigurationList
	_jsii_.Get(
		j,
		"tlsAuthConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MailmanagerIngressPointIngressPointConfigurationOutputReference) TlsAuthConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tlsAuthConfigurationInput",
		&returns,
	)
	return returns
}


func NewMailmanagerIngressPointIngressPointConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) MailmanagerIngressPointIngressPointConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewMailmanagerIngressPointIngressPointConfigurationOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_MailmanagerIngressPointIngressPointConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-aws.mailmanagerIngressPoint.MailmanagerIngressPointIngressPointConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewMailmanagerIngressPointIngressPointConfigurationOutputReference_Override(m MailmanagerIngressPointIngressPointConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.mailmanagerIngressPoint.MailmanagerIngressPointIngressPointConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		m,
	)
}

func (j *jsiiProxy_MailmanagerIngressPointIngressPointConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MailmanagerIngressPointIngressPointConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MailmanagerIngressPointIngressPointConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MailmanagerIngressPointIngressPointConfigurationOutputReference)SetSecretArn(val *string) {
	if err := j.validateSetSecretArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretArn",
		val,
	)
}

func (j *jsiiProxy_MailmanagerIngressPointIngressPointConfigurationOutputReference)SetSmtpPasswordWo(val *string) {
	if err := j.validateSetSmtpPasswordWoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"smtpPasswordWo",
		val,
	)
}

func (j *jsiiProxy_MailmanagerIngressPointIngressPointConfigurationOutputReference)SetSmtpPasswordWoVersion(val *float64) {
	if err := j.validateSetSmtpPasswordWoVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"smtpPasswordWoVersion",
		val,
	)
}

func (j *jsiiProxy_MailmanagerIngressPointIngressPointConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MailmanagerIngressPointIngressPointConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MailmanagerIngressPointIngressPointConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MailmanagerIngressPointIngressPointConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MailmanagerIngressPointIngressPointConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (m *jsiiProxy_MailmanagerIngressPointIngressPointConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MailmanagerIngressPointIngressPointConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MailmanagerIngressPointIngressPointConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MailmanagerIngressPointIngressPointConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MailmanagerIngressPointIngressPointConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MailmanagerIngressPointIngressPointConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MailmanagerIngressPointIngressPointConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MailmanagerIngressPointIngressPointConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MailmanagerIngressPointIngressPointConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (m *jsiiProxy_MailmanagerIngressPointIngressPointConfigurationOutputReference) PutTlsAuthConfiguration(value interface{}) {
	if err := m.validatePutTlsAuthConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putTlsAuthConfiguration",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MailmanagerIngressPointIngressPointConfigurationOutputReference) ResetSecretArn() {
	_jsii_.InvokeVoid(
		m,
		"resetSecretArn",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MailmanagerIngressPointIngressPointConfigurationOutputReference) ResetSmtpPasswordWo() {
	_jsii_.InvokeVoid(
		m,
		"resetSmtpPasswordWo",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MailmanagerIngressPointIngressPointConfigurationOutputReference) ResetSmtpPasswordWoVersion() {
	_jsii_.InvokeVoid(
		m,
		"resetSmtpPasswordWoVersion",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MailmanagerIngressPointIngressPointConfigurationOutputReference) ResetTlsAuthConfiguration() {
	_jsii_.InvokeVoid(
		m,
		"resetTlsAuthConfiguration",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MailmanagerIngressPointIngressPointConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (m *jsiiProxy_MailmanagerIngressPointIngressPointConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

