// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerlabelingjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v22/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v22/sagemakerlabelingjob/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.31.0/docs/resources/sagemaker_labeling_job aws_sagemaker_labeling_job}.
type SagemakerLabelingJob interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
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
	FailureReason() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HumanTaskConfig() SagemakerLabelingJobHumanTaskConfigList
	HumanTaskConfigInput() interface{}
	InputConfig() SagemakerLabelingJobInputConfigList
	InputConfigInput() interface{}
	JobReferenceCode() *string
	LabelAttributeName() *string
	SetLabelAttributeName(val *string)
	LabelAttributeNameInput() *string
	LabelCategoryConfigS3Uri() *string
	SetLabelCategoryConfigS3Uri(val *string)
	LabelCategoryConfigS3UriInput() *string
	LabelCounters() SagemakerLabelingJobLabelCountersList
	LabelingJobAlgorithmsConfig() SagemakerLabelingJobLabelingJobAlgorithmsConfigList
	LabelingJobAlgorithmsConfigInput() interface{}
	LabelingJobArn() *string
	LabelingJobName() *string
	SetLabelingJobName(val *string)
	LabelingJobNameInput() *string
	LabelingJobStatus() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	OutputConfig() SagemakerLabelingJobOutputConfigList
	OutputConfigInput() interface{}
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	RoleArn() *string
	SetRoleArn(val *string)
	RoleArnInput() *string
	StoppingConditions() SagemakerLabelingJobStoppingConditionsList
	StoppingConditionsInput() interface{}
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsAll() cdktf.StringMap
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
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
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
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
	PutHumanTaskConfig(value interface{})
	PutInputConfig(value interface{})
	PutLabelingJobAlgorithmsConfig(value interface{})
	PutOutputConfig(value interface{})
	PutStoppingConditions(value interface{})
	ResetHumanTaskConfig()
	ResetInputConfig()
	ResetLabelCategoryConfigS3Uri()
	ResetLabelingJobAlgorithmsConfig()
	ResetOutputConfig()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRegion()
	ResetStoppingConditions()
	ResetTags()
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
}

// The jsii proxy struct for SagemakerLabelingJob
type jsiiProxy_SagemakerLabelingJob struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SagemakerLabelingJob) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerLabelingJob) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerLabelingJob) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerLabelingJob) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerLabelingJob) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerLabelingJob) FailureReason() *string {
	var returns *string
	_jsii_.Get(
		j,
		"failureReason",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerLabelingJob) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerLabelingJob) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerLabelingJob) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerLabelingJob) HumanTaskConfig() SagemakerLabelingJobHumanTaskConfigList {
	var returns SagemakerLabelingJobHumanTaskConfigList
	_jsii_.Get(
		j,
		"humanTaskConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerLabelingJob) HumanTaskConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"humanTaskConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerLabelingJob) InputConfig() SagemakerLabelingJobInputConfigList {
	var returns SagemakerLabelingJobInputConfigList
	_jsii_.Get(
		j,
		"inputConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerLabelingJob) InputConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inputConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerLabelingJob) JobReferenceCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobReferenceCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerLabelingJob) LabelAttributeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"labelAttributeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerLabelingJob) LabelAttributeNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"labelAttributeNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerLabelingJob) LabelCategoryConfigS3Uri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"labelCategoryConfigS3Uri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerLabelingJob) LabelCategoryConfigS3UriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"labelCategoryConfigS3UriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerLabelingJob) LabelCounters() SagemakerLabelingJobLabelCountersList {
	var returns SagemakerLabelingJobLabelCountersList
	_jsii_.Get(
		j,
		"labelCounters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerLabelingJob) LabelingJobAlgorithmsConfig() SagemakerLabelingJobLabelingJobAlgorithmsConfigList {
	var returns SagemakerLabelingJobLabelingJobAlgorithmsConfigList
	_jsii_.Get(
		j,
		"labelingJobAlgorithmsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerLabelingJob) LabelingJobAlgorithmsConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"labelingJobAlgorithmsConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerLabelingJob) LabelingJobArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"labelingJobArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerLabelingJob) LabelingJobName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"labelingJobName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerLabelingJob) LabelingJobNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"labelingJobNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerLabelingJob) LabelingJobStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"labelingJobStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerLabelingJob) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerLabelingJob) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerLabelingJob) OutputConfig() SagemakerLabelingJobOutputConfigList {
	var returns SagemakerLabelingJobOutputConfigList
	_jsii_.Get(
		j,
		"outputConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerLabelingJob) OutputConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"outputConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerLabelingJob) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerLabelingJob) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerLabelingJob) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerLabelingJob) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerLabelingJob) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerLabelingJob) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerLabelingJob) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerLabelingJob) StoppingConditions() SagemakerLabelingJobStoppingConditionsList {
	var returns SagemakerLabelingJobStoppingConditionsList
	_jsii_.Get(
		j,
		"stoppingConditions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerLabelingJob) StoppingConditionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stoppingConditionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerLabelingJob) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerLabelingJob) TagsAll() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerLabelingJob) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerLabelingJob) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerLabelingJob) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerLabelingJob) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.31.0/docs/resources/sagemaker_labeling_job aws_sagemaker_labeling_job} Resource.
func NewSagemakerLabelingJob(scope constructs.Construct, id *string, config *SagemakerLabelingJobConfig) SagemakerLabelingJob {
	_init_.Initialize()

	if err := validateNewSagemakerLabelingJobParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_SagemakerLabelingJob{}

	_jsii_.Create(
		"@cdktn/provider-aws.sagemakerLabelingJob.SagemakerLabelingJob",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.31.0/docs/resources/sagemaker_labeling_job aws_sagemaker_labeling_job} Resource.
func NewSagemakerLabelingJob_Override(s SagemakerLabelingJob, scope constructs.Construct, id *string, config *SagemakerLabelingJobConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.sagemakerLabelingJob.SagemakerLabelingJob",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SagemakerLabelingJob)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_SagemakerLabelingJob)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SagemakerLabelingJob)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SagemakerLabelingJob)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_SagemakerLabelingJob)SetLabelAttributeName(val *string) {
	if err := j.validateSetLabelAttributeNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labelAttributeName",
		val,
	)
}

func (j *jsiiProxy_SagemakerLabelingJob)SetLabelCategoryConfigS3Uri(val *string) {
	if err := j.validateSetLabelCategoryConfigS3UriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labelCategoryConfigS3Uri",
		val,
	)
}

func (j *jsiiProxy_SagemakerLabelingJob)SetLabelingJobName(val *string) {
	if err := j.validateSetLabelingJobNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labelingJobName",
		val,
	)
}

func (j *jsiiProxy_SagemakerLabelingJob)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SagemakerLabelingJob)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SagemakerLabelingJob)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_SagemakerLabelingJob)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_SagemakerLabelingJob)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_SagemakerLabelingJob)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

// Generates CDKTF code for importing a SagemakerLabelingJob resource upon running "cdktf plan <stack-name>".
func SagemakerLabelingJob_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateSagemakerLabelingJob_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-aws.sagemakerLabelingJob.SagemakerLabelingJob",
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
func SagemakerLabelingJob_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSagemakerLabelingJob_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-aws.sagemakerLabelingJob.SagemakerLabelingJob",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SagemakerLabelingJob_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSagemakerLabelingJob_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-aws.sagemakerLabelingJob.SagemakerLabelingJob",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SagemakerLabelingJob_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSagemakerLabelingJob_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-aws.sagemakerLabelingJob.SagemakerLabelingJob",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SagemakerLabelingJob_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-aws.sagemakerLabelingJob.SagemakerLabelingJob",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_SagemakerLabelingJob) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_SagemakerLabelingJob) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_SagemakerLabelingJob) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerLabelingJob) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerLabelingJob) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerLabelingJob) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerLabelingJob) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerLabelingJob) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerLabelingJob) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerLabelingJob) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerLabelingJob) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerLabelingJob) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerLabelingJob) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_SagemakerLabelingJob) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerLabelingJob) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SagemakerLabelingJob) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_SagemakerLabelingJob) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SagemakerLabelingJob) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SagemakerLabelingJob) PutHumanTaskConfig(value interface{}) {
	if err := s.validatePutHumanTaskConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putHumanTaskConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerLabelingJob) PutInputConfig(value interface{}) {
	if err := s.validatePutInputConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putInputConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerLabelingJob) PutLabelingJobAlgorithmsConfig(value interface{}) {
	if err := s.validatePutLabelingJobAlgorithmsConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putLabelingJobAlgorithmsConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerLabelingJob) PutOutputConfig(value interface{}) {
	if err := s.validatePutOutputConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putOutputConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerLabelingJob) PutStoppingConditions(value interface{}) {
	if err := s.validatePutStoppingConditionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putStoppingConditions",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerLabelingJob) ResetHumanTaskConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetHumanTaskConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerLabelingJob) ResetInputConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetInputConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerLabelingJob) ResetLabelCategoryConfigS3Uri() {
	_jsii_.InvokeVoid(
		s,
		"resetLabelCategoryConfigS3Uri",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerLabelingJob) ResetLabelingJobAlgorithmsConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetLabelingJobAlgorithmsConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerLabelingJob) ResetOutputConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetOutputConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerLabelingJob) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerLabelingJob) ResetRegion() {
	_jsii_.InvokeVoid(
		s,
		"resetRegion",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerLabelingJob) ResetStoppingConditions() {
	_jsii_.InvokeVoid(
		s,
		"resetStoppingConditions",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerLabelingJob) ResetTags() {
	_jsii_.InvokeVoid(
		s,
		"resetTags",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerLabelingJob) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerLabelingJob) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerLabelingJob) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerLabelingJob) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerLabelingJob) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerLabelingJob) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

