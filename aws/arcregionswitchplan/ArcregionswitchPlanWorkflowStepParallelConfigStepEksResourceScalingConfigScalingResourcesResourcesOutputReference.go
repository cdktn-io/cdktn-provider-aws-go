// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package arcregionswitchplan

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v22/jsii"

	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v22/arcregionswitchplan/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigScalingResourcesResourcesOutputReference interface {
	cdktf.ComplexObject
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
	HpaName() *string
	SetHpaName(val *string)
	HpaNameInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	Namespace() *string
	SetNamespace(val *string)
	NamespaceInput() *string
	ResourceName() *string
	SetResourceName(val *string)
	ResourceNameInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	ResetHpaName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigScalingResourcesResourcesOutputReference
type jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigScalingResourcesResourcesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigScalingResourcesResourcesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigScalingResourcesResourcesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigScalingResourcesResourcesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigScalingResourcesResourcesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigScalingResourcesResourcesOutputReference) HpaName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hpaName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigScalingResourcesResourcesOutputReference) HpaNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hpaNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigScalingResourcesResourcesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigScalingResourcesResourcesOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigScalingResourcesResourcesOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigScalingResourcesResourcesOutputReference) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigScalingResourcesResourcesOutputReference) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigScalingResourcesResourcesOutputReference) ResourceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigScalingResourcesResourcesOutputReference) ResourceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigScalingResourcesResourcesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigScalingResourcesResourcesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigScalingResourcesResourcesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigScalingResourcesResourcesOutputReference {
	_init_.Initialize()

	if err := validateNewArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigScalingResourcesResourcesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigScalingResourcesResourcesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-aws.arcregionswitchPlan.ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigScalingResourcesResourcesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigScalingResourcesResourcesOutputReference_Override(a ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigScalingResourcesResourcesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.arcregionswitchPlan.ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigScalingResourcesResourcesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigScalingResourcesResourcesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigScalingResourcesResourcesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigScalingResourcesResourcesOutputReference)SetHpaName(val *string) {
	if err := j.validateSetHpaNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hpaName",
		val,
	)
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigScalingResourcesResourcesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigScalingResourcesResourcesOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigScalingResourcesResourcesOutputReference)SetNamespace(val *string) {
	if err := j.validateSetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigScalingResourcesResourcesOutputReference)SetResourceName(val *string) {
	if err := j.validateSetResourceNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceName",
		val,
	)
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigScalingResourcesResourcesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigScalingResourcesResourcesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigScalingResourcesResourcesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigScalingResourcesResourcesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigScalingResourcesResourcesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigScalingResourcesResourcesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigScalingResourcesResourcesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigScalingResourcesResourcesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigScalingResourcesResourcesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigScalingResourcesResourcesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigScalingResourcesResourcesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigScalingResourcesResourcesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigScalingResourcesResourcesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigScalingResourcesResourcesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigScalingResourcesResourcesOutputReference) ResetHpaName() {
	_jsii_.InvokeVoid(
		a,
		"resetHpaName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigScalingResourcesResourcesOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigScalingResourcesResourcesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

