// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package resiliencehubv2inputsource

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/jsii"

	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/resiliencehubv2inputsource/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Resiliencehubv2InputSourceResourceConfigurationOutputReference interface {
	cdktn.ComplexObject
	CfnStackArn() *string
	SetCfnStackArn(val *string)
	CfnStackArnInput() *string
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
	DesignFileS3Url() *string
	SetDesignFileS3Url(val *string)
	DesignFileS3UrlInput() *string
	Eks() Resiliencehubv2InputSourceResourceConfigurationEksList
	EksInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	ResourceTag() Resiliencehubv2InputSourceResourceConfigurationResourceTagList
	ResourceTagInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TfStateFileUrl() *string
	SetTfStateFileUrl(val *string)
	TfStateFileUrlInput() *string
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktn.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	PutEks(value interface{})
	PutResourceTag(value interface{})
	ResetCfnStackArn()
	ResetDesignFileS3Url()
	ResetEks()
	ResetResourceTag()
	ResetTfStateFileUrl()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Resiliencehubv2InputSourceResourceConfigurationOutputReference
type jsiiProxy_Resiliencehubv2InputSourceResourceConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_Resiliencehubv2InputSourceResourceConfigurationOutputReference) CfnStackArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnStackArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resiliencehubv2InputSourceResourceConfigurationOutputReference) CfnStackArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnStackArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resiliencehubv2InputSourceResourceConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resiliencehubv2InputSourceResourceConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resiliencehubv2InputSourceResourceConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resiliencehubv2InputSourceResourceConfigurationOutputReference) DesignFileS3Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"designFileS3Url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resiliencehubv2InputSourceResourceConfigurationOutputReference) DesignFileS3UrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"designFileS3UrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resiliencehubv2InputSourceResourceConfigurationOutputReference) Eks() Resiliencehubv2InputSourceResourceConfigurationEksList {
	var returns Resiliencehubv2InputSourceResourceConfigurationEksList
	_jsii_.Get(
		j,
		"eks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resiliencehubv2InputSourceResourceConfigurationOutputReference) EksInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"eksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resiliencehubv2InputSourceResourceConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resiliencehubv2InputSourceResourceConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resiliencehubv2InputSourceResourceConfigurationOutputReference) ResourceTag() Resiliencehubv2InputSourceResourceConfigurationResourceTagList {
	var returns Resiliencehubv2InputSourceResourceConfigurationResourceTagList
	_jsii_.Get(
		j,
		"resourceTag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resiliencehubv2InputSourceResourceConfigurationOutputReference) ResourceTagInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceTagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resiliencehubv2InputSourceResourceConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resiliencehubv2InputSourceResourceConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resiliencehubv2InputSourceResourceConfigurationOutputReference) TfStateFileUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tfStateFileUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resiliencehubv2InputSourceResourceConfigurationOutputReference) TfStateFileUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tfStateFileUrlInput",
		&returns,
	)
	return returns
}


func NewResiliencehubv2InputSourceResourceConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) Resiliencehubv2InputSourceResourceConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewResiliencehubv2InputSourceResourceConfigurationOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_Resiliencehubv2InputSourceResourceConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-aws.resiliencehubv2InputSource.Resiliencehubv2InputSourceResourceConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewResiliencehubv2InputSourceResourceConfigurationOutputReference_Override(r Resiliencehubv2InputSourceResourceConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.resiliencehubv2InputSource.Resiliencehubv2InputSourceResourceConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		r,
	)
}

func (j *jsiiProxy_Resiliencehubv2InputSourceResourceConfigurationOutputReference)SetCfnStackArn(val *string) {
	if err := j.validateSetCfnStackArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cfnStackArn",
		val,
	)
}

func (j *jsiiProxy_Resiliencehubv2InputSourceResourceConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Resiliencehubv2InputSourceResourceConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Resiliencehubv2InputSourceResourceConfigurationOutputReference)SetDesignFileS3Url(val *string) {
	if err := j.validateSetDesignFileS3UrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"designFileS3Url",
		val,
	)
}

func (j *jsiiProxy_Resiliencehubv2InputSourceResourceConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Resiliencehubv2InputSourceResourceConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Resiliencehubv2InputSourceResourceConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_Resiliencehubv2InputSourceResourceConfigurationOutputReference)SetTfStateFileUrl(val *string) {
	if err := j.validateSetTfStateFileUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tfStateFileUrl",
		val,
	)
}

func (r *jsiiProxy_Resiliencehubv2InputSourceResourceConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Resiliencehubv2InputSourceResourceConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := r.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Resiliencehubv2InputSourceResourceConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := r.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Resiliencehubv2InputSourceResourceConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := r.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		r,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Resiliencehubv2InputSourceResourceConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := r.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Resiliencehubv2InputSourceResourceConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := r.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Resiliencehubv2InputSourceResourceConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := r.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		r,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Resiliencehubv2InputSourceResourceConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := r.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		r,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Resiliencehubv2InputSourceResourceConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := r.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Resiliencehubv2InputSourceResourceConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := r.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		r,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Resiliencehubv2InputSourceResourceConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Resiliencehubv2InputSourceResourceConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := r.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Resiliencehubv2InputSourceResourceConfigurationOutputReference) PutEks(value interface{}) {
	if err := r.validatePutEksParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putEks",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_Resiliencehubv2InputSourceResourceConfigurationOutputReference) PutResourceTag(value interface{}) {
	if err := r.validatePutResourceTagParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putResourceTag",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_Resiliencehubv2InputSourceResourceConfigurationOutputReference) ResetCfnStackArn() {
	_jsii_.InvokeVoid(
		r,
		"resetCfnStackArn",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Resiliencehubv2InputSourceResourceConfigurationOutputReference) ResetDesignFileS3Url() {
	_jsii_.InvokeVoid(
		r,
		"resetDesignFileS3Url",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Resiliencehubv2InputSourceResourceConfigurationOutputReference) ResetEks() {
	_jsii_.InvokeVoid(
		r,
		"resetEks",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Resiliencehubv2InputSourceResourceConfigurationOutputReference) ResetResourceTag() {
	_jsii_.InvokeVoid(
		r,
		"resetResourceTag",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Resiliencehubv2InputSourceResourceConfigurationOutputReference) ResetTfStateFileUrl() {
	_jsii_.InvokeVoid(
		r,
		"resetTfStateFileUrl",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Resiliencehubv2InputSourceResourceConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := r.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		r,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Resiliencehubv2InputSourceResourceConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

