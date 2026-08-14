// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ephemeralawsstswebidentitytoken

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/ephemeralawsstswebidentitytoken/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/ephemeral-resources/sts_web_identity_token aws_sts_web_identity_token}.
type EphemeralAwsStsWebIdentityToken interface {
	cdktn.TerraformEphemeralResource
	Audience() *[]*string
	SetAudience(val *[]*string)
	AudienceInput() *[]*string
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
	DurationSeconds() *float64
	SetDurationSeconds(val *float64)
	DurationSecondsInput() *float64
	Expiration() *string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformEphemeralResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformEphemeralResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	SigningAlgorithm() *string
	SetSigningAlgorithm(val *string)
	SigningAlgorithmInput() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	WebIdentityToken() *string
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
	ResetDurationSeconds()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetTags()
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

// The jsii proxy struct for EphemeralAwsStsWebIdentityToken
type jsiiProxy_EphemeralAwsStsWebIdentityToken struct {
	internal.Type__cdktnTerraformEphemeralResource
}

func (j *jsiiProxy_EphemeralAwsStsWebIdentityToken) Audience() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"audience",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAwsStsWebIdentityToken) AudienceInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"audienceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAwsStsWebIdentityToken) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAwsStsWebIdentityToken) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAwsStsWebIdentityToken) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAwsStsWebIdentityToken) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAwsStsWebIdentityToken) DurationSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"durationSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAwsStsWebIdentityToken) DurationSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"durationSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAwsStsWebIdentityToken) Expiration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expiration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAwsStsWebIdentityToken) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAwsStsWebIdentityToken) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAwsStsWebIdentityToken) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAwsStsWebIdentityToken) Lifecycle() *cdktn.TerraformEphemeralResourceLifecycle {
	var returns *cdktn.TerraformEphemeralResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAwsStsWebIdentityToken) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAwsStsWebIdentityToken) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAwsStsWebIdentityToken) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAwsStsWebIdentityToken) SigningAlgorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signingAlgorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAwsStsWebIdentityToken) SigningAlgorithmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signingAlgorithmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAwsStsWebIdentityToken) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAwsStsWebIdentityToken) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAwsStsWebIdentityToken) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAwsStsWebIdentityToken) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAwsStsWebIdentityToken) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAwsStsWebIdentityToken) WebIdentityToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webIdentityToken",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/ephemeral-resources/sts_web_identity_token aws_sts_web_identity_token} Ephemeral Resource.
func NewEphemeralAwsStsWebIdentityToken(scope constructs.Construct, id *string, config *EphemeralAwsStsWebIdentityTokenConfig) EphemeralAwsStsWebIdentityToken {
	_init_.Initialize()

	if err := validateNewEphemeralAwsStsWebIdentityTokenParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_EphemeralAwsStsWebIdentityToken{}

	_jsii_.Create(
		"@cdktn/provider-aws.ephemeralAwsStsWebIdentityToken.EphemeralAwsStsWebIdentityToken",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.60.0/docs/ephemeral-resources/sts_web_identity_token aws_sts_web_identity_token} Ephemeral Resource.
func NewEphemeralAwsStsWebIdentityToken_Override(e EphemeralAwsStsWebIdentityToken, scope constructs.Construct, id *string, config *EphemeralAwsStsWebIdentityTokenConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.ephemeralAwsStsWebIdentityToken.EphemeralAwsStsWebIdentityToken",
		[]interface{}{scope, id, config},
		e,
	)
}

func (j *jsiiProxy_EphemeralAwsStsWebIdentityToken)SetAudience(val *[]*string) {
	if err := j.validateSetAudienceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"audience",
		val,
	)
}

func (j *jsiiProxy_EphemeralAwsStsWebIdentityToken)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_EphemeralAwsStsWebIdentityToken)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_EphemeralAwsStsWebIdentityToken)SetDurationSeconds(val *float64) {
	if err := j.validateSetDurationSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"durationSeconds",
		val,
	)
}

func (j *jsiiProxy_EphemeralAwsStsWebIdentityToken)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_EphemeralAwsStsWebIdentityToken)SetLifecycle(val *cdktn.TerraformEphemeralResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_EphemeralAwsStsWebIdentityToken)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_EphemeralAwsStsWebIdentityToken)SetSigningAlgorithm(val *string) {
	if err := j.validateSetSigningAlgorithmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"signingAlgorithm",
		val,
	)
}

func (j *jsiiProxy_EphemeralAwsStsWebIdentityToken)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
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
func EphemeralAwsStsWebIdentityToken_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEphemeralAwsStsWebIdentityToken_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-aws.ephemeralAwsStsWebIdentityToken.EphemeralAwsStsWebIdentityToken",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func EphemeralAwsStsWebIdentityToken_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEphemeralAwsStsWebIdentityToken_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-aws.ephemeralAwsStsWebIdentityToken.EphemeralAwsStsWebIdentityToken",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func EphemeralAwsStsWebIdentityToken_IsTerraformEphemeralResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEphemeralAwsStsWebIdentityToken_IsTerraformEphemeralResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-aws.ephemeralAwsStsWebIdentityToken.EphemeralAwsStsWebIdentityToken",
		"isTerraformEphemeralResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func EphemeralAwsStsWebIdentityToken_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-aws.ephemeralAwsStsWebIdentityToken.EphemeralAwsStsWebIdentityToken",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (e *jsiiProxy_EphemeralAwsStsWebIdentityToken) AddOverride(path *string, value interface{}) {
	if err := e.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (e *jsiiProxy_EphemeralAwsStsWebIdentityToken) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_EphemeralAwsStsWebIdentityToken) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (e *jsiiProxy_EphemeralAwsStsWebIdentityToken) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_EphemeralAwsStsWebIdentityToken) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_EphemeralAwsStsWebIdentityToken) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_EphemeralAwsStsWebIdentityToken) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_EphemeralAwsStsWebIdentityToken) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_EphemeralAwsStsWebIdentityToken) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_EphemeralAwsStsWebIdentityToken) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_EphemeralAwsStsWebIdentityToken) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (e *jsiiProxy_EphemeralAwsStsWebIdentityToken) OverrideLogicalId(newLogicalId *string) {
	if err := e.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (e *jsiiProxy_EphemeralAwsStsWebIdentityToken) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := e.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (e *jsiiProxy_EphemeralAwsStsWebIdentityToken) ResetDurationSeconds() {
	_jsii_.InvokeVoid(
		e,
		"resetDurationSeconds",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EphemeralAwsStsWebIdentityToken) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EphemeralAwsStsWebIdentityToken) ResetTags() {
	_jsii_.InvokeVoid(
		e,
		"resetTags",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EphemeralAwsStsWebIdentityToken) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralAwsStsWebIdentityToken) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralAwsStsWebIdentityToken) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralAwsStsWebIdentityToken) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralAwsStsWebIdentityToken) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralAwsStsWebIdentityToken) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralAwsStsWebIdentityToken) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

