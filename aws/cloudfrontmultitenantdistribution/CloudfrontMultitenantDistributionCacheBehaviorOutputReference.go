// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudfrontmultitenantdistribution

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v21/jsii"

	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v21/cloudfrontmultitenantdistribution/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudfrontMultitenantDistributionCacheBehaviorOutputReference interface {
	cdktf.ComplexObject
	AllowedMethods() CloudfrontMultitenantDistributionCacheBehaviorAllowedMethodsList
	AllowedMethodsInput() interface{}
	CachePolicyId() *string
	SetCachePolicyId(val *string)
	CachePolicyIdInput() *string
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
	Compress() interface{}
	SetCompress(val interface{})
	CompressInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	FieldLevelEncryptionId() *string
	SetFieldLevelEncryptionId(val *string)
	FieldLevelEncryptionIdInput() *string
	// Experimental.
	Fqn() *string
	FunctionAssociation() CloudfrontMultitenantDistributionCacheBehaviorFunctionAssociationList
	FunctionAssociationInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LambdaFunctionAssociation() CloudfrontMultitenantDistributionCacheBehaviorLambdaFunctionAssociationList
	LambdaFunctionAssociationInput() interface{}
	OriginRequestPolicyId() *string
	SetOriginRequestPolicyId(val *string)
	OriginRequestPolicyIdInput() *string
	PathPattern() *string
	SetPathPattern(val *string)
	PathPatternInput() *string
	RealtimeLogConfigArn() *string
	SetRealtimeLogConfigArn(val *string)
	RealtimeLogConfigArnInput() *string
	ResponseHeadersPolicyId() *string
	SetResponseHeadersPolicyId(val *string)
	ResponseHeadersPolicyIdInput() *string
	TargetOriginId() *string
	SetTargetOriginId(val *string)
	TargetOriginIdInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TrustedKeyGroups() CloudfrontMultitenantDistributionCacheBehaviorTrustedKeyGroupsList
	TrustedKeyGroupsInput() interface{}
	ViewerProtocolPolicy() *string
	SetViewerProtocolPolicy(val *string)
	ViewerProtocolPolicyInput() *string
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
	PutAllowedMethods(value interface{})
	PutFunctionAssociation(value interface{})
	PutLambdaFunctionAssociation(value interface{})
	PutTrustedKeyGroups(value interface{})
	ResetAllowedMethods()
	ResetCachePolicyId()
	ResetCompress()
	ResetFieldLevelEncryptionId()
	ResetFunctionAssociation()
	ResetLambdaFunctionAssociation()
	ResetOriginRequestPolicyId()
	ResetRealtimeLogConfigArn()
	ResetResponseHeadersPolicyId()
	ResetTrustedKeyGroups()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CloudfrontMultitenantDistributionCacheBehaviorOutputReference
type jsiiProxy_CloudfrontMultitenantDistributionCacheBehaviorOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionCacheBehaviorOutputReference) AllowedMethods() CloudfrontMultitenantDistributionCacheBehaviorAllowedMethodsList {
	var returns CloudfrontMultitenantDistributionCacheBehaviorAllowedMethodsList
	_jsii_.Get(
		j,
		"allowedMethods",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionCacheBehaviorOutputReference) AllowedMethodsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowedMethodsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionCacheBehaviorOutputReference) CachePolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cachePolicyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionCacheBehaviorOutputReference) CachePolicyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cachePolicyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionCacheBehaviorOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionCacheBehaviorOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionCacheBehaviorOutputReference) Compress() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"compress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionCacheBehaviorOutputReference) CompressInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"compressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionCacheBehaviorOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionCacheBehaviorOutputReference) FieldLevelEncryptionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fieldLevelEncryptionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionCacheBehaviorOutputReference) FieldLevelEncryptionIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fieldLevelEncryptionIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionCacheBehaviorOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionCacheBehaviorOutputReference) FunctionAssociation() CloudfrontMultitenantDistributionCacheBehaviorFunctionAssociationList {
	var returns CloudfrontMultitenantDistributionCacheBehaviorFunctionAssociationList
	_jsii_.Get(
		j,
		"functionAssociation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionCacheBehaviorOutputReference) FunctionAssociationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"functionAssociationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionCacheBehaviorOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionCacheBehaviorOutputReference) LambdaFunctionAssociation() CloudfrontMultitenantDistributionCacheBehaviorLambdaFunctionAssociationList {
	var returns CloudfrontMultitenantDistributionCacheBehaviorLambdaFunctionAssociationList
	_jsii_.Get(
		j,
		"lambdaFunctionAssociation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionCacheBehaviorOutputReference) LambdaFunctionAssociationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lambdaFunctionAssociationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionCacheBehaviorOutputReference) OriginRequestPolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originRequestPolicyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionCacheBehaviorOutputReference) OriginRequestPolicyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originRequestPolicyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionCacheBehaviorOutputReference) PathPattern() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathPattern",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionCacheBehaviorOutputReference) PathPatternInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathPatternInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionCacheBehaviorOutputReference) RealtimeLogConfigArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"realtimeLogConfigArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionCacheBehaviorOutputReference) RealtimeLogConfigArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"realtimeLogConfigArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionCacheBehaviorOutputReference) ResponseHeadersPolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseHeadersPolicyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionCacheBehaviorOutputReference) ResponseHeadersPolicyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseHeadersPolicyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionCacheBehaviorOutputReference) TargetOriginId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetOriginId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionCacheBehaviorOutputReference) TargetOriginIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetOriginIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionCacheBehaviorOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionCacheBehaviorOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionCacheBehaviorOutputReference) TrustedKeyGroups() CloudfrontMultitenantDistributionCacheBehaviorTrustedKeyGroupsList {
	var returns CloudfrontMultitenantDistributionCacheBehaviorTrustedKeyGroupsList
	_jsii_.Get(
		j,
		"trustedKeyGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionCacheBehaviorOutputReference) TrustedKeyGroupsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"trustedKeyGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionCacheBehaviorOutputReference) ViewerProtocolPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"viewerProtocolPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionCacheBehaviorOutputReference) ViewerProtocolPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"viewerProtocolPolicyInput",
		&returns,
	)
	return returns
}


func NewCloudfrontMultitenantDistributionCacheBehaviorOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) CloudfrontMultitenantDistributionCacheBehaviorOutputReference {
	_init_.Initialize()

	if err := validateNewCloudfrontMultitenantDistributionCacheBehaviorOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudfrontMultitenantDistributionCacheBehaviorOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-aws.cloudfrontMultitenantDistribution.CloudfrontMultitenantDistributionCacheBehaviorOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewCloudfrontMultitenantDistributionCacheBehaviorOutputReference_Override(c CloudfrontMultitenantDistributionCacheBehaviorOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.cloudfrontMultitenantDistribution.CloudfrontMultitenantDistributionCacheBehaviorOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionCacheBehaviorOutputReference)SetCachePolicyId(val *string) {
	if err := j.validateSetCachePolicyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cachePolicyId",
		val,
	)
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionCacheBehaviorOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionCacheBehaviorOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionCacheBehaviorOutputReference)SetCompress(val interface{}) {
	if err := j.validateSetCompressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"compress",
		val,
	)
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionCacheBehaviorOutputReference)SetFieldLevelEncryptionId(val *string) {
	if err := j.validateSetFieldLevelEncryptionIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fieldLevelEncryptionId",
		val,
	)
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionCacheBehaviorOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionCacheBehaviorOutputReference)SetOriginRequestPolicyId(val *string) {
	if err := j.validateSetOriginRequestPolicyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"originRequestPolicyId",
		val,
	)
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionCacheBehaviorOutputReference)SetPathPattern(val *string) {
	if err := j.validateSetPathPatternParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pathPattern",
		val,
	)
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionCacheBehaviorOutputReference)SetRealtimeLogConfigArn(val *string) {
	if err := j.validateSetRealtimeLogConfigArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"realtimeLogConfigArn",
		val,
	)
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionCacheBehaviorOutputReference)SetResponseHeadersPolicyId(val *string) {
	if err := j.validateSetResponseHeadersPolicyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"responseHeadersPolicyId",
		val,
	)
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionCacheBehaviorOutputReference)SetTargetOriginId(val *string) {
	if err := j.validateSetTargetOriginIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetOriginId",
		val,
	)
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionCacheBehaviorOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionCacheBehaviorOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CloudfrontMultitenantDistributionCacheBehaviorOutputReference)SetViewerProtocolPolicy(val *string) {
	if err := j.validateSetViewerProtocolPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"viewerProtocolPolicy",
		val,
	)
}

func (c *jsiiProxy_CloudfrontMultitenantDistributionCacheBehaviorOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontMultitenantDistributionCacheBehaviorOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CloudfrontMultitenantDistributionCacheBehaviorOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudfrontMultitenantDistributionCacheBehaviorOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CloudfrontMultitenantDistributionCacheBehaviorOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CloudfrontMultitenantDistributionCacheBehaviorOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CloudfrontMultitenantDistributionCacheBehaviorOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CloudfrontMultitenantDistributionCacheBehaviorOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CloudfrontMultitenantDistributionCacheBehaviorOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CloudfrontMultitenantDistributionCacheBehaviorOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CloudfrontMultitenantDistributionCacheBehaviorOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontMultitenantDistributionCacheBehaviorOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudfrontMultitenantDistributionCacheBehaviorOutputReference) PutAllowedMethods(value interface{}) {
	if err := c.validatePutAllowedMethodsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAllowedMethods",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontMultitenantDistributionCacheBehaviorOutputReference) PutFunctionAssociation(value interface{}) {
	if err := c.validatePutFunctionAssociationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putFunctionAssociation",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontMultitenantDistributionCacheBehaviorOutputReference) PutLambdaFunctionAssociation(value interface{}) {
	if err := c.validatePutLambdaFunctionAssociationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putLambdaFunctionAssociation",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontMultitenantDistributionCacheBehaviorOutputReference) PutTrustedKeyGroups(value interface{}) {
	if err := c.validatePutTrustedKeyGroupsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTrustedKeyGroups",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontMultitenantDistributionCacheBehaviorOutputReference) ResetAllowedMethods() {
	_jsii_.InvokeVoid(
		c,
		"resetAllowedMethods",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontMultitenantDistributionCacheBehaviorOutputReference) ResetCachePolicyId() {
	_jsii_.InvokeVoid(
		c,
		"resetCachePolicyId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontMultitenantDistributionCacheBehaviorOutputReference) ResetCompress() {
	_jsii_.InvokeVoid(
		c,
		"resetCompress",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontMultitenantDistributionCacheBehaviorOutputReference) ResetFieldLevelEncryptionId() {
	_jsii_.InvokeVoid(
		c,
		"resetFieldLevelEncryptionId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontMultitenantDistributionCacheBehaviorOutputReference) ResetFunctionAssociation() {
	_jsii_.InvokeVoid(
		c,
		"resetFunctionAssociation",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontMultitenantDistributionCacheBehaviorOutputReference) ResetLambdaFunctionAssociation() {
	_jsii_.InvokeVoid(
		c,
		"resetLambdaFunctionAssociation",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontMultitenantDistributionCacheBehaviorOutputReference) ResetOriginRequestPolicyId() {
	_jsii_.InvokeVoid(
		c,
		"resetOriginRequestPolicyId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontMultitenantDistributionCacheBehaviorOutputReference) ResetRealtimeLogConfigArn() {
	_jsii_.InvokeVoid(
		c,
		"resetRealtimeLogConfigArn",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontMultitenantDistributionCacheBehaviorOutputReference) ResetResponseHeadersPolicyId() {
	_jsii_.InvokeVoid(
		c,
		"resetResponseHeadersPolicyId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontMultitenantDistributionCacheBehaviorOutputReference) ResetTrustedKeyGroups() {
	_jsii_.InvokeVoid(
		c,
		"resetTrustedKeyGroups",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontMultitenantDistributionCacheBehaviorOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := c.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontMultitenantDistributionCacheBehaviorOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

