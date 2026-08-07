// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mailmanagerruleset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/jsii"

	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/mailmanagerruleset/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MailmanagerRuleSetRuleActionOutputReference interface {
	cdktn.ComplexObject
	AddHeader() MailmanagerRuleSetRuleActionAddHeaderList
	AddHeaderInput() interface{}
	Archive() MailmanagerRuleSetRuleActionArchiveList
	ArchiveInput() interface{}
	Bounce() MailmanagerRuleSetRuleActionBounceList
	BounceInput() interface{}
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
	DeliverToMailbox() MailmanagerRuleSetRuleActionDeliverToMailboxList
	DeliverToMailboxInput() interface{}
	DeliverToQBusiness() MailmanagerRuleSetRuleActionDeliverToQBusinessList
	DeliverToQBusinessInput() interface{}
	Drop() MailmanagerRuleSetRuleActionDropList
	DropInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	InvokeLambda() MailmanagerRuleSetRuleActionInvokeLambdaList
	InvokeLambdaInput() interface{}
	PublishToSns() MailmanagerRuleSetRuleActionPublishToSnsList
	PublishToSnsInput() interface{}
	Relay() MailmanagerRuleSetRuleActionRelayList
	RelayInput() interface{}
	ReplaceRecipient() MailmanagerRuleSetRuleActionReplaceRecipientList
	ReplaceRecipientInput() interface{}
	Send() MailmanagerRuleSetRuleActionSendList
	SendInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	WriteToS3() MailmanagerRuleSetRuleActionWriteToS3List
	WriteToS3Input() interface{}
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
	PutAddHeader(value interface{})
	PutArchive(value interface{})
	PutBounce(value interface{})
	PutDeliverToMailbox(value interface{})
	PutDeliverToQBusiness(value interface{})
	PutDrop(value interface{})
	PutInvokeLambda(value interface{})
	PutPublishToSns(value interface{})
	PutRelay(value interface{})
	PutReplaceRecipient(value interface{})
	PutSend(value interface{})
	PutWriteToS3(value interface{})
	ResetAddHeader()
	ResetArchive()
	ResetBounce()
	ResetDeliverToMailbox()
	ResetDeliverToQBusiness()
	ResetDrop()
	ResetInvokeLambda()
	ResetPublishToSns()
	ResetRelay()
	ResetReplaceRecipient()
	ResetSend()
	ResetWriteToS3()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MailmanagerRuleSetRuleActionOutputReference
type jsiiProxy_MailmanagerRuleSetRuleActionOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_MailmanagerRuleSetRuleActionOutputReference) AddHeader() MailmanagerRuleSetRuleActionAddHeaderList {
	var returns MailmanagerRuleSetRuleActionAddHeaderList
	_jsii_.Get(
		j,
		"addHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MailmanagerRuleSetRuleActionOutputReference) AddHeaderInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"addHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MailmanagerRuleSetRuleActionOutputReference) Archive() MailmanagerRuleSetRuleActionArchiveList {
	var returns MailmanagerRuleSetRuleActionArchiveList
	_jsii_.Get(
		j,
		"archive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MailmanagerRuleSetRuleActionOutputReference) ArchiveInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"archiveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MailmanagerRuleSetRuleActionOutputReference) Bounce() MailmanagerRuleSetRuleActionBounceList {
	var returns MailmanagerRuleSetRuleActionBounceList
	_jsii_.Get(
		j,
		"bounce",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MailmanagerRuleSetRuleActionOutputReference) BounceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bounceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MailmanagerRuleSetRuleActionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MailmanagerRuleSetRuleActionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MailmanagerRuleSetRuleActionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MailmanagerRuleSetRuleActionOutputReference) DeliverToMailbox() MailmanagerRuleSetRuleActionDeliverToMailboxList {
	var returns MailmanagerRuleSetRuleActionDeliverToMailboxList
	_jsii_.Get(
		j,
		"deliverToMailbox",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MailmanagerRuleSetRuleActionOutputReference) DeliverToMailboxInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deliverToMailboxInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MailmanagerRuleSetRuleActionOutputReference) DeliverToQBusiness() MailmanagerRuleSetRuleActionDeliverToQBusinessList {
	var returns MailmanagerRuleSetRuleActionDeliverToQBusinessList
	_jsii_.Get(
		j,
		"deliverToQBusiness",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MailmanagerRuleSetRuleActionOutputReference) DeliverToQBusinessInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deliverToQBusinessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MailmanagerRuleSetRuleActionOutputReference) Drop() MailmanagerRuleSetRuleActionDropList {
	var returns MailmanagerRuleSetRuleActionDropList
	_jsii_.Get(
		j,
		"drop",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MailmanagerRuleSetRuleActionOutputReference) DropInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dropInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MailmanagerRuleSetRuleActionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MailmanagerRuleSetRuleActionOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MailmanagerRuleSetRuleActionOutputReference) InvokeLambda() MailmanagerRuleSetRuleActionInvokeLambdaList {
	var returns MailmanagerRuleSetRuleActionInvokeLambdaList
	_jsii_.Get(
		j,
		"invokeLambda",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MailmanagerRuleSetRuleActionOutputReference) InvokeLambdaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"invokeLambdaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MailmanagerRuleSetRuleActionOutputReference) PublishToSns() MailmanagerRuleSetRuleActionPublishToSnsList {
	var returns MailmanagerRuleSetRuleActionPublishToSnsList
	_jsii_.Get(
		j,
		"publishToSns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MailmanagerRuleSetRuleActionOutputReference) PublishToSnsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publishToSnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MailmanagerRuleSetRuleActionOutputReference) Relay() MailmanagerRuleSetRuleActionRelayList {
	var returns MailmanagerRuleSetRuleActionRelayList
	_jsii_.Get(
		j,
		"relay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MailmanagerRuleSetRuleActionOutputReference) RelayInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"relayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MailmanagerRuleSetRuleActionOutputReference) ReplaceRecipient() MailmanagerRuleSetRuleActionReplaceRecipientList {
	var returns MailmanagerRuleSetRuleActionReplaceRecipientList
	_jsii_.Get(
		j,
		"replaceRecipient",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MailmanagerRuleSetRuleActionOutputReference) ReplaceRecipientInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"replaceRecipientInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MailmanagerRuleSetRuleActionOutputReference) Send() MailmanagerRuleSetRuleActionSendList {
	var returns MailmanagerRuleSetRuleActionSendList
	_jsii_.Get(
		j,
		"send",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MailmanagerRuleSetRuleActionOutputReference) SendInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sendInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MailmanagerRuleSetRuleActionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MailmanagerRuleSetRuleActionOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MailmanagerRuleSetRuleActionOutputReference) WriteToS3() MailmanagerRuleSetRuleActionWriteToS3List {
	var returns MailmanagerRuleSetRuleActionWriteToS3List
	_jsii_.Get(
		j,
		"writeToS3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MailmanagerRuleSetRuleActionOutputReference) WriteToS3Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"writeToS3Input",
		&returns,
	)
	return returns
}


func NewMailmanagerRuleSetRuleActionOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) MailmanagerRuleSetRuleActionOutputReference {
	_init_.Initialize()

	if err := validateNewMailmanagerRuleSetRuleActionOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_MailmanagerRuleSetRuleActionOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-aws.mailmanagerRuleSet.MailmanagerRuleSetRuleActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewMailmanagerRuleSetRuleActionOutputReference_Override(m MailmanagerRuleSetRuleActionOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.mailmanagerRuleSet.MailmanagerRuleSetRuleActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		m,
	)
}

func (j *jsiiProxy_MailmanagerRuleSetRuleActionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MailmanagerRuleSetRuleActionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MailmanagerRuleSetRuleActionOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MailmanagerRuleSetRuleActionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MailmanagerRuleSetRuleActionOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MailmanagerRuleSetRuleActionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MailmanagerRuleSetRuleActionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MailmanagerRuleSetRuleActionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (m *jsiiProxy_MailmanagerRuleSetRuleActionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MailmanagerRuleSetRuleActionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MailmanagerRuleSetRuleActionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MailmanagerRuleSetRuleActionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MailmanagerRuleSetRuleActionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MailmanagerRuleSetRuleActionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MailmanagerRuleSetRuleActionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MailmanagerRuleSetRuleActionOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MailmanagerRuleSetRuleActionOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (m *jsiiProxy_MailmanagerRuleSetRuleActionOutputReference) PutAddHeader(value interface{}) {
	if err := m.validatePutAddHeaderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putAddHeader",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MailmanagerRuleSetRuleActionOutputReference) PutArchive(value interface{}) {
	if err := m.validatePutArchiveParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putArchive",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MailmanagerRuleSetRuleActionOutputReference) PutBounce(value interface{}) {
	if err := m.validatePutBounceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putBounce",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MailmanagerRuleSetRuleActionOutputReference) PutDeliverToMailbox(value interface{}) {
	if err := m.validatePutDeliverToMailboxParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putDeliverToMailbox",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MailmanagerRuleSetRuleActionOutputReference) PutDeliverToQBusiness(value interface{}) {
	if err := m.validatePutDeliverToQBusinessParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putDeliverToQBusiness",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MailmanagerRuleSetRuleActionOutputReference) PutDrop(value interface{}) {
	if err := m.validatePutDropParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putDrop",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MailmanagerRuleSetRuleActionOutputReference) PutInvokeLambda(value interface{}) {
	if err := m.validatePutInvokeLambdaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putInvokeLambda",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MailmanagerRuleSetRuleActionOutputReference) PutPublishToSns(value interface{}) {
	if err := m.validatePutPublishToSnsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putPublishToSns",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MailmanagerRuleSetRuleActionOutputReference) PutRelay(value interface{}) {
	if err := m.validatePutRelayParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putRelay",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MailmanagerRuleSetRuleActionOutputReference) PutReplaceRecipient(value interface{}) {
	if err := m.validatePutReplaceRecipientParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putReplaceRecipient",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MailmanagerRuleSetRuleActionOutputReference) PutSend(value interface{}) {
	if err := m.validatePutSendParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putSend",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MailmanagerRuleSetRuleActionOutputReference) PutWriteToS3(value interface{}) {
	if err := m.validatePutWriteToS3Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putWriteToS3",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MailmanagerRuleSetRuleActionOutputReference) ResetAddHeader() {
	_jsii_.InvokeVoid(
		m,
		"resetAddHeader",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MailmanagerRuleSetRuleActionOutputReference) ResetArchive() {
	_jsii_.InvokeVoid(
		m,
		"resetArchive",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MailmanagerRuleSetRuleActionOutputReference) ResetBounce() {
	_jsii_.InvokeVoid(
		m,
		"resetBounce",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MailmanagerRuleSetRuleActionOutputReference) ResetDeliverToMailbox() {
	_jsii_.InvokeVoid(
		m,
		"resetDeliverToMailbox",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MailmanagerRuleSetRuleActionOutputReference) ResetDeliverToQBusiness() {
	_jsii_.InvokeVoid(
		m,
		"resetDeliverToQBusiness",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MailmanagerRuleSetRuleActionOutputReference) ResetDrop() {
	_jsii_.InvokeVoid(
		m,
		"resetDrop",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MailmanagerRuleSetRuleActionOutputReference) ResetInvokeLambda() {
	_jsii_.InvokeVoid(
		m,
		"resetInvokeLambda",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MailmanagerRuleSetRuleActionOutputReference) ResetPublishToSns() {
	_jsii_.InvokeVoid(
		m,
		"resetPublishToSns",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MailmanagerRuleSetRuleActionOutputReference) ResetRelay() {
	_jsii_.InvokeVoid(
		m,
		"resetRelay",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MailmanagerRuleSetRuleActionOutputReference) ResetReplaceRecipient() {
	_jsii_.InvokeVoid(
		m,
		"resetReplaceRecipient",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MailmanagerRuleSetRuleActionOutputReference) ResetSend() {
	_jsii_.InvokeVoid(
		m,
		"resetSend",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MailmanagerRuleSetRuleActionOutputReference) ResetWriteToS3() {
	_jsii_.InvokeVoid(
		m,
		"resetWriteToS3",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MailmanagerRuleSetRuleActionOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (m *jsiiProxy_MailmanagerRuleSetRuleActionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

