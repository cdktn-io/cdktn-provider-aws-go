// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package prometheusanomalydetector

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/jsii"

	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/prometheusanomalydetector/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromBelowOutputReference interface {
	cdktn.ComplexObject
	Amount() *float64
	SetAmount(val *float64)
	AmountInput() *float64
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Ratio() *float64
	SetRatio(val *float64)
	RatioInput() *float64
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
	ResetAmount()
	ResetRatio()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromBelowOutputReference
type jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromBelowOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromBelowOutputReference) Amount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"amount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromBelowOutputReference) AmountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"amountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromBelowOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromBelowOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromBelowOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromBelowOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromBelowOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromBelowOutputReference) Ratio() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ratio",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromBelowOutputReference) RatioInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ratioInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromBelowOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromBelowOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromBelowOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromBelowOutputReference {
	_init_.Initialize()

	if err := validateNewPrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromBelowOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromBelowOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-aws.prometheusAnomalyDetector.PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromBelowOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewPrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromBelowOutputReference_Override(p PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromBelowOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.prometheusAnomalyDetector.PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromBelowOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		p,
	)
}

func (j *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromBelowOutputReference)SetAmount(val *float64) {
	if err := j.validateSetAmountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"amount",
		val,
	)
}

func (j *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromBelowOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromBelowOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromBelowOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromBelowOutputReference)SetRatio(val *float64) {
	if err := j.validateSetRatioParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ratio",
		val,
	)
}

func (j *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromBelowOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromBelowOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromBelowOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromBelowOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromBelowOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromBelowOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromBelowOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromBelowOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromBelowOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromBelowOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromBelowOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromBelowOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromBelowOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromBelowOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromBelowOutputReference) ResetAmount() {
	_jsii_.InvokeVoid(
		p,
		"resetAmount",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromBelowOutputReference) ResetRatio() {
	_jsii_.InvokeVoid(
		p,
		"resetRatio",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromBelowOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (p *jsiiProxy_PrometheusAnomalyDetectorConfigurationRandomCutForestIgnoreNearExpectedFromBelowOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

