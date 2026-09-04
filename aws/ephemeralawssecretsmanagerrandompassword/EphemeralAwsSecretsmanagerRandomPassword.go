// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ephemeralawssecretsmanagerrandompassword

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/ephemeralawssecretsmanagerrandompassword/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/ephemeral-resources/secretsmanager_random_password aws_secretsmanager_random_password}.
type EphemeralAwsSecretsmanagerRandomPassword interface {
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
	ExcludeCharacters() *string
	SetExcludeCharacters(val *string)
	ExcludeCharactersInput() *string
	ExcludeLowercase() interface{}
	SetExcludeLowercase(val interface{})
	ExcludeLowercaseInput() interface{}
	ExcludeNumbers() interface{}
	SetExcludeNumbers(val interface{})
	ExcludeNumbersInput() interface{}
	ExcludePunctuation() interface{}
	SetExcludePunctuation(val interface{})
	ExcludePunctuationInput() interface{}
	ExcludeUppercase() interface{}
	SetExcludeUppercase(val interface{})
	ExcludeUppercaseInput() interface{}
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	IncludeSpace() interface{}
	SetIncludeSpace(val interface{})
	IncludeSpaceInput() interface{}
	// Experimental.
	Lifecycle() *cdktn.TerraformEphemeralResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformEphemeralResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	PasswordLength() *float64
	SetPasswordLength(val *float64)
	PasswordLengthInput() *float64
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	RandomPassword() *string
	// Experimental.
	RawOverrides() interface{}
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	RequireEachIncludedType() interface{}
	SetRequireEachIncludedType(val interface{})
	RequireEachIncludedTypeInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
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
	ResetExcludeCharacters()
	ResetExcludeLowercase()
	ResetExcludeNumbers()
	ResetExcludePunctuation()
	ResetExcludeUppercase()
	ResetIncludeSpace()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPasswordLength()
	ResetRegion()
	ResetRequireEachIncludedType()
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

// The jsii proxy struct for EphemeralAwsSecretsmanagerRandomPassword
type jsiiProxy_EphemeralAwsSecretsmanagerRandomPassword struct {
	internal.Type__cdktnTerraformEphemeralResource
}

func (j *jsiiProxy_EphemeralAwsSecretsmanagerRandomPassword) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAwsSecretsmanagerRandomPassword) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAwsSecretsmanagerRandomPassword) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAwsSecretsmanagerRandomPassword) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAwsSecretsmanagerRandomPassword) ExcludeCharacters() *string {
	var returns *string
	_jsii_.Get(
		j,
		"excludeCharacters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAwsSecretsmanagerRandomPassword) ExcludeCharactersInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"excludeCharactersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAwsSecretsmanagerRandomPassword) ExcludeLowercase() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludeLowercase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAwsSecretsmanagerRandomPassword) ExcludeLowercaseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludeLowercaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAwsSecretsmanagerRandomPassword) ExcludeNumbers() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludeNumbers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAwsSecretsmanagerRandomPassword) ExcludeNumbersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludeNumbersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAwsSecretsmanagerRandomPassword) ExcludePunctuation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludePunctuation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAwsSecretsmanagerRandomPassword) ExcludePunctuationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludePunctuationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAwsSecretsmanagerRandomPassword) ExcludeUppercase() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludeUppercase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAwsSecretsmanagerRandomPassword) ExcludeUppercaseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludeUppercaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAwsSecretsmanagerRandomPassword) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAwsSecretsmanagerRandomPassword) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAwsSecretsmanagerRandomPassword) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAwsSecretsmanagerRandomPassword) IncludeSpace() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeSpace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAwsSecretsmanagerRandomPassword) IncludeSpaceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeSpaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAwsSecretsmanagerRandomPassword) Lifecycle() *cdktn.TerraformEphemeralResourceLifecycle {
	var returns *cdktn.TerraformEphemeralResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAwsSecretsmanagerRandomPassword) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAwsSecretsmanagerRandomPassword) PasswordLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"passwordLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAwsSecretsmanagerRandomPassword) PasswordLengthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"passwordLengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAwsSecretsmanagerRandomPassword) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAwsSecretsmanagerRandomPassword) RandomPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"randomPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAwsSecretsmanagerRandomPassword) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAwsSecretsmanagerRandomPassword) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAwsSecretsmanagerRandomPassword) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAwsSecretsmanagerRandomPassword) RequireEachIncludedType() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireEachIncludedType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAwsSecretsmanagerRandomPassword) RequireEachIncludedTypeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireEachIncludedTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAwsSecretsmanagerRandomPassword) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAwsSecretsmanagerRandomPassword) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralAwsSecretsmanagerRandomPassword) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/ephemeral-resources/secretsmanager_random_password aws_secretsmanager_random_password} Ephemeral Resource.
func NewEphemeralAwsSecretsmanagerRandomPassword(scope constructs.Construct, id *string, config *EphemeralAwsSecretsmanagerRandomPasswordConfig) EphemeralAwsSecretsmanagerRandomPassword {
	_init_.Initialize()

	if err := validateNewEphemeralAwsSecretsmanagerRandomPasswordParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_EphemeralAwsSecretsmanagerRandomPassword{}

	_jsii_.Create(
		"@cdktn/provider-aws.ephemeralAwsSecretsmanagerRandomPassword.EphemeralAwsSecretsmanagerRandomPassword",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.63.0/docs/ephemeral-resources/secretsmanager_random_password aws_secretsmanager_random_password} Ephemeral Resource.
func NewEphemeralAwsSecretsmanagerRandomPassword_Override(e EphemeralAwsSecretsmanagerRandomPassword, scope constructs.Construct, id *string, config *EphemeralAwsSecretsmanagerRandomPasswordConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.ephemeralAwsSecretsmanagerRandomPassword.EphemeralAwsSecretsmanagerRandomPassword",
		[]interface{}{scope, id, config},
		e,
	)
}

func (j *jsiiProxy_EphemeralAwsSecretsmanagerRandomPassword)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_EphemeralAwsSecretsmanagerRandomPassword)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_EphemeralAwsSecretsmanagerRandomPassword)SetExcludeCharacters(val *string) {
	if err := j.validateSetExcludeCharactersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludeCharacters",
		val,
	)
}

func (j *jsiiProxy_EphemeralAwsSecretsmanagerRandomPassword)SetExcludeLowercase(val interface{}) {
	if err := j.validateSetExcludeLowercaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludeLowercase",
		val,
	)
}

func (j *jsiiProxy_EphemeralAwsSecretsmanagerRandomPassword)SetExcludeNumbers(val interface{}) {
	if err := j.validateSetExcludeNumbersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludeNumbers",
		val,
	)
}

func (j *jsiiProxy_EphemeralAwsSecretsmanagerRandomPassword)SetExcludePunctuation(val interface{}) {
	if err := j.validateSetExcludePunctuationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludePunctuation",
		val,
	)
}

func (j *jsiiProxy_EphemeralAwsSecretsmanagerRandomPassword)SetExcludeUppercase(val interface{}) {
	if err := j.validateSetExcludeUppercaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludeUppercase",
		val,
	)
}

func (j *jsiiProxy_EphemeralAwsSecretsmanagerRandomPassword)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_EphemeralAwsSecretsmanagerRandomPassword)SetIncludeSpace(val interface{}) {
	if err := j.validateSetIncludeSpaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeSpace",
		val,
	)
}

func (j *jsiiProxy_EphemeralAwsSecretsmanagerRandomPassword)SetLifecycle(val *cdktn.TerraformEphemeralResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_EphemeralAwsSecretsmanagerRandomPassword)SetPasswordLength(val *float64) {
	if err := j.validateSetPasswordLengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passwordLength",
		val,
	)
}

func (j *jsiiProxy_EphemeralAwsSecretsmanagerRandomPassword)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_EphemeralAwsSecretsmanagerRandomPassword)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_EphemeralAwsSecretsmanagerRandomPassword)SetRequireEachIncludedType(val interface{}) {
	if err := j.validateSetRequireEachIncludedTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireEachIncludedType",
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
func EphemeralAwsSecretsmanagerRandomPassword_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEphemeralAwsSecretsmanagerRandomPassword_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-aws.ephemeralAwsSecretsmanagerRandomPassword.EphemeralAwsSecretsmanagerRandomPassword",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func EphemeralAwsSecretsmanagerRandomPassword_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEphemeralAwsSecretsmanagerRandomPassword_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-aws.ephemeralAwsSecretsmanagerRandomPassword.EphemeralAwsSecretsmanagerRandomPassword",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func EphemeralAwsSecretsmanagerRandomPassword_IsTerraformEphemeralResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEphemeralAwsSecretsmanagerRandomPassword_IsTerraformEphemeralResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-aws.ephemeralAwsSecretsmanagerRandomPassword.EphemeralAwsSecretsmanagerRandomPassword",
		"isTerraformEphemeralResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func EphemeralAwsSecretsmanagerRandomPassword_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-aws.ephemeralAwsSecretsmanagerRandomPassword.EphemeralAwsSecretsmanagerRandomPassword",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (e *jsiiProxy_EphemeralAwsSecretsmanagerRandomPassword) AddOverride(path *string, value interface{}) {
	if err := e.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (e *jsiiProxy_EphemeralAwsSecretsmanagerRandomPassword) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_EphemeralAwsSecretsmanagerRandomPassword) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (e *jsiiProxy_EphemeralAwsSecretsmanagerRandomPassword) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_EphemeralAwsSecretsmanagerRandomPassword) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_EphemeralAwsSecretsmanagerRandomPassword) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_EphemeralAwsSecretsmanagerRandomPassword) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_EphemeralAwsSecretsmanagerRandomPassword) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_EphemeralAwsSecretsmanagerRandomPassword) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_EphemeralAwsSecretsmanagerRandomPassword) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_EphemeralAwsSecretsmanagerRandomPassword) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (e *jsiiProxy_EphemeralAwsSecretsmanagerRandomPassword) OverrideLogicalId(newLogicalId *string) {
	if err := e.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (e *jsiiProxy_EphemeralAwsSecretsmanagerRandomPassword) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := e.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (e *jsiiProxy_EphemeralAwsSecretsmanagerRandomPassword) ResetExcludeCharacters() {
	_jsii_.InvokeVoid(
		e,
		"resetExcludeCharacters",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EphemeralAwsSecretsmanagerRandomPassword) ResetExcludeLowercase() {
	_jsii_.InvokeVoid(
		e,
		"resetExcludeLowercase",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EphemeralAwsSecretsmanagerRandomPassword) ResetExcludeNumbers() {
	_jsii_.InvokeVoid(
		e,
		"resetExcludeNumbers",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EphemeralAwsSecretsmanagerRandomPassword) ResetExcludePunctuation() {
	_jsii_.InvokeVoid(
		e,
		"resetExcludePunctuation",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EphemeralAwsSecretsmanagerRandomPassword) ResetExcludeUppercase() {
	_jsii_.InvokeVoid(
		e,
		"resetExcludeUppercase",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EphemeralAwsSecretsmanagerRandomPassword) ResetIncludeSpace() {
	_jsii_.InvokeVoid(
		e,
		"resetIncludeSpace",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EphemeralAwsSecretsmanagerRandomPassword) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EphemeralAwsSecretsmanagerRandomPassword) ResetPasswordLength() {
	_jsii_.InvokeVoid(
		e,
		"resetPasswordLength",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EphemeralAwsSecretsmanagerRandomPassword) ResetRegion() {
	_jsii_.InvokeVoid(
		e,
		"resetRegion",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EphemeralAwsSecretsmanagerRandomPassword) ResetRequireEachIncludedType() {
	_jsii_.InvokeVoid(
		e,
		"resetRequireEachIncludedType",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EphemeralAwsSecretsmanagerRandomPassword) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralAwsSecretsmanagerRandomPassword) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralAwsSecretsmanagerRandomPassword) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralAwsSecretsmanagerRandomPassword) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralAwsSecretsmanagerRandomPassword) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralAwsSecretsmanagerRandomPassword) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralAwsSecretsmanagerRandomPassword) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

