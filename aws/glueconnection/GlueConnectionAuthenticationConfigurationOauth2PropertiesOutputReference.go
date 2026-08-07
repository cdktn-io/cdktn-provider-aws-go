// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package glueconnection

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/jsii"

	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/glueconnection/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GlueConnectionAuthenticationConfigurationOauth2PropertiesOutputReference interface {
	cdktn.ComplexObject
	AuthorizationCodeProperties() GlueConnectionAuthenticationConfigurationOauth2PropertiesAuthorizationCodePropertiesOutputReference
	AuthorizationCodePropertiesInput() *GlueConnectionAuthenticationConfigurationOauth2PropertiesAuthorizationCodeProperties
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
	InternalValue() *GlueConnectionAuthenticationConfigurationOauth2Properties
	SetInternalValue(val *GlueConnectionAuthenticationConfigurationOauth2Properties)
	Oauth2ClientApplication() GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2ClientApplicationOutputReference
	Oauth2ClientApplicationInput() *GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2ClientApplication
	Oauth2Credentials() GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2CredentialsOutputReference
	Oauth2CredentialsInput() *GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2Credentials
	Oauth2GrantType() *string
	SetOauth2GrantType(val *string)
	Oauth2GrantTypeInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TokenUrl() *string
	SetTokenUrl(val *string)
	TokenUrlInput() *string
	TokenUrlParametersMap() *map[string]*string
	SetTokenUrlParametersMap(val *map[string]*string)
	TokenUrlParametersMapInput() *map[string]*string
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
	PutAuthorizationCodeProperties(value *GlueConnectionAuthenticationConfigurationOauth2PropertiesAuthorizationCodeProperties)
	PutOauth2ClientApplication(value *GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2ClientApplication)
	PutOauth2Credentials(value *GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2Credentials)
	ResetAuthorizationCodeProperties()
	ResetOauth2ClientApplication()
	ResetOauth2Credentials()
	ResetOauth2GrantType()
	ResetTokenUrl()
	ResetTokenUrlParametersMap()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GlueConnectionAuthenticationConfigurationOauth2PropertiesOutputReference
type jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOutputReference) AuthorizationCodeProperties() GlueConnectionAuthenticationConfigurationOauth2PropertiesAuthorizationCodePropertiesOutputReference {
	var returns GlueConnectionAuthenticationConfigurationOauth2PropertiesAuthorizationCodePropertiesOutputReference
	_jsii_.Get(
		j,
		"authorizationCodeProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOutputReference) AuthorizationCodePropertiesInput() *GlueConnectionAuthenticationConfigurationOauth2PropertiesAuthorizationCodeProperties {
	var returns *GlueConnectionAuthenticationConfigurationOauth2PropertiesAuthorizationCodeProperties
	_jsii_.Get(
		j,
		"authorizationCodePropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOutputReference) InternalValue() *GlueConnectionAuthenticationConfigurationOauth2Properties {
	var returns *GlueConnectionAuthenticationConfigurationOauth2Properties
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOutputReference) Oauth2ClientApplication() GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2ClientApplicationOutputReference {
	var returns GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2ClientApplicationOutputReference
	_jsii_.Get(
		j,
		"oauth2ClientApplication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOutputReference) Oauth2ClientApplicationInput() *GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2ClientApplication {
	var returns *GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2ClientApplication
	_jsii_.Get(
		j,
		"oauth2ClientApplicationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOutputReference) Oauth2Credentials() GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2CredentialsOutputReference {
	var returns GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2CredentialsOutputReference
	_jsii_.Get(
		j,
		"oauth2Credentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOutputReference) Oauth2CredentialsInput() *GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2Credentials {
	var returns *GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2Credentials
	_jsii_.Get(
		j,
		"oauth2CredentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOutputReference) Oauth2GrantType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauth2GrantType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOutputReference) Oauth2GrantTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauth2GrantTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOutputReference) TokenUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOutputReference) TokenUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOutputReference) TokenUrlParametersMap() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tokenUrlParametersMap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOutputReference) TokenUrlParametersMapInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tokenUrlParametersMapInput",
		&returns,
	)
	return returns
}


func NewGlueConnectionAuthenticationConfigurationOauth2PropertiesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GlueConnectionAuthenticationConfigurationOauth2PropertiesOutputReference {
	_init_.Initialize()

	if err := validateNewGlueConnectionAuthenticationConfigurationOauth2PropertiesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-aws.glueConnection.GlueConnectionAuthenticationConfigurationOauth2PropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGlueConnectionAuthenticationConfigurationOauth2PropertiesOutputReference_Override(g GlueConnectionAuthenticationConfigurationOauth2PropertiesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.glueConnection.GlueConnectionAuthenticationConfigurationOauth2PropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOutputReference)SetInternalValue(val *GlueConnectionAuthenticationConfigurationOauth2Properties) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOutputReference)SetOauth2GrantType(val *string) {
	if err := j.validateSetOauth2GrantTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oauth2GrantType",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOutputReference)SetTokenUrl(val *string) {
	if err := j.validateSetTokenUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenUrl",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOutputReference)SetTokenUrlParametersMap(val *map[string]*string) {
	if err := j.validateSetTokenUrlParametersMapParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenUrlParametersMap",
		val,
	)
}

func (g *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOutputReference) PutAuthorizationCodeProperties(value *GlueConnectionAuthenticationConfigurationOauth2PropertiesAuthorizationCodeProperties) {
	if err := g.validatePutAuthorizationCodePropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAuthorizationCodeProperties",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOutputReference) PutOauth2ClientApplication(value *GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2ClientApplication) {
	if err := g.validatePutOauth2ClientApplicationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putOauth2ClientApplication",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOutputReference) PutOauth2Credentials(value *GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2Credentials) {
	if err := g.validatePutOauth2CredentialsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putOauth2Credentials",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOutputReference) ResetAuthorizationCodeProperties() {
	_jsii_.InvokeVoid(
		g,
		"resetAuthorizationCodeProperties",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOutputReference) ResetOauth2ClientApplication() {
	_jsii_.InvokeVoid(
		g,
		"resetOauth2ClientApplication",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOutputReference) ResetOauth2Credentials() {
	_jsii_.InvokeVoid(
		g,
		"resetOauth2Credentials",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOutputReference) ResetOauth2GrantType() {
	_jsii_.InvokeVoid(
		g,
		"resetOauth2GrantType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOutputReference) ResetTokenUrl() {
	_jsii_.InvokeVoid(
		g,
		"resetTokenUrl",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOutputReference) ResetTokenUrlParametersMap() {
	_jsii_.InvokeVoid(
		g,
		"resetTokenUrlParametersMap",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

