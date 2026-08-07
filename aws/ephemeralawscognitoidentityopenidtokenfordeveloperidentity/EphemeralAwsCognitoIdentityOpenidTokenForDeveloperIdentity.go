// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ephemeralawscognitoidentityopenidtokenfordeveloperidentity

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/ephemeralawscognitoidentityopenidtokenfordeveloperidentity/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.58.0/docs/ephemeral-resources/cognito_identity_openid_token_for_developer_identity aws_cognito_identity_openid_token_for_developer_identity}.
type EphemeralAwsCognitoIdentityOpenidTokenForDeveloperIdentity interface {
	cdktn.TerraformEphemeralResource
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	IdentityId() *string
	SetIdentityId(val *string)
	IdentityIdInput() *string
	IdentityPoolId() *string
	SetIdentityPoolId(val *string)
	IdentityPoolIdInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformEphemeralResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformEphemeralResourceLifecycle)
	Logins() *map[string]*string
	SetLogins(val *map[string]*string)
	LoginsInput() *map[string]*string
	// The tree node.
	Node() constructs.Node
	PrincipalTags() *map[string]*string
	SetPrincipalTags(val *map[string]*string)
	PrincipalTagsInput() *map[string]*string
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Token() *string
	TokenDuration() *float64
	SetTokenDuration(val *float64)
	TokenDurationInput() *float64
	// Experimental.
	AddOverride(path *string, value interface{})
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
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Registers a synth-time validation that the project's declared targetVersions admit the given provider-protocol feature family.
	//
	// Called by generated provider bindings when a versioned feature is
	// structurally in use - the element's existence in the construct tree
	// already implies the feature is used, e.g. constructing a
	// `TerraformEphemeralResource` at all - so, unlike
	// `_registerResolveDiscoveredProviderFeatureUsage`, this registration is
	// never deactivated by `_resetResolveDiscoveredProviderFeatureUsage`. Not
	// intended to be called directly by user code. Lives on `TerraformElement`
	// (rather than `TerraformResource`) so it covers any element subclass
	// that needs it.
	// Experimental.
	RegisterProviderFeatureUsage(feature cdktn.ProviderFeature)
	ResetIdentityId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPrincipalTags()
	ResetRegion()
	ResetTokenDuration()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this ephemeral resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
	// Applies one or more mixins to this construct.
	//
	// Mixins are applied in order. The list of constructs is captured at the
	// start of the call, so constructs added by a mixin will not be visited.
	// Use multiple `with()` calls if subsequent mixins should apply to added
	// constructs.
	//
	// Returns: This construct for chaining.
	With(mixins ...constructs.IMixin) constructs.IConstruct
}

// The jsii proxy struct for EphemeralAwsCognitoIdentityOpenidTokenForDeveloperIdentity
type jsiiProxy_EphemeralAwsCognitoIdentityOpenidTokenForDeveloperIdentity struct {
	internal.Type__cdktnTerraformEphemeralResource
}

func (j *jsiiProxy_EphemeralAwsCognitoIdentityOpenidTokenForDeveloperIdentity) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAwsCognitoIdentityOpenidTokenForDeveloperIdentity) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAwsCognitoIdentityOpenidTokenForDeveloperIdentity) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAwsCognitoIdentityOpenidTokenForDeveloperIdentity) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAwsCognitoIdentityOpenidTokenForDeveloperIdentity) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAwsCognitoIdentityOpenidTokenForDeveloperIdentity) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAwsCognitoIdentityOpenidTokenForDeveloperIdentity) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAwsCognitoIdentityOpenidTokenForDeveloperIdentity) IdentityId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAwsCognitoIdentityOpenidTokenForDeveloperIdentity) IdentityIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAwsCognitoIdentityOpenidTokenForDeveloperIdentity) IdentityPoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityPoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAwsCognitoIdentityOpenidTokenForDeveloperIdentity) IdentityPoolIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityPoolIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAwsCognitoIdentityOpenidTokenForDeveloperIdentity) Lifecycle() *cdktn.TerraformEphemeralResourceLifecycle {
	var returns *cdktn.TerraformEphemeralResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAwsCognitoIdentityOpenidTokenForDeveloperIdentity) Logins() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"logins",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAwsCognitoIdentityOpenidTokenForDeveloperIdentity) LoginsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"loginsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAwsCognitoIdentityOpenidTokenForDeveloperIdentity) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAwsCognitoIdentityOpenidTokenForDeveloperIdentity) PrincipalTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"principalTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAwsCognitoIdentityOpenidTokenForDeveloperIdentity) PrincipalTagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"principalTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAwsCognitoIdentityOpenidTokenForDeveloperIdentity) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAwsCognitoIdentityOpenidTokenForDeveloperIdentity) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAwsCognitoIdentityOpenidTokenForDeveloperIdentity) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAwsCognitoIdentityOpenidTokenForDeveloperIdentity) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAwsCognitoIdentityOpenidTokenForDeveloperIdentity) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAwsCognitoIdentityOpenidTokenForDeveloperIdentity) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAwsCognitoIdentityOpenidTokenForDeveloperIdentity) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAwsCognitoIdentityOpenidTokenForDeveloperIdentity) Token() *string {
	var returns *string
	_jsii_.Get(
		j,
		"token",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAwsCognitoIdentityOpenidTokenForDeveloperIdentity) TokenDuration() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAwsCognitoIdentityOpenidTokenForDeveloperIdentity) TokenDurationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenDurationInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.58.0/docs/ephemeral-resources/cognito_identity_openid_token_for_developer_identity aws_cognito_identity_openid_token_for_developer_identity} Ephemeral Resource.
func NewEphemeralAwsCognitoIdentityOpenidTokenForDeveloperIdentity(scope constructs.Construct, id *string, config *EphemeralAwsCognitoIdentityOpenidTokenForDeveloperIdentityConfig) EphemeralAwsCognitoIdentityOpenidTokenForDeveloperIdentity {
	_init_.Initialize()

	if err := validateNewEphemeralAwsCognitoIdentityOpenidTokenForDeveloperIdentityParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_EphemeralAwsCognitoIdentityOpenidTokenForDeveloperIdentity{}

	_jsii_.Create(
		"@cdktn/provider-aws.ephemeralAwsCognitoIdentityOpenidTokenForDeveloperIdentity.EphemeralAwsCognitoIdentityOpenidTokenForDeveloperIdentity",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.58.0/docs/ephemeral-resources/cognito_identity_openid_token_for_developer_identity aws_cognito_identity_openid_token_for_developer_identity} Ephemeral Resource.
func NewEphemeralAwsCognitoIdentityOpenidTokenForDeveloperIdentity_Override(e EphemeralAwsCognitoIdentityOpenidTokenForDeveloperIdentity, scope constructs.Construct, id *string, config *EphemeralAwsCognitoIdentityOpenidTokenForDeveloperIdentityConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.ephemeralAwsCognitoIdentityOpenidTokenForDeveloperIdentity.EphemeralAwsCognitoIdentityOpenidTokenForDeveloperIdentity",
		[]interface{}{scope, id, config},
		e,
	)
}

func (j *jsiiProxy_EphemeralAwsCognitoIdentityOpenidTokenForDeveloperIdentity)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_EphemeralAwsCognitoIdentityOpenidTokenForDeveloperIdentity)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_EphemeralAwsCognitoIdentityOpenidTokenForDeveloperIdentity)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_EphemeralAwsCognitoIdentityOpenidTokenForDeveloperIdentity)SetIdentityId(val *string) {
	if err := j.validateSetIdentityIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"identityId",
		val,
	)
}

func (j *jsiiProxy_EphemeralAwsCognitoIdentityOpenidTokenForDeveloperIdentity)SetIdentityPoolId(val *string) {
	if err := j.validateSetIdentityPoolIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"identityPoolId",
		val,
	)
}

func (j *jsiiProxy_EphemeralAwsCognitoIdentityOpenidTokenForDeveloperIdentity)SetLifecycle(val *cdktn.TerraformEphemeralResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_EphemeralAwsCognitoIdentityOpenidTokenForDeveloperIdentity)SetLogins(val *map[string]*string) {
	if err := j.validateSetLoginsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logins",
		val,
	)
}

func (j *jsiiProxy_EphemeralAwsCognitoIdentityOpenidTokenForDeveloperIdentity)SetPrincipalTags(val *map[string]*string) {
	if err := j.validateSetPrincipalTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"principalTags",
		val,
	)
}

func (j *jsiiProxy_EphemeralAwsCognitoIdentityOpenidTokenForDeveloperIdentity)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_EphemeralAwsCognitoIdentityOpenidTokenForDeveloperIdentity)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_EphemeralAwsCognitoIdentityOpenidTokenForDeveloperIdentity)SetTokenDuration(val *float64) {
	if err := j.validateSetTokenDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenDuration",
		val,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func EphemeralAwsCognitoIdentityOpenidTokenForDeveloperIdentity_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEphemeralAwsCognitoIdentityOpenidTokenForDeveloperIdentity_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-aws.ephemeralAwsCognitoIdentityOpenidTokenForDeveloperIdentity.EphemeralAwsCognitoIdentityOpenidTokenForDeveloperIdentity",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func EphemeralAwsCognitoIdentityOpenidTokenForDeveloperIdentity_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEphemeralAwsCognitoIdentityOpenidTokenForDeveloperIdentity_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-aws.ephemeralAwsCognitoIdentityOpenidTokenForDeveloperIdentity.EphemeralAwsCognitoIdentityOpenidTokenForDeveloperIdentity",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func EphemeralAwsCognitoIdentityOpenidTokenForDeveloperIdentity_IsTerraformEphemeralResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEphemeralAwsCognitoIdentityOpenidTokenForDeveloperIdentity_IsTerraformEphemeralResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-aws.ephemeralAwsCognitoIdentityOpenidTokenForDeveloperIdentity.EphemeralAwsCognitoIdentityOpenidTokenForDeveloperIdentity",
		"isTerraformEphemeralResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func EphemeralAwsCognitoIdentityOpenidTokenForDeveloperIdentity_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-aws.ephemeralAwsCognitoIdentityOpenidTokenForDeveloperIdentity.EphemeralAwsCognitoIdentityOpenidTokenForDeveloperIdentity",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (e *jsiiProxy_EphemeralAwsCognitoIdentityOpenidTokenForDeveloperIdentity) AddOverride(path *string, value interface{}) {
	if err := e.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (e *jsiiProxy_EphemeralAwsCognitoIdentityOpenidTokenForDeveloperIdentity) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralAwsCognitoIdentityOpenidTokenForDeveloperIdentity) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralAwsCognitoIdentityOpenidTokenForDeveloperIdentity) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralAwsCognitoIdentityOpenidTokenForDeveloperIdentity) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralAwsCognitoIdentityOpenidTokenForDeveloperIdentity) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralAwsCognitoIdentityOpenidTokenForDeveloperIdentity) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralAwsCognitoIdentityOpenidTokenForDeveloperIdentity) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralAwsCognitoIdentityOpenidTokenForDeveloperIdentity) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralAwsCognitoIdentityOpenidTokenForDeveloperIdentity) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralAwsCognitoIdentityOpenidTokenForDeveloperIdentity) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralAwsCognitoIdentityOpenidTokenForDeveloperIdentity) OverrideLogicalId(newLogicalId *string) {
	if err := e.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (e *jsiiProxy_EphemeralAwsCognitoIdentityOpenidTokenForDeveloperIdentity) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := e.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (e *jsiiProxy_EphemeralAwsCognitoIdentityOpenidTokenForDeveloperIdentity) ResetIdentityId() {
	_jsii_.InvokeVoid(
		e,
		"resetIdentityId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EphemeralAwsCognitoIdentityOpenidTokenForDeveloperIdentity) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EphemeralAwsCognitoIdentityOpenidTokenForDeveloperIdentity) ResetPrincipalTags() {
	_jsii_.InvokeVoid(
		e,
		"resetPrincipalTags",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EphemeralAwsCognitoIdentityOpenidTokenForDeveloperIdentity) ResetRegion() {
	_jsii_.InvokeVoid(
		e,
		"resetRegion",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EphemeralAwsCognitoIdentityOpenidTokenForDeveloperIdentity) ResetTokenDuration() {
	_jsii_.InvokeVoid(
		e,
		"resetTokenDuration",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EphemeralAwsCognitoIdentityOpenidTokenForDeveloperIdentity) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralAwsCognitoIdentityOpenidTokenForDeveloperIdentity) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralAwsCognitoIdentityOpenidTokenForDeveloperIdentity) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralAwsCognitoIdentityOpenidTokenForDeveloperIdentity) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralAwsCognitoIdentityOpenidTokenForDeveloperIdentity) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralAwsCognitoIdentityOpenidTokenForDeveloperIdentity) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralAwsCognitoIdentityOpenidTokenForDeveloperIdentity) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		e,
		"with",
		args,
		&returns,
	)

	return returns
}

