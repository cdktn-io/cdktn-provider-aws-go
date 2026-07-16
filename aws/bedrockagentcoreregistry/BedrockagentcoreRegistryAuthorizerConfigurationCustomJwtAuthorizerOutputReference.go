// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreregistry

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v24/jsii"

	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v24/bedrockagentcoreregistry/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BedrockagentcoreRegistryAuthorizerConfigurationCustomJwtAuthorizerOutputReference interface {
	cdktn.ComplexObject
	AllowedAudience() *[]*string
	SetAllowedAudience(val *[]*string)
	AllowedAudienceInput() *[]*string
	AllowedClients() *[]*string
	SetAllowedClients(val *[]*string)
	AllowedClientsInput() *[]*string
	AllowedScopes() *[]*string
	SetAllowedScopes(val *[]*string)
	AllowedScopesInput() *[]*string
	AllowedWorkloadConfiguration() BedrockagentcoreRegistryAuthorizerConfigurationCustomJwtAuthorizerAllowedWorkloadConfigurationList
	AllowedWorkloadConfigurationInput() interface{}
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
	CustomClaim() BedrockagentcoreRegistryAuthorizerConfigurationCustomJwtAuthorizerCustomClaimList
	CustomClaimInput() interface{}
	DiscoveryUrl() *string
	SetDiscoveryUrl(val *string)
	DiscoveryUrlInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	PrivateEndpoint() BedrockagentcoreRegistryAuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointList
	PrivateEndpointInput() interface{}
	PrivateEndpointOverrides() BedrockagentcoreRegistryAuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointOverridesList
	PrivateEndpointOverridesInput() interface{}
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
	PutAllowedWorkloadConfiguration(value interface{})
	PutCustomClaim(value interface{})
	PutPrivateEndpoint(value interface{})
	PutPrivateEndpointOverrides(value interface{})
	ResetAllowedAudience()
	ResetAllowedClients()
	ResetAllowedScopes()
	ResetAllowedWorkloadConfiguration()
	ResetCustomClaim()
	ResetPrivateEndpoint()
	ResetPrivateEndpointOverrides()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BedrockagentcoreRegistryAuthorizerConfigurationCustomJwtAuthorizerOutputReference
type jsiiProxy_BedrockagentcoreRegistryAuthorizerConfigurationCustomJwtAuthorizerOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_BedrockagentcoreRegistryAuthorizerConfigurationCustomJwtAuthorizerOutputReference) AllowedAudience() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedAudience",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreRegistryAuthorizerConfigurationCustomJwtAuthorizerOutputReference) AllowedAudienceInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedAudienceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreRegistryAuthorizerConfigurationCustomJwtAuthorizerOutputReference) AllowedClients() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedClients",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreRegistryAuthorizerConfigurationCustomJwtAuthorizerOutputReference) AllowedClientsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedClientsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreRegistryAuthorizerConfigurationCustomJwtAuthorizerOutputReference) AllowedScopes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedScopes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreRegistryAuthorizerConfigurationCustomJwtAuthorizerOutputReference) AllowedScopesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedScopesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreRegistryAuthorizerConfigurationCustomJwtAuthorizerOutputReference) AllowedWorkloadConfiguration() BedrockagentcoreRegistryAuthorizerConfigurationCustomJwtAuthorizerAllowedWorkloadConfigurationList {
	var returns BedrockagentcoreRegistryAuthorizerConfigurationCustomJwtAuthorizerAllowedWorkloadConfigurationList
	_jsii_.Get(
		j,
		"allowedWorkloadConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreRegistryAuthorizerConfigurationCustomJwtAuthorizerOutputReference) AllowedWorkloadConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowedWorkloadConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreRegistryAuthorizerConfigurationCustomJwtAuthorizerOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreRegistryAuthorizerConfigurationCustomJwtAuthorizerOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreRegistryAuthorizerConfigurationCustomJwtAuthorizerOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreRegistryAuthorizerConfigurationCustomJwtAuthorizerOutputReference) CustomClaim() BedrockagentcoreRegistryAuthorizerConfigurationCustomJwtAuthorizerCustomClaimList {
	var returns BedrockagentcoreRegistryAuthorizerConfigurationCustomJwtAuthorizerCustomClaimList
	_jsii_.Get(
		j,
		"customClaim",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreRegistryAuthorizerConfigurationCustomJwtAuthorizerOutputReference) CustomClaimInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customClaimInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreRegistryAuthorizerConfigurationCustomJwtAuthorizerOutputReference) DiscoveryUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"discoveryUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreRegistryAuthorizerConfigurationCustomJwtAuthorizerOutputReference) DiscoveryUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"discoveryUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreRegistryAuthorizerConfigurationCustomJwtAuthorizerOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreRegistryAuthorizerConfigurationCustomJwtAuthorizerOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreRegistryAuthorizerConfigurationCustomJwtAuthorizerOutputReference) PrivateEndpoint() BedrockagentcoreRegistryAuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointList {
	var returns BedrockagentcoreRegistryAuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointList
	_jsii_.Get(
		j,
		"privateEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreRegistryAuthorizerConfigurationCustomJwtAuthorizerOutputReference) PrivateEndpointInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privateEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreRegistryAuthorizerConfigurationCustomJwtAuthorizerOutputReference) PrivateEndpointOverrides() BedrockagentcoreRegistryAuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointOverridesList {
	var returns BedrockagentcoreRegistryAuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointOverridesList
	_jsii_.Get(
		j,
		"privateEndpointOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreRegistryAuthorizerConfigurationCustomJwtAuthorizerOutputReference) PrivateEndpointOverridesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privateEndpointOverridesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreRegistryAuthorizerConfigurationCustomJwtAuthorizerOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreRegistryAuthorizerConfigurationCustomJwtAuthorizerOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBedrockagentcoreRegistryAuthorizerConfigurationCustomJwtAuthorizerOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) BedrockagentcoreRegistryAuthorizerConfigurationCustomJwtAuthorizerOutputReference {
	_init_.Initialize()

	if err := validateNewBedrockagentcoreRegistryAuthorizerConfigurationCustomJwtAuthorizerOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_BedrockagentcoreRegistryAuthorizerConfigurationCustomJwtAuthorizerOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-aws.bedrockagentcoreRegistry.BedrockagentcoreRegistryAuthorizerConfigurationCustomJwtAuthorizerOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewBedrockagentcoreRegistryAuthorizerConfigurationCustomJwtAuthorizerOutputReference_Override(b BedrockagentcoreRegistryAuthorizerConfigurationCustomJwtAuthorizerOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.bedrockagentcoreRegistry.BedrockagentcoreRegistryAuthorizerConfigurationCustomJwtAuthorizerOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		b,
	)
}

func (j *jsiiProxy_BedrockagentcoreRegistryAuthorizerConfigurationCustomJwtAuthorizerOutputReference)SetAllowedAudience(val *[]*string) {
	if err := j.validateSetAllowedAudienceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedAudience",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreRegistryAuthorizerConfigurationCustomJwtAuthorizerOutputReference)SetAllowedClients(val *[]*string) {
	if err := j.validateSetAllowedClientsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedClients",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreRegistryAuthorizerConfigurationCustomJwtAuthorizerOutputReference)SetAllowedScopes(val *[]*string) {
	if err := j.validateSetAllowedScopesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedScopes",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreRegistryAuthorizerConfigurationCustomJwtAuthorizerOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreRegistryAuthorizerConfigurationCustomJwtAuthorizerOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreRegistryAuthorizerConfigurationCustomJwtAuthorizerOutputReference)SetDiscoveryUrl(val *string) {
	if err := j.validateSetDiscoveryUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"discoveryUrl",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreRegistryAuthorizerConfigurationCustomJwtAuthorizerOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreRegistryAuthorizerConfigurationCustomJwtAuthorizerOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreRegistryAuthorizerConfigurationCustomJwtAuthorizerOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BedrockagentcoreRegistryAuthorizerConfigurationCustomJwtAuthorizerOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreRegistryAuthorizerConfigurationCustomJwtAuthorizerOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (b *jsiiProxy_BedrockagentcoreRegistryAuthorizerConfigurationCustomJwtAuthorizerOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (b *jsiiProxy_BedrockagentcoreRegistryAuthorizerConfigurationCustomJwtAuthorizerOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (b *jsiiProxy_BedrockagentcoreRegistryAuthorizerConfigurationCustomJwtAuthorizerOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (b *jsiiProxy_BedrockagentcoreRegistryAuthorizerConfigurationCustomJwtAuthorizerOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (b *jsiiProxy_BedrockagentcoreRegistryAuthorizerConfigurationCustomJwtAuthorizerOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (b *jsiiProxy_BedrockagentcoreRegistryAuthorizerConfigurationCustomJwtAuthorizerOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (b *jsiiProxy_BedrockagentcoreRegistryAuthorizerConfigurationCustomJwtAuthorizerOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (b *jsiiProxy_BedrockagentcoreRegistryAuthorizerConfigurationCustomJwtAuthorizerOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (b *jsiiProxy_BedrockagentcoreRegistryAuthorizerConfigurationCustomJwtAuthorizerOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreRegistryAuthorizerConfigurationCustomJwtAuthorizerOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (b *jsiiProxy_BedrockagentcoreRegistryAuthorizerConfigurationCustomJwtAuthorizerOutputReference) PutAllowedWorkloadConfiguration(value interface{}) {
	if err := b.validatePutAllowedWorkloadConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putAllowedWorkloadConfiguration",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentcoreRegistryAuthorizerConfigurationCustomJwtAuthorizerOutputReference) PutCustomClaim(value interface{}) {
	if err := b.validatePutCustomClaimParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putCustomClaim",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentcoreRegistryAuthorizerConfigurationCustomJwtAuthorizerOutputReference) PutPrivateEndpoint(value interface{}) {
	if err := b.validatePutPrivateEndpointParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putPrivateEndpoint",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentcoreRegistryAuthorizerConfigurationCustomJwtAuthorizerOutputReference) PutPrivateEndpointOverrides(value interface{}) {
	if err := b.validatePutPrivateEndpointOverridesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putPrivateEndpointOverrides",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentcoreRegistryAuthorizerConfigurationCustomJwtAuthorizerOutputReference) ResetAllowedAudience() {
	_jsii_.InvokeVoid(
		b,
		"resetAllowedAudience",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreRegistryAuthorizerConfigurationCustomJwtAuthorizerOutputReference) ResetAllowedClients() {
	_jsii_.InvokeVoid(
		b,
		"resetAllowedClients",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreRegistryAuthorizerConfigurationCustomJwtAuthorizerOutputReference) ResetAllowedScopes() {
	_jsii_.InvokeVoid(
		b,
		"resetAllowedScopes",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreRegistryAuthorizerConfigurationCustomJwtAuthorizerOutputReference) ResetAllowedWorkloadConfiguration() {
	_jsii_.InvokeVoid(
		b,
		"resetAllowedWorkloadConfiguration",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreRegistryAuthorizerConfigurationCustomJwtAuthorizerOutputReference) ResetCustomClaim() {
	_jsii_.InvokeVoid(
		b,
		"resetCustomClaim",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreRegistryAuthorizerConfigurationCustomJwtAuthorizerOutputReference) ResetPrivateEndpoint() {
	_jsii_.InvokeVoid(
		b,
		"resetPrivateEndpoint",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreRegistryAuthorizerConfigurationCustomJwtAuthorizerOutputReference) ResetPrivateEndpointOverrides() {
	_jsii_.InvokeVoid(
		b,
		"resetPrivateEndpointOverrides",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreRegistryAuthorizerConfigurationCustomJwtAuthorizerOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (b *jsiiProxy_BedrockagentcoreRegistryAuthorizerConfigurationCustomJwtAuthorizerOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

