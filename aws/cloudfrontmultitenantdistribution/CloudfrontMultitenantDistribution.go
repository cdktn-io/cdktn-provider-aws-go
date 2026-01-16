// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudfrontmultitenantdistribution

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v21/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v21/cloudfrontmultitenantdistribution/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_multitenant_distribution aws_cloudfront_multitenant_distribution}.
type CloudfrontMultitenantDistribution interface {
	cdktf.TerraformResource
	ActiveTrustedKeyGroups() CloudfrontMultitenantDistributionActiveTrustedKeyGroupsList
	ActiveTrustedKeyGroupsInput() interface{}
	Arn() *string
	CacheBehavior() CloudfrontMultitenantDistributionCacheBehaviorList
	CacheBehaviorInput() interface{}
	CallerReference() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Comment() *string
	SetComment(val *string)
	CommentInput() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	ConnectionMode() *string
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CustomErrorResponse() CloudfrontMultitenantDistributionCustomErrorResponseList
	CustomErrorResponseInput() interface{}
	DefaultCacheBehavior() CloudfrontMultitenantDistributionDefaultCacheBehaviorList
	DefaultCacheBehaviorInput() interface{}
	DefaultRootObject() *string
	SetDefaultRootObject(val *string)
	DefaultRootObjectInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DomainName() *string
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	Etag() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HttpVersion() *string
	SetHttpVersion(val *string)
	HttpVersionInput() *string
	Id() *string
	InProgressInvalidationBatches() *float64
	LastModifiedTime() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	Origin() CloudfrontMultitenantDistributionOriginList
	OriginGroup() CloudfrontMultitenantDistributionOriginGroupList
	OriginGroupInput() interface{}
	OriginInput() interface{}
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
	Restrictions() CloudfrontMultitenantDistributionRestrictionsList
	RestrictionsInput() interface{}
	Status() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsAll() cdktf.StringMap
	TagsInput() *map[string]*string
	TenantConfig() CloudfrontMultitenantDistributionTenantConfigList
	TenantConfigInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() CloudfrontMultitenantDistributionTimeoutsOutputReference
	TimeoutsInput() interface{}
	ViewerCertificate() CloudfrontMultitenantDistributionViewerCertificateList
	ViewerCertificateInput() interface{}
	WebAclId() *string
	SetWebAclId(val *string)
	WebAclIdInput() *string
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
	PutActiveTrustedKeyGroups(value interface{})
	PutCacheBehavior(value interface{})
	PutCustomErrorResponse(value interface{})
	PutDefaultCacheBehavior(value interface{})
	PutOrigin(value interface{})
	PutOriginGroup(value interface{})
	PutRestrictions(value interface{})
	PutTenantConfig(value interface{})
	PutTimeouts(value *CloudfrontMultitenantDistributionTimeouts)
	PutViewerCertificate(value interface{})
	ResetActiveTrustedKeyGroups()
	ResetCacheBehavior()
	ResetCustomErrorResponse()
	ResetDefaultCacheBehavior()
	ResetDefaultRootObject()
	ResetHttpVersion()
	ResetOrigin()
	ResetOriginGroup()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRestrictions()
	ResetTags()
	ResetTenantConfig()
	ResetTimeouts()
	ResetViewerCertificate()
	ResetWebAclId()
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

// The jsii proxy struct for CloudfrontMultitenantDistribution
type jsiiProxy_CloudfrontMultitenantDistribution struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_CloudfrontMultitenantDistribution) ActiveTrustedKeyGroups() CloudfrontMultitenantDistributionActiveTrustedKeyGroupsList {
	var returns CloudfrontMultitenantDistributionActiveTrustedKeyGroupsList
	_jsii_.Get(
		j,
		"activeTrustedKeyGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistribution) ActiveTrustedKeyGroupsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"activeTrustedKeyGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistribution) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistribution) CacheBehavior() CloudfrontMultitenantDistributionCacheBehaviorList {
	var returns CloudfrontMultitenantDistributionCacheBehaviorList
	_jsii_.Get(
		j,
		"cacheBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistribution) CacheBehaviorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cacheBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistribution) CallerReference() *string {
	var returns *string
	_jsii_.Get(
		j,
		"callerReference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistribution) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistribution) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistribution) CommentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistribution) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistribution) ConnectionMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistribution) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistribution) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistribution) CustomErrorResponse() CloudfrontMultitenantDistributionCustomErrorResponseList {
	var returns CloudfrontMultitenantDistributionCustomErrorResponseList
	_jsii_.Get(
		j,
		"customErrorResponse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistribution) CustomErrorResponseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customErrorResponseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistribution) DefaultCacheBehavior() CloudfrontMultitenantDistributionDefaultCacheBehaviorList {
	var returns CloudfrontMultitenantDistributionDefaultCacheBehaviorList
	_jsii_.Get(
		j,
		"defaultCacheBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistribution) DefaultCacheBehaviorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultCacheBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistribution) DefaultRootObject() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultRootObject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistribution) DefaultRootObjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultRootObjectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistribution) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistribution) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistribution) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistribution) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistribution) Etag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistribution) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistribution) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistribution) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistribution) HttpVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistribution) HttpVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistribution) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistribution) InProgressInvalidationBatches() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"inProgressInvalidationBatches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistribution) LastModifiedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastModifiedTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistribution) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistribution) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistribution) Origin() CloudfrontMultitenantDistributionOriginList {
	var returns CloudfrontMultitenantDistributionOriginList
	_jsii_.Get(
		j,
		"origin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistribution) OriginGroup() CloudfrontMultitenantDistributionOriginGroupList {
	var returns CloudfrontMultitenantDistributionOriginGroupList
	_jsii_.Get(
		j,
		"originGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistribution) OriginGroupInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"originGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistribution) OriginInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"originInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistribution) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistribution) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistribution) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistribution) Restrictions() CloudfrontMultitenantDistributionRestrictionsList {
	var returns CloudfrontMultitenantDistributionRestrictionsList
	_jsii_.Get(
		j,
		"restrictions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistribution) RestrictionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"restrictionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistribution) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistribution) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistribution) TagsAll() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistribution) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistribution) TenantConfig() CloudfrontMultitenantDistributionTenantConfigList {
	var returns CloudfrontMultitenantDistributionTenantConfigList
	_jsii_.Get(
		j,
		"tenantConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistribution) TenantConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tenantConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistribution) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistribution) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistribution) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistribution) Timeouts() CloudfrontMultitenantDistributionTimeoutsOutputReference {
	var returns CloudfrontMultitenantDistributionTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistribution) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistribution) ViewerCertificate() CloudfrontMultitenantDistributionViewerCertificateList {
	var returns CloudfrontMultitenantDistributionViewerCertificateList
	_jsii_.Get(
		j,
		"viewerCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistribution) ViewerCertificateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"viewerCertificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistribution) WebAclId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webAclId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistribution) WebAclIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webAclIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_multitenant_distribution aws_cloudfront_multitenant_distribution} Resource.
func NewCloudfrontMultitenantDistribution(scope constructs.Construct, id *string, config *CloudfrontMultitenantDistributionConfig) CloudfrontMultitenantDistribution {
	_init_.Initialize()

	if err := validateNewCloudfrontMultitenantDistributionParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudfrontMultitenantDistribution{}

	_jsii_.Create(
		"@cdktn/provider-aws.cloudfrontMultitenantDistribution.CloudfrontMultitenantDistribution",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.28.0/docs/resources/cloudfront_multitenant_distribution aws_cloudfront_multitenant_distribution} Resource.
func NewCloudfrontMultitenantDistribution_Override(c CloudfrontMultitenantDistribution, scope constructs.Construct, id *string, config *CloudfrontMultitenantDistributionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.cloudfrontMultitenantDistribution.CloudfrontMultitenantDistribution",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_CloudfrontMultitenantDistribution)SetComment(val *string) {
	if err := j.validateSetCommentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"comment",
		val,
	)
}

func (j *jsiiProxy_CloudfrontMultitenantDistribution)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_CloudfrontMultitenantDistribution)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_CloudfrontMultitenantDistribution)SetDefaultRootObject(val *string) {
	if err := j.validateSetDefaultRootObjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultRootObject",
		val,
	)
}

func (j *jsiiProxy_CloudfrontMultitenantDistribution)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_CloudfrontMultitenantDistribution)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_CloudfrontMultitenantDistribution)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_CloudfrontMultitenantDistribution)SetHttpVersion(val *string) {
	if err := j.validateSetHttpVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpVersion",
		val,
	)
}

func (j *jsiiProxy_CloudfrontMultitenantDistribution)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_CloudfrontMultitenantDistribution)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_CloudfrontMultitenantDistribution)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_CloudfrontMultitenantDistribution)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_CloudfrontMultitenantDistribution)SetWebAclId(val *string) {
	if err := j.validateSetWebAclIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"webAclId",
		val,
	)
}

// Generates CDKTF code for importing a CloudfrontMultitenantDistribution resource upon running "cdktf plan <stack-name>".
func CloudfrontMultitenantDistribution_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateCloudfrontMultitenantDistribution_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-aws.cloudfrontMultitenantDistribution.CloudfrontMultitenantDistribution",
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
func CloudfrontMultitenantDistribution_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCloudfrontMultitenantDistribution_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-aws.cloudfrontMultitenantDistribution.CloudfrontMultitenantDistribution",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CloudfrontMultitenantDistribution_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCloudfrontMultitenantDistribution_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-aws.cloudfrontMultitenantDistribution.CloudfrontMultitenantDistribution",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CloudfrontMultitenantDistribution_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCloudfrontMultitenantDistribution_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-aws.cloudfrontMultitenantDistribution.CloudfrontMultitenantDistribution",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CloudfrontMultitenantDistribution_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-aws.cloudfrontMultitenantDistribution.CloudfrontMultitenantDistribution",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CloudfrontMultitenantDistribution) AddMoveTarget(moveTarget *string) {
	if err := c.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (c *jsiiProxy_CloudfrontMultitenantDistribution) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CloudfrontMultitenantDistribution) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontMultitenantDistribution) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontMultitenantDistribution) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontMultitenantDistribution) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontMultitenantDistribution) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontMultitenantDistribution) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontMultitenantDistribution) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontMultitenantDistribution) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontMultitenantDistribution) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontMultitenantDistribution) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontMultitenantDistribution) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := c.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (c *jsiiProxy_CloudfrontMultitenantDistribution) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontMultitenantDistribution) MoveFromId(id *string) {
	if err := c.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveFromId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_CloudfrontMultitenantDistribution) MoveTo(moveTarget *string, index interface{}) {
	if err := c.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (c *jsiiProxy_CloudfrontMultitenantDistribution) MoveToId(id *string) {
	if err := c.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveToId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_CloudfrontMultitenantDistribution) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CloudfrontMultitenantDistribution) PutActiveTrustedKeyGroups(value interface{}) {
	if err := c.validatePutActiveTrustedKeyGroupsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putActiveTrustedKeyGroups",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontMultitenantDistribution) PutCacheBehavior(value interface{}) {
	if err := c.validatePutCacheBehaviorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putCacheBehavior",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontMultitenantDistribution) PutCustomErrorResponse(value interface{}) {
	if err := c.validatePutCustomErrorResponseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putCustomErrorResponse",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontMultitenantDistribution) PutDefaultCacheBehavior(value interface{}) {
	if err := c.validatePutDefaultCacheBehaviorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putDefaultCacheBehavior",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontMultitenantDistribution) PutOrigin(value interface{}) {
	if err := c.validatePutOriginParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putOrigin",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontMultitenantDistribution) PutOriginGroup(value interface{}) {
	if err := c.validatePutOriginGroupParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putOriginGroup",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontMultitenantDistribution) PutRestrictions(value interface{}) {
	if err := c.validatePutRestrictionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRestrictions",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontMultitenantDistribution) PutTenantConfig(value interface{}) {
	if err := c.validatePutTenantConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTenantConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontMultitenantDistribution) PutTimeouts(value *CloudfrontMultitenantDistributionTimeouts) {
	if err := c.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontMultitenantDistribution) PutViewerCertificate(value interface{}) {
	if err := c.validatePutViewerCertificateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putViewerCertificate",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontMultitenantDistribution) ResetActiveTrustedKeyGroups() {
	_jsii_.InvokeVoid(
		c,
		"resetActiveTrustedKeyGroups",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontMultitenantDistribution) ResetCacheBehavior() {
	_jsii_.InvokeVoid(
		c,
		"resetCacheBehavior",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontMultitenantDistribution) ResetCustomErrorResponse() {
	_jsii_.InvokeVoid(
		c,
		"resetCustomErrorResponse",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontMultitenantDistribution) ResetDefaultCacheBehavior() {
	_jsii_.InvokeVoid(
		c,
		"resetDefaultCacheBehavior",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontMultitenantDistribution) ResetDefaultRootObject() {
	_jsii_.InvokeVoid(
		c,
		"resetDefaultRootObject",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontMultitenantDistribution) ResetHttpVersion() {
	_jsii_.InvokeVoid(
		c,
		"resetHttpVersion",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontMultitenantDistribution) ResetOrigin() {
	_jsii_.InvokeVoid(
		c,
		"resetOrigin",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontMultitenantDistribution) ResetOriginGroup() {
	_jsii_.InvokeVoid(
		c,
		"resetOriginGroup",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontMultitenantDistribution) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontMultitenantDistribution) ResetRestrictions() {
	_jsii_.InvokeVoid(
		c,
		"resetRestrictions",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontMultitenantDistribution) ResetTags() {
	_jsii_.InvokeVoid(
		c,
		"resetTags",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontMultitenantDistribution) ResetTenantConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetTenantConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontMultitenantDistribution) ResetTimeouts() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontMultitenantDistribution) ResetViewerCertificate() {
	_jsii_.InvokeVoid(
		c,
		"resetViewerCertificate",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontMultitenantDistribution) ResetWebAclId() {
	_jsii_.InvokeVoid(
		c,
		"resetWebAclId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontMultitenantDistribution) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontMultitenantDistribution) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontMultitenantDistribution) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontMultitenantDistribution) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontMultitenantDistribution) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontMultitenantDistribution) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

