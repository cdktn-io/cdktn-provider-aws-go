// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentknowledgebase

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v24/jsii"

	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v24/bedrockagentknowledgebase/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BedrockagentKnowledgeBaseKnowledgeBaseConfigurationManagedKnowledgeBaseConfigurationOutputReference interface {
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
	EmbeddingModelArn() *string
	SetEmbeddingModelArn(val *string)
	EmbeddingModelArnInput() *string
	EmbeddingModelConfiguration() BedrockagentKnowledgeBaseKnowledgeBaseConfigurationManagedKnowledgeBaseConfigurationEmbeddingModelConfigurationList
	EmbeddingModelConfigurationInput() interface{}
	EmbeddingModelType() *string
	SetEmbeddingModelType(val *string)
	EmbeddingModelTypeInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	ServerSideEncryptionConfiguration() BedrockagentKnowledgeBaseKnowledgeBaseConfigurationManagedKnowledgeBaseConfigurationServerSideEncryptionConfigurationList
	ServerSideEncryptionConfigurationInput() interface{}
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
	PutEmbeddingModelConfiguration(value interface{})
	PutServerSideEncryptionConfiguration(value interface{})
	ResetEmbeddingModelArn()
	ResetEmbeddingModelConfiguration()
	ResetEmbeddingModelType()
	ResetServerSideEncryptionConfiguration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BedrockagentKnowledgeBaseKnowledgeBaseConfigurationManagedKnowledgeBaseConfigurationOutputReference
type jsiiProxy_BedrockagentKnowledgeBaseKnowledgeBaseConfigurationManagedKnowledgeBaseConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_BedrockagentKnowledgeBaseKnowledgeBaseConfigurationManagedKnowledgeBaseConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentKnowledgeBaseKnowledgeBaseConfigurationManagedKnowledgeBaseConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentKnowledgeBaseKnowledgeBaseConfigurationManagedKnowledgeBaseConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentKnowledgeBaseKnowledgeBaseConfigurationManagedKnowledgeBaseConfigurationOutputReference) EmbeddingModelArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"embeddingModelArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentKnowledgeBaseKnowledgeBaseConfigurationManagedKnowledgeBaseConfigurationOutputReference) EmbeddingModelArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"embeddingModelArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentKnowledgeBaseKnowledgeBaseConfigurationManagedKnowledgeBaseConfigurationOutputReference) EmbeddingModelConfiguration() BedrockagentKnowledgeBaseKnowledgeBaseConfigurationManagedKnowledgeBaseConfigurationEmbeddingModelConfigurationList {
	var returns BedrockagentKnowledgeBaseKnowledgeBaseConfigurationManagedKnowledgeBaseConfigurationEmbeddingModelConfigurationList
	_jsii_.Get(
		j,
		"embeddingModelConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentKnowledgeBaseKnowledgeBaseConfigurationManagedKnowledgeBaseConfigurationOutputReference) EmbeddingModelConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"embeddingModelConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentKnowledgeBaseKnowledgeBaseConfigurationManagedKnowledgeBaseConfigurationOutputReference) EmbeddingModelType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"embeddingModelType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentKnowledgeBaseKnowledgeBaseConfigurationManagedKnowledgeBaseConfigurationOutputReference) EmbeddingModelTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"embeddingModelTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentKnowledgeBaseKnowledgeBaseConfigurationManagedKnowledgeBaseConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentKnowledgeBaseKnowledgeBaseConfigurationManagedKnowledgeBaseConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentKnowledgeBaseKnowledgeBaseConfigurationManagedKnowledgeBaseConfigurationOutputReference) ServerSideEncryptionConfiguration() BedrockagentKnowledgeBaseKnowledgeBaseConfigurationManagedKnowledgeBaseConfigurationServerSideEncryptionConfigurationList {
	var returns BedrockagentKnowledgeBaseKnowledgeBaseConfigurationManagedKnowledgeBaseConfigurationServerSideEncryptionConfigurationList
	_jsii_.Get(
		j,
		"serverSideEncryptionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentKnowledgeBaseKnowledgeBaseConfigurationManagedKnowledgeBaseConfigurationOutputReference) ServerSideEncryptionConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serverSideEncryptionConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentKnowledgeBaseKnowledgeBaseConfigurationManagedKnowledgeBaseConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentKnowledgeBaseKnowledgeBaseConfigurationManagedKnowledgeBaseConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBedrockagentKnowledgeBaseKnowledgeBaseConfigurationManagedKnowledgeBaseConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) BedrockagentKnowledgeBaseKnowledgeBaseConfigurationManagedKnowledgeBaseConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewBedrockagentKnowledgeBaseKnowledgeBaseConfigurationManagedKnowledgeBaseConfigurationOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_BedrockagentKnowledgeBaseKnowledgeBaseConfigurationManagedKnowledgeBaseConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-aws.bedrockagentKnowledgeBase.BedrockagentKnowledgeBaseKnowledgeBaseConfigurationManagedKnowledgeBaseConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewBedrockagentKnowledgeBaseKnowledgeBaseConfigurationManagedKnowledgeBaseConfigurationOutputReference_Override(b BedrockagentKnowledgeBaseKnowledgeBaseConfigurationManagedKnowledgeBaseConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.bedrockagentKnowledgeBase.BedrockagentKnowledgeBaseKnowledgeBaseConfigurationManagedKnowledgeBaseConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		b,
	)
}

func (j *jsiiProxy_BedrockagentKnowledgeBaseKnowledgeBaseConfigurationManagedKnowledgeBaseConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BedrockagentKnowledgeBaseKnowledgeBaseConfigurationManagedKnowledgeBaseConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BedrockagentKnowledgeBaseKnowledgeBaseConfigurationManagedKnowledgeBaseConfigurationOutputReference)SetEmbeddingModelArn(val *string) {
	if err := j.validateSetEmbeddingModelArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"embeddingModelArn",
		val,
	)
}

func (j *jsiiProxy_BedrockagentKnowledgeBaseKnowledgeBaseConfigurationManagedKnowledgeBaseConfigurationOutputReference)SetEmbeddingModelType(val *string) {
	if err := j.validateSetEmbeddingModelTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"embeddingModelType",
		val,
	)
}

func (j *jsiiProxy_BedrockagentKnowledgeBaseKnowledgeBaseConfigurationManagedKnowledgeBaseConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BedrockagentKnowledgeBaseKnowledgeBaseConfigurationManagedKnowledgeBaseConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BedrockagentKnowledgeBaseKnowledgeBaseConfigurationManagedKnowledgeBaseConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BedrockagentKnowledgeBaseKnowledgeBaseConfigurationManagedKnowledgeBaseConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentKnowledgeBaseKnowledgeBaseConfigurationManagedKnowledgeBaseConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (b *jsiiProxy_BedrockagentKnowledgeBaseKnowledgeBaseConfigurationManagedKnowledgeBaseConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (b *jsiiProxy_BedrockagentKnowledgeBaseKnowledgeBaseConfigurationManagedKnowledgeBaseConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (b *jsiiProxy_BedrockagentKnowledgeBaseKnowledgeBaseConfigurationManagedKnowledgeBaseConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (b *jsiiProxy_BedrockagentKnowledgeBaseKnowledgeBaseConfigurationManagedKnowledgeBaseConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (b *jsiiProxy_BedrockagentKnowledgeBaseKnowledgeBaseConfigurationManagedKnowledgeBaseConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (b *jsiiProxy_BedrockagentKnowledgeBaseKnowledgeBaseConfigurationManagedKnowledgeBaseConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (b *jsiiProxy_BedrockagentKnowledgeBaseKnowledgeBaseConfigurationManagedKnowledgeBaseConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (b *jsiiProxy_BedrockagentKnowledgeBaseKnowledgeBaseConfigurationManagedKnowledgeBaseConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (b *jsiiProxy_BedrockagentKnowledgeBaseKnowledgeBaseConfigurationManagedKnowledgeBaseConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentKnowledgeBaseKnowledgeBaseConfigurationManagedKnowledgeBaseConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (b *jsiiProxy_BedrockagentKnowledgeBaseKnowledgeBaseConfigurationManagedKnowledgeBaseConfigurationOutputReference) PutEmbeddingModelConfiguration(value interface{}) {
	if err := b.validatePutEmbeddingModelConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putEmbeddingModelConfiguration",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentKnowledgeBaseKnowledgeBaseConfigurationManagedKnowledgeBaseConfigurationOutputReference) PutServerSideEncryptionConfiguration(value interface{}) {
	if err := b.validatePutServerSideEncryptionConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putServerSideEncryptionConfiguration",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentKnowledgeBaseKnowledgeBaseConfigurationManagedKnowledgeBaseConfigurationOutputReference) ResetEmbeddingModelArn() {
	_jsii_.InvokeVoid(
		b,
		"resetEmbeddingModelArn",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentKnowledgeBaseKnowledgeBaseConfigurationManagedKnowledgeBaseConfigurationOutputReference) ResetEmbeddingModelConfiguration() {
	_jsii_.InvokeVoid(
		b,
		"resetEmbeddingModelConfiguration",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentKnowledgeBaseKnowledgeBaseConfigurationManagedKnowledgeBaseConfigurationOutputReference) ResetEmbeddingModelType() {
	_jsii_.InvokeVoid(
		b,
		"resetEmbeddingModelType",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentKnowledgeBaseKnowledgeBaseConfigurationManagedKnowledgeBaseConfigurationOutputReference) ResetServerSideEncryptionConfiguration() {
	_jsii_.InvokeVoid(
		b,
		"resetServerSideEncryptionConfiguration",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentKnowledgeBaseKnowledgeBaseConfigurationManagedKnowledgeBaseConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (b *jsiiProxy_BedrockagentKnowledgeBaseKnowledgeBaseConfigurationManagedKnowledgeBaseConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

