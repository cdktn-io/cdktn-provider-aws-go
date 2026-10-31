// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawssagemakerhubcontent

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v26/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v26/dataawssagemakerhubcontent/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.71.0/docs/data-sources/sagemaker_hub_content aws_sagemaker_hub_content}.
type DataAwsSagemakerHubContent interface {
	cdktn.TerraformDataSource
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreationTime() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DocumentSchemaVersion() *string
	FailureReason() *string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HubArn() *string
	HubContentArn() *string
	HubContentDependency() DataAwsSagemakerHubContentHubContentDependencyList
	HubContentDescription() *string
	HubContentDisplayName() *string
	HubContentDocument() *string
	HubContentMarkdown() *string
	HubContentName() *string
	SetHubContentName(val *string)
	HubContentNameInput() *string
	HubContentSearchKeywords() *[]*string
	HubContentStatus() *string
	HubContentType() *string
	SetHubContentType(val *string)
	HubContentTypeInput() *string
	HubContentVersion() *string
	SetHubContentVersion(val *string)
	HubContentVersionInput() *string
	HubName() *string
	SetHubName(val *string)
	HubNameInput() *string
	LastModifiedTime() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	ReferenceMinVersion() *string
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	SagemakerPublicHubContentArn() *string
	SupportStatus() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	VersionedHubContentArn() *string
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
	ResetHubContentVersion()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRegion()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Adds this resource to the terraform JSON output.
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

// The jsii proxy struct for DataAwsSagemakerHubContent
type jsiiProxy_DataAwsSagemakerHubContent struct {
	internal.Type__cdktnTerraformDataSource
}

func (j *jsiiProxy_DataAwsSagemakerHubContent) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSagemakerHubContent) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSagemakerHubContent) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSagemakerHubContent) CreationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSagemakerHubContent) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSagemakerHubContent) DocumentSchemaVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentSchemaVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSagemakerHubContent) FailureReason() *string {
	var returns *string
	_jsii_.Get(
		j,
		"failureReason",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSagemakerHubContent) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSagemakerHubContent) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSagemakerHubContent) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSagemakerHubContent) HubArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hubArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSagemakerHubContent) HubContentArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hubContentArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSagemakerHubContent) HubContentDependency() DataAwsSagemakerHubContentHubContentDependencyList {
	var returns DataAwsSagemakerHubContentHubContentDependencyList
	_jsii_.Get(
		j,
		"hubContentDependency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSagemakerHubContent) HubContentDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hubContentDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSagemakerHubContent) HubContentDisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hubContentDisplayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSagemakerHubContent) HubContentDocument() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hubContentDocument",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSagemakerHubContent) HubContentMarkdown() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hubContentMarkdown",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSagemakerHubContent) HubContentName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hubContentName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSagemakerHubContent) HubContentNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hubContentNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSagemakerHubContent) HubContentSearchKeywords() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"hubContentSearchKeywords",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSagemakerHubContent) HubContentStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hubContentStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSagemakerHubContent) HubContentType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hubContentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSagemakerHubContent) HubContentTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hubContentTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSagemakerHubContent) HubContentVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hubContentVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSagemakerHubContent) HubContentVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hubContentVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSagemakerHubContent) HubName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hubName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSagemakerHubContent) HubNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hubNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSagemakerHubContent) LastModifiedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastModifiedTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSagemakerHubContent) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSagemakerHubContent) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSagemakerHubContent) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSagemakerHubContent) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSagemakerHubContent) ReferenceMinVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"referenceMinVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSagemakerHubContent) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSagemakerHubContent) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSagemakerHubContent) SagemakerPublicHubContentArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sagemakerPublicHubContentArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSagemakerHubContent) SupportStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"supportStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSagemakerHubContent) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSagemakerHubContent) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSagemakerHubContent) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSagemakerHubContent) VersionedHubContentArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionedHubContentArn",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.71.0/docs/data-sources/sagemaker_hub_content aws_sagemaker_hub_content} Data Source.
func NewDataAwsSagemakerHubContent(scope constructs.Construct, id *string, config *DataAwsSagemakerHubContentConfig) DataAwsSagemakerHubContent {
	_init_.Initialize()

	if err := validateNewDataAwsSagemakerHubContentParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsSagemakerHubContent{}

	_jsii_.Create(
		"@cdktn/provider-aws.dataAwsSagemakerHubContent.DataAwsSagemakerHubContent",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.71.0/docs/data-sources/sagemaker_hub_content aws_sagemaker_hub_content} Data Source.
func NewDataAwsSagemakerHubContent_Override(d DataAwsSagemakerHubContent, scope constructs.Construct, id *string, config *DataAwsSagemakerHubContentConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.dataAwsSagemakerHubContent.DataAwsSagemakerHubContent",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsSagemakerHubContent)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsSagemakerHubContent)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsSagemakerHubContent)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataAwsSagemakerHubContent)SetHubContentName(val *string) {
	if err := j.validateSetHubContentNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hubContentName",
		val,
	)
}

func (j *jsiiProxy_DataAwsSagemakerHubContent)SetHubContentType(val *string) {
	if err := j.validateSetHubContentTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hubContentType",
		val,
	)
}

func (j *jsiiProxy_DataAwsSagemakerHubContent)SetHubContentVersion(val *string) {
	if err := j.validateSetHubContentVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hubContentVersion",
		val,
	)
}

func (j *jsiiProxy_DataAwsSagemakerHubContent)SetHubName(val *string) {
	if err := j.validateSetHubNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hubName",
		val,
	)
}

func (j *jsiiProxy_DataAwsSagemakerHubContent)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsSagemakerHubContent)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataAwsSagemakerHubContent)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

// Generates CDKTN code for importing a DataAwsSagemakerHubContent resource upon running "cdktn plan <stack-name>".
func DataAwsSagemakerHubContent_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateDataAwsSagemakerHubContent_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-aws.dataAwsSagemakerHubContent.DataAwsSagemakerHubContent",
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
func DataAwsSagemakerHubContent_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsSagemakerHubContent_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-aws.dataAwsSagemakerHubContent.DataAwsSagemakerHubContent",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAwsSagemakerHubContent_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsSagemakerHubContent_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-aws.dataAwsSagemakerHubContent.DataAwsSagemakerHubContent",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAwsSagemakerHubContent_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsSagemakerHubContent_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-aws.dataAwsSagemakerHubContent.DataAwsSagemakerHubContent",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsSagemakerHubContent_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-aws.dataAwsSagemakerHubContent.DataAwsSagemakerHubContent",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataAwsSagemakerHubContent) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataAwsSagemakerHubContent) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsSagemakerHubContent) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsSagemakerHubContent) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsSagemakerHubContent) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsSagemakerHubContent) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsSagemakerHubContent) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsSagemakerHubContent) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsSagemakerHubContent) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsSagemakerHubContent) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsSagemakerHubContent) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsSagemakerHubContent) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataAwsSagemakerHubContent) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := d.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (d *jsiiProxy_DataAwsSagemakerHubContent) ResetHubContentVersion() {
	_jsii_.InvokeVoid(
		d,
		"resetHubContentVersion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsSagemakerHubContent) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsSagemakerHubContent) ResetRegion() {
	_jsii_.InvokeVoid(
		d,
		"resetRegion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsSagemakerHubContent) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsSagemakerHubContent) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsSagemakerHubContent) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsSagemakerHubContent) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsSagemakerHubContent) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsSagemakerHubContent) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsSagemakerHubContent) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		d,
		"with",
		args,
		&returns,
	)

	return returns
}

