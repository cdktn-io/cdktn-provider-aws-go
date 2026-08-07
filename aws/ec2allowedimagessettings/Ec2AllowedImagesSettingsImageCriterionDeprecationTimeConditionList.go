// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2allowedimagessettings

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/jsii"

	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/ec2allowedimagessettings/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Ec2AllowedImagesSettingsImageCriterionDeprecationTimeConditionList interface {
	cdktn.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	WrapsSet() *bool
	// Experimental.
	SetWrapsSet(val *bool)
	// Creating an iterator for this complex list.
	//
	// The list will be converted into a map with the mapKeyAttributeName as the key.
	// Experimental.
	AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) Ec2AllowedImagesSettingsImageCriterionDeprecationTimeConditionOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Ec2AllowedImagesSettingsImageCriterionDeprecationTimeConditionList
type jsiiProxy_Ec2AllowedImagesSettingsImageCriterionDeprecationTimeConditionList struct {
	internal.Type__cdktnComplexList
}

func (j *jsiiProxy_Ec2AllowedImagesSettingsImageCriterionDeprecationTimeConditionList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2AllowedImagesSettingsImageCriterionDeprecationTimeConditionList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2AllowedImagesSettingsImageCriterionDeprecationTimeConditionList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2AllowedImagesSettingsImageCriterionDeprecationTimeConditionList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2AllowedImagesSettingsImageCriterionDeprecationTimeConditionList) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2AllowedImagesSettingsImageCriterionDeprecationTimeConditionList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewEc2AllowedImagesSettingsImageCriterionDeprecationTimeConditionList(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) Ec2AllowedImagesSettingsImageCriterionDeprecationTimeConditionList {
	_init_.Initialize()

	if err := validateNewEc2AllowedImagesSettingsImageCriterionDeprecationTimeConditionListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_Ec2AllowedImagesSettingsImageCriterionDeprecationTimeConditionList{}

	_jsii_.Create(
		"@cdktn/provider-aws.ec2AllowedImagesSettings.Ec2AllowedImagesSettingsImageCriterionDeprecationTimeConditionList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewEc2AllowedImagesSettingsImageCriterionDeprecationTimeConditionList_Override(e Ec2AllowedImagesSettingsImageCriterionDeprecationTimeConditionList, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.ec2AllowedImagesSettings.Ec2AllowedImagesSettingsImageCriterionDeprecationTimeConditionList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		e,
	)
}

func (j *jsiiProxy_Ec2AllowedImagesSettingsImageCriterionDeprecationTimeConditionList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Ec2AllowedImagesSettingsImageCriterionDeprecationTimeConditionList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Ec2AllowedImagesSettingsImageCriterionDeprecationTimeConditionList)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_Ec2AllowedImagesSettingsImageCriterionDeprecationTimeConditionList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (e *jsiiProxy_Ec2AllowedImagesSettingsImageCriterionDeprecationTimeConditionList) AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator {
	if err := e.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktn.DynamicListTerraformIterator

	_jsii_.Invoke(
		e,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2AllowedImagesSettingsImageCriterionDeprecationTimeConditionList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2AllowedImagesSettingsImageCriterionDeprecationTimeConditionList) Get(index *float64) Ec2AllowedImagesSettingsImageCriterionDeprecationTimeConditionOutputReference {
	if err := e.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns Ec2AllowedImagesSettingsImageCriterionDeprecationTimeConditionOutputReference

	_jsii_.Invoke(
		e,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2AllowedImagesSettingsImageCriterionDeprecationTimeConditionList) Resolve(context cdktn.IResolveContext) interface{} {
	if err := e.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2AllowedImagesSettingsImageCriterionDeprecationTimeConditionList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

