// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dynamodbglobalsecondaryindex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v24/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v24/dynamodbglobalsecondaryindex/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.47.0/docs/resources/dynamodb_global_secondary_index aws_dynamodb_global_secondary_index}.
type DynamodbGlobalSecondaryIndex interface {
	cdktn.TerraformResource
	Arn() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
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
	IndexName() *string
	SetIndexName(val *string)
	IndexNameInput() *string
	KeySchema() DynamodbGlobalSecondaryIndexKeySchemaList
	KeySchemaInput() interface{}
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	OnDemandThroughput() DynamodbGlobalSecondaryIndexOnDemandThroughputList
	OnDemandThroughputInput() interface{}
	Projection() DynamodbGlobalSecondaryIndexProjectionList
	ProjectionInput() interface{}
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	ProvisionedThroughput() DynamodbGlobalSecondaryIndexProvisionedThroughputList
	ProvisionedThroughputInput() interface{}
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	TableName() *string
	SetTableName(val *string)
	TableNameInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() DynamodbGlobalSecondaryIndexTimeoutsOutputReference
	TimeoutsInput() interface{}
	WarmThroughput() DynamodbGlobalSecondaryIndexWarmThroughputOutputReference
	WarmThroughputInput() interface{}
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
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
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
	PutKeySchema(value interface{})
	PutOnDemandThroughput(value interface{})
	PutProjection(value interface{})
	PutProvisionedThroughput(value interface{})
	PutTimeouts(value *DynamodbGlobalSecondaryIndexTimeouts)
	PutWarmThroughput(value *DynamodbGlobalSecondaryIndexWarmThroughput)
	ResetKeySchema()
	ResetOnDemandThroughput()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProjection()
	ResetProvisionedThroughput()
	ResetRegion()
	ResetTimeouts()
	ResetWarmThroughput()
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

// The jsii proxy struct for DynamodbGlobalSecondaryIndex
type jsiiProxy_DynamodbGlobalSecondaryIndex struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_DynamodbGlobalSecondaryIndex) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalSecondaryIndex) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalSecondaryIndex) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalSecondaryIndex) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalSecondaryIndex) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalSecondaryIndex) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalSecondaryIndex) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalSecondaryIndex) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalSecondaryIndex) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalSecondaryIndex) IndexName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"indexName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalSecondaryIndex) IndexNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"indexNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalSecondaryIndex) KeySchema() DynamodbGlobalSecondaryIndexKeySchemaList {
	var returns DynamodbGlobalSecondaryIndexKeySchemaList
	_jsii_.Get(
		j,
		"keySchema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalSecondaryIndex) KeySchemaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"keySchemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalSecondaryIndex) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalSecondaryIndex) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalSecondaryIndex) OnDemandThroughput() DynamodbGlobalSecondaryIndexOnDemandThroughputList {
	var returns DynamodbGlobalSecondaryIndexOnDemandThroughputList
	_jsii_.Get(
		j,
		"onDemandThroughput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalSecondaryIndex) OnDemandThroughputInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"onDemandThroughputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalSecondaryIndex) Projection() DynamodbGlobalSecondaryIndexProjectionList {
	var returns DynamodbGlobalSecondaryIndexProjectionList
	_jsii_.Get(
		j,
		"projection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalSecondaryIndex) ProjectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"projectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalSecondaryIndex) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalSecondaryIndex) ProvisionedThroughput() DynamodbGlobalSecondaryIndexProvisionedThroughputList {
	var returns DynamodbGlobalSecondaryIndexProvisionedThroughputList
	_jsii_.Get(
		j,
		"provisionedThroughput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalSecondaryIndex) ProvisionedThroughputInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"provisionedThroughputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalSecondaryIndex) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalSecondaryIndex) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalSecondaryIndex) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalSecondaryIndex) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalSecondaryIndex) TableName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalSecondaryIndex) TableNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalSecondaryIndex) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalSecondaryIndex) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalSecondaryIndex) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalSecondaryIndex) Timeouts() DynamodbGlobalSecondaryIndexTimeoutsOutputReference {
	var returns DynamodbGlobalSecondaryIndexTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalSecondaryIndex) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalSecondaryIndex) WarmThroughput() DynamodbGlobalSecondaryIndexWarmThroughputOutputReference {
	var returns DynamodbGlobalSecondaryIndexWarmThroughputOutputReference
	_jsii_.Get(
		j,
		"warmThroughput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalSecondaryIndex) WarmThroughputInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"warmThroughputInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.47.0/docs/resources/dynamodb_global_secondary_index aws_dynamodb_global_secondary_index} Resource.
func NewDynamodbGlobalSecondaryIndex(scope constructs.Construct, id *string, config *DynamodbGlobalSecondaryIndexConfig) DynamodbGlobalSecondaryIndex {
	_init_.Initialize()

	if err := validateNewDynamodbGlobalSecondaryIndexParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DynamodbGlobalSecondaryIndex{}

	_jsii_.Create(
		"@cdktn/provider-aws.dynamodbGlobalSecondaryIndex.DynamodbGlobalSecondaryIndex",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.47.0/docs/resources/dynamodb_global_secondary_index aws_dynamodb_global_secondary_index} Resource.
func NewDynamodbGlobalSecondaryIndex_Override(d DynamodbGlobalSecondaryIndex, scope constructs.Construct, id *string, config *DynamodbGlobalSecondaryIndexConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.dynamodbGlobalSecondaryIndex.DynamodbGlobalSecondaryIndex",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DynamodbGlobalSecondaryIndex)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DynamodbGlobalSecondaryIndex)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DynamodbGlobalSecondaryIndex)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DynamodbGlobalSecondaryIndex)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DynamodbGlobalSecondaryIndex)SetIndexName(val *string) {
	if err := j.validateSetIndexNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"indexName",
		val,
	)
}

func (j *jsiiProxy_DynamodbGlobalSecondaryIndex)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DynamodbGlobalSecondaryIndex)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DynamodbGlobalSecondaryIndex)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_DynamodbGlobalSecondaryIndex)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_DynamodbGlobalSecondaryIndex)SetTableName(val *string) {
	if err := j.validateSetTableNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tableName",
		val,
	)
}

// Generates CDKTN code for importing a DynamodbGlobalSecondaryIndex resource upon running "cdktn plan <stack-name>".
func DynamodbGlobalSecondaryIndex_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateDynamodbGlobalSecondaryIndex_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-aws.dynamodbGlobalSecondaryIndex.DynamodbGlobalSecondaryIndex",
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
func DynamodbGlobalSecondaryIndex_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDynamodbGlobalSecondaryIndex_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-aws.dynamodbGlobalSecondaryIndex.DynamodbGlobalSecondaryIndex",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DynamodbGlobalSecondaryIndex_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDynamodbGlobalSecondaryIndex_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-aws.dynamodbGlobalSecondaryIndex.DynamodbGlobalSecondaryIndex",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DynamodbGlobalSecondaryIndex_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDynamodbGlobalSecondaryIndex_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-aws.dynamodbGlobalSecondaryIndex.DynamodbGlobalSecondaryIndex",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DynamodbGlobalSecondaryIndex_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-aws.dynamodbGlobalSecondaryIndex.DynamodbGlobalSecondaryIndex",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DynamodbGlobalSecondaryIndex) AddMoveTarget(moveTarget *string) {
	if err := d.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (d *jsiiProxy_DynamodbGlobalSecondaryIndex) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DynamodbGlobalSecondaryIndex) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DynamodbGlobalSecondaryIndex) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DynamodbGlobalSecondaryIndex) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DynamodbGlobalSecondaryIndex) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DynamodbGlobalSecondaryIndex) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DynamodbGlobalSecondaryIndex) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DynamodbGlobalSecondaryIndex) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DynamodbGlobalSecondaryIndex) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DynamodbGlobalSecondaryIndex) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DynamodbGlobalSecondaryIndex) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamodbGlobalSecondaryIndex) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := d.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (d *jsiiProxy_DynamodbGlobalSecondaryIndex) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DynamodbGlobalSecondaryIndex) MoveFromId(id *string) {
	if err := d.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveFromId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DynamodbGlobalSecondaryIndex) MoveTo(moveTarget *string, index interface{}) {
	if err := d.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (d *jsiiProxy_DynamodbGlobalSecondaryIndex) MoveToId(id *string) {
	if err := d.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveToId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DynamodbGlobalSecondaryIndex) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DynamodbGlobalSecondaryIndex) PutKeySchema(value interface{}) {
	if err := d.validatePutKeySchemaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putKeySchema",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DynamodbGlobalSecondaryIndex) PutOnDemandThroughput(value interface{}) {
	if err := d.validatePutOnDemandThroughputParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putOnDemandThroughput",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DynamodbGlobalSecondaryIndex) PutProjection(value interface{}) {
	if err := d.validatePutProjectionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putProjection",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DynamodbGlobalSecondaryIndex) PutProvisionedThroughput(value interface{}) {
	if err := d.validatePutProvisionedThroughputParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putProvisionedThroughput",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DynamodbGlobalSecondaryIndex) PutTimeouts(value *DynamodbGlobalSecondaryIndexTimeouts) {
	if err := d.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DynamodbGlobalSecondaryIndex) PutWarmThroughput(value *DynamodbGlobalSecondaryIndexWarmThroughput) {
	if err := d.validatePutWarmThroughputParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putWarmThroughput",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DynamodbGlobalSecondaryIndex) ResetKeySchema() {
	_jsii_.InvokeVoid(
		d,
		"resetKeySchema",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbGlobalSecondaryIndex) ResetOnDemandThroughput() {
	_jsii_.InvokeVoid(
		d,
		"resetOnDemandThroughput",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbGlobalSecondaryIndex) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbGlobalSecondaryIndex) ResetProjection() {
	_jsii_.InvokeVoid(
		d,
		"resetProjection",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbGlobalSecondaryIndex) ResetProvisionedThroughput() {
	_jsii_.InvokeVoid(
		d,
		"resetProvisionedThroughput",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbGlobalSecondaryIndex) ResetRegion() {
	_jsii_.InvokeVoid(
		d,
		"resetRegion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbGlobalSecondaryIndex) ResetTimeouts() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbGlobalSecondaryIndex) ResetWarmThroughput() {
	_jsii_.InvokeVoid(
		d,
		"resetWarmThroughput",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbGlobalSecondaryIndex) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamodbGlobalSecondaryIndex) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamodbGlobalSecondaryIndex) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamodbGlobalSecondaryIndex) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamodbGlobalSecondaryIndex) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamodbGlobalSecondaryIndex) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamodbGlobalSecondaryIndex) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

