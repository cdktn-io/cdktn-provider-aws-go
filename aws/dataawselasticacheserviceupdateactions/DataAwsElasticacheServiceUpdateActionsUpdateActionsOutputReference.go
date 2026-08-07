// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawselasticacheserviceupdateactions

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/jsii"

	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/dataawselasticacheserviceupdateactions/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAwsElasticacheServiceUpdateActionsUpdateActionsOutputReference interface {
	cdktn.ComplexObject
	CacheClusterId() *string
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
	Engine() *string
	EstimatedUpdateTime() *string
	// Experimental.
	Fqn() *string
	InternalValue() *DataAwsElasticacheServiceUpdateActionsUpdateActions
	SetInternalValue(val *DataAwsElasticacheServiceUpdateActionsUpdateActions)
	RecommendedApplyByDate() *string
	ReleaseDate() *string
	ReplicationGroupId() *string
	ServiceUpdateName() *string
	ServiceUpdateSeverity() *string
	ServiceUpdateStatus() *string
	ServiceUpdateType() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	UpdateActionStatus() *string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataAwsElasticacheServiceUpdateActionsUpdateActionsOutputReference
type jsiiProxy_DataAwsElasticacheServiceUpdateActionsUpdateActionsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsElasticacheServiceUpdateActionsUpdateActionsOutputReference) CacheClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacheClusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticacheServiceUpdateActionsUpdateActionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticacheServiceUpdateActionsUpdateActionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticacheServiceUpdateActionsUpdateActionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticacheServiceUpdateActionsUpdateActionsOutputReference) Engine() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticacheServiceUpdateActionsUpdateActionsOutputReference) EstimatedUpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"estimatedUpdateTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticacheServiceUpdateActionsUpdateActionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticacheServiceUpdateActionsUpdateActionsOutputReference) InternalValue() *DataAwsElasticacheServiceUpdateActionsUpdateActions {
	var returns *DataAwsElasticacheServiceUpdateActionsUpdateActions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticacheServiceUpdateActionsUpdateActionsOutputReference) RecommendedApplyByDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recommendedApplyByDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticacheServiceUpdateActionsUpdateActionsOutputReference) ReleaseDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"releaseDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticacheServiceUpdateActionsUpdateActionsOutputReference) ReplicationGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticacheServiceUpdateActionsUpdateActionsOutputReference) ServiceUpdateName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceUpdateName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticacheServiceUpdateActionsUpdateActionsOutputReference) ServiceUpdateSeverity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceUpdateSeverity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticacheServiceUpdateActionsUpdateActionsOutputReference) ServiceUpdateStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceUpdateStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticacheServiceUpdateActionsUpdateActionsOutputReference) ServiceUpdateType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceUpdateType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticacheServiceUpdateActionsUpdateActionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticacheServiceUpdateActionsUpdateActionsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticacheServiceUpdateActionsUpdateActionsOutputReference) UpdateActionStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateActionStatus",
		&returns,
	)
	return returns
}


func NewDataAwsElasticacheServiceUpdateActionsUpdateActionsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataAwsElasticacheServiceUpdateActionsUpdateActionsOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsElasticacheServiceUpdateActionsUpdateActionsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsElasticacheServiceUpdateActionsUpdateActionsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-aws.dataAwsElasticacheServiceUpdateActions.DataAwsElasticacheServiceUpdateActionsUpdateActionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataAwsElasticacheServiceUpdateActionsUpdateActionsOutputReference_Override(d DataAwsElasticacheServiceUpdateActionsUpdateActionsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.dataAwsElasticacheServiceUpdateActions.DataAwsElasticacheServiceUpdateActionsUpdateActionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsElasticacheServiceUpdateActionsUpdateActionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsElasticacheServiceUpdateActionsUpdateActionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsElasticacheServiceUpdateActionsUpdateActionsOutputReference)SetInternalValue(val *DataAwsElasticacheServiceUpdateActionsUpdateActions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsElasticacheServiceUpdateActionsUpdateActionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsElasticacheServiceUpdateActionsUpdateActionsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsElasticacheServiceUpdateActionsUpdateActionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsElasticacheServiceUpdateActionsUpdateActionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsElasticacheServiceUpdateActionsUpdateActionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsElasticacheServiceUpdateActionsUpdateActionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsElasticacheServiceUpdateActionsUpdateActionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsElasticacheServiceUpdateActionsUpdateActionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsElasticacheServiceUpdateActionsUpdateActionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsElasticacheServiceUpdateActionsUpdateActionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsElasticacheServiceUpdateActionsUpdateActionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsElasticacheServiceUpdateActionsUpdateActionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsElasticacheServiceUpdateActionsUpdateActionsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsElasticacheServiceUpdateActionsUpdateActionsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsElasticacheServiceUpdateActionsUpdateActionsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsElasticacheServiceUpdateActionsUpdateActionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

