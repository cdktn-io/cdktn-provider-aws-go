// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoregatewaytarget

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/jsii"

	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/bedrockagentcoregatewaytarget/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BedrockagentcoreGatewayTargetTargetConfigurationHttpPassthroughOutputReference interface {
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
	Endpoint() *string
	SetEndpoint(val *string)
	EndpointInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	ProtocolType() *string
	SetProtocolType(val *string)
	ProtocolTypeInput() *string
	Schema() BedrockagentcoreGatewayTargetTargetConfigurationHttpPassthroughSchemaList
	SchemaInput() interface{}
	StaticQueryParameterConflictResolution() *string
	SetStaticQueryParameterConflictResolution(val *string)
	StaticQueryParameterConflictResolutionInput() *string
	StaticQueryParameters() *map[string]*string
	SetStaticQueryParameters(val *map[string]*string)
	StaticQueryParametersInput() *map[string]*string
	StickinessConfiguration() BedrockagentcoreGatewayTargetTargetConfigurationHttpPassthroughStickinessConfigurationList
	StickinessConfigurationInput() interface{}
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
	PutSchema(value interface{})
	PutStickinessConfiguration(value interface{})
	ResetSchema()
	ResetStaticQueryParameterConflictResolution()
	ResetStaticQueryParameters()
	ResetStickinessConfiguration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BedrockagentcoreGatewayTargetTargetConfigurationHttpPassthroughOutputReference
type jsiiProxy_BedrockagentcoreGatewayTargetTargetConfigurationHttpPassthroughOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_BedrockagentcoreGatewayTargetTargetConfigurationHttpPassthroughOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreGatewayTargetTargetConfigurationHttpPassthroughOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreGatewayTargetTargetConfigurationHttpPassthroughOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreGatewayTargetTargetConfigurationHttpPassthroughOutputReference) Endpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreGatewayTargetTargetConfigurationHttpPassthroughOutputReference) EndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreGatewayTargetTargetConfigurationHttpPassthroughOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreGatewayTargetTargetConfigurationHttpPassthroughOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreGatewayTargetTargetConfigurationHttpPassthroughOutputReference) ProtocolType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreGatewayTargetTargetConfigurationHttpPassthroughOutputReference) ProtocolTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreGatewayTargetTargetConfigurationHttpPassthroughOutputReference) Schema() BedrockagentcoreGatewayTargetTargetConfigurationHttpPassthroughSchemaList {
	var returns BedrockagentcoreGatewayTargetTargetConfigurationHttpPassthroughSchemaList
	_jsii_.Get(
		j,
		"schema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreGatewayTargetTargetConfigurationHttpPassthroughOutputReference) SchemaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"schemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreGatewayTargetTargetConfigurationHttpPassthroughOutputReference) StaticQueryParameterConflictResolution() *string {
	var returns *string
	_jsii_.Get(
		j,
		"staticQueryParameterConflictResolution",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreGatewayTargetTargetConfigurationHttpPassthroughOutputReference) StaticQueryParameterConflictResolutionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"staticQueryParameterConflictResolutionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreGatewayTargetTargetConfigurationHttpPassthroughOutputReference) StaticQueryParameters() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"staticQueryParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreGatewayTargetTargetConfigurationHttpPassthroughOutputReference) StaticQueryParametersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"staticQueryParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreGatewayTargetTargetConfigurationHttpPassthroughOutputReference) StickinessConfiguration() BedrockagentcoreGatewayTargetTargetConfigurationHttpPassthroughStickinessConfigurationList {
	var returns BedrockagentcoreGatewayTargetTargetConfigurationHttpPassthroughStickinessConfigurationList
	_jsii_.Get(
		j,
		"stickinessConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreGatewayTargetTargetConfigurationHttpPassthroughOutputReference) StickinessConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stickinessConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreGatewayTargetTargetConfigurationHttpPassthroughOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreGatewayTargetTargetConfigurationHttpPassthroughOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBedrockagentcoreGatewayTargetTargetConfigurationHttpPassthroughOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) BedrockagentcoreGatewayTargetTargetConfigurationHttpPassthroughOutputReference {
	_init_.Initialize()

	if err := validateNewBedrockagentcoreGatewayTargetTargetConfigurationHttpPassthroughOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_BedrockagentcoreGatewayTargetTargetConfigurationHttpPassthroughOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-aws.bedrockagentcoreGatewayTarget.BedrockagentcoreGatewayTargetTargetConfigurationHttpPassthroughOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewBedrockagentcoreGatewayTargetTargetConfigurationHttpPassthroughOutputReference_Override(b BedrockagentcoreGatewayTargetTargetConfigurationHttpPassthroughOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.bedrockagentcoreGatewayTarget.BedrockagentcoreGatewayTargetTargetConfigurationHttpPassthroughOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		b,
	)
}

func (j *jsiiProxy_BedrockagentcoreGatewayTargetTargetConfigurationHttpPassthroughOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreGatewayTargetTargetConfigurationHttpPassthroughOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreGatewayTargetTargetConfigurationHttpPassthroughOutputReference)SetEndpoint(val *string) {
	if err := j.validateSetEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpoint",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreGatewayTargetTargetConfigurationHttpPassthroughOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreGatewayTargetTargetConfigurationHttpPassthroughOutputReference)SetProtocolType(val *string) {
	if err := j.validateSetProtocolTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocolType",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreGatewayTargetTargetConfigurationHttpPassthroughOutputReference)SetStaticQueryParameterConflictResolution(val *string) {
	if err := j.validateSetStaticQueryParameterConflictResolutionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"staticQueryParameterConflictResolution",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreGatewayTargetTargetConfigurationHttpPassthroughOutputReference)SetStaticQueryParameters(val *map[string]*string) {
	if err := j.validateSetStaticQueryParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"staticQueryParameters",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreGatewayTargetTargetConfigurationHttpPassthroughOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreGatewayTargetTargetConfigurationHttpPassthroughOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BedrockagentcoreGatewayTargetTargetConfigurationHttpPassthroughOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreGatewayTargetTargetConfigurationHttpPassthroughOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (b *jsiiProxy_BedrockagentcoreGatewayTargetTargetConfigurationHttpPassthroughOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (b *jsiiProxy_BedrockagentcoreGatewayTargetTargetConfigurationHttpPassthroughOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (b *jsiiProxy_BedrockagentcoreGatewayTargetTargetConfigurationHttpPassthroughOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (b *jsiiProxy_BedrockagentcoreGatewayTargetTargetConfigurationHttpPassthroughOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (b *jsiiProxy_BedrockagentcoreGatewayTargetTargetConfigurationHttpPassthroughOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (b *jsiiProxy_BedrockagentcoreGatewayTargetTargetConfigurationHttpPassthroughOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (b *jsiiProxy_BedrockagentcoreGatewayTargetTargetConfigurationHttpPassthroughOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (b *jsiiProxy_BedrockagentcoreGatewayTargetTargetConfigurationHttpPassthroughOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (b *jsiiProxy_BedrockagentcoreGatewayTargetTargetConfigurationHttpPassthroughOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreGatewayTargetTargetConfigurationHttpPassthroughOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (b *jsiiProxy_BedrockagentcoreGatewayTargetTargetConfigurationHttpPassthroughOutputReference) PutSchema(value interface{}) {
	if err := b.validatePutSchemaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putSchema",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentcoreGatewayTargetTargetConfigurationHttpPassthroughOutputReference) PutStickinessConfiguration(value interface{}) {
	if err := b.validatePutStickinessConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putStickinessConfiguration",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentcoreGatewayTargetTargetConfigurationHttpPassthroughOutputReference) ResetSchema() {
	_jsii_.InvokeVoid(
		b,
		"resetSchema",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreGatewayTargetTargetConfigurationHttpPassthroughOutputReference) ResetStaticQueryParameterConflictResolution() {
	_jsii_.InvokeVoid(
		b,
		"resetStaticQueryParameterConflictResolution",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreGatewayTargetTargetConfigurationHttpPassthroughOutputReference) ResetStaticQueryParameters() {
	_jsii_.InvokeVoid(
		b,
		"resetStaticQueryParameters",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreGatewayTargetTargetConfigurationHttpPassthroughOutputReference) ResetStickinessConfiguration() {
	_jsii_.InvokeVoid(
		b,
		"resetStickinessConfiguration",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreGatewayTargetTargetConfigurationHttpPassthroughOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (b *jsiiProxy_BedrockagentcoreGatewayTargetTargetConfigurationHttpPassthroughOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

