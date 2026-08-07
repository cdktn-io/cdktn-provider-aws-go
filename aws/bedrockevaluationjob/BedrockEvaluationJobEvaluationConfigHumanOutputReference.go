// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockevaluationjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/jsii"

	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/bedrockevaluationjob/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BedrockEvaluationJobEvaluationConfigHumanOutputReference interface {
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
	CustomMetric() BedrockEvaluationJobEvaluationConfigHumanCustomMetricList
	CustomMetricInput() interface{}
	DatasetMetricConfig() BedrockEvaluationJobEvaluationConfigHumanDatasetMetricConfigList
	DatasetMetricConfigInput() interface{}
	// Experimental.
	Fqn() *string
	HumanWorkflowConfig() BedrockEvaluationJobEvaluationConfigHumanHumanWorkflowConfigList
	HumanWorkflowConfigInput() interface{}
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
	PutCustomMetric(value interface{})
	PutDatasetMetricConfig(value interface{})
	PutHumanWorkflowConfig(value interface{})
	ResetCustomMetric()
	ResetDatasetMetricConfig()
	ResetHumanWorkflowConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BedrockEvaluationJobEvaluationConfigHumanOutputReference
type jsiiProxy_BedrockEvaluationJobEvaluationConfigHumanOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_BedrockEvaluationJobEvaluationConfigHumanOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockEvaluationJobEvaluationConfigHumanOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockEvaluationJobEvaluationConfigHumanOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockEvaluationJobEvaluationConfigHumanOutputReference) CustomMetric() BedrockEvaluationJobEvaluationConfigHumanCustomMetricList {
	var returns BedrockEvaluationJobEvaluationConfigHumanCustomMetricList
	_jsii_.Get(
		j,
		"customMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockEvaluationJobEvaluationConfigHumanOutputReference) CustomMetricInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customMetricInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockEvaluationJobEvaluationConfigHumanOutputReference) DatasetMetricConfig() BedrockEvaluationJobEvaluationConfigHumanDatasetMetricConfigList {
	var returns BedrockEvaluationJobEvaluationConfigHumanDatasetMetricConfigList
	_jsii_.Get(
		j,
		"datasetMetricConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockEvaluationJobEvaluationConfigHumanOutputReference) DatasetMetricConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"datasetMetricConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockEvaluationJobEvaluationConfigHumanOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockEvaluationJobEvaluationConfigHumanOutputReference) HumanWorkflowConfig() BedrockEvaluationJobEvaluationConfigHumanHumanWorkflowConfigList {
	var returns BedrockEvaluationJobEvaluationConfigHumanHumanWorkflowConfigList
	_jsii_.Get(
		j,
		"humanWorkflowConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockEvaluationJobEvaluationConfigHumanOutputReference) HumanWorkflowConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"humanWorkflowConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockEvaluationJobEvaluationConfigHumanOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockEvaluationJobEvaluationConfigHumanOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockEvaluationJobEvaluationConfigHumanOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBedrockEvaluationJobEvaluationConfigHumanOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) BedrockEvaluationJobEvaluationConfigHumanOutputReference {
	_init_.Initialize()

	if err := validateNewBedrockEvaluationJobEvaluationConfigHumanOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_BedrockEvaluationJobEvaluationConfigHumanOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-aws.bedrockEvaluationJob.BedrockEvaluationJobEvaluationConfigHumanOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewBedrockEvaluationJobEvaluationConfigHumanOutputReference_Override(b BedrockEvaluationJobEvaluationConfigHumanOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.bedrockEvaluationJob.BedrockEvaluationJobEvaluationConfigHumanOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		b,
	)
}

func (j *jsiiProxy_BedrockEvaluationJobEvaluationConfigHumanOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BedrockEvaluationJobEvaluationConfigHumanOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BedrockEvaluationJobEvaluationConfigHumanOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BedrockEvaluationJobEvaluationConfigHumanOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BedrockEvaluationJobEvaluationConfigHumanOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BedrockEvaluationJobEvaluationConfigHumanOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockEvaluationJobEvaluationConfigHumanOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := b.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockEvaluationJobEvaluationConfigHumanOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := b.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockEvaluationJobEvaluationConfigHumanOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := b.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockEvaluationJobEvaluationConfigHumanOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := b.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockEvaluationJobEvaluationConfigHumanOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := b.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockEvaluationJobEvaluationConfigHumanOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := b.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockEvaluationJobEvaluationConfigHumanOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := b.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockEvaluationJobEvaluationConfigHumanOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := b.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockEvaluationJobEvaluationConfigHumanOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := b.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockEvaluationJobEvaluationConfigHumanOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockEvaluationJobEvaluationConfigHumanOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockEvaluationJobEvaluationConfigHumanOutputReference) PutCustomMetric(value interface{}) {
	if err := b.validatePutCustomMetricParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putCustomMetric",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockEvaluationJobEvaluationConfigHumanOutputReference) PutDatasetMetricConfig(value interface{}) {
	if err := b.validatePutDatasetMetricConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putDatasetMetricConfig",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockEvaluationJobEvaluationConfigHumanOutputReference) PutHumanWorkflowConfig(value interface{}) {
	if err := b.validatePutHumanWorkflowConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putHumanWorkflowConfig",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockEvaluationJobEvaluationConfigHumanOutputReference) ResetCustomMetric() {
	_jsii_.InvokeVoid(
		b,
		"resetCustomMetric",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockEvaluationJobEvaluationConfigHumanOutputReference) ResetDatasetMetricConfig() {
	_jsii_.InvokeVoid(
		b,
		"resetDatasetMetricConfig",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockEvaluationJobEvaluationConfigHumanOutputReference) ResetHumanWorkflowConfig() {
	_jsii_.InvokeVoid(
		b,
		"resetHumanWorkflowConfig",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockEvaluationJobEvaluationConfigHumanOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (b *jsiiProxy_BedrockEvaluationJobEvaluationConfigHumanOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

