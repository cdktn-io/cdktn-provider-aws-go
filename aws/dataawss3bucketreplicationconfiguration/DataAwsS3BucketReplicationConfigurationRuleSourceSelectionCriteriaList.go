// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawss3bucketreplicationconfiguration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v22/jsii"

	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v22/dataawss3bucketreplicationconfiguration/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAwsS3BucketReplicationConfigurationRuleSourceSelectionCriteriaList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Creating an iterator for this complex list.
	//
	// The list will be converted into a map with the mapKeyAttributeName as the key.
	// Experimental.
	AllWithMapKey(mapKeyAttributeName *string) cdktf.DynamicListTerraformIterator
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) DataAwsS3BucketReplicationConfigurationRuleSourceSelectionCriteriaOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataAwsS3BucketReplicationConfigurationRuleSourceSelectionCriteriaList
type jsiiProxy_DataAwsS3BucketReplicationConfigurationRuleSourceSelectionCriteriaList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_DataAwsS3BucketReplicationConfigurationRuleSourceSelectionCriteriaList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsS3BucketReplicationConfigurationRuleSourceSelectionCriteriaList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsS3BucketReplicationConfigurationRuleSourceSelectionCriteriaList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsS3BucketReplicationConfigurationRuleSourceSelectionCriteriaList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsS3BucketReplicationConfigurationRuleSourceSelectionCriteriaList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewDataAwsS3BucketReplicationConfigurationRuleSourceSelectionCriteriaList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) DataAwsS3BucketReplicationConfigurationRuleSourceSelectionCriteriaList {
	_init_.Initialize()

	if err := validateNewDataAwsS3BucketReplicationConfigurationRuleSourceSelectionCriteriaListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsS3BucketReplicationConfigurationRuleSourceSelectionCriteriaList{}

	_jsii_.Create(
		"@cdktn/provider-aws.dataAwsS3BucketReplicationConfiguration.DataAwsS3BucketReplicationConfigurationRuleSourceSelectionCriteriaList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewDataAwsS3BucketReplicationConfigurationRuleSourceSelectionCriteriaList_Override(d DataAwsS3BucketReplicationConfigurationRuleSourceSelectionCriteriaList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.dataAwsS3BucketReplicationConfiguration.DataAwsS3BucketReplicationConfigurationRuleSourceSelectionCriteriaList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsS3BucketReplicationConfigurationRuleSourceSelectionCriteriaList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsS3BucketReplicationConfigurationRuleSourceSelectionCriteriaList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataAwsS3BucketReplicationConfigurationRuleSourceSelectionCriteriaList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (d *jsiiProxy_DataAwsS3BucketReplicationConfigurationRuleSourceSelectionCriteriaList) AllWithMapKey(mapKeyAttributeName *string) cdktf.DynamicListTerraformIterator {
	if err := d.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktf.DynamicListTerraformIterator

	_jsii_.Invoke(
		d,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsS3BucketReplicationConfigurationRuleSourceSelectionCriteriaList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsS3BucketReplicationConfigurationRuleSourceSelectionCriteriaList) Get(index *float64) DataAwsS3BucketReplicationConfigurationRuleSourceSelectionCriteriaOutputReference {
	if err := d.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns DataAwsS3BucketReplicationConfigurationRuleSourceSelectionCriteriaOutputReference

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsS3BucketReplicationConfigurationRuleSourceSelectionCriteriaList) Resolve(context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsS3BucketReplicationConfigurationRuleSourceSelectionCriteriaList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

