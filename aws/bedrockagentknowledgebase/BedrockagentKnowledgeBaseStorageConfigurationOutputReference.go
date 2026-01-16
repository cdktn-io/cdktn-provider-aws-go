// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockagentknowledgebase

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v21/jsii"

	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v21/bedrockagentknowledgebase/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BedrockagentKnowledgeBaseStorageConfigurationOutputReference interface {
	cdktf.ComplexObject
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MongoDbAtlasConfiguration() BedrockagentKnowledgeBaseStorageConfigurationMongoDbAtlasConfigurationList
	MongoDbAtlasConfigurationInput() interface{}
	NeptuneAnalyticsConfiguration() BedrockagentKnowledgeBaseStorageConfigurationNeptuneAnalyticsConfigurationList
	NeptuneAnalyticsConfigurationInput() interface{}
	OpensearchManagedClusterConfiguration() BedrockagentKnowledgeBaseStorageConfigurationOpensearchManagedClusterConfigurationList
	OpensearchManagedClusterConfigurationInput() interface{}
	OpensearchServerlessConfiguration() BedrockagentKnowledgeBaseStorageConfigurationOpensearchServerlessConfigurationList
	OpensearchServerlessConfigurationInput() interface{}
	PineconeConfiguration() BedrockagentKnowledgeBaseStorageConfigurationPineconeConfigurationList
	PineconeConfigurationInput() interface{}
	RdsConfiguration() BedrockagentKnowledgeBaseStorageConfigurationRdsConfigurationList
	RdsConfigurationInput() interface{}
	RedisEnterpriseCloudConfiguration() BedrockagentKnowledgeBaseStorageConfigurationRedisEnterpriseCloudConfigurationList
	RedisEnterpriseCloudConfigurationInput() interface{}
	S3VectorsConfiguration() BedrockagentKnowledgeBaseStorageConfigurationS3VectorsConfigurationList
	S3VectorsConfigurationInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeInput() *string
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	PutMongoDbAtlasConfiguration(value interface{})
	PutNeptuneAnalyticsConfiguration(value interface{})
	PutOpensearchManagedClusterConfiguration(value interface{})
	PutOpensearchServerlessConfiguration(value interface{})
	PutPineconeConfiguration(value interface{})
	PutRdsConfiguration(value interface{})
	PutRedisEnterpriseCloudConfiguration(value interface{})
	PutS3VectorsConfiguration(value interface{})
	ResetMongoDbAtlasConfiguration()
	ResetNeptuneAnalyticsConfiguration()
	ResetOpensearchManagedClusterConfiguration()
	ResetOpensearchServerlessConfiguration()
	ResetPineconeConfiguration()
	ResetRdsConfiguration()
	ResetRedisEnterpriseCloudConfiguration()
	ResetS3VectorsConfiguration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BedrockagentKnowledgeBaseStorageConfigurationOutputReference
type jsiiProxy_BedrockagentKnowledgeBaseStorageConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BedrockagentKnowledgeBaseStorageConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentKnowledgeBaseStorageConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentKnowledgeBaseStorageConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentKnowledgeBaseStorageConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentKnowledgeBaseStorageConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentKnowledgeBaseStorageConfigurationOutputReference) MongoDbAtlasConfiguration() BedrockagentKnowledgeBaseStorageConfigurationMongoDbAtlasConfigurationList {
	var returns BedrockagentKnowledgeBaseStorageConfigurationMongoDbAtlasConfigurationList
	_jsii_.Get(
		j,
		"mongoDbAtlasConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentKnowledgeBaseStorageConfigurationOutputReference) MongoDbAtlasConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mongoDbAtlasConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentKnowledgeBaseStorageConfigurationOutputReference) NeptuneAnalyticsConfiguration() BedrockagentKnowledgeBaseStorageConfigurationNeptuneAnalyticsConfigurationList {
	var returns BedrockagentKnowledgeBaseStorageConfigurationNeptuneAnalyticsConfigurationList
	_jsii_.Get(
		j,
		"neptuneAnalyticsConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentKnowledgeBaseStorageConfigurationOutputReference) NeptuneAnalyticsConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"neptuneAnalyticsConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentKnowledgeBaseStorageConfigurationOutputReference) OpensearchManagedClusterConfiguration() BedrockagentKnowledgeBaseStorageConfigurationOpensearchManagedClusterConfigurationList {
	var returns BedrockagentKnowledgeBaseStorageConfigurationOpensearchManagedClusterConfigurationList
	_jsii_.Get(
		j,
		"opensearchManagedClusterConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentKnowledgeBaseStorageConfigurationOutputReference) OpensearchManagedClusterConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"opensearchManagedClusterConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentKnowledgeBaseStorageConfigurationOutputReference) OpensearchServerlessConfiguration() BedrockagentKnowledgeBaseStorageConfigurationOpensearchServerlessConfigurationList {
	var returns BedrockagentKnowledgeBaseStorageConfigurationOpensearchServerlessConfigurationList
	_jsii_.Get(
		j,
		"opensearchServerlessConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentKnowledgeBaseStorageConfigurationOutputReference) OpensearchServerlessConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"opensearchServerlessConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentKnowledgeBaseStorageConfigurationOutputReference) PineconeConfiguration() BedrockagentKnowledgeBaseStorageConfigurationPineconeConfigurationList {
	var returns BedrockagentKnowledgeBaseStorageConfigurationPineconeConfigurationList
	_jsii_.Get(
		j,
		"pineconeConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentKnowledgeBaseStorageConfigurationOutputReference) PineconeConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pineconeConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentKnowledgeBaseStorageConfigurationOutputReference) RdsConfiguration() BedrockagentKnowledgeBaseStorageConfigurationRdsConfigurationList {
	var returns BedrockagentKnowledgeBaseStorageConfigurationRdsConfigurationList
	_jsii_.Get(
		j,
		"rdsConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentKnowledgeBaseStorageConfigurationOutputReference) RdsConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rdsConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentKnowledgeBaseStorageConfigurationOutputReference) RedisEnterpriseCloudConfiguration() BedrockagentKnowledgeBaseStorageConfigurationRedisEnterpriseCloudConfigurationList {
	var returns BedrockagentKnowledgeBaseStorageConfigurationRedisEnterpriseCloudConfigurationList
	_jsii_.Get(
		j,
		"redisEnterpriseCloudConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentKnowledgeBaseStorageConfigurationOutputReference) RedisEnterpriseCloudConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"redisEnterpriseCloudConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentKnowledgeBaseStorageConfigurationOutputReference) S3VectorsConfiguration() BedrockagentKnowledgeBaseStorageConfigurationS3VectorsConfigurationList {
	var returns BedrockagentKnowledgeBaseStorageConfigurationS3VectorsConfigurationList
	_jsii_.Get(
		j,
		"s3VectorsConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentKnowledgeBaseStorageConfigurationOutputReference) S3VectorsConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"s3VectorsConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentKnowledgeBaseStorageConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentKnowledgeBaseStorageConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentKnowledgeBaseStorageConfigurationOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentKnowledgeBaseStorageConfigurationOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewBedrockagentKnowledgeBaseStorageConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) BedrockagentKnowledgeBaseStorageConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewBedrockagentKnowledgeBaseStorageConfigurationOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_BedrockagentKnowledgeBaseStorageConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-aws.bedrockagentKnowledgeBase.BedrockagentKnowledgeBaseStorageConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewBedrockagentKnowledgeBaseStorageConfigurationOutputReference_Override(b BedrockagentKnowledgeBaseStorageConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.bedrockagentKnowledgeBase.BedrockagentKnowledgeBaseStorageConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		b,
	)
}

func (j *jsiiProxy_BedrockagentKnowledgeBaseStorageConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BedrockagentKnowledgeBaseStorageConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BedrockagentKnowledgeBaseStorageConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BedrockagentKnowledgeBaseStorageConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BedrockagentKnowledgeBaseStorageConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_BedrockagentKnowledgeBaseStorageConfigurationOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (b *jsiiProxy_BedrockagentKnowledgeBaseStorageConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentKnowledgeBaseStorageConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (b *jsiiProxy_BedrockagentKnowledgeBaseStorageConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := b.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentKnowledgeBaseStorageConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (b *jsiiProxy_BedrockagentKnowledgeBaseStorageConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (b *jsiiProxy_BedrockagentKnowledgeBaseStorageConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (b *jsiiProxy_BedrockagentKnowledgeBaseStorageConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (b *jsiiProxy_BedrockagentKnowledgeBaseStorageConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (b *jsiiProxy_BedrockagentKnowledgeBaseStorageConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (b *jsiiProxy_BedrockagentKnowledgeBaseStorageConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (b *jsiiProxy_BedrockagentKnowledgeBaseStorageConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentKnowledgeBaseStorageConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentKnowledgeBaseStorageConfigurationOutputReference) PutMongoDbAtlasConfiguration(value interface{}) {
	if err := b.validatePutMongoDbAtlasConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putMongoDbAtlasConfiguration",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentKnowledgeBaseStorageConfigurationOutputReference) PutNeptuneAnalyticsConfiguration(value interface{}) {
	if err := b.validatePutNeptuneAnalyticsConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putNeptuneAnalyticsConfiguration",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentKnowledgeBaseStorageConfigurationOutputReference) PutOpensearchManagedClusterConfiguration(value interface{}) {
	if err := b.validatePutOpensearchManagedClusterConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putOpensearchManagedClusterConfiguration",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentKnowledgeBaseStorageConfigurationOutputReference) PutOpensearchServerlessConfiguration(value interface{}) {
	if err := b.validatePutOpensearchServerlessConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putOpensearchServerlessConfiguration",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentKnowledgeBaseStorageConfigurationOutputReference) PutPineconeConfiguration(value interface{}) {
	if err := b.validatePutPineconeConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putPineconeConfiguration",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentKnowledgeBaseStorageConfigurationOutputReference) PutRdsConfiguration(value interface{}) {
	if err := b.validatePutRdsConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putRdsConfiguration",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentKnowledgeBaseStorageConfigurationOutputReference) PutRedisEnterpriseCloudConfiguration(value interface{}) {
	if err := b.validatePutRedisEnterpriseCloudConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putRedisEnterpriseCloudConfiguration",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentKnowledgeBaseStorageConfigurationOutputReference) PutS3VectorsConfiguration(value interface{}) {
	if err := b.validatePutS3VectorsConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putS3VectorsConfiguration",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentKnowledgeBaseStorageConfigurationOutputReference) ResetMongoDbAtlasConfiguration() {
	_jsii_.InvokeVoid(
		b,
		"resetMongoDbAtlasConfiguration",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentKnowledgeBaseStorageConfigurationOutputReference) ResetNeptuneAnalyticsConfiguration() {
	_jsii_.InvokeVoid(
		b,
		"resetNeptuneAnalyticsConfiguration",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentKnowledgeBaseStorageConfigurationOutputReference) ResetOpensearchManagedClusterConfiguration() {
	_jsii_.InvokeVoid(
		b,
		"resetOpensearchManagedClusterConfiguration",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentKnowledgeBaseStorageConfigurationOutputReference) ResetOpensearchServerlessConfiguration() {
	_jsii_.InvokeVoid(
		b,
		"resetOpensearchServerlessConfiguration",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentKnowledgeBaseStorageConfigurationOutputReference) ResetPineconeConfiguration() {
	_jsii_.InvokeVoid(
		b,
		"resetPineconeConfiguration",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentKnowledgeBaseStorageConfigurationOutputReference) ResetRdsConfiguration() {
	_jsii_.InvokeVoid(
		b,
		"resetRdsConfiguration",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentKnowledgeBaseStorageConfigurationOutputReference) ResetRedisEnterpriseCloudConfiguration() {
	_jsii_.InvokeVoid(
		b,
		"resetRedisEnterpriseCloudConfiguration",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentKnowledgeBaseStorageConfigurationOutputReference) ResetS3VectorsConfiguration() {
	_jsii_.InvokeVoid(
		b,
		"resetS3VectorsConfiguration",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentKnowledgeBaseStorageConfigurationOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := b.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentKnowledgeBaseStorageConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

