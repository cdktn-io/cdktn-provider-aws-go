// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package prometheusanomalydetector

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/jsii"

	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/prometheusanomalydetector/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromAboveList interface {
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
	Get(index *float64) PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromAboveOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromAboveList
type jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromAboveList struct {
	internal.Type__cdktnComplexList
}

func (j *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromAboveList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromAboveList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromAboveList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromAboveList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromAboveList) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromAboveList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewPrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromAboveList(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromAboveList {
	_init_.Initialize()

	if err := validateNewPrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromAboveListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromAboveList{}

	_jsii_.Create(
		"@cdktn/provider-aws.prometheusAnomalyDetector.PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromAboveList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewPrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromAboveList_Override(p PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromAboveList, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.prometheusAnomalyDetector.PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromAboveList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		p,
	)
}

func (j *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromAboveList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromAboveList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromAboveList)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromAboveList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (p *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromAboveList) AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator {
	if err := p.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktn.DynamicListTerraformIterator

	_jsii_.Invoke(
		p,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromAboveList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromAboveList) Get(index *float64) PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromAboveOutputReference {
	if err := p.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromAboveOutputReference

	_jsii_.Invoke(
		p,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromAboveList) Resolve(context cdktn.IResolveContext) interface{} {
	if err := p.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		p,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromAboveList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

