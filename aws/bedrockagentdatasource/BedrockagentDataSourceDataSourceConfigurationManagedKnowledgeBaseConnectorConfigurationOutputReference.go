// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentdatasource

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v24/jsii"

	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v24/bedrockagentdatasource/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BedrockagentDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationOutputReference interface {
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
	ConnectorParameters() *string
	SetConnectorParameters(val *string)
	ConnectorParametersInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DeletionProtectionConfiguration() BedrockagentDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationDeletionProtectionConfigurationList
	DeletionProtectionConfigurationInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MediaExtractionConfiguration() BedrockagentDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationMediaExtractionConfigurationList
	MediaExtractionConfigurationInput() interface{}
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
	PutDeletionProtectionConfiguration(value interface{})
	PutMediaExtractionConfiguration(value interface{})
	ResetConnectorParameters()
	ResetDeletionProtectionConfiguration()
	ResetMediaExtractionConfiguration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BedrockagentDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationOutputReference
type jsiiProxy_BedrockagentDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_BedrockagentDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationOutputReference) ConnectorParameters() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectorParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationOutputReference) ConnectorParametersInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectorParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationOutputReference) DeletionProtectionConfiguration() BedrockagentDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationDeletionProtectionConfigurationList {
	var returns BedrockagentDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationDeletionProtectionConfigurationList
	_jsii_.Get(
		j,
		"deletionProtectionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationOutputReference) DeletionProtectionConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtectionConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationOutputReference) MediaExtractionConfiguration() BedrockagentDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationMediaExtractionConfigurationList {
	var returns BedrockagentDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationMediaExtractionConfigurationList
	_jsii_.Get(
		j,
		"mediaExtractionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationOutputReference) MediaExtractionConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mediaExtractionConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBedrockagentDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) BedrockagentDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewBedrockagentDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_BedrockagentDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-aws.bedrockagentDataSource.BedrockagentDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewBedrockagentDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationOutputReference_Override(b BedrockagentDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.bedrockagentDataSource.BedrockagentDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		b,
	)
}

func (j *jsiiProxy_BedrockagentDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BedrockagentDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BedrockagentDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationOutputReference)SetConnectorParameters(val *string) {
	if err := j.validateSetConnectorParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectorParameters",
		val,
	)
}

func (j *jsiiProxy_BedrockagentDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BedrockagentDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BedrockagentDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BedrockagentDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (b *jsiiProxy_BedrockagentDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (b *jsiiProxy_BedrockagentDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (b *jsiiProxy_BedrockagentDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (b *jsiiProxy_BedrockagentDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (b *jsiiProxy_BedrockagentDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (b *jsiiProxy_BedrockagentDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (b *jsiiProxy_BedrockagentDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (b *jsiiProxy_BedrockagentDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (b *jsiiProxy_BedrockagentDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (b *jsiiProxy_BedrockagentDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationOutputReference) PutDeletionProtectionConfiguration(value interface{}) {
	if err := b.validatePutDeletionProtectionConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putDeletionProtectionConfiguration",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationOutputReference) PutMediaExtractionConfiguration(value interface{}) {
	if err := b.validatePutMediaExtractionConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putMediaExtractionConfiguration",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationOutputReference) ResetConnectorParameters() {
	_jsii_.InvokeVoid(
		b,
		"resetConnectorParameters",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationOutputReference) ResetDeletionProtectionConfiguration() {
	_jsii_.InvokeVoid(
		b,
		"resetDeletionProtectionConfiguration",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationOutputReference) ResetMediaExtractionConfiguration() {
	_jsii_.InvokeVoid(
		b,
		"resetMediaExtractionConfiguration",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (b *jsiiProxy_BedrockagentDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

