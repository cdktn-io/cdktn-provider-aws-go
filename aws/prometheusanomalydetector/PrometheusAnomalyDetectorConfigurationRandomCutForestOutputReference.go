// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package prometheusanomalydetector

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/jsii"

	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/prometheusanomalydetector/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PrometheusAnomalyDetectorConfigurationRandomCutForestOutputReference interface {
	cdktn.ComplexObject
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
	IgnoreNearExpectedFromAbove() PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromAboveList
	IgnoreNearExpectedFromAboveInput() interface{}
	IgnoreNearExpectedFromBelow() PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromBelowList
	IgnoreNearExpectedFromBelowInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Query() *string
	SetQuery(val *string)
	QueryInput() *string
	SampleSize() *float64
	SetSampleSize(val *float64)
	SampleSizeInput() *float64
	ShingleSize() *float64
	SetShingleSize(val *float64)
	ShingleSizeInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
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
	PutIgnoreNearExpectedFromAbove(value interface{})
	PutIgnoreNearExpectedFromBelow(value interface{})
	ResetIgnoreNearExpectedFromAbove()
	ResetIgnoreNearExpectedFromBelow()
	ResetSampleSize()
	ResetShingleSize()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PrometheusAnomalyDetectorConfigurationRandomCutForestOutputReference
type jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestOutputReference) IgnoreNearExpectedFromAbove() PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromAboveList {
	var returns PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromAboveList
	_jsii_.Get(
		j,
		"ignoreNearExpectedFromAbove",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestOutputReference) IgnoreNearExpectedFromAboveInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreNearExpectedFromAboveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestOutputReference) IgnoreNearExpectedFromBelow() PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromBelowList {
	var returns PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromBelowList
	_jsii_.Get(
		j,
		"ignoreNearExpectedFromBelow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestOutputReference) IgnoreNearExpectedFromBelowInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreNearExpectedFromBelowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestOutputReference) Query() *string {
	var returns *string
	_jsii_.Get(
		j,
		"query",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestOutputReference) QueryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestOutputReference) SampleSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sampleSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestOutputReference) SampleSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sampleSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestOutputReference) ShingleSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"shingleSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestOutputReference) ShingleSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"shingleSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPrometheusAnomalyDetectorConfigurationRandomCutForestOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) PrometheusAnomalyDetectorConfigurationRandomCutForestOutputReference {
	_init_.Initialize()

	if err := validateNewPrometheusAnomalyDetectorConfigurationRandomCutForestOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-aws.prometheusAnomalyDetector.PrometheusAnomalyDetectorConfigurationRandomCutForestOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewPrometheusAnomalyDetectorConfigurationRandomCutForestOutputReference_Override(p PrometheusAnomalyDetectorConfigurationRandomCutForestOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.prometheusAnomalyDetector.PrometheusAnomalyDetectorConfigurationRandomCutForestOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		p,
	)
}

func (j *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestOutputReference)SetQuery(val *string) {
	if err := j.validateSetQueryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"query",
		val,
	)
}

func (j *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestOutputReference)SetSampleSize(val *float64) {
	if err := j.validateSetSampleSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sampleSize",
		val,
	)
}

func (j *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestOutputReference)SetShingleSize(val *float64) {
	if err := j.validateSetShingleSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shingleSize",
		val,
	)
}

func (j *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestOutputReference) PutIgnoreNearExpectedFromAbove(value interface{}) {
	if err := p.validatePutIgnoreNearExpectedFromAboveParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putIgnoreNearExpectedFromAbove",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestOutputReference) PutIgnoreNearExpectedFromBelow(value interface{}) {
	if err := p.validatePutIgnoreNearExpectedFromBelowParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putIgnoreNearExpectedFromBelow",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestOutputReference) ResetIgnoreNearExpectedFromAbove() {
	_jsii_.InvokeVoid(
		p,
		"resetIgnoreNearExpectedFromAbove",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestOutputReference) ResetIgnoreNearExpectedFromBelow() {
	_jsii_.InvokeVoid(
		p,
		"resetIgnoreNearExpectedFromBelow",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestOutputReference) ResetSampleSize() {
	_jsii_.InvokeVoid(
		p,
		"resetSampleSize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestOutputReference) ResetShingleSize() {
	_jsii_.InvokeVoid(
		p,
		"resetShingleSize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (p *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

