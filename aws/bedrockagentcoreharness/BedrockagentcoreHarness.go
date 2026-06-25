// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreharness

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v24/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v24/bedrockagentcoreharness/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/bedrockagentcore_harness aws_bedrockagentcore_harness}.
type BedrockagentcoreHarness interface {
	cdktn.TerraformResource
	AllowedTools() *[]*string
	SetAllowedTools(val *[]*string)
	AllowedToolsInput() *[]*string
	Arn() *string
	AuthorizerConfiguration() BedrockagentcoreHarnessAuthorizerConfigurationList
	AuthorizerConfigurationInput() interface{}
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
	Environment() BedrockagentcoreHarnessEnvironmentList
	EnvironmentArtifact() BedrockagentcoreHarnessEnvironmentArtifactList
	EnvironmentArtifactInput() interface{}
	EnvironmentInput() interface{}
	EnvironmentVariables() *map[string]*string
	SetEnvironmentVariables(val *map[string]*string)
	EnvironmentVariablesInput() *map[string]*string
	ExecutionRoleArn() *string
	SetExecutionRoleArn(val *string)
	ExecutionRoleArnInput() *string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HarnessId() *string
	HarnessName() *string
	SetHarnessName(val *string)
	HarnessNameInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	MaxIterations() *float64
	SetMaxIterations(val *float64)
	MaxIterationsInput() *float64
	MaxTokens() *float64
	SetMaxTokens(val *float64)
	MaxTokensInput() *float64
	Memory() BedrockagentcoreHarnessMemoryList
	MemoryInput() interface{}
	Model() BedrockagentcoreHarnessModelList
	ModelInput() interface{}
	// The tree node.
	Node() constructs.Node
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
	Skill() BedrockagentcoreHarnessSkillList
	SkillInput() interface{}
	SystemPrompt() BedrockagentcoreHarnessSystemPromptList
	SystemPromptInput() interface{}
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsAll() cdktn.StringMap
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() BedrockagentcoreHarnessTimeoutsOutputReference
	TimeoutSeconds() *float64
	SetTimeoutSeconds(val *float64)
	TimeoutSecondsInput() *float64
	TimeoutsInput() interface{}
	Tool() BedrockagentcoreHarnessToolList
	ToolInput() interface{}
	Truncation() BedrockagentcoreHarnessTruncationList
	TruncationInput() interface{}
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
	PutAuthorizerConfiguration(value interface{})
	PutEnvironment(value interface{})
	PutEnvironmentArtifact(value interface{})
	PutMemory(value interface{})
	PutModel(value interface{})
	PutSkill(value interface{})
	PutSystemPrompt(value interface{})
	PutTimeouts(value *BedrockagentcoreHarnessTimeouts)
	PutTool(value interface{})
	PutTruncation(value interface{})
	ResetAllowedTools()
	ResetAuthorizerConfiguration()
	ResetEnvironment()
	ResetEnvironmentArtifact()
	ResetEnvironmentVariables()
	ResetMaxIterations()
	ResetMaxTokens()
	ResetMemory()
	ResetModel()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRegion()
	ResetSkill()
	ResetSystemPrompt()
	ResetTags()
	ResetTimeouts()
	ResetTimeoutSeconds()
	ResetTool()
	ResetTruncation()
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

// The jsii proxy struct for BedrockagentcoreHarness
type jsiiProxy_BedrockagentcoreHarness struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_BedrockagentcoreHarness) AllowedTools() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedTools",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarness) AllowedToolsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedToolsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarness) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarness) AuthorizerConfiguration() BedrockagentcoreHarnessAuthorizerConfigurationList {
	var returns BedrockagentcoreHarnessAuthorizerConfigurationList
	_jsii_.Get(
		j,
		"authorizerConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarness) AuthorizerConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authorizerConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarness) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarness) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarness) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarness) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarness) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarness) Environment() BedrockagentcoreHarnessEnvironmentList {
	var returns BedrockagentcoreHarnessEnvironmentList
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarness) EnvironmentArtifact() BedrockagentcoreHarnessEnvironmentArtifactList {
	var returns BedrockagentcoreHarnessEnvironmentArtifactList
	_jsii_.Get(
		j,
		"environmentArtifact",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarness) EnvironmentArtifactInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"environmentArtifactInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarness) EnvironmentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"environmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarness) EnvironmentVariables() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"environmentVariables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarness) EnvironmentVariablesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"environmentVariablesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarness) ExecutionRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarness) ExecutionRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarness) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarness) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarness) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarness) HarnessId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"harnessId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarness) HarnessName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"harnessName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarness) HarnessNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"harnessNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarness) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarness) MaxIterations() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxIterations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarness) MaxIterationsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxIterationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarness) MaxTokens() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxTokens",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarness) MaxTokensInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxTokensInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarness) Memory() BedrockagentcoreHarnessMemoryList {
	var returns BedrockagentcoreHarnessMemoryList
	_jsii_.Get(
		j,
		"memory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarness) MemoryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"memoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarness) Model() BedrockagentcoreHarnessModelList {
	var returns BedrockagentcoreHarnessModelList
	_jsii_.Get(
		j,
		"model",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarness) ModelInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"modelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarness) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarness) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarness) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarness) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarness) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarness) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarness) Skill() BedrockagentcoreHarnessSkillList {
	var returns BedrockagentcoreHarnessSkillList
	_jsii_.Get(
		j,
		"skill",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarness) SkillInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skillInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarness) SystemPrompt() BedrockagentcoreHarnessSystemPromptList {
	var returns BedrockagentcoreHarnessSystemPromptList
	_jsii_.Get(
		j,
		"systemPrompt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarness) SystemPromptInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"systemPromptInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarness) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarness) TagsAll() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarness) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarness) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarness) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarness) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarness) Timeouts() BedrockagentcoreHarnessTimeoutsOutputReference {
	var returns BedrockagentcoreHarnessTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarness) TimeoutSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarness) TimeoutSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarness) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarness) Tool() BedrockagentcoreHarnessToolList {
	var returns BedrockagentcoreHarnessToolList
	_jsii_.Get(
		j,
		"tool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarness) ToolInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"toolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarness) Truncation() BedrockagentcoreHarnessTruncationList {
	var returns BedrockagentcoreHarnessTruncationList
	_jsii_.Get(
		j,
		"truncation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarness) TruncationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"truncationInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/bedrockagentcore_harness aws_bedrockagentcore_harness} Resource.
func NewBedrockagentcoreHarness(scope constructs.Construct, id *string, config *BedrockagentcoreHarnessConfig) BedrockagentcoreHarness {
	_init_.Initialize()

	if err := validateNewBedrockagentcoreHarnessParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_BedrockagentcoreHarness{}

	_jsii_.Create(
		"@cdktn/provider-aws.bedrockagentcoreHarness.BedrockagentcoreHarness",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.52.0/docs/resources/bedrockagentcore_harness aws_bedrockagentcore_harness} Resource.
func NewBedrockagentcoreHarness_Override(b BedrockagentcoreHarness, scope constructs.Construct, id *string, config *BedrockagentcoreHarnessConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.bedrockagentcoreHarness.BedrockagentcoreHarness",
		[]interface{}{scope, id, config},
		b,
	)
}

func (j *jsiiProxy_BedrockagentcoreHarness)SetAllowedTools(val *[]*string) {
	if err := j.validateSetAllowedToolsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedTools",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreHarness)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreHarness)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreHarness)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreHarness)SetEnvironmentVariables(val *map[string]*string) {
	if err := j.validateSetEnvironmentVariablesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"environmentVariables",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreHarness)SetExecutionRoleArn(val *string) {
	if err := j.validateSetExecutionRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executionRoleArn",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreHarness)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreHarness)SetHarnessName(val *string) {
	if err := j.validateSetHarnessNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"harnessName",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreHarness)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreHarness)SetMaxIterations(val *float64) {
	if err := j.validateSetMaxIterationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxIterations",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreHarness)SetMaxTokens(val *float64) {
	if err := j.validateSetMaxTokensParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxTokens",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreHarness)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreHarness)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreHarness)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreHarness)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreHarness)SetTimeoutSeconds(val *float64) {
	if err := j.validateSetTimeoutSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeoutSeconds",
		val,
	)
}

// Generates CDKTN code for importing a BedrockagentcoreHarness resource upon running "cdktn plan <stack-name>".
func BedrockagentcoreHarness_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateBedrockagentcoreHarness_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-aws.bedrockagentcoreHarness.BedrockagentcoreHarness",
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
func BedrockagentcoreHarness_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBedrockagentcoreHarness_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-aws.bedrockagentcoreHarness.BedrockagentcoreHarness",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func BedrockagentcoreHarness_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBedrockagentcoreHarness_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-aws.bedrockagentcoreHarness.BedrockagentcoreHarness",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func BedrockagentcoreHarness_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBedrockagentcoreHarness_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-aws.bedrockagentcoreHarness.BedrockagentcoreHarness",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func BedrockagentcoreHarness_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-aws.bedrockagentcoreHarness.BedrockagentcoreHarness",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (b *jsiiProxy_BedrockagentcoreHarness) AddMoveTarget(moveTarget *string) {
	if err := b.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (b *jsiiProxy_BedrockagentcoreHarness) AddOverride(path *string, value interface{}) {
	if err := b.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (b *jsiiProxy_BedrockagentcoreHarness) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := b.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreHarness) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := b.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreHarness) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := b.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreHarness) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := b.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreHarness) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := b.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreHarness) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := b.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreHarness) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := b.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreHarness) GetStringAttribute(terraformAttribute *string) *string {
	if err := b.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreHarness) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := b.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreHarness) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreHarness) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := b.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (b *jsiiProxy_BedrockagentcoreHarness) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreHarness) MoveFromId(id *string) {
	if err := b.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"moveFromId",
		[]interface{}{id},
	)
}

func (b *jsiiProxy_BedrockagentcoreHarness) MoveTo(moveTarget *string, index interface{}) {
	if err := b.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (b *jsiiProxy_BedrockagentcoreHarness) MoveToId(id *string) {
	if err := b.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"moveToId",
		[]interface{}{id},
	)
}

func (b *jsiiProxy_BedrockagentcoreHarness) OverrideLogicalId(newLogicalId *string) {
	if err := b.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (b *jsiiProxy_BedrockagentcoreHarness) PutAuthorizerConfiguration(value interface{}) {
	if err := b.validatePutAuthorizerConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putAuthorizerConfiguration",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentcoreHarness) PutEnvironment(value interface{}) {
	if err := b.validatePutEnvironmentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putEnvironment",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentcoreHarness) PutEnvironmentArtifact(value interface{}) {
	if err := b.validatePutEnvironmentArtifactParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putEnvironmentArtifact",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentcoreHarness) PutMemory(value interface{}) {
	if err := b.validatePutMemoryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putMemory",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentcoreHarness) PutModel(value interface{}) {
	if err := b.validatePutModelParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putModel",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentcoreHarness) PutSkill(value interface{}) {
	if err := b.validatePutSkillParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putSkill",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentcoreHarness) PutSystemPrompt(value interface{}) {
	if err := b.validatePutSystemPromptParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putSystemPrompt",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentcoreHarness) PutTimeouts(value *BedrockagentcoreHarnessTimeouts) {
	if err := b.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentcoreHarness) PutTool(value interface{}) {
	if err := b.validatePutToolParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putTool",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentcoreHarness) PutTruncation(value interface{}) {
	if err := b.validatePutTruncationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putTruncation",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentcoreHarness) ResetAllowedTools() {
	_jsii_.InvokeVoid(
		b,
		"resetAllowedTools",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreHarness) ResetAuthorizerConfiguration() {
	_jsii_.InvokeVoid(
		b,
		"resetAuthorizerConfiguration",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreHarness) ResetEnvironment() {
	_jsii_.InvokeVoid(
		b,
		"resetEnvironment",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreHarness) ResetEnvironmentArtifact() {
	_jsii_.InvokeVoid(
		b,
		"resetEnvironmentArtifact",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreHarness) ResetEnvironmentVariables() {
	_jsii_.InvokeVoid(
		b,
		"resetEnvironmentVariables",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreHarness) ResetMaxIterations() {
	_jsii_.InvokeVoid(
		b,
		"resetMaxIterations",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreHarness) ResetMaxTokens() {
	_jsii_.InvokeVoid(
		b,
		"resetMaxTokens",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreHarness) ResetMemory() {
	_jsii_.InvokeVoid(
		b,
		"resetMemory",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreHarness) ResetModel() {
	_jsii_.InvokeVoid(
		b,
		"resetModel",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreHarness) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		b,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreHarness) ResetRegion() {
	_jsii_.InvokeVoid(
		b,
		"resetRegion",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreHarness) ResetSkill() {
	_jsii_.InvokeVoid(
		b,
		"resetSkill",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreHarness) ResetSystemPrompt() {
	_jsii_.InvokeVoid(
		b,
		"resetSystemPrompt",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreHarness) ResetTags() {
	_jsii_.InvokeVoid(
		b,
		"resetTags",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreHarness) ResetTimeouts() {
	_jsii_.InvokeVoid(
		b,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreHarness) ResetTimeoutSeconds() {
	_jsii_.InvokeVoid(
		b,
		"resetTimeoutSeconds",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreHarness) ResetTool() {
	_jsii_.InvokeVoid(
		b,
		"resetTool",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreHarness) ResetTruncation() {
	_jsii_.InvokeVoid(
		b,
		"resetTruncation",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreHarness) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreHarness) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreHarness) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreHarness) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreHarness) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreHarness) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreHarness) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		b,
		"with",
		args,
		&returns,
	)

	return returns
}

