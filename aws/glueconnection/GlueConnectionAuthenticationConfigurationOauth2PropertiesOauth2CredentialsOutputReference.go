// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package glueconnection

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v24/jsii"

	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v24/glueconnection/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2CredentialsOutputReference interface {
	cdktn.ComplexObject
	AccessToken() *string
	SetAccessToken(val *string)
	AccessTokenInput() *string
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
	InternalValue() *GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2Credentials
	SetInternalValue(val *GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2Credentials)
	JwtToken() *string
	SetJwtToken(val *string)
	JwtTokenInput() *string
	RefreshToken() *string
	SetRefreshToken(val *string)
	RefreshTokenInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	UserManagedClientApplicationClientSecret() *string
	SetUserManagedClientApplicationClientSecret(val *string)
	UserManagedClientApplicationClientSecretInput() *string
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
	ResetAccessToken()
	ResetJwtToken()
	ResetRefreshToken()
	ResetUserManagedClientApplicationClientSecret()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2CredentialsOutputReference
type jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2CredentialsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2CredentialsOutputReference) AccessToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2CredentialsOutputReference) AccessTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2CredentialsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2CredentialsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2CredentialsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2CredentialsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2CredentialsOutputReference) InternalValue() *GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2Credentials {
	var returns *GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2Credentials
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2CredentialsOutputReference) JwtToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jwtToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2CredentialsOutputReference) JwtTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jwtTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2CredentialsOutputReference) RefreshToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"refreshToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2CredentialsOutputReference) RefreshTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"refreshTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2CredentialsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2CredentialsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2CredentialsOutputReference) UserManagedClientApplicationClientSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userManagedClientApplicationClientSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2CredentialsOutputReference) UserManagedClientApplicationClientSecretInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userManagedClientApplicationClientSecretInput",
		&returns,
	)
	return returns
}


func NewGlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2CredentialsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2CredentialsOutputReference {
	_init_.Initialize()

	if err := validateNewGlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2CredentialsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2CredentialsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-aws.glueConnection.GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2CredentialsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2CredentialsOutputReference_Override(g GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2CredentialsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.glueConnection.GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2CredentialsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2CredentialsOutputReference)SetAccessToken(val *string) {
	if err := j.validateSetAccessTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessToken",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2CredentialsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2CredentialsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2CredentialsOutputReference)SetInternalValue(val *GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2Credentials) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2CredentialsOutputReference)SetJwtToken(val *string) {
	if err := j.validateSetJwtTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jwtToken",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2CredentialsOutputReference)SetRefreshToken(val *string) {
	if err := j.validateSetRefreshTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"refreshToken",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2CredentialsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2CredentialsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2CredentialsOutputReference)SetUserManagedClientApplicationClientSecret(val *string) {
	if err := j.validateSetUserManagedClientApplicationClientSecretParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userManagedClientApplicationClientSecret",
		val,
	)
}

func (g *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2CredentialsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2CredentialsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2CredentialsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2CredentialsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2CredentialsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2CredentialsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2CredentialsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2CredentialsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2CredentialsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2CredentialsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2CredentialsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2CredentialsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2CredentialsOutputReference) ResetAccessToken() {
	_jsii_.InvokeVoid(
		g,
		"resetAccessToken",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2CredentialsOutputReference) ResetJwtToken() {
	_jsii_.InvokeVoid(
		g,
		"resetJwtToken",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2CredentialsOutputReference) ResetRefreshToken() {
	_jsii_.InvokeVoid(
		g,
		"resetRefreshToken",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2CredentialsOutputReference) ResetUserManagedClientApplicationClientSecret() {
	_jsii_.InvokeVoid(
		g,
		"resetUserManagedClientApplicationClientSecret",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2CredentialsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2CredentialsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

