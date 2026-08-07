// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreharness

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/jsii"

	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/bedrockagentcoreharness/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BedrockagentcoreHarnessToolConfigOutputReference interface {
	cdktn.ComplexObject
	AgentcoreBrowser() BedrockagentcoreHarnessToolConfigAgentcoreBrowserList
	AgentcoreBrowserInput() interface{}
	AgentcoreCodeInterpreter() BedrockagentcoreHarnessToolConfigAgentcoreCodeInterpreterList
	AgentcoreCodeInterpreterInput() interface{}
	AgentcoreGateway() BedrockagentcoreHarnessToolConfigAgentcoreGatewayList
	AgentcoreGatewayInput() interface{}
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
	InlineFunction() BedrockagentcoreHarnessToolConfigInlineFunctionList
	InlineFunctionInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	RemoteMcp() BedrockagentcoreHarnessToolConfigRemoteMcpList
	RemoteMcpInput() interface{}
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
	PutAgentcoreBrowser(value interface{})
	PutAgentcoreCodeInterpreter(value interface{})
	PutAgentcoreGateway(value interface{})
	PutInlineFunction(value interface{})
	PutRemoteMcp(value interface{})
	ResetAgentcoreBrowser()
	ResetAgentcoreCodeInterpreter()
	ResetAgentcoreGateway()
	ResetInlineFunction()
	ResetRemoteMcp()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BedrockagentcoreHarnessToolConfigOutputReference
type jsiiProxy_BedrockagentcoreHarnessToolConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_BedrockagentcoreHarnessToolConfigOutputReference) AgentcoreBrowser() BedrockagentcoreHarnessToolConfigAgentcoreBrowserList {
	var returns BedrockagentcoreHarnessToolConfigAgentcoreBrowserList
	_jsii_.Get(
		j,
		"agentcoreBrowser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarnessToolConfigOutputReference) AgentcoreBrowserInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"agentcoreBrowserInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarnessToolConfigOutputReference) AgentcoreCodeInterpreter() BedrockagentcoreHarnessToolConfigAgentcoreCodeInterpreterList {
	var returns BedrockagentcoreHarnessToolConfigAgentcoreCodeInterpreterList
	_jsii_.Get(
		j,
		"agentcoreCodeInterpreter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarnessToolConfigOutputReference) AgentcoreCodeInterpreterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"agentcoreCodeInterpreterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarnessToolConfigOutputReference) AgentcoreGateway() BedrockagentcoreHarnessToolConfigAgentcoreGatewayList {
	var returns BedrockagentcoreHarnessToolConfigAgentcoreGatewayList
	_jsii_.Get(
		j,
		"agentcoreGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarnessToolConfigOutputReference) AgentcoreGatewayInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"agentcoreGatewayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarnessToolConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarnessToolConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarnessToolConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarnessToolConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarnessToolConfigOutputReference) InlineFunction() BedrockagentcoreHarnessToolConfigInlineFunctionList {
	var returns BedrockagentcoreHarnessToolConfigInlineFunctionList
	_jsii_.Get(
		j,
		"inlineFunction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarnessToolConfigOutputReference) InlineFunctionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inlineFunctionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarnessToolConfigOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarnessToolConfigOutputReference) RemoteMcp() BedrockagentcoreHarnessToolConfigRemoteMcpList {
	var returns BedrockagentcoreHarnessToolConfigRemoteMcpList
	_jsii_.Get(
		j,
		"remoteMcp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarnessToolConfigOutputReference) RemoteMcpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"remoteMcpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarnessToolConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarnessToolConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBedrockagentcoreHarnessToolConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) BedrockagentcoreHarnessToolConfigOutputReference {
	_init_.Initialize()

	if err := validateNewBedrockagentcoreHarnessToolConfigOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_BedrockagentcoreHarnessToolConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-aws.bedrockagentcoreHarness.BedrockagentcoreHarnessToolConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewBedrockagentcoreHarnessToolConfigOutputReference_Override(b BedrockagentcoreHarnessToolConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.bedrockagentcoreHarness.BedrockagentcoreHarnessToolConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		b,
	)
}

func (j *jsiiProxy_BedrockagentcoreHarnessToolConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreHarnessToolConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreHarnessToolConfigOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreHarnessToolConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreHarnessToolConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BedrockagentcoreHarnessToolConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreHarnessToolConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := b.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreHarnessToolConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := b.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreHarnessToolConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := b.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreHarnessToolConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := b.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreHarnessToolConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := b.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreHarnessToolConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := b.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreHarnessToolConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := b.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreHarnessToolConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := b.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreHarnessToolConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := b.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreHarnessToolConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreHarnessToolConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreHarnessToolConfigOutputReference) PutAgentcoreBrowser(value interface{}) {
	if err := b.validatePutAgentcoreBrowserParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putAgentcoreBrowser",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentcoreHarnessToolConfigOutputReference) PutAgentcoreCodeInterpreter(value interface{}) {
	if err := b.validatePutAgentcoreCodeInterpreterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putAgentcoreCodeInterpreter",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentcoreHarnessToolConfigOutputReference) PutAgentcoreGateway(value interface{}) {
	if err := b.validatePutAgentcoreGatewayParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putAgentcoreGateway",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentcoreHarnessToolConfigOutputReference) PutInlineFunction(value interface{}) {
	if err := b.validatePutInlineFunctionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putInlineFunction",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentcoreHarnessToolConfigOutputReference) PutRemoteMcp(value interface{}) {
	if err := b.validatePutRemoteMcpParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putRemoteMcp",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentcoreHarnessToolConfigOutputReference) ResetAgentcoreBrowser() {
	_jsii_.InvokeVoid(
		b,
		"resetAgentcoreBrowser",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreHarnessToolConfigOutputReference) ResetAgentcoreCodeInterpreter() {
	_jsii_.InvokeVoid(
		b,
		"resetAgentcoreCodeInterpreter",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreHarnessToolConfigOutputReference) ResetAgentcoreGateway() {
	_jsii_.InvokeVoid(
		b,
		"resetAgentcoreGateway",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreHarnessToolConfigOutputReference) ResetInlineFunction() {
	_jsii_.InvokeVoid(
		b,
		"resetInlineFunction",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreHarnessToolConfigOutputReference) ResetRemoteMcp() {
	_jsii_.InvokeVoid(
		b,
		"resetRemoteMcp",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreHarnessToolConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := b.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreHarnessToolConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

