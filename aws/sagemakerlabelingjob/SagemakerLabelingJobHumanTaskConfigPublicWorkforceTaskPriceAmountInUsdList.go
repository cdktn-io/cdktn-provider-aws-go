// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerlabelingjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/jsii"

	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/sagemakerlabelingjob/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SagemakerLabelingJobHumanTaskConfigPublicWorkforceTaskPriceAmountInUsdList interface {
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
	Get(index *float64) SagemakerLabelingJobHumanTaskConfigPublicWorkforceTaskPriceAmountInUsdOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SagemakerLabelingJobHumanTaskConfigPublicWorkforceTaskPriceAmountInUsdList
type jsiiProxy_SagemakerLabelingJobHumanTaskConfigPublicWorkforceTaskPriceAmountInUsdList struct {
	internal.Type__cdktnComplexList
}

func (j *jsiiProxy_SagemakerLabelingJobHumanTaskConfigPublicWorkforceTaskPriceAmountInUsdList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerLabelingJobHumanTaskConfigPublicWorkforceTaskPriceAmountInUsdList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerLabelingJobHumanTaskConfigPublicWorkforceTaskPriceAmountInUsdList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerLabelingJobHumanTaskConfigPublicWorkforceTaskPriceAmountInUsdList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerLabelingJobHumanTaskConfigPublicWorkforceTaskPriceAmountInUsdList) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerLabelingJobHumanTaskConfigPublicWorkforceTaskPriceAmountInUsdList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewSagemakerLabelingJobHumanTaskConfigPublicWorkforceTaskPriceAmountInUsdList(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) SagemakerLabelingJobHumanTaskConfigPublicWorkforceTaskPriceAmountInUsdList {
	_init_.Initialize()

	if err := validateNewSagemakerLabelingJobHumanTaskConfigPublicWorkforceTaskPriceAmountInUsdListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_SagemakerLabelingJobHumanTaskConfigPublicWorkforceTaskPriceAmountInUsdList{}

	_jsii_.Create(
		"@cdktn/provider-aws.sagemakerLabelingJob.SagemakerLabelingJobHumanTaskConfigPublicWorkforceTaskPriceAmountInUsdList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewSagemakerLabelingJobHumanTaskConfigPublicWorkforceTaskPriceAmountInUsdList_Override(s SagemakerLabelingJobHumanTaskConfigPublicWorkforceTaskPriceAmountInUsdList, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.sagemakerLabelingJob.SagemakerLabelingJobHumanTaskConfigPublicWorkforceTaskPriceAmountInUsdList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		s,
	)
}

func (j *jsiiProxy_SagemakerLabelingJobHumanTaskConfigPublicWorkforceTaskPriceAmountInUsdList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SagemakerLabelingJobHumanTaskConfigPublicWorkforceTaskPriceAmountInUsdList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SagemakerLabelingJobHumanTaskConfigPublicWorkforceTaskPriceAmountInUsdList)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_SagemakerLabelingJobHumanTaskConfigPublicWorkforceTaskPriceAmountInUsdList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (s *jsiiProxy_SagemakerLabelingJobHumanTaskConfigPublicWorkforceTaskPriceAmountInUsdList) AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator {
	if err := s.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktn.DynamicListTerraformIterator

	_jsii_.Invoke(
		s,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerLabelingJobHumanTaskConfigPublicWorkforceTaskPriceAmountInUsdList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerLabelingJobHumanTaskConfigPublicWorkforceTaskPriceAmountInUsdList) Get(index *float64) SagemakerLabelingJobHumanTaskConfigPublicWorkforceTaskPriceAmountInUsdOutputReference {
	if err := s.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns SagemakerLabelingJobHumanTaskConfigPublicWorkforceTaskPriceAmountInUsdOutputReference

	_jsii_.Invoke(
		s,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerLabelingJobHumanTaskConfigPublicWorkforceTaskPriceAmountInUsdList) Resolve(context cdktn.IResolveContext) interface{} {
	if err := s.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerLabelingJobHumanTaskConfigPublicWorkforceTaskPriceAmountInUsdList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

