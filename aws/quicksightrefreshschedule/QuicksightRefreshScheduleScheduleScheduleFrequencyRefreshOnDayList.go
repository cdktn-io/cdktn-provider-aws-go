// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksightrefreshschedule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v23/jsii"

	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v23/quicksightrefreshschedule/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type QuicksightRefreshScheduleScheduleScheduleFrequencyRefreshOnDayList interface {
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
	Get(index *float64) QuicksightRefreshScheduleScheduleScheduleFrequencyRefreshOnDayOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for QuicksightRefreshScheduleScheduleScheduleFrequencyRefreshOnDayList
type jsiiProxy_QuicksightRefreshScheduleScheduleScheduleFrequencyRefreshOnDayList struct {
	internal.Type__cdktnComplexList
}

func (j *jsiiProxy_QuicksightRefreshScheduleScheduleScheduleFrequencyRefreshOnDayList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightRefreshScheduleScheduleScheduleFrequencyRefreshOnDayList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightRefreshScheduleScheduleScheduleFrequencyRefreshOnDayList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightRefreshScheduleScheduleScheduleFrequencyRefreshOnDayList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightRefreshScheduleScheduleScheduleFrequencyRefreshOnDayList) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightRefreshScheduleScheduleScheduleFrequencyRefreshOnDayList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewQuicksightRefreshScheduleScheduleScheduleFrequencyRefreshOnDayList(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) QuicksightRefreshScheduleScheduleScheduleFrequencyRefreshOnDayList {
	_init_.Initialize()

	if err := validateNewQuicksightRefreshScheduleScheduleScheduleFrequencyRefreshOnDayListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_QuicksightRefreshScheduleScheduleScheduleFrequencyRefreshOnDayList{}

	_jsii_.Create(
		"@cdktn/provider-aws.quicksightRefreshSchedule.QuicksightRefreshScheduleScheduleScheduleFrequencyRefreshOnDayList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewQuicksightRefreshScheduleScheduleScheduleFrequencyRefreshOnDayList_Override(q QuicksightRefreshScheduleScheduleScheduleFrequencyRefreshOnDayList, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.quicksightRefreshSchedule.QuicksightRefreshScheduleScheduleScheduleFrequencyRefreshOnDayList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		q,
	)
}

func (j *jsiiProxy_QuicksightRefreshScheduleScheduleScheduleFrequencyRefreshOnDayList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_QuicksightRefreshScheduleScheduleScheduleFrequencyRefreshOnDayList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_QuicksightRefreshScheduleScheduleScheduleFrequencyRefreshOnDayList)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_QuicksightRefreshScheduleScheduleScheduleFrequencyRefreshOnDayList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (q *jsiiProxy_QuicksightRefreshScheduleScheduleScheduleFrequencyRefreshOnDayList) AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator {
	if err := q.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktn.DynamicListTerraformIterator

	_jsii_.Invoke(
		q,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightRefreshScheduleScheduleScheduleFrequencyRefreshOnDayList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightRefreshScheduleScheduleScheduleFrequencyRefreshOnDayList) Get(index *float64) QuicksightRefreshScheduleScheduleScheduleFrequencyRefreshOnDayOutputReference {
	if err := q.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns QuicksightRefreshScheduleScheduleScheduleFrequencyRefreshOnDayOutputReference

	_jsii_.Invoke(
		q,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightRefreshScheduleScheduleScheduleFrequencyRefreshOnDayList) Resolve(context cdktn.IResolveContext) interface{} {
	if err := q.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		q,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightRefreshScheduleScheduleScheduleFrequencyRefreshOnDayList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

