// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreharness

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/jsii"

	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/bedrockagentcoreharness/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BedrockagentcoreHarnessEnvironmentActualAgentcoreRuntimeEnvironmentOutputReference interface {
	cdktn.ComplexObject
	AgentRuntimeArn() *string
	AgentRuntimeId() *string
	AgentRuntimeName() *string
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
	FilesystemConfiguration() BedrockagentcoreHarnessEnvironmentActualAgentcoreRuntimeEnvironmentFilesystemConfigurationList
	// Experimental.
	Fqn() *string
	InternalValue() *BedrockagentcoreHarnessEnvironmentActualAgentcoreRuntimeEnvironment
	SetInternalValue(val *BedrockagentcoreHarnessEnvironmentActualAgentcoreRuntimeEnvironment)
	LifecycleConfiguration() BedrockagentcoreHarnessEnvironmentActualAgentcoreRuntimeEnvironmentLifecycleConfigurationList
	NetworkConfiguration() BedrockagentcoreHarnessEnvironmentActualAgentcoreRuntimeEnvironmentNetworkConfigurationList
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

// The jsii proxy struct for BedrockagentcoreHarnessEnvironmentActualAgentcoreRuntimeEnvironmentOutputReference
type jsiiProxy_BedrockagentcoreHarnessEnvironmentActualAgentcoreRuntimeEnvironmentOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_BedrockagentcoreHarnessEnvironmentActualAgentcoreRuntimeEnvironmentOutputReference) AgentRuntimeArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentRuntimeArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarnessEnvironmentActualAgentcoreRuntimeEnvironmentOutputReference) AgentRuntimeId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentRuntimeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarnessEnvironmentActualAgentcoreRuntimeEnvironmentOutputReference) AgentRuntimeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentRuntimeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarnessEnvironmentActualAgentcoreRuntimeEnvironmentOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarnessEnvironmentActualAgentcoreRuntimeEnvironmentOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarnessEnvironmentActualAgentcoreRuntimeEnvironmentOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarnessEnvironmentActualAgentcoreRuntimeEnvironmentOutputReference) FilesystemConfiguration() BedrockagentcoreHarnessEnvironmentActualAgentcoreRuntimeEnvironmentFilesystemConfigurationList {
	var returns BedrockagentcoreHarnessEnvironmentActualAgentcoreRuntimeEnvironmentFilesystemConfigurationList
	_jsii_.Get(
		j,
		"filesystemConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarnessEnvironmentActualAgentcoreRuntimeEnvironmentOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarnessEnvironmentActualAgentcoreRuntimeEnvironmentOutputReference) InternalValue() *BedrockagentcoreHarnessEnvironmentActualAgentcoreRuntimeEnvironment {
	var returns *BedrockagentcoreHarnessEnvironmentActualAgentcoreRuntimeEnvironment
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarnessEnvironmentActualAgentcoreRuntimeEnvironmentOutputReference) LifecycleConfiguration() BedrockagentcoreHarnessEnvironmentActualAgentcoreRuntimeEnvironmentLifecycleConfigurationList {
	var returns BedrockagentcoreHarnessEnvironmentActualAgentcoreRuntimeEnvironmentLifecycleConfigurationList
	_jsii_.Get(
		j,
		"lifecycleConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarnessEnvironmentActualAgentcoreRuntimeEnvironmentOutputReference) NetworkConfiguration() BedrockagentcoreHarnessEnvironmentActualAgentcoreRuntimeEnvironmentNetworkConfigurationList {
	var returns BedrockagentcoreHarnessEnvironmentActualAgentcoreRuntimeEnvironmentNetworkConfigurationList
	_jsii_.Get(
		j,
		"networkConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarnessEnvironmentActualAgentcoreRuntimeEnvironmentOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarnessEnvironmentActualAgentcoreRuntimeEnvironmentOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBedrockagentcoreHarnessEnvironmentActualAgentcoreRuntimeEnvironmentOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) BedrockagentcoreHarnessEnvironmentActualAgentcoreRuntimeEnvironmentOutputReference {
	_init_.Initialize()

	if err := validateNewBedrockagentcoreHarnessEnvironmentActualAgentcoreRuntimeEnvironmentOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_BedrockagentcoreHarnessEnvironmentActualAgentcoreRuntimeEnvironmentOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-aws.bedrockagentcoreHarness.BedrockagentcoreHarnessEnvironmentActualAgentcoreRuntimeEnvironmentOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewBedrockagentcoreHarnessEnvironmentActualAgentcoreRuntimeEnvironmentOutputReference_Override(b BedrockagentcoreHarnessEnvironmentActualAgentcoreRuntimeEnvironmentOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.bedrockagentcoreHarness.BedrockagentcoreHarnessEnvironmentActualAgentcoreRuntimeEnvironmentOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		b,
	)
}

func (j *jsiiProxy_BedrockagentcoreHarnessEnvironmentActualAgentcoreRuntimeEnvironmentOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreHarnessEnvironmentActualAgentcoreRuntimeEnvironmentOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreHarnessEnvironmentActualAgentcoreRuntimeEnvironmentOutputReference)SetInternalValue(val *BedrockagentcoreHarnessEnvironmentActualAgentcoreRuntimeEnvironment) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreHarnessEnvironmentActualAgentcoreRuntimeEnvironmentOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreHarnessEnvironmentActualAgentcoreRuntimeEnvironmentOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BedrockagentcoreHarnessEnvironmentActualAgentcoreRuntimeEnvironmentOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreHarnessEnvironmentActualAgentcoreRuntimeEnvironmentOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (b *jsiiProxy_BedrockagentcoreHarnessEnvironmentActualAgentcoreRuntimeEnvironmentOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (b *jsiiProxy_BedrockagentcoreHarnessEnvironmentActualAgentcoreRuntimeEnvironmentOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (b *jsiiProxy_BedrockagentcoreHarnessEnvironmentActualAgentcoreRuntimeEnvironmentOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (b *jsiiProxy_BedrockagentcoreHarnessEnvironmentActualAgentcoreRuntimeEnvironmentOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (b *jsiiProxy_BedrockagentcoreHarnessEnvironmentActualAgentcoreRuntimeEnvironmentOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (b *jsiiProxy_BedrockagentcoreHarnessEnvironmentActualAgentcoreRuntimeEnvironmentOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (b *jsiiProxy_BedrockagentcoreHarnessEnvironmentActualAgentcoreRuntimeEnvironmentOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (b *jsiiProxy_BedrockagentcoreHarnessEnvironmentActualAgentcoreRuntimeEnvironmentOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (b *jsiiProxy_BedrockagentcoreHarnessEnvironmentActualAgentcoreRuntimeEnvironmentOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreHarnessEnvironmentActualAgentcoreRuntimeEnvironmentOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (b *jsiiProxy_BedrockagentcoreHarnessEnvironmentActualAgentcoreRuntimeEnvironmentOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (b *jsiiProxy_BedrockagentcoreHarnessEnvironmentActualAgentcoreRuntimeEnvironmentOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

