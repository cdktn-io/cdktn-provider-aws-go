// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package arcregionswitchplan

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/jsii"

	"github.com/cdktn-io/cdktn-provider-aws-go/aws/v25/arcregionswitchplan/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigOutputReference interface {
	cdktn.ComplexObject
	CapacityMonitoringApproach() *string
	SetCapacityMonitoringApproach(val *string)
	CapacityMonitoringApproachInput() *string
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
	EksClusters() ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigEksClustersList
	EksClustersInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	KubernetesResourceType() ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigKubernetesResourceTypeList
	KubernetesResourceTypeInput() interface{}
	ScalingResources() ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigScalingResourcesList
	ScalingResourcesInput() interface{}
	TargetPercent() *float64
	SetTargetPercent(val *float64)
	TargetPercentInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TimeoutMinutes() *float64
	SetTimeoutMinutes(val *float64)
	TimeoutMinutesInput() *float64
	Ungraceful() ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigUngracefulList
	UngracefulInput() interface{}
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
	PutEksClusters(value interface{})
	PutKubernetesResourceType(value interface{})
	PutScalingResources(value interface{})
	PutUngraceful(value interface{})
	ResetEksClusters()
	ResetKubernetesResourceType()
	ResetScalingResources()
	ResetTimeoutMinutes()
	ResetUngraceful()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigOutputReference
type jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigOutputReference) CapacityMonitoringApproach() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacityMonitoringApproach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigOutputReference) CapacityMonitoringApproachInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacityMonitoringApproachInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigOutputReference) EksClusters() ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigEksClustersList {
	var returns ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigEksClustersList
	_jsii_.Get(
		j,
		"eksClusters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigOutputReference) EksClustersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"eksClustersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigOutputReference) KubernetesResourceType() ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigKubernetesResourceTypeList {
	var returns ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigKubernetesResourceTypeList
	_jsii_.Get(
		j,
		"kubernetesResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigOutputReference) KubernetesResourceTypeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kubernetesResourceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigOutputReference) ScalingResources() ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigScalingResourcesList {
	var returns ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigScalingResourcesList
	_jsii_.Get(
		j,
		"scalingResources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigOutputReference) ScalingResourcesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scalingResourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigOutputReference) TargetPercent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetPercent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigOutputReference) TargetPercentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetPercentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigOutputReference) TimeoutMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigOutputReference) TimeoutMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigOutputReference) Ungraceful() ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigUngracefulList {
	var returns ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigUngracefulList
	_jsii_.Get(
		j,
		"ungraceful",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigOutputReference) UngracefulInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ungracefulInput",
		&returns,
	)
	return returns
}


func NewArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigOutputReference {
	_init_.Initialize()

	if err := validateNewArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-aws.arcregionswitchPlan.ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigOutputReference_Override(a ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.arcregionswitchPlan.ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigOutputReference)SetCapacityMonitoringApproach(val *string) {
	if err := j.validateSetCapacityMonitoringApproachParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"capacityMonitoringApproach",
		val,
	)
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigOutputReference)SetTargetPercent(val *float64) {
	if err := j.validateSetTargetPercentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetPercent",
		val,
	)
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigOutputReference)SetTimeoutMinutes(val *float64) {
	if err := j.validateSetTimeoutMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeoutMinutes",
		val,
	)
}

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigOutputReference) PutEksClusters(value interface{}) {
	if err := a.validatePutEksClustersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEksClusters",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigOutputReference) PutKubernetesResourceType(value interface{}) {
	if err := a.validatePutKubernetesResourceTypeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putKubernetesResourceType",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigOutputReference) PutScalingResources(value interface{}) {
	if err := a.validatePutScalingResourcesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putScalingResources",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigOutputReference) PutUngraceful(value interface{}) {
	if err := a.validatePutUngracefulParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putUngraceful",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigOutputReference) ResetEksClusters() {
	_jsii_.InvokeVoid(
		a,
		"resetEksClusters",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigOutputReference) ResetKubernetesResourceType() {
	_jsii_.InvokeVoid(
		a,
		"resetKubernetesResourceType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigOutputReference) ResetScalingResources() {
	_jsii_.InvokeVoid(
		a,
		"resetScalingResources",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigOutputReference) ResetTimeoutMinutes() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeoutMinutes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigOutputReference) ResetUngraceful() {
	_jsii_.InvokeVoid(
		a,
		"resetUngraceful",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_ArcregionswitchPlanWorkflowStepParallelConfigStepEksResourceScalingConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

