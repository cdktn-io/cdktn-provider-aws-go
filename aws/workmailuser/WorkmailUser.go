// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package workmailuser

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/workmailuser/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workmail_user aws_workmail_user}.
type WorkmailUser interface {
	cdktn.TerraformResource
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	City() *string
	SetCity(val *string)
	CityInput() *string
	Company() *string
	SetCompany(val *string)
	CompanyInput() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	Country() *string
	SetCountry(val *string)
	CountryInput() *string
	Department() *string
	SetDepartment(val *string)
	DepartmentInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DisabledDate() *string
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	Email() *string
	SetEmail(val *string)
	EmailInput() *string
	EnabledDate() *string
	FirstName() *string
	SetFirstName(val *string)
	FirstNameInput() *string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HiddenFromGlobalAddressList() interface{}
	SetHiddenFromGlobalAddressList(val interface{})
	HiddenFromGlobalAddressListInput() interface{}
	IdentityProviderIdentityStoreId() *string
	IdentityProviderUserId() *string
	SetIdentityProviderUserId(val *string)
	IdentityProviderUserIdInput() *string
	Initials() *string
	SetInitials(val *string)
	InitialsInput() *string
	JobTitle() *string
	SetJobTitle(val *string)
	JobTitleInput() *string
	LastName() *string
	SetLastName(val *string)
	LastNameInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	MailboxDeprovisionedDate() *string
	MailboxProvisionedDate() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	Office() *string
	SetOffice(val *string)
	OfficeInput() *string
	OrganizationId() *string
	SetOrganizationId(val *string)
	OrganizationIdInput() *string
	Password() *string
	SetPassword(val *string)
	PasswordInput() *string
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	State() *string
	Street() *string
	SetStreet(val *string)
	StreetInput() *string
	Telephone() *string
	SetTelephone(val *string)
	TelephoneInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	UserId() *string
	UserRole() *string
	SetUserRole(val *string)
	UserRoleInput() *string
	ZipCode() *string
	SetZipCode(val *string)
	ZipCodeInput() *string
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
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
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktn.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	// Wraps a write-only attribute's already-mapped value so that `ProviderFeature.WRITE_ONLY_ATTRIBUTES` usage is registered at *resolve* time instead of at mutation time (setter/constructor). Called by generated bindings from `synthesizeAttributes()` and `synthesizeHclAttributes()`, e.g. `secret_key_wo: this.markWriteOnlyAttribute(cdktn.stringToTerraform(this._secretKeyWo))`; not intended to be called directly.
	//
	// `undefined` passes through completely unchanged, so the existing
	// undefined-filtering that omits unset attributes from synthesized
	// output (see `resolve()` in `tokens/private/resolve.ts`, and the
	// `value.value !== undefined` filter in generated
	// `synthesizeHclAttributes()`) keeps working untouched. `null` is also
	// passed through unchanged: it already renders as an explicit
	// null-out and must not arm the validation either.
	//
	// Any other value - including one that will itself resolve to nothing
	// (e.g. a `Lazy`/`IResolvable` producer with no value to contribute) -
	// is wrapped in a token whose `resolve()` defers to the real resolver
	// first and registers usage only if what comes back is not
	// `null`/`undefined`; the resolved value is then returned unchanged,
	// so what actually renders is untouched by this wrapper. A producer
	// that resolves to `undefined` therefore neither registers usage nor
	// leaves anything behind in the synthesized attribute - the omission
	// behaves exactly as if the attribute had never been set.
	//
	// Registration goes through `_registerResolveDiscoveredProviderFeatureUsage`
	// rather than `registerProviderFeatureUsage`: usage here is only known at
	// resolve time, and a given element can be resolved across many
	// synthesis passes over its lifetime (repeated `app.synth()` calls,
	// tests reusing a construct tree), so it must represent only the CURRENT
	// pass rather than accumulate forever. Every validation-enabled entry
	// point (`App.synth`; `Testing.synth`/`synthHcl` with validations;
	// `StackSynthesizer.synthesize`) runs a prepare step that deactivates any
	// stale registration and then resolves every element's `toTerraform()`
	// before that same entry point's validations run - see
	// `TerraformStack._runPreparingResolve` - so whatever this closure
	// (re-)registers during that prepare step is always visible to the
	// validation that reads it afterwards, and nothing left over from an
	// earlier pass leaks into the current one.
	// Experimental.
	MarkWriteOnlyAttribute(value interface{}) interface{}
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using its instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
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
	ResetCity()
	ResetCompany()
	ResetCountry()
	ResetDepartment()
	ResetFirstName()
	ResetHiddenFromGlobalAddressList()
	ResetIdentityProviderUserId()
	ResetInitials()
	ResetJobTitle()
	ResetLastName()
	ResetOffice()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPassword()
	ResetRegion()
	ResetStreet()
	ResetTelephone()
	ResetUserRole()
	ResetZipCode()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
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

// The jsii proxy struct for WorkmailUser
type jsiiProxy_WorkmailUser struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_WorkmailUser) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkmailUser) City() *string {
	var returns *string
	_jsii_.Get(
		j,
		"city",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkmailUser) CityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkmailUser) Company() *string {
	var returns *string
	_jsii_.Get(
		j,
		"company",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkmailUser) CompanyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"companyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkmailUser) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkmailUser) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkmailUser) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkmailUser) Country() *string {
	var returns *string
	_jsii_.Get(
		j,
		"country",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkmailUser) CountryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"countryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkmailUser) Department() *string {
	var returns *string
	_jsii_.Get(
		j,
		"department",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkmailUser) DepartmentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"departmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkmailUser) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkmailUser) DisabledDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"disabledDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkmailUser) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkmailUser) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkmailUser) Email() *string {
	var returns *string
	_jsii_.Get(
		j,
		"email",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkmailUser) EmailInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkmailUser) EnabledDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enabledDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkmailUser) FirstName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firstName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkmailUser) FirstNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firstNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkmailUser) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkmailUser) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkmailUser) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkmailUser) HiddenFromGlobalAddressList() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hiddenFromGlobalAddressList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkmailUser) HiddenFromGlobalAddressListInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hiddenFromGlobalAddressListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkmailUser) IdentityProviderIdentityStoreId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityProviderIdentityStoreId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkmailUser) IdentityProviderUserId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityProviderUserId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkmailUser) IdentityProviderUserIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityProviderUserIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkmailUser) Initials() *string {
	var returns *string
	_jsii_.Get(
		j,
		"initials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkmailUser) InitialsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"initialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkmailUser) JobTitle() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobTitle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkmailUser) JobTitleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobTitleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkmailUser) LastName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkmailUser) LastNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkmailUser) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkmailUser) MailboxDeprovisionedDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mailboxDeprovisionedDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkmailUser) MailboxProvisionedDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mailboxProvisionedDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkmailUser) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkmailUser) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkmailUser) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkmailUser) Office() *string {
	var returns *string
	_jsii_.Get(
		j,
		"office",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkmailUser) OfficeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"officeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkmailUser) OrganizationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkmailUser) OrganizationIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkmailUser) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkmailUser) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkmailUser) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkmailUser) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkmailUser) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkmailUser) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkmailUser) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkmailUser) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkmailUser) Street() *string {
	var returns *string
	_jsii_.Get(
		j,
		"street",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkmailUser) StreetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkmailUser) Telephone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"telephone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkmailUser) TelephoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"telephoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkmailUser) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkmailUser) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkmailUser) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkmailUser) UserId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkmailUser) UserRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkmailUser) UserRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkmailUser) ZipCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zipCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkmailUser) ZipCodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zipCodeInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workmail_user aws_workmail_user} Resource.
func NewWorkmailUser(scope constructs.Construct, id *string, config *WorkmailUserConfig) WorkmailUser {
	_init_.Initialize()

	if err := validateNewWorkmailUserParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_WorkmailUser{}

	_jsii_.Create(
		"@cdktn/provider-aws.workmailUser.WorkmailUser",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workmail_user aws_workmail_user} Resource.
func NewWorkmailUser_Override(w WorkmailUser, scope constructs.Construct, id *string, config *WorkmailUserConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.workmailUser.WorkmailUser",
		[]interface{}{scope, id, config},
		w,
	)
}

func (j *jsiiProxy_WorkmailUser)SetCity(val *string) {
	if err := j.validateSetCityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"city",
		val,
	)
}

func (j *jsiiProxy_WorkmailUser)SetCompany(val *string) {
	if err := j.validateSetCompanyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"company",
		val,
	)
}

func (j *jsiiProxy_WorkmailUser)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_WorkmailUser)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_WorkmailUser)SetCountry(val *string) {
	if err := j.validateSetCountryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"country",
		val,
	)
}

func (j *jsiiProxy_WorkmailUser)SetDepartment(val *string) {
	if err := j.validateSetDepartmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"department",
		val,
	)
}

func (j *jsiiProxy_WorkmailUser)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_WorkmailUser)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_WorkmailUser)SetEmail(val *string) {
	if err := j.validateSetEmailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"email",
		val,
	)
}

func (j *jsiiProxy_WorkmailUser)SetFirstName(val *string) {
	if err := j.validateSetFirstNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"firstName",
		val,
	)
}

func (j *jsiiProxy_WorkmailUser)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_WorkmailUser)SetHiddenFromGlobalAddressList(val interface{}) {
	if err := j.validateSetHiddenFromGlobalAddressListParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hiddenFromGlobalAddressList",
		val,
	)
}

func (j *jsiiProxy_WorkmailUser)SetIdentityProviderUserId(val *string) {
	if err := j.validateSetIdentityProviderUserIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"identityProviderUserId",
		val,
	)
}

func (j *jsiiProxy_WorkmailUser)SetInitials(val *string) {
	if err := j.validateSetInitialsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"initials",
		val,
	)
}

func (j *jsiiProxy_WorkmailUser)SetJobTitle(val *string) {
	if err := j.validateSetJobTitleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jobTitle",
		val,
	)
}

func (j *jsiiProxy_WorkmailUser)SetLastName(val *string) {
	if err := j.validateSetLastNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lastName",
		val,
	)
}

func (j *jsiiProxy_WorkmailUser)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_WorkmailUser)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_WorkmailUser)SetOffice(val *string) {
	if err := j.validateSetOfficeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"office",
		val,
	)
}

func (j *jsiiProxy_WorkmailUser)SetOrganizationId(val *string) {
	if err := j.validateSetOrganizationIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"organizationId",
		val,
	)
}

func (j *jsiiProxy_WorkmailUser)SetPassword(val *string) {
	if err := j.validateSetPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_WorkmailUser)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_WorkmailUser)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_WorkmailUser)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_WorkmailUser)SetStreet(val *string) {
	if err := j.validateSetStreetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"street",
		val,
	)
}

func (j *jsiiProxy_WorkmailUser)SetTelephone(val *string) {
	if err := j.validateSetTelephoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"telephone",
		val,
	)
}

func (j *jsiiProxy_WorkmailUser)SetUserRole(val *string) {
	if err := j.validateSetUserRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userRole",
		val,
	)
}

func (j *jsiiProxy_WorkmailUser)SetZipCode(val *string) {
	if err := j.validateSetZipCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zipCode",
		val,
	)
}

// Generates CDKTN code for importing a WorkmailUser resource upon running "cdktn plan <stack-name>".
func WorkmailUser_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateWorkmailUser_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-aws.workmailUser.WorkmailUser",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
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
func WorkmailUser_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateWorkmailUser_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-aws.workmailUser.WorkmailUser",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func WorkmailUser_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateWorkmailUser_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-aws.workmailUser.WorkmailUser",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func WorkmailUser_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateWorkmailUser_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-aws.workmailUser.WorkmailUser",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func WorkmailUser_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-aws.workmailUser.WorkmailUser",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (w *jsiiProxy_WorkmailUser) AddMoveTarget(moveTarget *string) {
	if err := w.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (w *jsiiProxy_WorkmailUser) AddOverride(path *string, value interface{}) {
	if err := w.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (w *jsiiProxy_WorkmailUser) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := w.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkmailUser) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := w.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkmailUser) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := w.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkmailUser) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := w.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkmailUser) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := w.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkmailUser) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := w.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkmailUser) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := w.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkmailUser) GetStringAttribute(terraformAttribute *string) *string {
	if err := w.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkmailUser) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := w.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkmailUser) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkmailUser) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := w.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (w *jsiiProxy_WorkmailUser) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := w.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkmailUser) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := w.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		w,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkmailUser) MoveFromId(id *string) {
	if err := w.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"moveFromId",
		[]interface{}{id},
	)
}

func (w *jsiiProxy_WorkmailUser) MoveTo(moveTarget *string, index interface{}) {
	if err := w.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (w *jsiiProxy_WorkmailUser) MoveToId(id *string) {
	if err := w.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"moveToId",
		[]interface{}{id},
	)
}

func (w *jsiiProxy_WorkmailUser) OverrideLogicalId(newLogicalId *string) {
	if err := w.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (w *jsiiProxy_WorkmailUser) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := w.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (w *jsiiProxy_WorkmailUser) ResetCity() {
	_jsii_.InvokeVoid(
		w,
		"resetCity",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkmailUser) ResetCompany() {
	_jsii_.InvokeVoid(
		w,
		"resetCompany",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkmailUser) ResetCountry() {
	_jsii_.InvokeVoid(
		w,
		"resetCountry",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkmailUser) ResetDepartment() {
	_jsii_.InvokeVoid(
		w,
		"resetDepartment",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkmailUser) ResetFirstName() {
	_jsii_.InvokeVoid(
		w,
		"resetFirstName",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkmailUser) ResetHiddenFromGlobalAddressList() {
	_jsii_.InvokeVoid(
		w,
		"resetHiddenFromGlobalAddressList",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkmailUser) ResetIdentityProviderUserId() {
	_jsii_.InvokeVoid(
		w,
		"resetIdentityProviderUserId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkmailUser) ResetInitials() {
	_jsii_.InvokeVoid(
		w,
		"resetInitials",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkmailUser) ResetJobTitle() {
	_jsii_.InvokeVoid(
		w,
		"resetJobTitle",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkmailUser) ResetLastName() {
	_jsii_.InvokeVoid(
		w,
		"resetLastName",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkmailUser) ResetOffice() {
	_jsii_.InvokeVoid(
		w,
		"resetOffice",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkmailUser) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		w,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkmailUser) ResetPassword() {
	_jsii_.InvokeVoid(
		w,
		"resetPassword",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkmailUser) ResetRegion() {
	_jsii_.InvokeVoid(
		w,
		"resetRegion",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkmailUser) ResetStreet() {
	_jsii_.InvokeVoid(
		w,
		"resetStreet",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkmailUser) ResetTelephone() {
	_jsii_.InvokeVoid(
		w,
		"resetTelephone",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkmailUser) ResetUserRole() {
	_jsii_.InvokeVoid(
		w,
		"resetUserRole",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkmailUser) ResetZipCode() {
	_jsii_.InvokeVoid(
		w,
		"resetZipCode",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkmailUser) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkmailUser) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkmailUser) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkmailUser) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkmailUser) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkmailUser) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkmailUser) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		w,
		"with",
		args,
		&returns,
	)

	return returns
}

