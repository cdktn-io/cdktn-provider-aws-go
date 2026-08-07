// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockevaluationjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/jsii"

	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/bedrockevaluationjob/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BedrockEvaluationJobEvaluationConfigHumanDatasetMetricConfigDatasetDatasetLocationList interface {
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
	Get(index *float64) BedrockEvaluationJobEvaluationConfigHumanDatasetMetricConfigDatasetDatasetLocationOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BedrockEvaluationJobEvaluationConfigHumanDatasetMetricConfigDatasetDatasetLocationList
type jsiiProxy_BedrockEvaluationJobEvaluationConfigHumanDatasetMetricConfigDatasetDatasetLocationList struct {
	internal.Type__cdktnComplexList
}

func (j *jsiiProxy_BedrockEvaluationJobEvaluationConfigHumanDatasetMetricConfigDatasetDatasetLocationList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockEvaluationJobEvaluationConfigHumanDatasetMetricConfigDatasetDatasetLocationList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockEvaluationJobEvaluationConfigHumanDatasetMetricConfigDatasetDatasetLocationList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockEvaluationJobEvaluationConfigHumanDatasetMetricConfigDatasetDatasetLocationList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockEvaluationJobEvaluationConfigHumanDatasetMetricConfigDatasetDatasetLocationList) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockEvaluationJobEvaluationConfigHumanDatasetMetricConfigDatasetDatasetLocationList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewBedrockEvaluationJobEvaluationConfigHumanDatasetMetricConfigDatasetDatasetLocationList(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) BedrockEvaluationJobEvaluationConfigHumanDatasetMetricConfigDatasetDatasetLocationList {
	_init_.Initialize()

	if err := validateNewBedrockEvaluationJobEvaluationConfigHumanDatasetMetricConfigDatasetDatasetLocationListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_BedrockEvaluationJobEvaluationConfigHumanDatasetMetricConfigDatasetDatasetLocationList{}

	_jsii_.Create(
		"@cdktn/provider-aws.bedrockEvaluationJob.BedrockEvaluationJobEvaluationConfigHumanDatasetMetricConfigDatasetDatasetLocationList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewBedrockEvaluationJobEvaluationConfigHumanDatasetMetricConfigDatasetDatasetLocationList_Override(b BedrockEvaluationJobEvaluationConfigHumanDatasetMetricConfigDatasetDatasetLocationList, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.bedrockEvaluationJob.BedrockEvaluationJobEvaluationConfigHumanDatasetMetricConfigDatasetDatasetLocationList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		b,
	)
}

func (j *jsiiProxy_BedrockEvaluationJobEvaluationConfigHumanDatasetMetricConfigDatasetDatasetLocationList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BedrockEvaluationJobEvaluationConfigHumanDatasetMetricConfigDatasetDatasetLocationList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BedrockEvaluationJobEvaluationConfigHumanDatasetMetricConfigDatasetDatasetLocationList)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_BedrockEvaluationJobEvaluationConfigHumanDatasetMetricConfigDatasetDatasetLocationList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (b *jsiiProxy_BedrockEvaluationJobEvaluationConfigHumanDatasetMetricConfigDatasetDatasetLocationList) AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator {
	if err := b.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktn.DynamicListTerraformIterator

	_jsii_.Invoke(
		b,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockEvaluationJobEvaluationConfigHumanDatasetMetricConfigDatasetDatasetLocationList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockEvaluationJobEvaluationConfigHumanDatasetMetricConfigDatasetDatasetLocationList) Get(index *float64) BedrockEvaluationJobEvaluationConfigHumanDatasetMetricConfigDatasetDatasetLocationOutputReference {
	if err := b.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns BedrockEvaluationJobEvaluationConfigHumanDatasetMetricConfigDatasetDatasetLocationOutputReference

	_jsii_.Invoke(
		b,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockEvaluationJobEvaluationConfigHumanDatasetMetricConfigDatasetDatasetLocationList) Resolve(context cdktn.IResolveContext) interface{} {
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

func (b *jsiiProxy_BedrockEvaluationJobEvaluationConfigHumanDatasetMetricConfigDatasetDatasetLocationList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

