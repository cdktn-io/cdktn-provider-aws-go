// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package glueconnection

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/jsii"

	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/glueconnection/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2ClientApplicationOutputReference interface {
	cdktn.ComplexObject
	AwsManagedClientApplicationReference() *string
	SetAwsManagedClientApplicationReference(val *string)
	AwsManagedClientApplicationReferenceInput() *string
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
	InternalValue() *GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2ClientApplication
	SetInternalValue(val *GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2ClientApplication)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	UserManagedClientApplicationClientId() *string
	SetUserManagedClientApplicationClientId(val *string)
	UserManagedClientApplicationClientIdInput() *string
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
	ResetAwsManagedClientApplicationReference()
	ResetUserManagedClientApplicationClientId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2ClientApplicationOutputReference
type jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2ClientApplicationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2ClientApplicationOutputReference) AwsManagedClientApplicationReference() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsManagedClientApplicationReference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2ClientApplicationOutputReference) AwsManagedClientApplicationReferenceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsManagedClientApplicationReferenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2ClientApplicationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2ClientApplicationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2ClientApplicationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2ClientApplicationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2ClientApplicationOutputReference) InternalValue() *GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2ClientApplication {
	var returns *GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2ClientApplication
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2ClientApplicationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2ClientApplicationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2ClientApplicationOutputReference) UserManagedClientApplicationClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userManagedClientApplicationClientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2ClientApplicationOutputReference) UserManagedClientApplicationClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userManagedClientApplicationClientIdInput",
		&returns,
	)
	return returns
}


func NewGlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2ClientApplicationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2ClientApplicationOutputReference {
	_init_.Initialize()

	if err := validateNewGlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2ClientApplicationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2ClientApplicationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-aws.glueConnection.GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2ClientApplicationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2ClientApplicationOutputReference_Override(g GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2ClientApplicationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.glueConnection.GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2ClientApplicationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2ClientApplicationOutputReference)SetAwsManagedClientApplicationReference(val *string) {
	if err := j.validateSetAwsManagedClientApplicationReferenceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"awsManagedClientApplicationReference",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2ClientApplicationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2ClientApplicationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2ClientApplicationOutputReference)SetInternalValue(val *GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2ClientApplication) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2ClientApplicationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2ClientApplicationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2ClientApplicationOutputReference)SetUserManagedClientApplicationClientId(val *string) {
	if err := j.validateSetUserManagedClientApplicationClientIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userManagedClientApplicationClientId",
		val,
	)
}

func (g *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2ClientApplicationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2ClientApplicationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2ClientApplicationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2ClientApplicationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2ClientApplicationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2ClientApplicationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2ClientApplicationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2ClientApplicationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2ClientApplicationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2ClientApplicationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2ClientApplicationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2ClientApplicationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2ClientApplicationOutputReference) ResetAwsManagedClientApplicationReference() {
	_jsii_.InvokeVoid(
		g,
		"resetAwsManagedClientApplicationReference",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2ClientApplicationOutputReference) ResetUserManagedClientApplicationClientId() {
	_jsii_.InvokeVoid(
		g,
		"resetUserManagedClientApplicationClientId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2ClientApplicationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GlueConnectionAuthenticationConfigurationOauth2PropertiesOauth2ClientApplicationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

