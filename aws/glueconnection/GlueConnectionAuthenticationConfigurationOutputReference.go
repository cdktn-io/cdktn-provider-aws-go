// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package glueconnection

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/jsii"

	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/glueconnection/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GlueConnectionAuthenticationConfigurationOutputReference interface {
	cdktn.ComplexObject
	AuthenticationType() *string
	SetAuthenticationType(val *string)
	AuthenticationTypeInput() *string
	BasicAuthenticationCredentials() GlueConnectionAuthenticationConfigurationBasicAuthenticationCredentialsOutputReference
	BasicAuthenticationCredentialsInput() *GlueConnectionAuthenticationConfigurationBasicAuthenticationCredentials
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
	CustomAuthenticationCredentials() *map[string]*string
	SetCustomAuthenticationCredentials(val *map[string]*string)
	CustomAuthenticationCredentialsInput() *map[string]*string
	// Experimental.
	Fqn() *string
	InternalValue() *GlueConnectionAuthenticationConfiguration
	SetInternalValue(val *GlueConnectionAuthenticationConfiguration)
	KmsKeyArn() *string
	SetKmsKeyArn(val *string)
	KmsKeyArnInput() *string
	Oauth2Properties() GlueConnectionAuthenticationConfigurationOauth2PropertiesOutputReference
	Oauth2PropertiesInput() *GlueConnectionAuthenticationConfigurationOauth2Properties
	SecretArn() *string
	SetSecretArn(val *string)
	SecretArnInput() *string
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
	PutBasicAuthenticationCredentials(value *GlueConnectionAuthenticationConfigurationBasicAuthenticationCredentials)
	PutOauth2Properties(value *GlueConnectionAuthenticationConfigurationOauth2Properties)
	ResetBasicAuthenticationCredentials()
	ResetCustomAuthenticationCredentials()
	ResetKmsKeyArn()
	ResetOauth2Properties()
	ResetSecretArn()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GlueConnectionAuthenticationConfigurationOutputReference
type jsiiProxy_GlueConnectionAuthenticationConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GlueConnectionAuthenticationConfigurationOutputReference) AuthenticationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionAuthenticationConfigurationOutputReference) AuthenticationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionAuthenticationConfigurationOutputReference) BasicAuthenticationCredentials() GlueConnectionAuthenticationConfigurationBasicAuthenticationCredentialsOutputReference {
	var returns GlueConnectionAuthenticationConfigurationBasicAuthenticationCredentialsOutputReference
	_jsii_.Get(
		j,
		"basicAuthenticationCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionAuthenticationConfigurationOutputReference) BasicAuthenticationCredentialsInput() *GlueConnectionAuthenticationConfigurationBasicAuthenticationCredentials {
	var returns *GlueConnectionAuthenticationConfigurationBasicAuthenticationCredentials
	_jsii_.Get(
		j,
		"basicAuthenticationCredentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionAuthenticationConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionAuthenticationConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionAuthenticationConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionAuthenticationConfigurationOutputReference) CustomAuthenticationCredentials() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"customAuthenticationCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionAuthenticationConfigurationOutputReference) CustomAuthenticationCredentialsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"customAuthenticationCredentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionAuthenticationConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionAuthenticationConfigurationOutputReference) InternalValue() *GlueConnectionAuthenticationConfiguration {
	var returns *GlueConnectionAuthenticationConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionAuthenticationConfigurationOutputReference) KmsKeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionAuthenticationConfigurationOutputReference) KmsKeyArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionAuthenticationConfigurationOutputReference) Oauth2Properties() GlueConnectionAuthenticationConfigurationOauth2PropertiesOutputReference {
	var returns GlueConnectionAuthenticationConfigurationOauth2PropertiesOutputReference
	_jsii_.Get(
		j,
		"oauth2Properties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionAuthenticationConfigurationOutputReference) Oauth2PropertiesInput() *GlueConnectionAuthenticationConfigurationOauth2Properties {
	var returns *GlueConnectionAuthenticationConfigurationOauth2Properties
	_jsii_.Get(
		j,
		"oauth2PropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionAuthenticationConfigurationOutputReference) SecretArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionAuthenticationConfigurationOutputReference) SecretArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionAuthenticationConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionAuthenticationConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGlueConnectionAuthenticationConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GlueConnectionAuthenticationConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewGlueConnectionAuthenticationConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GlueConnectionAuthenticationConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-aws.glueConnection.GlueConnectionAuthenticationConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGlueConnectionAuthenticationConfigurationOutputReference_Override(g GlueConnectionAuthenticationConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.glueConnection.GlueConnectionAuthenticationConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GlueConnectionAuthenticationConfigurationOutputReference)SetAuthenticationType(val *string) {
	if err := j.validateSetAuthenticationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authenticationType",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionAuthenticationConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionAuthenticationConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionAuthenticationConfigurationOutputReference)SetCustomAuthenticationCredentials(val *map[string]*string) {
	if err := j.validateSetCustomAuthenticationCredentialsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customAuthenticationCredentials",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionAuthenticationConfigurationOutputReference)SetInternalValue(val *GlueConnectionAuthenticationConfiguration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionAuthenticationConfigurationOutputReference)SetKmsKeyArn(val *string) {
	if err := j.validateSetKmsKeyArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyArn",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionAuthenticationConfigurationOutputReference)SetSecretArn(val *string) {
	if err := j.validateSetSecretArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretArn",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionAuthenticationConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionAuthenticationConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GlueConnectionAuthenticationConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueConnectionAuthenticationConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueConnectionAuthenticationConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueConnectionAuthenticationConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueConnectionAuthenticationConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueConnectionAuthenticationConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueConnectionAuthenticationConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueConnectionAuthenticationConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueConnectionAuthenticationConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueConnectionAuthenticationConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueConnectionAuthenticationConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueConnectionAuthenticationConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueConnectionAuthenticationConfigurationOutputReference) PutBasicAuthenticationCredentials(value *GlueConnectionAuthenticationConfigurationBasicAuthenticationCredentials) {
	if err := g.validatePutBasicAuthenticationCredentialsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putBasicAuthenticationCredentials",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueConnectionAuthenticationConfigurationOutputReference) PutOauth2Properties(value *GlueConnectionAuthenticationConfigurationOauth2Properties) {
	if err := g.validatePutOauth2PropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putOauth2Properties",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueConnectionAuthenticationConfigurationOutputReference) ResetBasicAuthenticationCredentials() {
	_jsii_.InvokeVoid(
		g,
		"resetBasicAuthenticationCredentials",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnectionAuthenticationConfigurationOutputReference) ResetCustomAuthenticationCredentials() {
	_jsii_.InvokeVoid(
		g,
		"resetCustomAuthenticationCredentials",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnectionAuthenticationConfigurationOutputReference) ResetKmsKeyArn() {
	_jsii_.InvokeVoid(
		g,
		"resetKmsKeyArn",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnectionAuthenticationConfigurationOutputReference) ResetOauth2Properties() {
	_jsii_.InvokeVoid(
		g,
		"resetOauth2Properties",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnectionAuthenticationConfigurationOutputReference) ResetSecretArn() {
	_jsii_.InvokeVoid(
		g,
		"resetSecretArn",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnectionAuthenticationConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := g.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueConnectionAuthenticationConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

