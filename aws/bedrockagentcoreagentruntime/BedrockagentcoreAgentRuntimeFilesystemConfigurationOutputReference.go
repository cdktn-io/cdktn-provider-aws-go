// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreagentruntime

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v24/jsii"

	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v24/bedrockagentcoreagentruntime/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BedrockagentcoreAgentRuntimeFilesystemConfigurationOutputReference interface {
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
	EfsAccessPoint() BedrockagentcoreAgentRuntimeFilesystemConfigurationEfsAccessPointList
	EfsAccessPointInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	S3FilesAccessPoint() BedrockagentcoreAgentRuntimeFilesystemConfigurationS3FilesAccessPointList
	S3FilesAccessPointInput() interface{}
	SessionStorage() BedrockagentcoreAgentRuntimeFilesystemConfigurationSessionStorageList
	SessionStorageInput() interface{}
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
	PutEfsAccessPoint(value interface{})
	PutS3FilesAccessPoint(value interface{})
	PutSessionStorage(value interface{})
	ResetEfsAccessPoint()
	ResetS3FilesAccessPoint()
	ResetSessionStorage()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BedrockagentcoreAgentRuntimeFilesystemConfigurationOutputReference
type jsiiProxy_BedrockagentcoreAgentRuntimeFilesystemConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_BedrockagentcoreAgentRuntimeFilesystemConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreAgentRuntimeFilesystemConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreAgentRuntimeFilesystemConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreAgentRuntimeFilesystemConfigurationOutputReference) EfsAccessPoint() BedrockagentcoreAgentRuntimeFilesystemConfigurationEfsAccessPointList {
	var returns BedrockagentcoreAgentRuntimeFilesystemConfigurationEfsAccessPointList
	_jsii_.Get(
		j,
		"efsAccessPoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreAgentRuntimeFilesystemConfigurationOutputReference) EfsAccessPointInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"efsAccessPointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreAgentRuntimeFilesystemConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreAgentRuntimeFilesystemConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreAgentRuntimeFilesystemConfigurationOutputReference) S3FilesAccessPoint() BedrockagentcoreAgentRuntimeFilesystemConfigurationS3FilesAccessPointList {
	var returns BedrockagentcoreAgentRuntimeFilesystemConfigurationS3FilesAccessPointList
	_jsii_.Get(
		j,
		"s3FilesAccessPoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreAgentRuntimeFilesystemConfigurationOutputReference) S3FilesAccessPointInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"s3FilesAccessPointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreAgentRuntimeFilesystemConfigurationOutputReference) SessionStorage() BedrockagentcoreAgentRuntimeFilesystemConfigurationSessionStorageList {
	var returns BedrockagentcoreAgentRuntimeFilesystemConfigurationSessionStorageList
	_jsii_.Get(
		j,
		"sessionStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreAgentRuntimeFilesystemConfigurationOutputReference) SessionStorageInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sessionStorageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreAgentRuntimeFilesystemConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreAgentRuntimeFilesystemConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBedrockagentcoreAgentRuntimeFilesystemConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) BedrockagentcoreAgentRuntimeFilesystemConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewBedrockagentcoreAgentRuntimeFilesystemConfigurationOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_BedrockagentcoreAgentRuntimeFilesystemConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-aws.bedrockagentcoreAgentRuntime.BedrockagentcoreAgentRuntimeFilesystemConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewBedrockagentcoreAgentRuntimeFilesystemConfigurationOutputReference_Override(b BedrockagentcoreAgentRuntimeFilesystemConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.bedrockagentcoreAgentRuntime.BedrockagentcoreAgentRuntimeFilesystemConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		b,
	)
}

func (j *jsiiProxy_BedrockagentcoreAgentRuntimeFilesystemConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreAgentRuntimeFilesystemConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreAgentRuntimeFilesystemConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreAgentRuntimeFilesystemConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreAgentRuntimeFilesystemConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BedrockagentcoreAgentRuntimeFilesystemConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreAgentRuntimeFilesystemConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (b *jsiiProxy_BedrockagentcoreAgentRuntimeFilesystemConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (b *jsiiProxy_BedrockagentcoreAgentRuntimeFilesystemConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (b *jsiiProxy_BedrockagentcoreAgentRuntimeFilesystemConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (b *jsiiProxy_BedrockagentcoreAgentRuntimeFilesystemConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (b *jsiiProxy_BedrockagentcoreAgentRuntimeFilesystemConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (b *jsiiProxy_BedrockagentcoreAgentRuntimeFilesystemConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (b *jsiiProxy_BedrockagentcoreAgentRuntimeFilesystemConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (b *jsiiProxy_BedrockagentcoreAgentRuntimeFilesystemConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (b *jsiiProxy_BedrockagentcoreAgentRuntimeFilesystemConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreAgentRuntimeFilesystemConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (b *jsiiProxy_BedrockagentcoreAgentRuntimeFilesystemConfigurationOutputReference) PutEfsAccessPoint(value interface{}) {
	if err := b.validatePutEfsAccessPointParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putEfsAccessPoint",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentcoreAgentRuntimeFilesystemConfigurationOutputReference) PutS3FilesAccessPoint(value interface{}) {
	if err := b.validatePutS3FilesAccessPointParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putS3FilesAccessPoint",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentcoreAgentRuntimeFilesystemConfigurationOutputReference) PutSessionStorage(value interface{}) {
	if err := b.validatePutSessionStorageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putSessionStorage",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentcoreAgentRuntimeFilesystemConfigurationOutputReference) ResetEfsAccessPoint() {
	_jsii_.InvokeVoid(
		b,
		"resetEfsAccessPoint",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreAgentRuntimeFilesystemConfigurationOutputReference) ResetS3FilesAccessPoint() {
	_jsii_.InvokeVoid(
		b,
		"resetS3FilesAccessPoint",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreAgentRuntimeFilesystemConfigurationOutputReference) ResetSessionStorage() {
	_jsii_.InvokeVoid(
		b,
		"resetSessionStorage",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreAgentRuntimeFilesystemConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (b *jsiiProxy_BedrockagentcoreAgentRuntimeFilesystemConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

